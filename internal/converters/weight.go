package converters

import "errors"

func ConvertWeight(fromUnit, toUnit string, value float64) (float64, error) {
	conversions := map[string]map[string]float64{
		"g":     {"kg": 0.001, "lbs": 0.00220462},
		"kg":    {"g": 1000, "lbs": 2.20462},
		"lbs":   {"g": 453.592, "kg": 0.453592},
		"ounce": {"g": 28.3495, "kg": 0.0283495, "lbs": 0.0625},
		"ton":   {"kg": 1000, "lbs": 2204.62},
	}

	if toUnits, ok := conversions[fromUnit]; ok {
		if factor, ok := toUnits[toUnit]; ok {
			return value * factor, nil
		}
	}

	return 0, errors.New("unsupported weight unit conversion")
}
