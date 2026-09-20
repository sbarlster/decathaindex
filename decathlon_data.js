// The Decathlon Index — generated 2026-09-20
// 100 cities across 21 countries
const DATA = [
  // Austria
  { country: "Austria", rank: 1, city: "Vienna", drive_km: 0.7, drive_min: 2, walk_km: 0.5, walk_min: 6, straight_km: 0.4 },
  { country: "Austria", rank: 2, city: "Graz", drive_km: 8.6, drive_min: 16, walk_km: 7.9, walk_min: 106, straight_km: 6.5 },

  // Belgium
  { country: "Belgium", rank: 1, city: "Brussels", drive_km: 4.7, drive_min: 13, walk_km: 2.3, walk_min: 30, straight_km: 2.0 },
  { country: "Belgium", rank: 2, city: "Liege", drive_km: 3.2, drive_min: 7, walk_km: 2.5, walk_min: 33, straight_km: 2.1 },

  // Croatia
  { country: "Croatia", rank: 1, city: "Zagreb", drive_km: 7.1, drive_min: 10, walk_km: 6.6, walk_min: 84, straight_km: 5.4 },
  { country: "Croatia", rank: 2, city: "Split", drive_km: 6.6, drive_min: 10, walk_km: 6.3, walk_min: 84, straight_km: 4.7 },

  // Czech Republic
  { country: "Czech Republic", rank: 1, city: "Ostrava", drive_km: 3.2, drive_min: 7, walk_km: 2.8, walk_min: 37, straight_km: 1.9 },
  { country: "Czech Republic", rank: 2, city: "Prague", drive_km: 3.6, drive_min: 6, walk_km: 3.4, walk_min: 45, straight_km: 2.6 },
  { country: "Czech Republic", rank: 3, city: "Brno", drive_km: 8.1, drive_min: 11, walk_km: 8.6, walk_min: 118, straight_km: 6.2 },

  // France
  { country: "France", rank: 1, city: "Nice", drive_km: 1.0, drive_min: 2, walk_km: 0.3, walk_min: 5, straight_km: 0.3 },
  { country: "France", rank: 2, city: "Lyon", drive_km: 2.2, drive_min: 6, walk_km: 0.6, walk_min: 7, straight_km: 0.4 },
  { country: "France", rank: 3, city: "Lille", drive_km: 1.1, drive_min: 3, walk_km: 0.6, walk_min: 8, straight_km: 0.5 },
  { country: "France", rank: 4, city: "Strasbourg", drive_km: 1.3, drive_min: 3, walk_km: 0.9, walk_min: 12, straight_km: 0.6 },
  { country: "France", rank: 5, city: "Montpellier", drive_km: 2.2, drive_min: 6, walk_km: 1.0, walk_min: 13, straight_km: 0.7 },
  { country: "France", rank: 6, city: "Toulouse", drive_km: 1.6, drive_min: 4, walk_km: 1.1, walk_min: 15, straight_km: 0.9 },
  { country: "France", rank: 7, city: "Rennes", drive_km: 2.2, drive_min: 7, walk_km: 1.5, walk_min: 20, straight_km: 1.2 },
  { country: "France", rank: 8, city: "Grenoble", drive_km: 2.3, drive_min: 7, walk_km: 1.5, walk_min: 20, straight_km: 1.1 },
  { country: "France", rank: 9, city: "Nantes", drive_km: 1.9, drive_min: 5, walk_km: 1.8, walk_min: 24, straight_km: 1.5 },
  { country: "France", rank: 10, city: "Bordeaux", drive_km: 2.9, drive_min: 6, walk_km: 2.2, walk_min: 30, straight_km: 1.9 },
  { country: "France", rank: 11, city: "Marseille", drive_km: 3.7, drive_min: 8, walk_km: 2.6, walk_min: 35, straight_km: 1.5 },
  { country: "France", rank: 12, city: "Paris", drive_km: 6.8, drive_min: 17, walk_km: 6.5, walk_min: 86, straight_km: 5.8 },

  // Germany
  { country: "Germany", rank: 1, city: "Dresden", drive_km: 1.2, drive_min: 4, walk_km: 0.2, walk_min: 3, straight_km: 0.1 },
  { country: "Germany", rank: 2, city: "Munich", drive_km: 1.3, drive_min: 2, walk_km: 0.9, walk_min: 12, straight_km: 0.6 },
  { country: "Germany", rank: 3, city: "Cologne", drive_km: 1.6, drive_min: 3, walk_km: 1.0, walk_min: 14, straight_km: 0.8 },
  { country: "Germany", rank: 4, city: "Dusseldorf", drive_km: 2.1, drive_min: 5, walk_km: 1.3, walk_min: 16, straight_km: 1.0 },
  { country: "Germany", rank: 5, city: "Nuremberg", drive_km: 2.1, drive_min: 4, walk_km: 1.4, walk_min: 19, straight_km: 1.0 },
  { country: "Germany", rank: 6, city: "Stuttgart", drive_km: 3.8, drive_min: 9, walk_km: 1.5, walk_min: 20, straight_km: 1.2 },
  { country: "Germany", rank: 7, city: "Leipzig", drive_km: 2.3, drive_min: 6, walk_km: 1.6, walk_min: 21, straight_km: 1.2 },
  { country: "Germany", rank: 8, city: "Frankfurt", drive_km: 3.0, drive_min: 6, walk_km: 2.0, walk_min: 27, straight_km: 1.9 },
  { country: "Germany", rank: 9, city: "Berlin", drive_km: 5.4, drive_min: 13, walk_km: 5.2, walk_min: 70, straight_km: 4.7 },
  { country: "Germany", rank: 10, city: "Bremen", drive_km: 8.3, drive_min: 15, walk_km: 8.2, walk_min: 93, straight_km: 6.5 },
  { country: "Germany", rank: 11, city: "Hanover", drive_km: 10.0, drive_min: 13, walk_km: 8.9, walk_min: 117, straight_km: 7.3 },
  { country: "Germany", rank: 12, city: "Dortmund", drive_km: 12.0, drive_min: 20, walk_km: 9.9, walk_min: 132, straight_km: 7.9 },
  { country: "Germany", rank: 13, city: "Hamburg", drive_km: 12.5, drive_min: 20, walk_km: 13.5, walk_min: 178, straight_km: 10.7 },

  // Hungary
  { country: "Hungary", rank: 1, city: "Budapest", drive_km: 3.7, drive_min: 8, walk_km: 2.3, walk_min: 31, straight_km: 1.7 },
  { country: "Hungary", rank: 2, city: "Debrecen", drive_km: 4.5, drive_min: 8, walk_km: 4.8, walk_min: 63, straight_km: 3.7 },

  // Ireland
  { country: "Ireland", rank: 1, city: "Dublin", drive_km: 4.0, drive_min: 11, walk_km: 2.6, walk_min: 34, straight_km: 2.2 },

  // Italy
  { country: "Italy", rank: 1, city: "Naples", drive_km: 0.8, drive_min: 1, walk_km: 0.3, walk_min: 4, straight_km: 0.3 },
  { country: "Italy", rank: 2, city: "Turin", drive_km: 0.2, drive_min: 1, walk_km: 0.4, walk_min: 5, straight_km: 0.3 },
  { country: "Italy", rank: 3, city: "Florence", drive_km: 3.7, drive_min: 7, walk_km: 2.7, walk_min: 36, straight_km: 2.1 },
  { country: "Italy", rank: 4, city: "Milan", drive_km: 3.2, drive_min: 7, walk_km: 3.0, walk_min: 40, straight_km: 2.6 },
  { country: "Italy", rank: 5, city: "Genoa", drive_km: 5.8, drive_min: 9, walk_km: 3.4, walk_min: 45, straight_km: 2.2 },
  { country: "Italy", rank: 6, city: "Rome", drive_km: 3.8, drive_min: 6, walk_km: 4.0, walk_min: 54, straight_km: 3.4 },
  { country: "Italy", rank: 7, city: "Verona", drive_km: 6.1, drive_min: 9, walk_km: 5.5, walk_min: 71, straight_km: 4.6 },
  { country: "Italy", rank: 8, city: "Bologna", drive_km: 8.1, drive_min: 16, walk_km: 5.4, walk_min: 73, straight_km: 4.4 },
  { country: "Italy", rank: 9, city: "Bari", drive_km: 11.0, drive_min: 14, walk_km: 7.6, walk_min: 102, straight_km: 4.1 },
  { country: "Italy", rank: 10, city: "Venice", drive_km: 11.2, drive_min: 15, walk_km: 16.3, walk_min: 217, straight_km: 8.9 },

  // Latvia
  { country: "Latvia", rank: 1, city: "Riga", drive_km: 8.4, drive_min: 12, walk_km: 6.6, walk_min: 88, straight_km: 5.3 },

  // Lithuania
  { country: "Lithuania", rank: 1, city: "Vilnius", drive_km: 4.6, drive_min: 7, walk_km: 3.2, walk_min: 43, straight_km: 2.3 },

  // Netherlands
  { country: "Netherlands", rank: 1, city: "Eindhoven", drive_km: 1.2, drive_min: 3, walk_km: 0.4, walk_min: 5, straight_km: 0.2 },
  { country: "Netherlands", rank: 2, city: "Rotterdam", drive_km: 1.2, drive_min: 3, walk_km: 0.7, walk_min: 10, straight_km: 0.5 },
  { country: "Netherlands", rank: 3, city: "The Hague", drive_km: 1.8, drive_min: 4, walk_km: 1.3, walk_min: 18, straight_km: 0.8 },
  { country: "Netherlands", rank: 4, city: "Amsterdam", drive_km: 5.3, drive_min: 11, walk_km: 2.9, walk_min: 38, straight_km: 2.4 },
  { country: "Netherlands", rank: 5, city: "Utrecht", drive_km: 6.0, drive_min: 13, walk_km: 5.0, walk_min: 67, straight_km: 4.2 },

  // Poland
  { country: "Poland", rank: 1, city: "Warsaw", drive_km: 1.7, drive_min: 4, walk_km: 1.6, walk_min: 21, straight_km: 1.4 },
  { country: "Poland", rank: 2, city: "Lodz", drive_km: 3.7, drive_min: 8, walk_km: 2.5, walk_min: 33, straight_km: 1.8 },
  { country: "Poland", rank: 3, city: "Krakow", drive_km: 4.7, drive_min: 9, walk_km: 3.8, walk_min: 50, straight_km: 2.8 },
  { country: "Poland", rank: 4, city: "Wroclaw", drive_km: 5.7, drive_min: 12, walk_km: 5.3, walk_min: 70, straight_km: 4.1 },
  { country: "Poland", rank: 5, city: "Poznan", drive_km: 7.1, drive_min: 13, walk_km: 6.7, walk_min: 88, straight_km: 5.4 },
  { country: "Poland", rank: 6, city: "Gdansk", drive_km: 7.8, drive_min: 13, walk_km: 8.0, walk_min: 105, straight_km: 6.8 },

  // Portugal
  { country: "Portugal", rank: 1, city: "Lisbon", drive_km: 2.9, drive_min: 5, walk_km: 2.2, walk_min: 29, straight_km: 1.4 },
  { country: "Portugal", rank: 2, city: "Porto", drive_km: 3.4, drive_min: 5, walk_km: 2.3, walk_min: 31, straight_km: 2.0 },
  { country: "Portugal", rank: 3, city: "Coimbra", drive_km: 2.8, drive_min: 5, walk_km: 2.5, walk_min: 34, straight_km: 1.7 },
  { country: "Portugal", rank: 4, city: "Braga", drive_km: 4.4, drive_min: 7, walk_km: 3.1, walk_min: 41, straight_km: 2.5 },

  // Romania
  { country: "Romania", rank: 1, city: "Timisoara", drive_km: 5.3, drive_min: 9, walk_km: 5.3, walk_min: 71, straight_km: 4.3 },
  { country: "Romania", rank: 2, city: "Cluj-Napoca", drive_km: 6.8, drive_min: 11, walk_km: 6.5, walk_min: 87, straight_km: 5.3 },

  // Serbia
  { country: "Serbia", rank: 1, city: "Belgrade", drive_km: 6.8, drive_min: 12, walk_km: 7.1, walk_min: 95, straight_km: 5.4 },

  // Slovakia
  { country: "Slovakia", rank: 1, city: "Kosice", drive_km: 5.1, drive_min: 11, walk_km: 4.0, walk_min: 54, straight_km: 2.9 },
  { country: "Slovakia", rank: 2, city: "Bratislava", drive_km: 7.8, drive_min: 13, walk_km: 7.9, walk_min: 105, straight_km: 6.3 },

  // Slovenia
  { country: "Slovenia", rank: 1, city: "Ljubljana", drive_km: 5.7, drive_min: 9, walk_km: 6.9, walk_min: 91, straight_km: 4.7 },

  // Spain
  { country: "Spain", rank: 1, city: "Bilbao", drive_km: 0.3, drive_min: 1, walk_km: 0.3, walk_min: 4, straight_km: 0.2 },
  { country: "Spain", rank: 2, city: "Madrid", drive_km: 1.6, drive_min: 3, walk_km: 0.4, walk_min: 5, straight_km: 0.3 },
  { country: "Spain", rank: 3, city: "San Sebastian", drive_km: 0.6, drive_min: 2, walk_km: 0.5, walk_min: 7, straight_km: 0.4 },
  { country: "Spain", rank: 4, city: "Cordoba", drive_km: 1.8, drive_min: 3, walk_km: 1.0, walk_min: 13, straight_km: 0.7 },
  { country: "Spain", rank: 5, city: "Valladolid", drive_km: 2.4, drive_min: 6, walk_km: 1.2, walk_min: 16, straight_km: 1.0 },
  { country: "Spain", rank: 6, city: "Barcelona", drive_km: 2.0, drive_min: 5, walk_km: 1.5, walk_min: 20, straight_km: 1.3 },
  { country: "Spain", rank: 7, city: "Malaga", drive_km: 3.5, drive_min: 7, walk_km: 1.7, walk_min: 22, straight_km: 1.4 },
  { country: "Spain", rank: 8, city: "Valencia", drive_km: 2.4, drive_min: 4, walk_km: 1.7, walk_min: 23, straight_km: 1.5 },
  { country: "Spain", rank: 9, city: "Seville", drive_km: 5.8, drive_min: 12, walk_km: 1.9, walk_min: 25, straight_km: 1.7 },
  { country: "Spain", rank: 10, city: "A Coruna", drive_km: 3.8, drive_min: 10, walk_km: 2.7, walk_min: 36, straight_km: 2.3 },
  { country: "Spain", rank: 11, city: "Zaragoza", drive_km: 2.9, drive_min: 6, walk_km: 2.9, walk_min: 39, straight_km: 2.5 },
  { country: "Spain", rank: 12, city: "Alicante", drive_km: 5.7, drive_min: 11, walk_km: 5.4, walk_min: 72, straight_km: 3.6 },

  // Switzerland
  { country: "Switzerland", rank: 1, city: "Bern", drive_km: 1.0, drive_min: 2, walk_km: 0.3, walk_min: 4, straight_km: 0.2 },
  { country: "Switzerland", rank: 2, city: "Zurich", drive_km: 1.2, drive_min: 2, walk_km: 0.8, walk_min: 11, straight_km: 0.7 },
  { country: "Switzerland", rank: 3, city: "Geneva", drive_km: 2.4, drive_min: 5, walk_km: 1.9, walk_min: 26, straight_km: 1.7 },
  { country: "Switzerland", rank: 4, city: "Basel", drive_km: 2.5, drive_min: 5, walk_km: 2.2, walk_min: 30, straight_km: 1.8 },

  // United Kingdom
  { country: "United Kingdom", rank: 1, city: "Leeds", drive_km: 2.9, drive_min: 7, walk_km: 0.3, walk_min: 4, straight_km: 0.2 },
  { country: "United Kingdom", rank: 2, city: "Liverpool", drive_km: 2.7, drive_min: 6, walk_km: 0.7, walk_min: 9, straight_km: 0.6 },
  { country: "United Kingdom", rank: 3, city: "Southampton", drive_km: 1.5, drive_min: 4, walk_km: 0.9, walk_min: 12, straight_km: 0.6 },
  { country: "United Kingdom", rank: 4, city: "Sheffield", drive_km: 1.3, drive_min: 3, walk_km: 1.0, walk_min: 14, straight_km: 0.9 },
  { country: "United Kingdom", rank: 5, city: "Cambridge", drive_km: 2.9, drive_min: 6, walk_km: 1.9, walk_min: 26, straight_km: 1.4 },
  { country: "United Kingdom", rank: 6, city: "Manchester", drive_km: 3.8, drive_min: 7, walk_km: 3.5, walk_min: 47, straight_km: 2.6 },
  { country: "United Kingdom", rank: 7, city: "Belfast", drive_km: 6.9, drive_min: 9, walk_km: 6.6, walk_min: 87, straight_km: 5.0 },
  { country: "United Kingdom", rank: 8, city: "Cardiff", drive_km: 8.6, drive_min: 11, walk_km: 8.3, walk_min: 111, straight_km: 6.7 },
  { country: "United Kingdom", rank: 9, city: "Portsmouth", drive_km: 7.9, drive_min: 11, walk_km: 8.8, walk_min: 117, straight_km: 5.6 },
  { country: "United Kingdom", rank: 10, city: "Edinburgh", drive_km: 10.2, drive_min: 22, walk_km: 8.9, walk_min: 119, straight_km: 7.8 },
  { country: "United Kingdom", rank: 11, city: "Glasgow", drive_km: 8.0, drive_min: 10, walk_km: 9.1, walk_min: 122, straight_km: 6.5 },
  { country: "United Kingdom", rank: 12, city: "London", drive_km: 13.6, drive_min: 33, walk_km: 11.9, walk_min: 158, straight_km: 9.9 },
  { country: "United Kingdom", rank: 13, city: "Nottingham", drive_km: 14.1, drive_min: 21, walk_km: 14.1, walk_min: 188, straight_km: 10.9 },
  { country: "United Kingdom", rank: 14, city: "Birmingham", drive_km: 20.0, drive_min: 21, walk_km: 15.8, walk_min: 212, straight_km: 12.5 }
];
