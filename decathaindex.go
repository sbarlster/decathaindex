package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"sort"
	"time"
)

// --- Types ---

type NominatimResult struct {
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
}

type OSRMResponse struct {
	Code   string `json:"code"`
	Routes []struct {
		Distance float64 `json:"distance"` // metres
		Duration float64 `json:"duration"` // seconds
	} `json:"routes"`
}

type Location struct {
	Name string
	Lat  float64
	Lon  float64
}

// --- Cache ---

const cacheFile = "locations_cache.csv"

// cacheKey is "City|Type" e.g. "Edinburgh|station"
type cacheKey struct {
	City string
	Type string // "station" or "decathlon"
}

// notFoundSentinel marks a cache row as a confirmed "no result" rather
// than an actual location. Stored in the name field; lat/lon are "0.0".
const notFoundSentinel = "NOT_FOUND"

// loadCache reads the CSV cache file into memory, split into:
//   - found:    resolved locations
//   - notFound: keys previously confirmed to have no result
//
// CSV columns: city,type,name,lat,lon
// Returns empty maps if the file doesn't exist yet.
func loadCache(path string) (found map[cacheKey]Location, notFound map[cacheKey]bool, err error) {
	found = make(map[cacheKey]Location)
	notFound = make(map[cacheKey]bool)

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return found, notFound, nil // no cache yet — that's fine
		}
		return nil, nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	for i, rec := range records {
		if i == 0 && len(rec) > 0 && rec[0] == "city" {
			continue // skip header row
		}
		if len(rec) != 5 {
			continue // skip malformed row
		}

		key := cacheKey{City: rec[0], Type: rec[1]}

		if rec[2] == notFoundSentinel {
			notFound[key] = true
			continue
		}

		var lat, lon float64
		fmt.Sscanf(rec[3], "%f", &lat)
		fmt.Sscanf(rec[4], "%f", &lon)
		found[key] = Location{Name: rec[2], Lat: lat, Lon: lon}
	}

	return found, notFound, nil
}

