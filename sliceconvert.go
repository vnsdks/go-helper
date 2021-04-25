package go_helper

import (
	"reflect"
	"strings"
)

func SliceConvertInterface(slice interface{}) []interface{} {
	arrV := reflect.ValueOf(slice)
	if arrV.Kind() != reflect.Slice {
		panic("parameter should be a slice")
	}
	if arrV.IsNil() {
		return nil
	}
	ret := make([]interface{}, arrV.Len())
	for i := 0; i < arrV.Len(); i++ {
		ret[i] = arrV.Index(i).Interface()
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
