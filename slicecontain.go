package go_helper

import (
	"reflect"
	"strings"
)

func SliceContain(slice interface{}, element interface{}) bool {
	arrV := reflect.ValueOf(slice)
	if arrV.Kind() != reflect.Slice {
		return false
	}
	for i := 0; i < arrV.Len(); i++ {
		if reflect.DeepEqual(arrV.Index(i).Interface(), element) {
			return true
		}
	}

	return false
}

func SliceContainString(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}

func SliceContainStringInsensitive(slice []string, element string) bool {
	for _, v := range slice {
		if strings.EqualFold(v, element) {
			return true
		}
	}

	return false
}

func SliceContainInt(slice []int, element int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}

func SliceContainUint(slice []uint, element uint) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}

func SliceContainInt8(slice []int8, element int8) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}

func SliceContainUint8(slice []uint8, element uint8) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}

func SliceContainInt16(slice []int16, element int16) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}

func SliceContainUint16(slice []uint16, element uint16) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}

func SliceContainInt32(slice []int32, element int32) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}

func SliceContainUint32(slice []uint32, element uint32) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}

func SliceContainInt64(slice []int64, element int64) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}

func SliceContainUint64(slice []uint64, element uint64) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}

func SliceContainFloat32(slice []float32, element float32) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}

func SliceContainFloat64(slice []float64, element float64) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}