// appendToCache appends a single resolved location to the CSV cache file,
// writing a header first if the file is new.
func appendToCache(path string, key cacheKey, loc Location) error {
	needsHeader := false
	if _, err := os.Stat(path); os.IsNotExist(err) {
		needsHeader = true
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	if needsHeader {
		if err := w.Write([]string{"city", "type", "name", "lat", "lon"}); err != nil {
			return err
		}
	}

	return w.Write([]string{
		key.City,
		key.Type,
		loc.Name,
		fmt.Sprintf("%f", loc.Lat),
		fmt.Sprintf("%f", loc.Lon),
	})
}

// appendNotFoundToCache records a confirmed "no result" for a key, so
// future runs skip it without calling Nominatim again. To retry later
// (e.g. a new store opens), delete the corresponding row from the CSV.
func appendNotFoundToCache(path string, key cacheKey) error {
	needsHeader := false
	if _, err := os.Stat(path); os.IsNotExist(err) {
		needsHeader = true
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	if needsHeader {
		if err := w.Write([]string{"city", "type", "name", "lat", "lon"}); err != nil {
			return err
		}
	}

	return w.Write([]string{key.City, key.Type, notFoundSentinel, "0.0", "0.0"})
}

// --- Config ---

// City bundles everything needed to look up a city's station and
// Decathlon store, including the ISO 3166-1 alpha-2 country code used
// to scope the Nominatim search (e.g. "gb", "es").
type City struct {
	Name           string
	CountryCode    string
	StationQuery   string
	DecathlonQuery string
}

// countryName maps ISO 3166-1 alpha-2 codes to display names.
var countryName = map[string]string{
	"gb": "United Kingdom",
	"ie": "Ireland",
	"fr": "France",
	"es": "Spain",
	"pt": "Portugal",
	"de": "Germany",
	"it": "Italy",
	"be": "Belgium",
	"nl": "Netherlands",
	"ch": "Switzerland",
	"at": "Austria",
	"pl": "Poland",
	"cz": "Czech Republic",
	"ro": "Romania",
	"hu": "Hungary",
	"se": "Sweden",
	"gr": "Greece",
	"rs": "Serbia",
	"hr": "Croatia",
	"si": "Slovenia",
	"sk": "Slovakia",
	"lt": "Lithuania",
	"lv": "Latvia",
	"ee": "Estonia",
	"mt": "Malta",
}

var cities = []City{

	// --- United Kingdom (Scotland) ---
	{Name: "Edinburgh", CountryCode: "gb", StationQuery: "Edinburgh Waverley railway station, Edinburgh, Scotland", DecathlonQuery: "Decathlon Edinburgh, Hermiston Gait, Scotland"},
	{Name: "Glasgow", CountryCode: "gb", StationQuery: "Glasgow Central railway station, Glasgow, Scotland", DecathlonQuery: "Decathlon Glasgow, Scotland"},
	{Name: "Perth", CountryCode: "gb", StationQuery: "Perth railway station, Perth, Scotland", DecathlonQuery: "Decathlon Perth, Scotland"},
	{Name: "Stirling", CountryCode: "gb", StationQuery: "Stirling railway station, Stirling, Scotland", DecathlonQuery: "Decathlon Stirling, Scotland"},
	{Name: "Aberdeen", CountryCode: "gb", StationQuery: "Aberdeen railway station, Aberdeen, Scotland", DecathlonQuery: "Decathlon Aberdeen, Scotland"},
	{Name: "Dundee", CountryCode: "gb", StationQuery: "Dundee railway station, Dundee, Scotland", DecathlonQuery: "Decathlon Dundee, Scotland"},
	{Name: "Inverness", CountryCode: "gb", StationQuery: "Inverness railway station, Inverness, Scotland", DecathlonQuery: "Decathlon Inverness, Scotland"},

	// --- United Kingdom (England & Wales) ---
	{Name: "London", CountryCode: "gb", StationQuery: "London King's Cross railway station, London, England", DecathlonQuery: "Decathlon London, England"},
	{Name: "Manchester", CountryCode: "gb", StationQuery: "Manchester Piccadilly railway station, Manchester, England", DecathlonQuery: "Decathlon Manchester, England"},
	{Name: "Birmingham", CountryCode: "gb", StationQuery: "Birmingham New Street railway station, Birmingham, England", DecathlonQuery: "B&M Bargains, Wednesbury, England"},
	{Name: "Leeds", CountryCode: "gb", StationQuery: "Leeds railway station, Leeds, England", DecathlonQuery: "Decathlon Leeds, England"},
	{Name: "Bristol", CountryCode: "gb", StationQuery: "Bristol Temple Meads railway station, Bristol, England", DecathlonQuery: "Decathlon Bristol, England"},
	{Name: "Cardiff", CountryCode: "gb", StationQuery: "Cardiff Central railway station, Cardiff, Wales", DecathlonQuery: "Decathlon, Culverhouse Cross, St. Georges-super-Ely, Drope, Vale of Glamorgan, Wales"},
	{Name: "Southampton", CountryCode: "gb", StationQuery: "Southampton Central railway station, Southampton, England", DecathlonQuery: "Decathlon Southampton, England"},
	{Name: "Nottingham", CountryCode: "gb", StationQuery: "Nottingham railway station, Nottingham, England", DecathlonQuery: "Decathlon Nottingham, England"},
	{Name: "Norwich", CountryCode: "gb", StationQuery: "Norwich railway station, Norwich, England", DecathlonQuery: "Decathlon Norwich, England"},
	{Name: "Cambridge", CountryCode: "gb", StationQuery: "Cambridge railway station, Cambridge, England", DecathlonQuery: "Decathlon Cambridge, England"},
	{Name: "Newcastle", CountryCode: "gb", StationQuery: "Newcastle railway station, Newcastle, England", DecathlonQuery: "Decathlon, Dukesway, Team Valley Trading Estate, Lamesley, Whickham, Gateshead, Tyne and Wear, North East, England, NE11 0BD, United Kingdom"},
	{Name: "Liverpool", CountryCode: "gb", StationQuery: "Liverpool Lime Street railway station, Liverpool, England", DecathlonQuery: "Decathlon Liverpool, England"},
	{Name: "Sheffield", CountryCode: "gb", StationQuery: "Sheffield railway station, Sheffield, England", DecathlonQuery: "Decathlon Sheffield, England"},
	{Name: "Belfast", CountryCode: "gb", StationQuery: "Belfast Central railway station, Belfast, Northern Ireland", DecathlonQuery: "Decathlon, Airport Road West, Belfast City District, County Down, Northern Ireland, BT3 9EJ, United Kingdom"},
	{Name: "Portsmouth", CountryCode: "gb", StationQuery: "Portsmouth railway station, Portsmouth, England", DecathlonQuery: "Decathlon, Binnacle Way, Port Solent, Portsmouth, England, PO6 4FB, United Kingdom"},

	// --- Ireland ---
	{Name: "Dublin", CountryCode: "ie", StationQuery: "Dublin Heuston railway station, Dublin, Ireland", DecathlonQuery: "Decathlon Dublin, Ireland"},
	{Name: "Cork", CountryCode: "ie", StationQuery: "Cork Kent railway station, Cork, Ireland", DecathlonQuery: "Decathlon Cork, Ireland"},

	// --- France ---
	{Name: "Paris", CountryCode: "fr", StationQuery: "Paris Gare du Nord railway station, Paris, France", DecathlonQuery: "Decathlon Paris, France"},
	{Name: "Lyon", CountryCode: "fr", StationQuery: "Lyon Part-Dieu railway station, Lyon, France", DecathlonQuery: "Decathlon Lyon, France"},
	{Name: "Marseille", CountryCode: "fr", StationQuery: "Marseille Saint-Charles railway station, Marseille, France", DecathlonQuery: "Decathlon, Quai du Maroc, France"},
	{Name: "Toulouse", CountryCode: "fr", StationQuery: "Toulouse Matabiau railway station, Toulouse, France", DecathlonQuery: "Decathlon Toulouse, France"},
	{Name: "Bordeaux", CountryCode: "fr", StationQuery: "Bordeaux Saint-Jean railway station, Bordeaux, France", DecathlonQuery: "Decathlon, 130, Rue Sainte-Catherine, Bordeaux, Port of the Moon, Bordeaux Centre, Bordeaux, France"},
	{Name: "Lille", CountryCode: "fr", StationQuery: "Lille Flandres railway station, Lille, France", DecathlonQuery: "Decathlon, Rue de Béthune, Euralille, Lille-Centre, Lille, France"},
	{Name: "Nantes", CountryCode: "fr", StationQuery: "Nantes railway station, Nantes, France", DecathlonQuery: "Decathlon Nantes, France"},
	{Name: "Strasbourg", CountryCode: "fr", StationQuery: "Strasbourg railway station, Strasbourg, France", DecathlonQuery: "Decathlon Strasbourg, France"},
	{Name: "Rennes", CountryCode: "fr", StationQuery: "Rennes railway station, Rennes, France", DecathlonQuery: "Decathlon, 3 - 5 Rue le Bastard, 35000 Rennes, France"},
	{Name: "Nice", CountryCode: "fr", StationQuery: "Nice Ville railway station, Nice, France", DecathlonQuery: "Decathlon, Avenue Jean Médecin, Quartier Jean-Médecin, Nice, France"},
	{Name: "Montpellier", CountryCode: "fr", StationQuery: "Montpellier Saint-Roch railway station, Montpellier, France", DecathlonQuery: "Allée Anne-Marie de Backer, Comédie, Centre, Montpellier, France"},
	{Name: "Grenoble", CountryCode: "fr", StationQuery: "Grenoble railway station, Grenoble, France", DecathlonQuery: "Decathlon Grenoble, France"},

	// --- Spain ---
	{Name: "Madrid", CountryCode: "es", StationQuery: "Madrid Atocha railway station, Madrid, Spain", DecathlonQuery: "Decathlon, 1, Paseo de Santa María de la Cabeza, Palos de la Frontera, Arganzuela, Madrid, Spain"},
	{Name: "Barcelona", CountryCode: "es", StationQuery: "Barcelona Sants railway station, Barcelona, Spain", DecathlonQuery: "Decathlon Barcelona, Spain"},
	{Name: "Valencia", CountryCode: "es", StationQuery: "Valencia Joaquin Sorolla railway station, Valencia, Spain", DecathlonQuery: "Decathlon Valencia, Spain"},
	{Name: "Seville", CountryCode: "es", StationQuery: "Santa Justa railway station, Seville, Spain", DecathlonQuery: "Decathlon Seville, Spain"},
	{Name: "Zaragoza", CountryCode: "es", StationQuery: "Zaragoza Delicias railway station, Zaragoza, Spain", DecathlonQuery: "Decathlon, 10, Paseo de Sagasta, Zaragoza, Aragon, 50008, Spain"},
	{Name: "Malaga", CountryCode: "es", StationQuery: "Malaga Maria Zambrano railway station, Malaga, Spain", DecathlonQuery: "Decathlon, 3, Calle Strachan, Centro Histórico, Málaga, Málaga-Costa del Sol, Malaga, Spain"},
	{Name: "Bilbao", CountryCode: "es", StationQuery: "Bilbao Abando, Spain", DecathlonQuery: "Decathlon Bilbao, Spain"},
	{Name: "San Sebastian", CountryCode: "es", StationQuery: "San Sebastian railway station, Donostia, Spain", DecathlonQuery: "Decathlon City Donostia, Spain"},
	{Name: "Valladolid", CountryCode: "es", StationQuery: "Valladolid Campo Grande railway station, Valladolid, Spain", DecathlonQuery: "Decathlon Valladolid, Spain"},
	{Name: "Alicante", CountryCode: "es", StationQuery: "Alicante railway station, Alicante, Spain", DecathlonQuery: "Decathlon Alicante, Spain"},
	{Name: "Cordoba", CountryCode: "es", StationQuery: "Cordoba railway station, Cordoba, Spain", DecathlonQuery: "C. Concepción, 6, Centro, 14008 Córdoba, Spain"},
	{Name: "A Coruna", CountryCode: "es", StationQuery: "A Coruna railway station, A Coruna, Spain", DecathlonQuery: "Decathlon City, Rúa Real, Zalaeta, A Pescaría, A Coruña, Galicia, 15003, Spain"},

	// --- Portugal ---
	{Name: "Lisbon", CountryCode: "pt", StationQuery: "Lisboa Oriente railway station, Lisbon, Portugal", DecathlonQuery: "Decathlon Lisboa, Portugal"},
	{Name: "Porto", CountryCode: "pt", StationQuery: "Porto - São Bento, Porto, Portugal", DecathlonQuery: "Decathlon Porto, Portugal"},
	{Name: "Braga", CountryCode: "pt", StationQuery: "Braga railway station, Braga, Portugal", DecathlonQuery: "Decathlon Braga, Portugal"},
	{Name: "Coimbra", CountryCode: "pt", StationQuery: "Coimbra B railway station, Coimbra, Portugal", DecathlonQuery: "Decathlon Coimbra, Portugal"},

	// --- Germany ---
	{Name: "Berlin", CountryCode: "de", StationQuery: "Berlin Hauptbahnhof railway station, Berlin, Germany", DecathlonQuery: "Decathlon Berlin, Germany"},
	{Name: "Hamburg", CountryCode: "de", StationQuery: "Hamburg Hauptbahnhof railway station, Hamburg, Germany", DecathlonQuery: "Decathlon, 1, Mönckebergstraße, Altstadt, Hamburg-Mitte, Hamburg, 20095, Germany"},
	{Name: "Munich", CountryCode: "de", StationQuery: "Munich Hauptbahnhof railway station, Munich, Germany", DecathlonQuery: "Decathlon München, Germany"},
	{Name: "Cologne", CountryCode: "de", StationQuery: "Trankgasse 11, 50667 Köln, Germany", DecathlonQuery: "Decathlon, 80-90, Breite Straße, Neumarktviertel, Altstadt-Nord, Innenstadt, Cologne, North Rhine-Westphalia, 50667, Germany"},
	{Name: "Frankfurt", CountryCode: "de", StationQuery: "Frankfurt Hauptbahnhof railway station, Frankfurt, Germany", DecathlonQuery: "Decathlon Frankfurt, Germany"},
	{Name: "Stuttgart", CountryCode: "de", StationQuery: "Stuttgart Hauptbahnhof railway station, Stuttgart, Germany", DecathlonQuery: "Decathlon Stuttgart, Germany"},
	{Name: "Dusseldorf", CountryCode: "de", StationQuery: "Düsseldorf Hauptbahnhof railway station, Dusseldorf, Germany", DecathlonQuery: "Decathlon Düsseldorf, Germany"},
	{Name: "Dortmund", CountryCode: "de", StationQuery: "Dortmund Hbf, Paul-Winzen-Straße, Innenstadt Nord, Dortmund, North Rhine-Westphalia, 44147, Germany", DecathlonQuery: "Decathlon Dortmund, Germany"},
	{Name: "Leipzig", CountryCode: "de", StationQuery: "Leipzig Hauptbahnhof railway station, Leipzig, Germany", DecathlonQuery: "Decathlon, 44, Petersstraße, center, Mitte, Leipzig, Saxony, 04109, Germany"},
	{Name: "Dresden", CountryCode: "de", StationQuery: "Dresden Hauptbahnhof railway station, Dresden, Germany", DecathlonQuery: "Decathlon Dresden, Germany"},
	{Name: "Nuremberg", CountryCode: "de", StationQuery: "Nuremberg Hauptbahnhof railway station, Nuremberg, Germany", DecathlonQuery: "Decathlon Nürnberg, Germany"},
	{Name: "Hanover", CountryCode: "de", StationQuery: "Hanover Hauptbahnhof railway station, Hanover, Germany", DecathlonQuery: "Decathlon Hannover, Germany"},
	{Name: "Bremen", CountryCode: "de", StationQuery: "Bremen Hauptbahnhof railway station, Bremen, Germany", DecathlonQuery: "Decathlon, 1, Robert-Bosch-Straße, Brinkum, Stuhr, Landkreis Diepholz, Lower Saxony, 28816, Germany"},

	// --- Italy ---
	{Name: "Rome", CountryCode: "it", StationQuery: "Roma Termini railway station, Rome, Italy", DecathlonQuery: "Euronics, 44, Via Cesare Baronio, Municipio Roma VII, Rome, Roma Capitale, Lazio, 00179, Italy"},
	{Name: "Milan", CountryCode: "it", StationQuery: "Milano Centrale railway station, Milan, Italy", DecathlonQuery: "Decathlon Milano, Italy"},
	{Name: "Naples", CountryCode: "it", StationQuery: "Napoli Centrale railway station, Naples, Italy", DecathlonQuery: "P.za Giuseppe Garibaldi, 80142 Napoli NA, Italy"},
	{Name: "Turin", CountryCode: "it", StationQuery: "Torino Porta Nuova railway station, Turin, Italy", DecathlonQuery: "Decathlon, 85, Piazza Carlo Felice, Centro, Circoscrizione 1, Turin, Piedmont, 10123, Italy"},
	{Name: "Bologna", CountryCode: "it", StationQuery: "Bologna Centrale railway station, Bologna, Italy", DecathlonQuery: "Decathlon Bologna, Italy"},
	{Name: "Florence", CountryCode: "it", StationQuery: "Firenze Santa Maria Novella railway station, Florence, Italy", DecathlonQuery: "Decathlon Firenze, Italy"},
	{Name: "Genoa", CountryCode: "it", StationQuery: "Genova Piazza Principe railway station, Genoa, Italy", DecathlonQuery: "Decathlon Genova, Italy"},
	{Name: "Venice", CountryCode: "it", StationQuery: "Venezia Santa Lucia railway station, Venice, Italy", DecathlonQuery: "Decathlon Venezia, Italy"},
	{Name: "Verona", CountryCode: "it", StationQuery: "Verona Porta Nuova railway station, Verona, Italy", DecathlonQuery: "Decathlon, Strada Statale 434 Transpolesana, San Giovanni Lupatoto, Verona, Veneto, 37057, Italy"},
	{Name: "Bari", CountryCode: "it", StationQuery: "Bari Centrale railway station, Bari, Italy", DecathlonQuery: "Decathlon Bari, Italy"},

	// --- Belgium ---
	{Name: "Brussels", CountryCode: "be", StationQuery: "Brussels Midi railway station, Brussels, Belgium", DecathlonQuery: "Decathlon Bruxelles, Belgium"},
	{Name: "Antwerp", CountryCode: "be", StationQuery: "Antwerp Central railway station, Antwerp, Belgium", DecathlonQuery: "Decathlon Antwerpen, Belgium"},
	{Name: "Ghent", CountryCode: "be", StationQuery: "Ghent Sint-Pieters railway station, Ghent, Belgium", DecathlonQuery: "Decathlon Gent, Belgium"},
	{Name: "Liege", CountryCode: "be", StationQuery: "Liege Guillemins railway station, Liege, Belgium", DecathlonQuery: "Decathlon Liège, Belgium"},

	// --- Netherlands ---
	{Name: "Amsterdam", CountryCode: "nl", StationQuery: "Amsterdam Centraal railway station, Amsterdam, Netherlands", DecathlonQuery: "Decathlon Amsterdam, Netherlands"},
	{Name: "Rotterdam", CountryCode: "nl", StationQuery: "Rotterdam Centraal railway station, Rotterdam, Netherlands", DecathlonQuery: "Decathlon Rotterdam, Netherlands"},
	{Name: "The Hague", CountryCode: "nl", StationQuery: "Den Haag Centraal railway station, The Hague, Netherlands", DecathlonQuery: "Decathlon Den Haag, Netherlands"},
	{Name: "Utrecht", CountryCode: "nl", StationQuery: "Utrecht Centraal railway station, Utrecht, Netherlands", DecathlonQuery: "Decathlon Utrecht, Netherlands"},
	{Name: "Eindhoven", CountryCode: "nl", StationQuery: "Eindhoven railway station, Eindhoven, Netherlands", DecathlonQuery: "Decathlon Eindhoven, Netherlands"},

	// --- Switzerland ---
	{Name: "Zurich", CountryCode: "ch", StationQuery: "Zürich Hauptbahnhof railway station, Zurich, Switzerland", DecathlonQuery: "Decathlon Zürich, Switzerland"},
	{Name: "Geneva", CountryCode: "ch", StationQuery: "Geneva railway station, Geneva, Switzerland", DecathlonQuery: "Decathlon Genève, Switzerland"},
	{Name: "Bern", CountryCode: "ch", StationQuery: "Bern railway station, Bern, Switzerland", DecathlonQuery: "Decathlon Bern, Switzerland"},
	{Name: "Basel", CountryCode: "ch", StationQuery: "Basel SBB railway station, Basel, Switzerland", DecathlonQuery: "Decathlon Basel, Switzerland"},

	// --- Austria ---
	{Name: "Vienna", CountryCode: "at", StationQuery: "Wien Hauptbahnhof railway station, Vienna, Austria", DecathlonQuery: "Decathlon, 7-8, Columbusplatz, Neues Landgut, Katastralgemeinde Favoriten, Favoriten, Vienna, 1100, Austria"},
	{Name: "Graz", CountryCode: "at", StationQuery: "Graz Hauptbahnhof railway station, Graz, Austria", DecathlonQuery: "Decathlon Graz, Austria"},
	{Name: "Salzburg", CountryCode: "at", StationQuery: "Salzburg Hauptbahnhof railway station, Salzburg, Austria", DecathlonQuery: "Decathlon Salzburg, Austria"},
	{Name: "Innsbruck", CountryCode: "at", StationQuery: "Innsbruck Hauptbahnhof railway station, Innsbruck, Austria", DecathlonQuery: "Decathlon Innsbruck, Austria"},

	// --- Poland ---
	{Name: "Warsaw", CountryCode: "pl", StationQuery: "Warszawa Centralna railway station, Warsaw, Poland", DecathlonQuery: "Decathlon, 11, Polna, Koszyki, Śródmieście Południowe, Midtown, Warsaw, Masovian Voivodeship, 00-633, Poland"},
	{Name: "Krakow", CountryCode: "pl", StationQuery: "Kraków Główny railway station, Krakow, Poland", DecathlonQuery: "Decathlon Kraków, Poland"},
	{Name: "Wroclaw", CountryCode: "pl", StationQuery: "Wrocław Główny railway station, Wroclaw, Poland", DecathlonQuery: "Decathlon Wrocław, Poland"},
	{Name: "Poznan", CountryCode: "pl", StationQuery: "Poznań Główny railway station, Poznan, Poland", DecathlonQuery: "Decathlon Poznań, Poland"},
	{Name: "Gdansk", CountryCode: "pl", StationQuery: "Gdańsk Główny railway station, Gdansk, Poland", DecathlonQuery: "Decathlon Gdańsk, Poland"},
	{Name: "Lodz", CountryCode: "pl", StationQuery: "Łódź Fabryczna railway station, Lodz, Poland", DecathlonQuery: "Decathlon Łódź, Poland"},

	// --- Czech Republic ---
	{Name: "Prague", CountryCode: "cz", StationQuery: "Praha Hlavní Nádraží railway station, Prague, Czech Republic", DecathlonQuery: "Decathlon, 8, Plzeňská, Smíchov, Praha 5, obvod Praha 5, Prague, 150 00, Czechia"},
	{Name: "Brno", CountryCode: "cz", StationQuery: "Brno hlavní nádraží railway station, Brno, Czech Republic", DecathlonQuery: "Decathlon Brno, Czech Republic"},
	{Name: "Ostrava", CountryCode: "cz", StationQuery: "Ostrava hlavní nádraží railway station, Ostrava, Czech Republic", DecathlonQuery: "Decathlon, 3309/50, Varenská, Moravská Ostrava a Přívoz, Ostrava, okres Ostrava-město, Moravian-Silesian Region, 702 00, Czechia"},

	// --- Romania ---
	{Name: "Bucharest", CountryCode: "ro", StationQuery: "Gara de Nord, 1-3, Piața Gării de Nord, Matache, Gara de Nord, Sector 1, Bucharest, 010741, Romania", DecathlonQuery: "Decathlon, 23, Strada Ziduri Moși, Piața Obor, Obor, Sector 2, Bucharest, 023498, Romania"},
	{Name: "Cluj-Napoca", CountryCode: "ro", StationQuery: "Cluj-Napoca railway station, Cluj-Napoca, Romania", DecathlonQuery: "Decathlon Cluj-Napoca, Romania"},
	{Name: "Timisoara", CountryCode: "ro", StationQuery: "Timișoara Nord railway station, Timisoara, Romania", DecathlonQuery: "Decathlon Timișoara, Romania"},

	// --- Hungary ---
	{Name: "Budapest", CountryCode: "hu", StationQuery: "Budapest Keleti railway station, Budapest, Hungary", DecathlonQuery: "Decathlon Budapest, Hungary"},
	{Name: "Debrecen", CountryCode: "hu", StationQuery: "Debrecen railway station, Debrecen, Hungary", DecathlonQuery: "Decathlon Debrecen, Hungary"},

	// --- Sweden ---
	{Name: "Stockholm", CountryCode: "se", StationQuery: "Stockholm Central railway station, Stockholm, Sweden", DecathlonQuery: "Decathlon Stockholm, Sweden"},
	{Name: "Gothenburg", CountryCode: "se", StationQuery: "Gothenburg Central railway station, Gothenburg, Sweden", DecathlonQuery: "Decathlon Göteborg, Sweden"},
	{Name: "Malmo", CountryCode: "se", StationQuery: "Malmö Central railway station, Malmö, Sweden", DecathlonQuery: "Decathlon Malmö, Sweden"},

	// --- Greece ---
	{Name: "Athens", CountryCode: "gr", StationQuery: "Athens Larissa railway station, Athens, Greece", DecathlonQuery: "Decathlon Αθήνα, Greece"},
	{Name: "Thessaloniki", CountryCode: "gr", StationQuery: "Thessaloniki railway station, Thessaloniki, Greece", DecathlonQuery: "Decathlon Θεσσαλονίκη, Greece"},

	// --- Serbia ---
	{Name: "Belgrade", CountryCode: "rs", StationQuery: "Belgrade Centar railway station, Belgrade, Serbia", DecathlonQuery: "Decathlon Beograd, Serbia"},

	// --- Croatia ---
	{Name: "Zagreb", CountryCode: "hr", StationQuery: "Zagreb Glavni Kolodvor railway station, Zagreb, Croatia", DecathlonQuery: "Ulica Siniše Glavaševića, Mjesni odbor Trnava, Gradska četvrt Donja Dubrava, Zagreb, City of Zagreb, 10132, Croatia"},
	{Name: "Split", CountryCode: "hr", StationQuery: "Split railway station, Split, Croatia", DecathlonQuery: "Decathlon Split, Croatia"},

	// --- Slovenia ---
	{Name: "Ljubljana", CountryCode: "si", StationQuery: "Ljubljana railway station, Ljubljana, Slovenia", DecathlonQuery: "Decathlon Ljubljana, Slovenia"},

	// --- Slovakia ---
	{Name: "Bratislava", CountryCode: "sk", StationQuery: "Bratislava hlavná stanica railway station, Bratislava, Slovakia", DecathlonQuery: "Decathlon Bratislava, Slovakia"},
	{Name: "Kosice", CountryCode: "sk", StationQuery: "Košice railway station, Kosice, Slovakia", DecathlonQuery: "Decathlon Košice, Slovakia"},

	// --- Lithuania ---
	{Name: "Vilnius", CountryCode: "lt", StationQuery: "Vilnius railway station, Vilnius, Lithuania", DecathlonQuery: "Decathlon Vilnius, Lithuania"},

	// --- Latvia ---
	{Name: "Riga", CountryCode: "lv", StationQuery: "Riga railway station, Riga, Latvia", DecathlonQuery: "Decathlon Riga, Latvia"},

	// --- Estonia ---
	{Name: "Tallinn", CountryCode: "ee", StationQuery: "Tallinn Balti jaam railway station, Tallinn, Estonia", DecathlonQuery: "Decathlon Tallinn, Estonia"},

	// --- Malta ---
	{Name: "Valletta", CountryCode: "mt", StationQuery: "Valletta bus terminus, Valletta, Malta", DecathlonQuery: "Decathlon Malta"},
}

// --- Nominatim geocoding ---

func geocode(query string, countryCode string) (*Location, error) {
	base := "https://nominatim.openstreetmap.org/search"
	params := url.Values{}
	params.Set("q", query)
	params.Set("format", "json")
	params.Set("limit", "1")
	params.Set("countrycodes", countryCode)

	req, err := http.NewRequest("GET", base+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	// Nominatim requires a descriptive User-Agent
	req.Header.Set("User-Agent", "DecathlonDistanceScript/1.0 (educational use)")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var results []NominatimResult
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, nil // not found
	}

	var lat, lon float64
	fmt.Sscanf(results[0].Lat, "%f", &lat)
	fmt.Sscanf(results[0].Lon, "%f", &lon)

	return &Location{
		Name: results[0].DisplayName,
		Lat:  lat,
		Lon:  lon,
	}, nil
}

// geocodeCached checks the in-memory caches first:
//   - if key is in `found`, returns that location instantly
//   - if key is in `notFound`, returns (nil, true, nil) without calling Nominatim
//
// On a true cache miss it calls Nominatim, then persists the result
// (found or not-found) to the CSV cache file so it's terminal from then on.
// To force a retry for a specific city/type, delete its row from the CSV.
func geocodeCached(found map[cacheKey]Location, notFound map[cacheKey]bool, key cacheKey, query string, countryCode string) (loc *Location, fromCache bool, err error) {
	if cached, ok := found[key]; ok {
		return &cached, true, nil
	}
	if notFound[key] {
		return nil, true, nil // previously confirmed not found — skip silently
	}

	// Rate limit only applies to live Nominatim calls
	time.Sleep(1100 * time.Millisecond)

	result, err := geocode(query, countryCode)
	if err != nil {
		return nil, false, err
	}
	if result == nil {
		notFound[key] = true
		if err := appendNotFoundToCache(cacheFile, key); err != nil {
			fmt.Printf("warning: could not write not-found cache for %s/%s: %v\n", key.City, key.Type, err)
		}
		return nil, false, nil
	}

	found[key] = *result
	if err := appendToCache(cacheFile, key, *result); err != nil {
		// Non-fatal: warn but keep going with the in-memory value
		fmt.Printf("warning: could not write cache for %s/%s: %v\n", key.City, key.Type, err)
	}

	return result, false, nil
}

// --- OSRM road distance ---

// osrmServer returns the correct public OSRM demo server for a given
// travel profile. The main router.project-osrm.org server only serves
// the "driving" profile; walking needs the dedicated "routed-foot" server.
func osrmServer(profile string) string {
	switch profile {
	case "walking":
		return "https://routing.openstreetmap.de/routed-foot/route/v1/foot"
	default: // driving
		return "https://router.project-osrm.org/route/v1/driving"
	}
}

// routeDistance calls OSRM with the given profile ("driving" or "walking")
// and returns distance in km and duration in minutes.
func routeDistance(from, to Location, profile string) (distKm float64, durationMin float64, err error) {
	url := fmt.Sprintf(
		"%s/%f,%f;%f,%f?overview=false",
		osrmServer(profile), from.Lon, from.Lat, to.Lon, to.Lat,
	)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	var osrm OSRMResponse
	if err := json.Unmarshal(body, &osrm); err != nil {
		return 0, 0, err
	}
	if osrm.Code != "Ok" || len(osrm.Routes) == 0 {
		return 0, 0, fmt.Errorf("OSRM returned no route (code: %s)", osrm.Code)
	}

	distKm = osrm.Routes[0].Distance / 1000
	durationMin = osrm.Routes[0].Duration / 60
	return
}

// --- Haversine straight-line distance (for reference) ---

func haversine(from, to Location) float64 {
	const R = 6371.0
	lat1 := from.Lat * math.Pi / 180
	lat2 := to.Lat * math.Pi / 180
	dLat := (to.Lat - from.Lat) * math.Pi / 180
	dLon := (to.Lon - from.Lon) * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

// --- Results CSV ---

const resultsFile = "decathlon_distances.csv"
const resultsJSFile = "decathlon_data.js"

// Result holds the computed distances for one city.
type Result struct {
	City       string
	Country    string
	DriveKm    float64
	DriveMin   float64
	WalkKm     float64
	WalkMin    float64
	StraightKm float64
	Source     string
}

// sortResults returns a copy of results sorted by country name then walk time.
// This ordering is shared by both output functions so CSV and JS stay consistent.
func sortResults(results []Result) []Result {
	sorted := make([]Result, len(results))
	copy(sorted, results)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Country != sorted[j].Country {
			return sorted[i].Country < sorted[j].Country
		}
		return sorted[i].WalkMin < sorted[j].WalkMin
	})
	return sorted
}

// writeResultsCSV writes all computed results to a CSV file, overwriting
// any previous run's output. Skipped cities are not included.
func writeResultsCSV(path string, results []Result) error {
	// Sort by country name, then by drive time within country.
	sorted := make([]Result, len(results))
	copy(sorted, results)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Country != sorted[j].Country {
			return sorted[i].Country < sorted[j].Country
		}
		return sorted[i].WalkMin < sorted[j].WalkMin
	})

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	if err := w.Write([]string{
		"country", "rank_in_country", "city", "drive_km", "drive_min", "walk_km", "walk_min", "straight_km", "source",
	}); err != nil {
		return err
	}

	rank := 1
	prevCountry := ""

	for _, r := range results {
		if r.Country != prevCountry {
			rank = 1
			prevCountry = r.Country
		}

		if err := w.Write([]string{
			r.Country,
			fmt.Sprintf("%d", rank),
			r.City,
			fmt.Sprintf("%.1f", r.DriveKm),
			fmt.Sprintf("%.1f", r.DriveMin),
			fmt.Sprintf("%.1f", r.WalkKm),
			fmt.Sprintf("%.1f", r.WalkMin),
			fmt.Sprintf("%.1f", r.StraightKm),
			r.Source,
		}); err != nil {
			return err
		}
		rank++
	}

	return nil
}

