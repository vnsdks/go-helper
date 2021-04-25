package go_helper

import (
	"reflect"
	"strings"
)

func SliceSearch(slice interface{}, needle interface{}) (index int, isFounded bool) {
	arrV := reflect.ValueOf(slice)
	if arrV.Kind() != reflect.Slice {
		panic("parameter should be a slice")
	}
	for i := 0; i < arrV.Len(); i++ {
		element := arrV.Index(i).Interface()
		if reflect.DeepEqual(element, needle) {
			index = i
			isFounded = true
			return
		}
	}

	return
}

func SliceSearchString(slice []string, needle string) (index int, isFounded bool) {
	for i, element := range slice {
		if needle == element {
			index = i
			isFounded = true
			return
		}
	}

	return
}

func SliceSearchStringInsensitive(slice []string, needle string) (index int, isFounded bool) {
	for i, element := range slice {
		if strings.EqualFold(element, needle) {
			index = i
			isFounded = true
			return
		}
	}

	return
}

func SliceSearchInt(slice []int, needle int) (index int, isFounded bool) {
	for i, element := range slice {
		if needle == element {
			index = i
			isFounded = true
			return
		}
	}

	return
}

func SliceSearchUint(slice []uint, needle uint) (index int, isFounded bool) {
	for i, element := range slice {
		if needle == element {
			index = i
			isFounded = true
			return
		}
	}

	return
}

func SliceSearchInt8(slice []int8, needle int8) (index int, isFounded bool) {
	for i, element := range slice {
		if needle == element {
			index = i
			isFounded = true
			return
		}
	}

	return
}

func SliceSearchUint8(slice []uint8, needle uint8) (index int, isFounded bool) {
	for i, element := range slice {
		if needle == element {
			index = i
			isFounded = true
			return
		}
	}

	return
}

func SliceSearchInt16(slice []int16, needle int16) (index int, isFounded bool) {
	for i, element := range slice {
		if needle == element {
			index = i
			isFounded = true
			return
		}
	}

	return
}

func SliceSearchUint16(slice []uint16, needle uint16) (index int, isFounded bool) {
	for i, element := range slice {
		if needle == element {
			index = i
			isFounded = true
			return
		}
	}

	return
}

func SliceSearchInt32(slice []int32, needle int32) (index int, isFounded bool) {
	for i, element := range slice {
		if needle == element {
			index = i
			isFounded = true
			return
		}
	}

	return
}

func SliceSearchUint32(slice []uint32, needle uint32) (index int, isFounded bool) {
	for i, element := range slice {
		if needle == element {
			index = i
			isFounded = true
			return
		}
	}

	return
}

func SliceSearchInt64(slice []int64, needle int64) (index int, isFounded bool) {
	for i, element := range slice {
		if needle == element {
			index = i
			isFounded = true
			return
		}
	}

	return
}

func SliceSearchUint64(slice []uint64, needle uint64) (index int, isFounded bool) {
	for i, element := range slice {
		if needle == element {
			index = i
			isFounded = true
			return
		}
	}

	return
}

func SliceSearchFloat32(slice []float32, needle float32) (index int, isFounded bool) {
	for i, element := range slice {
		if needle == element {
			index = i
			isFounded = true
			return
		}
	}

	return
}

func SliceSearchFloat64(slice []float64, needle float64) (index int, isFounded bool) {
	for i, element := range slice {
		if needle == element {
			index = i
			isFounded = true
			return
		}
	}

	return
}
