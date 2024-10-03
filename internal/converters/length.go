package converters

import (
	"errors"
)

func ConvertLength(fromUnit, toUnit string, value float64) (float64, error) {
	conversions := map[string]map[string]float64{
		"mm":   {"cm": 0.1, "m": 0.001, "km": 1e-6},
		"cm":   {"mm": 10, "m": 0.01, "km": 1e-5},
		"m":    {"cm": 100, "mm": 1000, "km": 0.001},
		"km":   {"m": 1000, "cm": 100000, "mm": 1e6},
		"ft":   {"cm": 30.48, "m": 0.3048},
		"yard": {"mile": 1.0 / 1760},
		// Add more conversions here as needed
	}

	// Check if the conversion exists
	if toUnits, ok := conversions[fromUnit]; ok {
		if factor, ok := toUnits[toUnit]; ok {
			return value * factor, nil
		}
	}

	// If no valid conversion is found
	return 0, errors.New("unsupported length unit conversion")
}
