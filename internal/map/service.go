package _map

import (
	"encoding/json"
	"log"
	"wifi-radar-go/internal/external/overpass"
)

func GetWifiNetworks(ids []int64) OverpassElementListDTO {
	client := overpass.NewClient()
	query, err := overpass.WifisByIDs(ids)

	if err != nil {
		log.Println(err)
	}

	data, err := client.Query(query)

	if err != nil {
		log.Printf("Failed to query %s ERROR: %v", client.ApiURL, err)
	}

	elements := OverpassElementListDTO{}

	err = json.Unmarshal(data, &elements)

	if err != nil {
		log.Println(err)
	}

	return elements
}
