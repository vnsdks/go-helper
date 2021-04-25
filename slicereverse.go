package go_helper

import "reflect"

func SliceReserve(slice interface{}) (result []interface{}) {
	arrV := reflect.ValueOf(slice)
	if arrV.Kind() != reflect.Slice {
		return
	}
	for i := arrV.Len() - 1; i >= 0; i-- {
		result = append(result, arrV.Index(i).Interface())
	}

	return
}

func SliceReserveString(slice []string) (result []string) {
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}

	return
}

func SliceReserveInt(slice []int) (result []int) {
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}

	return
}

func SliceReserveUint(slice []uint) (result []uint) {
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}

	return
}

func SliceReserveInt8(slice []int8) (result []int8) {
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}

	return
}

func SliceReserveUint8(slice []uint8) (result []uint8) {
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}

	return
}

func SliceReserveInt16(slice []int16) (result []int16) {
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}

	return
}

func SliceReserveUint16(slice []uint16) (result []uint16) {
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}

	return
}

func SliceReserveInt32(slice []int32) (result []int32) {
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}

	return
}

func SliceReserveUint32(slice []uint32) (result []uint32) {
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}

	return
}

func SliceReserveInt64(slice []int64) (result []int64) {
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}

	return
}

func SliceReserveUint64(slice []uint64) (result []uint64) {
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}

	return
}

func SliceReserveFloat32(slice []float32) (result []float32) {
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}

	return
}

func SliceReserveFloat64(slice []float64) (result []float64) {
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}

	return
}