// writeResultsJS writes a JS file containing the DATA array for direct use
// in decathlon.html. Results are sorted and ranked identically to the CSV.
func writeResultsJS(path string, results []Result) error {
	sorted := sortResults(results)

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, "// The Decathlon Index — generated %s\n", time.Now().Format("2006-01-02"))
	fmt.Fprintf(f, "// %d cities across %d countries\n", len(sorted), countCountries(sorted))
	fmt.Fprintf(f, "const DATA = [\n")

	rank := 1
	prevCountry := ""
	for i, r := range sorted {
		if r.Country != prevCountry {
			if prevCountry != "" {
				fmt.Fprintf(f, "\n")
			}
			fmt.Fprintf(f, "  // %s\n", r.Country)
			rank = 1
			prevCountry = r.Country
		}
		comma := ","
		if i == len(sorted)-1 {
			comma = ""
		}
		fmt.Fprintf(f,
			"  { country: %q, rank: %d, city: %q, drive_km: %.1f, drive_min: %.0f, walk_km: %.1f, walk_min: %.0f, straight_km: %.1f }%s\n",
			r.Country, rank, r.City, r.DriveKm, r.DriveMin, r.WalkKm, r.WalkMin, r.StraightKm, comma,
		)
		rank++
	}

	fmt.Fprintf(f, "];\n")
	return nil
}

