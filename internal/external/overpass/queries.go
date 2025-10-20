package overpass

import "fmt"

func FreeWifiInCity(city string) string {
	return `[out:json];
node
  ["internet_access"="wlan"]
  ["internet_access:fee"="no"]
  ["addr:city"="` + city + `"];
out body;`
}

func FreeWifiInRadius(radius int, lat, lon string) string {
	return `[out:json];
node
  ["internet_access"="wlan"]
  ["internet_access:fee"="no"]
  (around:` + fmt.Sprintf("%d,%s,%s", radius, lat, lon) + `);
out body;`
}

func WifiCoordinates(city, street string, buildingNumber string) string {
	return `[out:json];
node
  ["internet_access"="wlan"]
  ["internet_access:fee"="no"]
  ["addr:city"="` + city + `"]
  ["addr:street"="` + street + `"]
  ["addr:housenumber"="` + buildingNumber + `"];
out body;`
}
