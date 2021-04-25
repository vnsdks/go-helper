package go_helper

import (
	"reflect"
	"strings"
)

func SliceDiff(slice1 interface{}, slices ...interface{}) (result []interface{}) {
	arrV := reflect.ValueOf(slice1)
	if arrV.Kind() != reflect.Slice {
		return
	}
	comparer := SliceMergeUnique(slices...)
	for i := 0; i < arrV.Len(); i++ {
		element := arrV.Index(i).Interface()
		if SliceContain(comparer, element) {
			continue
		}
		result = append(result, element)
	}

	return
}

func SliceDiffString(slice1 []string, slices ...[]string) (result []string) {
	checker := make(map[string]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			checker[element] = struct{}{}
		}
	}
	for _, element := range slice1 {
		if _, ok := checker[element]; ok {
			continue
		}
		result = append(result, element)
	}

	return
}

func SliceDiffStringInsensitive(slice1 []string, slices ...[]string) (result []string) {
	checker := make(map[string]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			checker[strings.ToLower(element)] = struct{}{}
		}
	}
	for _, element := range slice1 {
		if _, ok := checker[strings.ToLower(element)]; ok {
			continue
		}
		result = append(result, element)
	}

	return
}

func SliceDiffInt(slice1 []int, slices ...[]int) (result []int) {
	checker := make(map[int]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			checker[element] = struct{}{}
		}
	}
	for _, element := range slice1 {
		if _, ok := checker[element]; ok {
			continue
		}
		result = append(result, element)
	}

	return
}

func SliceDiffUint(slice1 []uint, slices ...[]uint) (result []uint) {
	checker := make(map[uint]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			checker[element] = struct{}{}
		}
	}
	for _, element := range slice1 {
		if _, ok := checker[element]; ok {
			continue
		}
		result = append(result, element)
	}

	return
}

func SliceDiffInt8(slice1 []int8, slices ...[]int8) (result []int8) {
	checker := make(map[int8]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			checker[element] = struct{}{}
		}
	}
	for _, element := range slice1 {
		if _, ok := checker[element]; ok {
			continue
		}
		result = append(result, element)
	}

	return
}

func SliceDiffUint8(slice1 []uint8, slices ...[]uint8) (result []uint8) {
	checker := make(map[uint8]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			checker[element] = struct{}{}
		}
	}
	for _, element := range slice1 {
		if _, ok := checker[element]; ok {
			continue
		}
		result = append(result, element)
	}

	return
}

func SliceDiffInt16(slice1 []int16, slices ...[]int16) (result []int16) {
	checker := make(map[int16]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			checker[element] = struct{}{}
		}
	}
	for _, element := range slice1 {
		if _, ok := checker[element]; ok {
			continue
		}
		result = append(result, element)
	}

	return
}

func SliceDiffUint16(slice1 []uint16, slices ...[]uint16) (result []uint16) {
	checker := make(map[uint16]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			checker[element] = struct{}{}
		}
	}
	for _, element := range slice1 {
		if _, ok := checker[element]; ok {
			continue
		}
		result = append(result, element)
	}

	return
}

func SliceDiffInt32(slice1 []int32, slices ...[]int32) (result []int32) {
	checker := make(map[int32]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			checker[element] = struct{}{}
		}
	}
	for _, element := range slice1 {
		if _, ok := checker[element]; ok {
			continue
		}
		result = append(result, element)
	}

	return
}

func SliceDiffUint32(slice1 []uint32, slices ...[]uint32) (result []uint32) {
	checker := make(map[uint32]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			checker[element] = struct{}{}
		}
	}
	for _, element := range slice1 {
		if _, ok := checker[element]; ok {
			continue
		}
		result = append(result, element)
	}

	return
}

func SliceDiffInt64(slice1 []int64, slices ...[]int64) (result []int64) {
	checker := make(map[int64]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			checker[element] = struct{}{}
		}
	}
	for _, element := range slice1 {
		if _, ok := checker[element]; ok {
			continue
		}
		result = append(result, element)
	}

	return
}

func SliceDiffUint64(slice1 []uint64, slices ...[]uint64) (result []uint64) {
	checker := make(map[uint64]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			checker[element] = struct{}{}
		}
	}
	for _, element := range slice1 {
		if _, ok := checker[element]; ok {
			continue
		}
		result = append(result, element)
	}

	return
}

func SliceDiffFloat32(slice1 []float32, slices ...[]float32) (result []float32) {
	checker := make(map[float32]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			checker[element] = struct{}{}
		}
	}
	for _, element := range slice1 {
		if _, ok := checker[element]; ok {
			continue
		}
		result = append(result, element)
	}

	return
}

func SliceDiffFloat64(slice1 []float64, slices ...[]float64) (result []float64) {
	checker := make(map[float64]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			checker[element] = struct{}{}
		}
	}
	for _, element := range slice1 {
		if _, ok := checker[element]; ok {
			continue
		}
		result = append(result, element)
	}

	return
}
