package internal

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func GetMetricByName(obj interface{}, fieldName string) reflect.Value {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct.Elem()
	if curStruct.Kind() != reflect.Struct {
		panic("not struct")
	}
	curField := curStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		panic("not found:" + fieldName)
	}
	return curField
}

func GetValueAndTypeFromReflection(value reflect.Value) (string, string, error) {
	switch value.Kind().String() {
	case "float64":
		return fmt.Sprintf("%f", value.Interface().(float64)), "gauge", nil
	case "uint64":
		return strconv.FormatUint(value.Interface().(uint64), 10), "counter", nil // https://go.dev/blog/laws-of-reflection
	case "uint32":
		return strconv.FormatUint(uint64(value.Interface().(uint32)), 10), "counter", nil // https://go.dev/blog/laws-of-reflection
	}
	return "", "", errors.New("unexpected value type") // что возвращать?
}
