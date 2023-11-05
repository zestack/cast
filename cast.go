package cast

import (
	"fmt"
	"reflect"
	"strings"
)

// FromType casts a string value to the given reflected type.
func FromType(s string, targetType reflect.Type) (any, error) {
	var typeName = targetType.String()

	if strings.HasPrefix(typeName, "[]") {
		itemType := typeName[2:]
		array := reflect.New(targetType).Elem()

		for _, v := range strings.Split(s, ",") {
			if item, err := FromString(strings.Trim(v, " \n\r"), itemType); err != nil {
				return array.Interface(), err
			} else {
				array = reflect.Append(array, reflect.ValueOf(item))
			}
		}

		return array.Interface(), nil
	}

	return FromString(s, typeName)
}

// FromString casts a string value to the given type name.
func FromString(s string, targetType string) (any, error) {
	switch targetType {
	case "int":
		return Int(s)
	case "int8":
		return Int8(s)
	case "int16":
		return Int16(s)
	case "int32":
		return Int32(s)
	case "int64":
		return Int64(s)
	case "uint":
		return Uint(s)
	case "uint8":
		return Uint8(s)
	case "uint16":
		return Uint16(s)
	case "uint32":
		return Uint32(s)
	case "uint64":
		return Uint64(s)
	case "bool":
		return Bool(s)
	case "float32":
		return Float32(s)
	case "float64":
		return Float64(s)
	case "string":
		return s, nil
	}
	return nil, fmt.Errorf("cast: type %v is not supported", targetType)
}
