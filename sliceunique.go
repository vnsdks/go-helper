package go_helper

import "reflect"

func SliceUnique(slice interface{}) (result []interface{}) {
	arrV := reflect.ValueOf(slice)
	if arrV.Kind() != reflect.Slice {
		panic("parameter should be a slice")
	}
	for i := 0; i < arrV.Len(); i++ {
		element := arrV.Index(i).Interface()
		if SliceContain(result, element) {
			continue
		}
		result = append(result, element)
	}

	return
}

func SliceUniqueString(slice []string) (result []string) {
	checker := make(map[string]struct{}, 0)
	for _, element := range slice {
		if _, ok := checker[element]; ok {
			continue
		}
		checker[element] = struct{}{}
		result = append(result, element)
	}

	return
}

func SliceUniqueInt(slice []int) (result []int) {
	checker := make(map[int]struct{}, 0)
	for _, element := range slice {
		if _, ok := checker[element]; ok {
			continue
		}
		checker[element] = struct{}{}
		result = append(result, element)
	}

	return
}

func SliceUniqueUint(slice []uint) (result []uint) {
	checker := make(map[uint]struct{}, 0)
	for _, element := range slice {
		if _, ok := checker[element]; ok {
			continue
		}
		checker[element] = struct{}{}
		result = append(result, element)
	}

	return
}

func SliceUniqueInt8(slice []int8) (result []int8) {
	checker := make(map[int8]struct{}, 0)
	for _, element := range slice {
		if _, ok := checker[element]; ok {
			continue
		}
		checker[element] = struct{}{}
		result = append(result, element)
	}

	return
}

func SliceUniqueUint8(slice []uint8) (result []uint8) {
	checker := make(map[uint8]struct{}, 0)
	for _, element := range slice {
		if _, ok := checker[element]; ok {
			continue
		}
		checker[element] = struct{}{}
		result = append(result, element)
	}

	return
}

func SliceUniqueInt16(slice []int16) (result []int16) {
	checker := make(map[int16]struct{}, 0)
	for _, element := range slice {
		if _, ok := checker[element]; ok {
			continue
		}
		checker[element] = struct{}{}
		result = append(result, element)
	}

	return
}

func SliceUniqueUint16(slice []uint16) (result []uint16) {
	checker := make(map[uint16]struct{}, 0)
	for _, element := range slice {
		if _, ok := checker[element]; ok {
			continue
		}
		checker[element] = struct{}{}
		result = append(result, element)
	}

	return
}

func SliceUniqueInt32(slice []int32) (result []int32) {
	checker := make(map[int32]struct{}, 0)
	for _, element := range slice {
		if _, ok := checker[element]; ok {
			continue
		}
		checker[element] = struct{}{}
		result = append(result, element)
	}

	return
}

func SliceUniqueUint32(slice []uint32) (result []uint32) {
	checker := make(map[uint32]struct{}, 0)
	for _, element := range slice {
		if _, ok := checker[element]; ok {
			continue
		}
		checker[element] = struct{}{}
		result = append(result, element)
	}

	return
}

func SliceUniqueInt64(slice []int64) (result []int64) {
	checker := make(map[int64]struct{}, 0)
	for _, element := range slice {
		if _, ok := checker[element]; ok {
			continue
		}
		checker[element] = struct{}{}
		result = append(result, element)
	}

	return
}

func SliceUniqueUint64(slice []uint64) (result []uint64) {
	checker := make(map[uint64]struct{}, 0)
	for _, element := range slice {
		if _, ok := checker[element]; ok {
			continue
		}
		checker[element] = struct{}{}
		result = append(result, element)
	}

	return
}

func SliceUniqueFloat32(slice []float32) (result []float32) {
	checker := make(map[float32]struct{}, 0)
	for _, element := range slice {
		if _, ok := checker[element]; ok {
			continue
		}
		checker[element] = struct{}{}
		result = append(result, element)
	}

	return
}

func SliceUniqueFloat64(slice []float64) (result []float64) {
	checker := make(map[float64]struct{}, 0)
	for _, element := range slice {
		if _, ok := checker[element]; ok {
			continue
		}
		checker[element] = struct{}{}
		result = append(result, element)
	}

	return
}