// countCountries returns the number of distinct countries in a result set.
func countCountries(results []Result) int {
	seen := make(map[string]bool)
	for _, r := range results {
		seen[r.Country] = true
	}
	return len(seen)
}

// --- Main ---

func main() {
	found, notFound, err := loadCache(cacheFile)
	if err != nil {
		fmt.Printf("warning: could not load cache file, starting fresh: %v\n", err)
		found = make(map[cacheKey]Location)
		notFound = make(map[cacheKey]bool)
	}

	fmt.Printf("%-12s  %10s  %11s  %10s  %11s  %13s  %s\n",
		"City", "Drive(km)", "Drive(min)", "Walk(km)", "Walk(min)", "Straight(km)", "Source")
	fmt.Println("----------------------------------------------------------------------------------")

	var results []Result

	for _, city := range cities {
		stationKey := cacheKey{City: city.Name, Type: "station"}
		station, stationCached, err := geocodeCached(found, notFound, stationKey, city.StationQuery, city.CountryCode)
		if err != nil {
			fmt.Printf("%-12s  Could not geocode station: %v\n", city.Name, err)
			continue
		}
		if station == nil {
			fmt.Printf("%-12s  Station not found (cached) — skipping\n", city.Name)
			continue
		}

		decathlonKey := cacheKey{City: city.Name, Type: "decathlon"}
		decathlon, decathlonCached, err := geocodeCached(found, notFound, decathlonKey, city.DecathlonQuery, city.CountryCode)
		if err != nil {
			fmt.Printf("%-12s  Could not geocode Decathlon: %v\n", city.Name, err)
			continue
		}
		if decathlon == nil {
			source := "live"
			if decathlonCached {
				source = "cache"
			}
			fmt.Printf("%-12s  No Decathlon found — skipping (%s)\n", city.Name, source)
			continue
		}

		driveKm, driveMin, err := routeDistance(*station, *decathlon, "driving")
		if err != nil {
			fmt.Printf("%-12s  Driving route error: %v\n", city.Name, err)
			continue
		}

		walkKm, walkMin, err := routeDistance(*station, *decathlon, "walking")
		if err != nil {
			fmt.Printf("%-12s  Walking route error: %v\n", city.Name, err)
			continue
		}

		straightKm := haversine(*station, *decathlon)

		source := "live"
		if stationCached && decathlonCached {
			source = "cache"
		} else if stationCached || decathlonCached {
			source = "mixed"
		}

		fmt.Printf("%-12s  %10.1f  %11.1f  %10.1f  %11.1f  %13.1f  %s\n",
			city.Name, driveKm, driveMin, walkKm, walkMin, straightKm, source)

		results = append(results, Result{
			Country:    countryName[city.CountryCode],
			City:       city.Name,
			DriveKm:    driveKm,
			DriveMin:   driveMin,
			WalkKm:     walkKm,
			WalkMin:    walkMin,
			StraightKm: straightKm,
			Source:     source,
		})
	}

	if len(results) > 0 {
		if err := writeResultsCSV(resultsFile, results); err != nil {
			fmt.Printf("warning: could not write results CSV: %v\n", err)
		} else {
			fmt.Printf("\nResults written to %s (%d cities)\n", resultsFile, len(results))
		}

		if err := writeResultsJS(resultsJSFile, results); err != nil {
			fmt.Printf("warning: could not write results JS: %v\n", err)
		} else {
			fmt.Printf("Results written to %s\n", resultsJSFile)
		}

	}
}
