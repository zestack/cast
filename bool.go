package cast

import "strconv"

func Bool(s string) (bool, error) {
	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return v, nil
}
