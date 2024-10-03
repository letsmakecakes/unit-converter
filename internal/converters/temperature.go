package converters

import "errors"

func ConvertTemperature(fromUnit, toUnit string, value float64) (float64, error) {
	switch fromUnit {
	case "C":
		switch toUnit {
		case "F":
			return (value * 9 / 5) + 32, nil
		case "K":
			return value + 273.15, nil
		}
	case "F":
		switch toUnit {
		case "C":
			return (value - 32) * 5 / 9, nil
		case "K":
			return (value-32)*5/9 + 273.15, nil
		}
	case "K":
		switch toUnit {
		case "C":
			return value - 273.15, nil
		case "F":
			return (value - 273.15) * 9 / 5, nil
		}
	}

	return 0, errors.New("unsupported temperature unit conversion")
}
