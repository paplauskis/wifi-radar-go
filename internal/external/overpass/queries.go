package overpass

import (
	"errors"
	"fmt"
	"strings"
)

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

func WifisByIDs(ids []int64) (string, error) {
	if len(ids) == 0 {
		return "", errors.New("No IDs provided")
	}
	var idString strings.Builder

	for i := 0; i < len(ids); i++ {
		if i != len(ids)-1 {
			idString.WriteString(fmt.Sprintf("%d, ", ids[i]))
		} else {
			idString.WriteString(fmt.Sprintf("%d", ids[i]))
		}
	}

	return `[out:json];
node(` + idString.String() + `)
out body;`, nil
}
