package go_helper

import (
	"reflect"
	"strings"
)

func SliceConvertInterface(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil
	}

	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func SliceConvertUpper(slice []string) []string {
	result := make([]string, len(slice))
	for i, element := range slice {
		result[i] = strings.ToUpper(element)
	}

	return result
}

func SliceConvertLower(slice []string) []string {
	result := make([]string, len(slice))
	for i, element := range slice {
		result[i] = strings.ToLower(element)
	}

	return result
}
