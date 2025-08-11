package altseasoncore

import "strconv"

// Method to convert Oi struct to OiConverted struct
func (oi Oi) ToOiFloat() (OiConverted, error) {
	long, err := strconv.ParseFloat(oi.Long, 64)
	if err != nil {
		return OiConverted{}, err
	}

	short, err := strconv.ParseFloat(oi.Short, 64)
	if err != nil {
		return OiConverted{}, err
	}

	maxOI, err := strconv.ParseFloat(oi.Max, 64)
	if err != nil {
		return OiConverted{}, err
	}

	return OiConverted{
		Long:  long,
		Short: short,
		Max:   maxOI,
	}, nil
}
