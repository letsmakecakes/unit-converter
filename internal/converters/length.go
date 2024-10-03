package converters

import "errors"

type LengthConverter struct {
	FromUnit string  `json:"from_unit"`
	ToUnit   string  `json:"to_unit"`
	Value    float64 `json:"value"`
}

func (lc LengthConverter) Convert() (float64, error) {
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
	if toUnits, ok := conversions[lc.FromUnit]; ok {
		if factor, ok := toUnits[lc.ToUnit]; ok {
			return lc.Value * factor, nil
		}
	}

	// If no valid conversion is found
	return 0, errors.New("unsupported length unit conversion")
}
