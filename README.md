## Introduction
You heard of the Bigmac Index, now we have the Decathlon Index, a simple measure of walkability in a city. Where does your city rank, a grass green walkable paradise or sewer brown car dominated concrete jungle?

This happened because I went on holiday, got to the city and it was raining and
needed an umbrella, stumbled upon the Decathlon store and thought WTAF it is literally right in the centre of town! In my city you have to travel for an hour
to the car filled bypass before you find a Decathlon store. And dont try cycling there cause you will be flattened on the urban motorways surrouding the retail
park.

## What is it
A simple dataset of 100 cities in Europe containing the distances from the main
train station to the closest Decathlon store. Time to walk and time to drive. Ranked from grass green to car swewer poo brown.

## How to generate a fresh dataset
Run the decadist.go file
```
go run .
```

This will connect to nominatim openstreetmap to return the geolocation of each
train station and Decathon store. It then uses the OSRM server to calculate the
walking and driving times and distances.

Two files are generated JS and CSV.