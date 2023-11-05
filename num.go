package cast

import (
	"fmt"
	"strconv"
)

func Uint(s string) (uint, error) {
	v, err := strconv.ParseUint(s, 0, strconv.IntSize)
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}

func Uint8(s string) (uint8, error) {
	v, err := strconv.ParseUint(s, 0, 8)
	if err != nil {
		return 0, err
	}
	return uint8(v), nil
}

func Uint16(s string) (uint16, error) {
	v, err := strconv.ParseUint(s, 0, 16)
	if err != nil {
		return 0, err
	}
	return uint16(v), nil
}

func Uint32(s string) (uint32, error) {
	v, err := strconv.ParseUint(s, 0, 32)
	if err != nil {
		return 0, err
	}
	return uint32(v), nil
}

func Uint64(s string) (uint64, error) {
	v, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func Int(s string) (int, error) {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		return 0, fmt.Errorf("cast: cannot cast `%v` to type `%v`", s, "int")
	}
	return int(v), nil
}

func Int8(s string) (int8, error) {
	v, err := strconv.ParseInt(s, 0, 8)
	if err != nil {
		return 0, err
	}
	return int8(v), nil
}

func Int16(s string) (int16, error) {
	v, err := strconv.ParseInt(s, 0, 16)
	if err != nil {
		return 0, err
	}
	return int16(v), nil
}

func Int32(s string) (int32, error) {
	v, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}

func Int64(s string) (int64, error) {
	v, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func Float32(s string) (float32, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return float32(v), nil
}

func Float64(s string) (float64, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return v, nil
}
