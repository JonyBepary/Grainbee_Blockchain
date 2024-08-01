package datatypes

import (
	"strconv"
	"strings"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var coordinates = assets.DataType{
	AcceptedFormats: []string{"string"},
	Description:     "A string representing coordinates in the format 'latitude,longitude'.",
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		coordStr, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("property must be a string", 400)
		}

		parts := strings.Split(coordStr, ",")
		if len(parts) != 2 {
			return "", nil, errors.NewCCError("coordinates must be in the format 'latitude,longitude'", 400)
		}
		latitude, err := strconv.ParseFloat(parts[0], 64)
		if err != nil || latitude < -90 || latitude > 90 {
			return "", nil, errors.NewCCError("invalid latitude", 400)
		}

		longitude, err := strconv.ParseFloat(parts[1], 64)
		if err != nil || longitude < -180 || longitude > 180 {
			return "", nil, errors.NewCCError("invalid longitude", 400)
		}
		coord := Coordinates{
			Latitude:  latitude,
			Longitude: longitude,
		}
		return "coordinates", coord, nil
	},
}
