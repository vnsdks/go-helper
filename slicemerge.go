package go_helper

import "reflect"

func SliceMerge(slices ...interface{}) (result []interface{}) {
	for _, slice := range slices {
		arrV := reflect.ValueOf(slice)
		if arrV.Kind() != reflect.Slice {
			panic("parameters should be slices")
		}
		for i := 0; i < arrV.Len(); i++ {
			result = append(result, arrV.Index(i).Interface())
		}
	}

	return
}

func SliceMergeUnique(slices ...interface{}) (result []interface{}) {
	for _, slice := range slices {
		arrV := reflect.ValueOf(slice)
		if arrV.Kind() != reflect.Slice {
			panic("parameters should be slices")
		}
		for i := 0; i < arrV.Len(); i++ {
			element := arrV.Index(i).Interface()
			if SliceContain(result, element) {
				continue
			}
			result = append(result, element)
		}
	}

	return
}

func SliceMergeString(slices ...[]string) (result []string) {
	for _, slice := range slices {
		result = append(result, slice...)
	}

	return
}

func SliceMergeStringUnique(slices ...[]string) (result []string) {
	checker := make(map[string]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			if _, ok := checker[element]; ok {
				continue
			}
			checker[element] = struct{}{}
			result = append(result, element)
		}
	}

	return
}

func SliceMergeInt(slices ...[]int) (result []int) {
	for _, slice := range slices {
		result = append(result, slice...)
	}

	return
}

func SliceMergeIntUnique(slices ...[]int) (result []int) {
	checker := make(map[int]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			if _, ok := checker[element]; ok {
				continue
			}
			checker[element] = struct{}{}
			result = append(result, element)
		}
	}

	return
}

func SliceMergeUint(slices ...[]uint) (result []uint) {
	for _, slice := range slices {
		result = append(result, slice...)
	}

	return
}

func SliceMergeUintUnique(slices ...[]uint) (result []uint) {
	checker := make(map[uint]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			if _, ok := checker[element]; ok {
				continue
			}
			checker[element] = struct{}{}
			result = append(result, element)
		}
	}

	return
}

func SliceMergeInt8(slices ...[]int8) (result []int8) {
	for _, slice := range slices {
		result = append(result, slice...)
	}

	return
}

func SliceMergeInt8Unique(slices ...[]int8) (result []int8) {
	checker := make(map[int8]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			if _, ok := checker[element]; ok {
				continue
			}
			checker[element] = struct{}{}
			result = append(result, element)
		}
	}

	return
}

func SliceMergeUint8(slices ...[]uint8) (result []uint8) {
	for _, slice := range slices {
		result = append(result, slice...)
	}

	return
}

func SliceMergeUint8Unique(slices ...[]uint8) (result []uint8) {
	checker := make(map[uint8]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			if _, ok := checker[element]; ok {
				continue
			}
			checker[element] = struct{}{}
			result = append(result, element)
		}
	}

	return
}

func SliceMergeInt16(slices ...[]int16) (result []int16) {
	for _, slice := range slices {
		result = append(result, slice...)
	}

	return
}

func SliceMergeInt16Unique(slices ...[]int16) (result []int16) {
	checker := make(map[int16]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			if _, ok := checker[element]; ok {
				continue
			}
			checker[element] = struct{}{}
			result = append(result, element)
		}
	}

	return
}

func SliceMergeUint16(slices ...[]uint16) (result []uint16) {
	for _, slice := range slices {
		result = append(result, slice...)
	}

	return
}

func SliceMergeUint16Unique(slices ...[]uint16) (result []uint16) {
	checker := make(map[uint16]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			if _, ok := checker[element]; ok {
				continue
			}
			checker[element] = struct{}{}
			result = append(result, element)
		}
	}

	return
}

func SliceMergeInt32(slices ...[]int32) (result []int32) {
	for _, slice := range slices {
		result = append(result, slice...)
	}

	return
}

func SliceMergeInt32Unique(slices ...[]int32) (result []int32) {
	checker := make(map[int32]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			if _, ok := checker[element]; ok {
				continue
			}
			checker[element] = struct{}{}
			result = append(result, element)
		}
	}

	return
}

func SliceMergeUint32(slices ...[]uint32) (result []uint32) {
	for _, slice := range slices {
		result = append(result, slice...)
	}

	return
}

func SliceMergeUint32Unique(slices ...[]uint32) (result []uint32) {
	checker := make(map[uint32]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			if _, ok := checker[element]; ok {
				continue
			}
			checker[element] = struct{}{}
			result = append(result, element)
		}
	}

	return
}

func SliceMergeInt64(slices ...[]int64) (result []int64) {
	for _, slice := range slices {
		result = append(result, slice...)
	}

	return
}

func SliceMergeInt64Unique(slices ...[]int64) (result []int64) {
	checker := make(map[int64]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			if _, ok := checker[element]; ok {
				continue
			}
			checker[element] = struct{}{}
			result = append(result, element)
		}
	}

	return
}

func SliceMergeUint64(slices ...[]uint64) (result []uint64) {
	for _, slice := range slices {
		result = append(result, slice...)
	}

	return
}

func SliceMergeUint64Unique(slices ...[]uint64) (result []uint64) {
	checker := make(map[uint64]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			if _, ok := checker[element]; ok {
				continue
			}
			checker[element] = struct{}{}
			result = append(result, element)
		}
	}

	return
}

func SliceMergeFloat32(slices ...[]float32) (result []float32) {
	for _, slice := range slices {
		result = append(result, slice...)
	}

	return
}

func SliceMergeFloat32Unique(slices ...[]float32) (result []float32) {
	checker := make(map[float32]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			if _, ok := checker[element]; ok {
				continue
			}
			checker[element] = struct{}{}
			result = append(result, element)
		}
	}

	return
}

func SliceMergeFloat64(slices ...[]float64) (result []float64) {
	for _, slice := range slices {
		result = append(result, slice...)
	}

	return
}

func SliceMergeFloat64Unique(slices ...[]float64) (result []float64) {
	checker := make(map[float64]struct{}, 0)
	for _, slice := range slices {
		for _, element := range slice {
			if _, ok := checker[element]; ok {
				continue
			}
			checker[element] = struct{}{}
			result = append(result, element)
		}
	}

	return
}
