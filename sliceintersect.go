package go_helper

import (
	"reflect"
)

func SliceIntersect(slices ...interface{}) (result []interface{}) {
	if len(slices) == 0 {
		return
	}
	arrV := reflect.ValueOf(slices[0])
	if arrV.Kind() != reflect.Slice {
		panic("parameters should be slices")
	}
	for i := 0; i < arrV.Len(); i++ {
		result = append(result, arrV.Index(i).Interface())
	}

	for i := 1; i < len(slices); i++ {
		arrV := reflect.ValueOf(slices[i])
		if arrV.Kind() != reflect.Slice {
			panic("parameters should be slices")
		}
		var newResult []interface{}
		for i := 0; i < arrV.Len(); i++ {
			element := arrV.Index(i).Interface()
			if !SliceContain(result, element) {
				continue
			}
			newResult = append(newResult, element)
		}
		result = newResult
	}

	return SliceUnique(result)
}

func SliceIntersectString(slices ...[]string) (result []string) {
	if len(slices) == 0 {
		return
	}
	for _, element := range slices[0] {
		result = append(result, element)
	}
	for i := 1; i < len(slices); i++ {
		var newResult []string
		for _, element := range slices[i] {
			if !SliceContainString(result, element) {
				continue
			}
			newResult = append(newResult, element)
		}
		result = newResult
	}

	return SliceUniqueString(result)
}

func SliceIntersectInt(slices ...[]int) (result []int) {
	if len(slices) == 0 {
		return
	}
	for _, element := range slices[0] {
		result = append(result, element)
	}
	for i := 1; i < len(slices); i++ {
		var newResult []int
		for _, element := range slices[i] {
			if !SliceContainInt(result, element) {
				continue
			}
			newResult = append(newResult, element)
		}
		result = newResult
	}

	return SliceUniqueInt(result)
}

func SliceIntersectUint(slices ...[]uint) (result []uint) {
	if len(slices) == 0 {
		return
	}
	for _, element := range slices[0] {
		result = append(result, element)
	}
	for i := 1; i < len(slices); i++ {
		var newResult []uint
		for _, element := range slices[i] {
			if !SliceContainUint(result, element) {
				continue
			}
			newResult = append(newResult, element)
		}
		result = newResult
	}

	return SliceUniqueUint(result)
}

func SliceIntersectInt8(slices ...[]int8) (result []int8) {
	if len(slices) == 0 {
		return
	}
	for _, element := range slices[0] {
		result = append(result, element)
	}
	for i := 1; i < len(slices); i++ {
		var newResult []int8
		for _, element := range slices[i] {
			if !SliceContainInt8(result, element) {
				continue
			}
			newResult = append(newResult, element)
		}
		result = newResult
	}

	return SliceUniqueInt8(result)
}

func SliceIntersectUint8(slices ...[]uint8) (result []uint8) {
	if len(slices) == 0 {
		return
	}
	for _, element := range slices[0] {
		result = append(result, element)
	}
	for i := 1; i < len(slices); i++ {
		var newResult []uint8
		for _, element := range slices[i] {
			if !SliceContainUint8(result, element) {
				continue
			}
			newResult = append(newResult, element)
		}
		result = newResult
	}

	return SliceUniqueUint8(result)
}

func SliceIntersectInt16(slices ...[]int16) (result []int16) {
	if len(slices) == 0 {
		return
	}
	for _, element := range slices[0] {
		result = append(result, element)
	}
	for i := 1; i < len(slices); i++ {
		var newResult []int16
		for _, element := range slices[i] {
			if !SliceContainInt16(result, element) {
				continue
			}
			newResult = append(newResult, element)
		}
		result = newResult
	}

	return SliceUniqueInt16(result)
}

func SliceIntersectUint16(slices ...[]uint16) (result []uint16) {
	if len(slices) == 0 {
		return
	}
	for _, element := range slices[0] {
		result = append(result, element)
	}
	for i := 1; i < len(slices); i++ {
		var newResult []uint16
		for _, element := range slices[i] {
			if !SliceContainUint16(result, element) {
				continue
			}
			newResult = append(newResult, element)
		}
		result = newResult
	}

	return SliceUniqueUint16(result)
}

func SliceIntersectInt32(slices ...[]int32) (result []int32) {
	if len(slices) == 0 {
		return
	}
	for _, element := range slices[0] {
		result = append(result, element)
	}
	for i := 1; i < len(slices); i++ {
		var newResult []int32
		for _, element := range slices[i] {
			if !SliceContainInt32(result, element) {
				continue
			}
			newResult = append(newResult, element)
		}
		result = newResult
	}

	return SliceUniqueInt32(result)
}

func SliceIntersectUint32(slices ...[]uint32) (result []uint32) {
	if len(slices) == 0 {
		return
	}
	for _, element := range slices[0] {
		result = append(result, element)
	}
	for i := 1; i < len(slices); i++ {
		var newResult []uint32
		for _, element := range slices[i] {
			if !SliceContainUint32(result, element) {
				continue
			}
			newResult = append(newResult, element)
		}
		result = newResult
	}

	return SliceUniqueUint32(result)
}

func SliceIntersectInt64(slices ...[]int64) (result []int64) {
	if len(slices) == 0 {
		return
	}
	for _, element := range slices[0] {
		result = append(result, element)
	}
	for i := 1; i < len(slices); i++ {
		var newResult []int64
		for _, element := range slices[i] {
			if !SliceContainInt64(result, element) {
				continue
			}
			newResult = append(newResult, element)
		}
		result = newResult
	}

	return SliceUniqueInt64(result)
}

func SliceIntersectUint64(slices ...[]uint64) (result []uint64) {
	if len(slices) == 0 {
		return
	}
	for _, element := range slices[0] {
		result = append(result, element)
	}
	for i := 1; i < len(slices); i++ {
		var newResult []uint64
		for _, element := range slices[i] {
			if !SliceContainUint64(result, element) {
				continue
			}
			newResult = append(newResult, element)
		}
		result = newResult
	}

	return SliceUniqueUint64(result)
}

func SliceIntersectFloat32(slices ...[]float32) (result []float32) {
	if len(slices) == 0 {
		return
	}
	for _, element := range slices[0] {
		result = append(result, element)
	}
	for i := 1; i < len(slices); i++ {
		var newResult []float32
		for _, element := range slices[i] {
			if !SliceContainFloat32(result, element) {
				continue
			}
			newResult = append(newResult, element)
		}
		result = newResult
	}

	return SliceUniqueFloat32(result)
}

func SliceIntersectFloat64(slices ...[]float64) (result []float64) {
	if len(slices) == 0 {
		return
	}
	for _, element := range slices[0] {
		result = append(result, element)
	}
	for i := 1; i < len(slices); i++ {
		var newResult []float64
		for _, element := range slices[i] {
			if !SliceContainFloat64(result, element) {
				continue
			}
			newResult = append(newResult, element)
		}
		result = newResult
	}

	return SliceUniqueFloat64(result)
}
