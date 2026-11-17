// The Decathlon Index — generated 2026-11-17
// 100 cities across 21 countries
const DATA = [
  // Austria
  { country: "Austria", rank: 1, city: "Vienna", drive_km: 0.7, drive_min: 2, walk_km: 0.5, walk_min: 6, straight_km: 0.4, lat: 48.184988, lon: 16.377939 },
  { country: "Austria", rank: 2, city: "Graz", drive_km: 8.6, drive_min: 16, walk_km: 7.9, walk_min: 106, straight_km: 6.5, lat: 47.072398, lon: 15.416517 },

  // Belgium
  { country: "Belgium", rank: 1, city: "Brussels", drive_km: 4.7, drive_min: 13, walk_km: 2.3, walk_min: 30, straight_km: 2.0, lat: 50.835775, lon: 4.336054 },
  { country: "Belgium", rank: 2, city: "Liege", drive_km: 3.2, drive_min: 7, walk_km: 2.5, walk_min: 34, straight_km: 2.1, lat: 50.624728, lon: 5.567065 },

  // Croatia
  { country: "Croatia", rank: 1, city: "Zagreb", drive_km: 7.1, drive_min: 10, walk_km: 6.6, walk_min: 84, straight_km: 5.4, lat: 45.804446, lon: 15.978833 },
  { country: "Croatia", rank: 2, city: "Split", drive_km: 6.6, drive_min: 10, walk_km: 6.3, walk_min: 84, straight_km: 4.7, lat: 43.504815, lon: 16.443044 },

  // Czech Republic
  { country: "Czech Republic", rank: 1, city: "Ostrava", drive_km: 3.2, drive_min: 7, walk_km: 2.8, walk_min: 37, straight_km: 1.9, lat: 49.850670, lon: 18.267072 },
  { country: "Czech Republic", rank: 2, city: "Prague", drive_km: 3.6, drive_min: 6, walk_km: 3.4, walk_min: 45, straight_km: 2.6, lat: 50.082963, lon: 14.436094 },
  { country: "Czech Republic", rank: 3, city: "Brno", drive_km: 12.3, drive_min: 15, walk_km: 8.3, walk_min: 113, straight_km: 6.2, lat: 49.191122, lon: 16.613112 },

  // France
  { country: "France", rank: 1, city: "Nice", drive_km: 1.5, drive_min: 4, walk_km: 0.3, walk_min: 4, straight_km: 0.3, lat: 43.704745, lon: 7.261598 },
  { country: "France", rank: 2, city: "Lyon", drive_km: 2.2, drive_min: 6, walk_km: 0.6, walk_min: 7, straight_km: 0.4, lat: 45.760573, lon: 4.860163 },
  { country: "France", rank: 3, city: "Lille", drive_km: 1.1, drive_min: 3, walk_km: 0.6, walk_min: 8, straight_km: 0.5, lat: 50.636380, lon: 3.070530 },
  { country: "France", rank: 4, city: "Strasbourg", drive_km: 1.3, drive_min: 3, walk_km: 0.9, walk_min: 12, straight_km: 0.6, lat: 48.585293, lon: 7.733925 },
  { country: "France", rank: 5, city: "Montpellier", drive_km: 2.2, drive_min: 6, walk_km: 1.0, walk_min: 13, straight_km: 0.7, lat: 43.604246, lon: 3.880322 },
  { country: "France", rank: 6, city: "Toulouse", drive_km: 1.6, drive_min: 4, walk_km: 1.1, walk_min: 15, straight_km: 0.9, lat: 43.611431, lon: 1.454210 },
  { country: "France", rank: 7, city: "Rennes", drive_km: 2.2, drive_min: 7, walk_km: 1.5, walk_min: 20, straight_km: 1.2, lat: 48.103193, lon: -1.672381 },
  { country: "France", rank: 8, city: "Grenoble", drive_km: 2.3, drive_min: 7, walk_km: 1.5, walk_min: 20, straight_km: 1.1, lat: 45.191307, lon: 5.714102 },
  { country: "France", rank: 9, city: "Nantes", drive_km: 1.9, drive_min: 5, walk_km: 1.8, walk_min: 24, straight_km: 1.5, lat: 47.217095, lon: -1.541423 },
  { country: "France", rank: 10, city: "Bordeaux", drive_km: 2.9, drive_min: 6, walk_km: 2.2, walk_min: 29, straight_km: 1.9, lat: 44.825523, lon: -0.555650 },
  { country: "France", rank: 11, city: "Marseille", drive_km: 3.7, drive_min: 8, walk_km: 2.6, walk_min: 35, straight_km: 1.5, lat: 43.303230, lon: 5.381642 },
  { country: "France", rank: 12, city: "Paris", drive_km: 6.8, drive_min: 17, walk_km: 6.4, walk_min: 86, straight_km: 5.8, lat: 48.880400, lon: 2.354973 },

  // Germany
  { country: "Germany", rank: 1, city: "Dresden", drive_km: 1.2, drive_min: 4, walk_km: 0.2, walk_min: 3, straight_km: 0.1, lat: 51.040245, lon: 13.732077 },
  { country: "Germany", rank: 2, city: "Munich", drive_km: 1.3, drive_min: 2, walk_km: 0.9, walk_min: 12, straight_km: 0.6, lat: 48.140725, lon: 11.556943 },
  { country: "Germany", rank: 3, city: "Cologne", drive_km: 1.6, drive_min: 3, walk_km: 1.0, walk_min: 14, straight_km: 0.8, lat: 50.943022, lon: 6.959091 },
  { country: "Germany", rank: 4, city: "Dusseldorf", drive_km: 2.1, drive_min: 5, walk_km: 1.3, walk_min: 16, straight_km: 1.0, lat: 51.219734, lon: 6.794302 },
  { country: "Germany", rank: 5, city: "Nuremberg", drive_km: 2.1, drive_min: 4, walk_km: 1.4, walk_min: 19, straight_km: 1.0, lat: 49.446203, lon: 11.081865 },
  { country: "Germany", rank: 6, city: "Stuttgart", drive_km: 2.7, drive_min: 7, walk_km: 1.5, walk_min: 20, straight_km: 1.2, lat: 48.785610, lon: 9.183396 },
  { country: "Germany", rank: 7, city: "Leipzig", drive_km: 2.3, drive_min: 6, walk_km: 1.6, walk_min: 21, straight_km: 1.2, lat: 51.346552, lon: 12.383086 },
  { country: "Germany", rank: 8, city: "Frankfurt", drive_km: 2.8, drive_min: 6, walk_km: 2.0, walk_min: 27, straight_km: 1.9, lat: 50.106654, lon: 8.662581 },
  { country: "Germany", rank: 9, city: "Berlin", drive_km: 5.4, drive_min: 12, walk_km: 5.2, walk_min: 70, straight_km: 4.7, lat: 52.524945, lon: 13.369661 },
  { country: "Germany", rank: 10, city: "Bremen", drive_km: 8.3, drive_min: 15, walk_km: 8.1, walk_min: 92, straight_km: 6.5, lat: 53.083297, lon: 8.813758 },
  { country: "Germany", rank: 11, city: "Hanover", drive_km: 10.0, drive_min: 13, walk_km: 8.9, walk_min: 117, straight_km: 7.3, lat: 52.376950, lon: 9.741653 },
  { country: "Germany", rank: 12, city: "Dortmund", drive_km: 12.0, drive_min: 20, walk_km: 9.9, walk_min: 132, straight_km: 7.9, lat: 51.517893, lon: 7.459283 },
  { country: "Germany", rank: 13, city: "Hamburg", drive_km: 12.5, drive_min: 20, walk_km: 13.3, walk_min: 176, straight_km: 10.7, lat: 53.553199, lon: 10.006436 },

  // Hungary
  { country: "Hungary", rank: 1, city: "Budapest", drive_km: 3.7, drive_min: 8, walk_km: 2.3, walk_min: 31, straight_km: 1.7, lat: 47.500412, lon: 19.083980 },
  { country: "Hungary", rank: 2, city: "Debrecen", drive_km: 4.5, drive_min: 8, walk_km: 4.8, walk_min: 63, straight_km: 3.7, lat: 47.519886, lon: 21.628763 },

  // Ireland
  { country: "Ireland", rank: 1, city: "Dublin", drive_km: 4.0, drive_min: 11, walk_km: 2.6, walk_min: 34, straight_km: 2.2, lat: 53.346502, lon: -6.292968 },

  // Italy
  { country: "Italy", rank: 1, city: "Naples", drive_km: 0.9, drive_min: 1, walk_km: 0.3, walk_min: 4, straight_km: 0.3, lat: 40.853024, lon: 14.273140 },
  { country: "Italy", rank: 2, city: "Turin", drive_km: 0.2, drive_min: 1, walk_km: 0.4, walk_min: 5, straight_km: 0.3, lat: 45.061203, lon: 7.677800 },
  { country: "Italy", rank: 3, city: "Florence", drive_km: 3.7, drive_min: 7, walk_km: 2.8, walk_min: 38, straight_km: 2.1, lat: 43.777572, lon: 11.247418 },
  { country: "Italy", rank: 4, city: "Milan", drive_km: 3.2, drive_min: 7, walk_km: 3.0, walk_min: 40, straight_km: 2.6, lat: 45.485879, lon: 9.204262 },
  { country: "Italy", rank: 5, city: "Genoa", drive_km: 5.8, drive_min: 9, walk_km: 3.4, walk_min: 45, straight_km: 2.2, lat: 44.417020, lon: 8.921130 },
  { country: "Italy", rank: 6, city: "Rome", drive_km: 3.8, drive_min: 6, walk_km: 4.0, walk_min: 54, straight_km: 3.4, lat: 41.900582, lon: 12.502527 },
  { country: "Italy", rank: 7, city: "Verona", drive_km: 6.1, drive_min: 10, walk_km: 5.5, walk_min: 71, straight_km: 4.6, lat: 45.429182, lon: 10.982371 },
  { country: "Italy", rank: 8, city: "Bologna", drive_km: 8.1, drive_min: 16, walk_km: 5.4, walk_min: 73, straight_km: 4.4, lat: 44.505878, lon: 11.343343 },
  { country: "Italy", rank: 9, city: "Bari", drive_km: 9.7, drive_min: 12, walk_km: 7.6, walk_min: 102, straight_km: 4.1, lat: 41.117672, lon: 16.870066 },
  { country: "Italy", rank: 10, city: "Venice", drive_km: 11.2, drive_min: 15, walk_km: 11.7, walk_min: 155, straight_km: 8.9, lat: 45.441075, lon: 12.321032 },

  // Latvia
  { country: "Latvia", rank: 1, city: "Riga", drive_km: 8.7, drive_min: 12, walk_km: 6.3, walk_min: 83, straight_km: 5.3, lat: 56.946400, lon: 24.122911 },

  // Lithuania
  { country: "Lithuania", rank: 1, city: "Vilnius", drive_km: 4.6, drive_min: 7, walk_km: 3.1, walk_min: 41, straight_km: 2.3, lat: 54.669877, lon: 25.284758 },

  // Netherlands
  { country: "Netherlands", rank: 1, city: "Eindhoven", drive_km: 1.2, drive_min: 3, walk_km: 0.4, walk_min: 5, straight_km: 0.2, lat: 51.442962, lon: 5.479527 },
  { country: "Netherlands", rank: 2, city: "Rotterdam", drive_km: 1.2, drive_min: 3, walk_km: 0.7, walk_min: 10, straight_km: 0.5, lat: 51.923968, lon: 4.469799 },
  { country: "Netherlands", rank: 3, city: "The Hague", drive_km: 1.8, drive_min: 4, walk_km: 1.3, walk_min: 17, straight_km: 0.8, lat: 52.080844, lon: 4.324549 },
  { country: "Netherlands", rank: 4, city: "Amsterdam", drive_km: 5.3, drive_min: 11, walk_km: 2.8, walk_min: 38, straight_km: 2.4, lat: 52.378901, lon: 4.900581 },
  { country: "Netherlands", rank: 5, city: "Utrecht", drive_km: 6.0, drive_min: 13, walk_km: 5.2, walk_min: 70, straight_km: 4.2, lat: 52.089393, lon: 5.109821 },

  // Poland
  { country: "Poland", rank: 1, city: "Warsaw", drive_km: 1.7, drive_min: 4, walk_km: 1.6, walk_min: 21, straight_km: 1.4, lat: 52.228821, lon: 21.003160 },
  { country: "Poland", rank: 2, city: "Lodz", drive_km: 3.7, drive_min: 8, walk_km: 2.5, walk_min: 33, straight_km: 1.8, lat: 51.769554, lon: 19.469611 },
  { country: "Poland", rank: 3, city: "Krakow", drive_km: 4.7, drive_min: 9, walk_km: 3.8, walk_min: 50, straight_km: 2.8, lat: 50.068424, lon: 19.947886 },
  { country: "Poland", rank: 4, city: "Wroclaw", drive_km: 5.2, drive_min: 11, walk_km: 5.3, walk_min: 70, straight_km: 4.1, lat: 51.097916, lon: 17.037951 },
  { country: "Poland", rank: 5, city: "Poznan", drive_km: 7.1, drive_min: 13, walk_km: 6.8, walk_min: 88, straight_km: 5.4, lat: 52.401627, lon: 16.911579 },
  { country: "Poland", rank: 6, city: "Gdansk", drive_km: 7.8, drive_min: 13, walk_km: 7.9, walk_min: 105, straight_km: 6.8, lat: 54.356434, lon: 18.644195 },

  // Portugal
  { country: "Portugal", rank: 1, city: "Lisbon", drive_km: 2.9, drive_min: 5, walk_km: 2.2, walk_min: 29, straight_km: 1.4, lat: 38.767821, lon: -9.099073 },
  { country: "Portugal", rank: 2, city: "Porto", drive_km: 3.4, drive_min: 5, walk_km: 2.3, walk_min: 31, straight_km: 2.0, lat: 41.145333, lon: -8.609703 },
  { country: "Portugal", rank: 3, city: "Coimbra", drive_km: 2.8, drive_min: 5, walk_km: 2.4, walk_min: 32, straight_km: 1.7, lat: 40.224728, lon: -8.440477 },
  { country: "Portugal", rank: 4, city: "Braga", drive_km: 4.4, drive_min: 7, walk_km: 3.1, walk_min: 41, straight_km: 2.5, lat: 41.548046, lon: -8.434464 },

  // Romania
  { country: "Romania", rank: 1, city: "Timisoara", drive_km: 5.3, drive_min: 9, walk_km: 5.3, walk_min: 71, straight_km: 4.3, lat: 45.751077, lon: 21.207174 },
  { country: "Romania", rank: 2, city: "Cluj-Napoca", drive_km: 6.8, drive_min: 11, walk_km: 6.5, walk_min: 87, straight_km: 5.3, lat: 46.784625, lon: 23.586169 },

  // Serbia
  { country: "Serbia", rank: 1, city: "Belgrade", drive_km: 6.8, drive_min: 12, walk_km: 7.1, walk_min: 95, straight_km: 5.4, lat: 44.793630, lon: 20.453128 },

  // Slovakia
  { country: "Slovakia", rank: 1, city: "Kosice", drive_km: 5.1, drive_min: 11, walk_km: 4.0, walk_min: 54, straight_km: 2.9, lat: 48.722020, lon: 21.269212 },
  { country: "Slovakia", rank: 2, city: "Bratislava", drive_km: 7.8, drive_min: 13, walk_km: 7.9, walk_min: 106, straight_km: 6.3, lat: 48.159030, lon: 17.106468 },

  // Slovenia
  { country: "Slovenia", rank: 1, city: "Ljubljana", drive_km: 5.7, drive_min: 9, walk_km: 6.4, walk_min: 84, straight_km: 4.7, lat: 46.058728, lon: 14.512197 },

  // Spain
  { country: "Spain", rank: 1, city: "Bilbao", drive_km: 0.3, drive_min: 1, walk_km: 0.3, walk_min: 4, straight_km: 0.2, lat: 43.260527, lon: -2.927645 },
  { country: "Spain", rank: 2, city: "Madrid", drive_km: 1.6, drive_min: 3, walk_km: 0.4, walk_min: 5, straight_km: 0.3, lat: 40.405684, lon: -3.690019 },
  { country: "Spain", rank: 3, city: "San Sebastian", drive_km: 0.6, drive_min: 2, walk_km: 0.5, walk_min: 7, straight_km: 0.4, lat: 43.317737, lon: -1.976930 },
  { country: "Spain", rank: 4, city: "Cordoba", drive_km: 1.8, drive_min: 3, walk_km: 1.0, walk_min: 13, straight_km: 0.7, lat: 37.888510, lon: -4.789516 },
  { country: "Spain", rank: 5, city: "Valladolid", drive_km: 2.3, drive_min: 5, walk_km: 1.2, walk_min: 16, straight_km: 1.0, lat: 41.642127, lon: -4.726951 },
  { country: "Spain", rank: 6, city: "Barcelona", drive_km: 2.1, drive_min: 6, walk_km: 1.5, walk_min: 20, straight_km: 1.3, lat: 41.379012, lon: 2.139960 },
  { country: "Spain", rank: 7, city: "Malaga", drive_km: 3.5, drive_min: 7, walk_km: 1.7, walk_min: 22, straight_km: 1.4, lat: 36.711620, lon: -4.432360 },
  { country: "Spain", rank: 8, city: "Valencia", drive_km: 2.4, drive_min: 4, walk_km: 1.7, walk_min: 23, straight_km: 1.5, lat: 39.460047, lon: -0.381208 },
  { country: "Spain", rank: 9, city: "Seville", drive_km: 5.8, drive_min: 12, walk_km: 1.9, walk_min: 25, straight_km: 1.7, lat: 37.392092, lon: -5.975206 },
  { country: "Spain", rank: 10, city: "A Coruna", drive_km: 3.8, drive_min: 10, walk_km: 2.7, walk_min: 36, straight_km: 2.3, lat: 43.350424, lon: -8.412083 },
  { country: "Spain", rank: 11, city: "Zaragoza", drive_km: 3.0, drive_min: 7, walk_km: 3.0, walk_min: 39, straight_km: 2.5, lat: 41.658573, lon: -0.911388 },
  { country: "Spain", rank: 12, city: "Alicante", drive_km: 5.7, drive_min: 11, walk_km: 5.4, walk_min: 72, straight_km: 3.6, lat: 38.344980, lon: -0.496061 },

  // Switzerland
  { country: "Switzerland", rank: 1, city: "Bern", drive_km: 1.0, drive_min: 2, walk_km: 0.3, walk_min: 4, straight_km: 0.2, lat: 46.948909, lon: 7.439955 },
  { country: "Switzerland", rank: 2, city: "Zurich", drive_km: 1.2, drive_min: 2, walk_km: 0.8, walk_min: 11, straight_km: 0.7, lat: 47.378101, lon: 8.539364 },
  { country: "Switzerland", rank: 3, city: "Geneva", drive_km: 2.4, drive_min: 5, walk_km: 1.9, walk_min: 26, straight_km: 1.7, lat: 46.210229, lon: 6.142622 },
  { country: "Switzerland", rank: 4, city: "Basel", drive_km: 2.5, drive_min: 5, walk_km: 2.2, walk_min: 30, straight_km: 1.8, lat: 47.547623, lon: 7.589642 },

  // United Kingdom
  { country: "United Kingdom", rank: 1, city: "Leeds", drive_km: 2.6, drive_min: 6, walk_km: 0.3, walk_min: 4, straight_km: 0.2, lat: 53.795236, lon: -1.547738 },
  { country: "United Kingdom", rank: 2, city: "Liverpool", drive_km: 2.7, drive_min: 6, walk_km: 0.7, walk_min: 9, straight_km: 0.6, lat: 53.407610, lon: -2.977432 },
  { country: "United Kingdom", rank: 3, city: "Southampton", drive_km: 1.5, drive_min: 4, walk_km: 0.9, walk_min: 12, straight_km: 0.6, lat: 50.907498, lon: -1.414229 },
  { country: "United Kingdom", rank: 4, city: "Sheffield", drive_km: 1.3, drive_min: 3, walk_km: 1.0, walk_min: 14, straight_km: 0.9, lat: 53.378371, lon: -1.462138 },
  { country: "United Kingdom", rank: 5, city: "Cambridge", drive_km: 2.9, drive_min: 6, walk_km: 1.9, walk_min: 26, straight_km: 1.4, lat: 52.194642, lon: 0.137527 },
  { country: "United Kingdom", rank: 6, city: "Manchester", drive_km: 3.8, drive_min: 7, walk_km: 3.5, walk_min: 47, straight_km: 2.6, lat: 53.477220, lon: -2.230140 },
  { country: "United Kingdom", rank: 7, city: "Belfast", drive_km: 6.9, drive_min: 9, walk_km: 6.6, walk_min: 87, straight_km: 5.0, lat: 54.595151, lon: -5.917290 },
  { country: "United Kingdom", rank: 8, city: "Cardiff", drive_km: 8.6, drive_min: 11, walk_km: 8.2, walk_min: 110, straight_km: 6.7, lat: 51.475549, lon: -3.179706 },
  { country: "United Kingdom", rank: 9, city: "Portsmouth", drive_km: 7.9, drive_min: 11, walk_km: 8.8, walk_min: 117, straight_km: 5.6, lat: 50.796703, lon: -1.108781 },
  { country: "United Kingdom", rank: 10, city: "Edinburgh", drive_km: 10.2, drive_min: 22, walk_km: 8.9, walk_min: 119, straight_km: 7.8, lat: 55.951902, lon: -3.190420 },
  { country: "United Kingdom", rank: 11, city: "Glasgow", drive_km: 8.0, drive_min: 10, walk_km: 9.2, walk_min: 123, straight_km: 6.5, lat: 55.858313, lon: -4.258436 },
  { country: "United Kingdom", rank: 12, city: "London", drive_km: 13.6, drive_min: 33, walk_km: 11.9, walk_min: 158, straight_km: 9.9, lat: 51.532073, lon: -0.123606 },
  { country: "United Kingdom", rank: 13, city: "Nottingham", drive_km: 14.1, drive_min: 21, walk_km: 14.1, walk_min: 188, straight_km: 10.9, lat: 52.947304, lon: -1.144956 },
  { country: "United Kingdom", rank: 14, city: "Birmingham", drive_km: 20.0, drive_min: 21, walk_km: 15.9, walk_min: 212, straight_km: 12.5, lat: 52.477646, lon: -1.898694 }
];
