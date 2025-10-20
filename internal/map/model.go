package _map

type OverpassElementListDTO struct {
	Version  float64              `json:"version"`
	Elements []OverpassElementDTO `json:"elements"`
}

type OverpassElementDTO struct {
	WifiID    int64          `json:"id"`
	Latitude  float64        `json:"lat"`
	Longitude float64        `json:"lon"`
	Tags      OverpassTagDTO `json:"tags"`
}

type OverpassTagDTO struct {
	City           string `json:"addr:city"`
	PlaceName      string `json:"name"`
	Street         string `json:"addr:street"`
	BuildingNumber string `json:"addr:housenumber"`
	Postcode       string `json:"addr:postcode"`
}

type Coordinates struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"lon"`
}
