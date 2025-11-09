package validation

import (
	"errors"
	"strconv"
)

func ValidateAddress(city, street, building string) (int, error) {
	if city == "" || street == "" || building == "" {
		return 0, errors.New("city, street or building number not specified")
	}

	buildingNumber, err := strconv.Atoi(building)
	if err != nil || buildingNumber < 1 {
		return 0, errors.New("invalid building number specified")
	}

	return buildingNumber, nil
}
