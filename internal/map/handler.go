package _map

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"wifi-radar-go/internal/external/openstreetmap"
	"wifi-radar-go/internal/external/overpass"
	"wifi-radar-go/internal/validation"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Search(context *gin.Context) {
	city := context.Query("city")
	radiusStr := context.Query("radius")

	if city == "" {
		context.JSON(http.StatusBadRequest, "City not specified.")
		return
	}

	if radiusStr == "" || radiusStr == "0" {
		elements := searchInCity(city)
		context.JSON(http.StatusOK, elements)
		return
	}

	radius, err := strconv.Atoi(radiusStr)
	if err != nil || radius > 100000 || radius < 0 {
		context.JSON(http.StatusBadRequest, "Invalid radius specified.")
		return
	}

	elements := searchInRadius(city, radius)
	context.JSON(http.StatusOK, elements)
}

func (h *Handler) GetCoordinates(context *gin.Context) {
	city := context.Query("city")
	street := context.Query("street")
	building := context.Query("buildingNumber")

	buildingNumber, err := validation.ValidateAddress(city, street, building)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
	}

	client := overpass.NewClient()
	query := overpass.WifiCoordinates(city, street, strconv.Itoa(buildingNumber))
	data, err := client.Query(query)

	var response OverpassElementListDTO
	err = json.Unmarshal(data, &response)

	if err != nil {
		log.Println(err)
	}

	if len(response.Elements) == 0 {
		context.JSON(http.StatusNotFound, "No wi-fi network was found with this address")
		return
	}

	coordinates := Coordinates{
		Longitude: strconv.FormatFloat(response.Elements[0].Longitude, 'f', 6, 64),
		Latitude:  strconv.FormatFloat(response.Elements[0].Latitude, 'f', 6, 64),
	}

	context.JSON(http.StatusOK, coordinates)
}

func searchInCity(city string) OverpassElementListDTO {
	client := overpass.NewClient()
	query := overpass.FreeWifiInCity(city)
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

func searchInRadius(city string, radius int) OverpassElementListDTO {
	coordinates := getCityCenterPoint(city)
	client := overpass.NewClient()
	query := overpass.FreeWifiInRadius(radius, coordinates.Latitude, coordinates.Longitude)
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

func getCityCenterPoint(city string) Coordinates {
	clientOSM := openstreetmap.NewClient(city)
	dataOSM, err := clientOSM.Query()

	if err != nil {
		log.Printf("Failed to query %s ERROR: %v", clientOSM.ApiURL, err)
	}

	var coordinates []Coordinates
	err = json.Unmarshal(dataOSM, &coordinates)

	if err != nil {
		log.Println(err)
	}

	return coordinates[0]
}
