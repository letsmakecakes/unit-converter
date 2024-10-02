package converters

import "errors"

type LengthConverter struct {
	FromUnit string  `json:"from_unit"`
	ToUnit   string  `json:"to_unit"`
	Value    float64 `json:"value"`
}

func (lc *LengthConverter) Convert() (float64, error) {
	switch lc.FromUnit {
	case "mm":
		if lc.ToUnit == "cm" {
			return lc.Value / 10, nil
		}
	case "cm":
		if lc.ToUnit == "m" {
			return lc.Value / 100, nil
		}
	case "m":
		if lc.ToUnit == "km" {
			return lc.Value / 1000, nil
		}
	case "ft":
		if lc.ToUnit == "cm" {
			return lc.Value * 30.48, nil
		}
	case "yard":
		if lc.ToUnit == "mile" {
			return lc.Value / 1760, nil
		}
	}

	return 0, errors.New("unsupported length unit conversion")
}
