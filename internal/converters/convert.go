package converters

type converter interface {
	Convert() (float64, error)
}

func GetResult(c converter) (float64, error) {
	value, err := c.Convert()
	if err != nil {
		return 0, nil
	}

	return value, err
}
