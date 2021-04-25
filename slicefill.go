package go_helper

import "reflect"

func SliceFillOverride(slice interface{}, startIndex int, count int, value interface{}) (totalFilled int) {
	arrV := reflect.ValueOf(slice)
	if arrV.Kind() != reflect.Slice {
		panic("parameter should be a slice")
	}
	vv := reflect.ValueOf(value)
	if vv.Type() != arrV.Type().Elem() {
		panic("value should have same type with slice elements")
	}
	for i := startIndex; i < arrV.Len(); i++ {
		if totalFilled >= count {
			break
		}
		totalFilled++
		arrV.Index(i).Set(vv)
	}

	return
}

func SliceFill(slice interface{}, startIndex int, count int, value interface{}) (result []interface{}) {
	arrV := reflect.ValueOf(slice)
	if arrV.Kind() != reflect.Slice {
		panic("parameter should be a slice")
	}
	vv := reflect.ValueOf(value)
	if vv.Type() != arrV.Type().Elem() {
		panic("value should have same type with slice elements")
	}
	for i := 0; i < arrV.Len(); i++ {
		newValue := arrV.Index(i).Interface()
		if i >= startIndex && i < startIndex+count {
			newValue = vv
		}
		result = append(result, newValue)
	}

	return
}

func SliceFillString(slice []string, startIndex int, count int, value string) (result []string) {
	for i := 0; i < len(slice); i++ {
		newValue := slice[i]
		if i >= startIndex && i < startIndex+count {
			newValue = value
		}
		result = append(result, newValue)
	}

	return
}

func SliceFillInt(slice []int, startIndex int, count int, value int) (result []int) {
	for i := 0; i < len(slice); i++ {
		newValue := slice[i]
		if i >= startIndex && i < startIndex+count {
			newValue = value
		}
		result = append(result, newValue)
	}

	return
}

func SliceFillUint(slice []uint, startIndex int, count int, value uint) (result []uint) {
	for i := 0; i < len(slice); i++ {
		newValue := slice[i]
		if i >= startIndex && i < startIndex+count {
			newValue = value
		}
		result = append(result, newValue)
	}

	return
}

func SliceFillInt8(slice []int8, startIndex int, count int, value int8) (result []int8) {
	for i := 0; i < len(slice); i++ {
		newValue := slice[i]
		if i >= startIndex && i < startIndex+count {
			newValue = value
		}
		result = append(result, newValue)
	}

	return
}

func SliceFillUint8(slice []uint8, startIndex int, count int, value uint8) (result []uint8) {
	for i := 0; i < len(slice); i++ {
		newValue := slice[i]
		if i >= startIndex && i < startIndex+count {
			newValue = value
		}
		result = append(result, newValue)
	}

	return
}

func SliceFillInt16(slice []int16, startIndex int, count int, value int16) (result []int16) {
	for i := 0; i < len(slice); i++ {
		newValue := slice[i]
		if i >= startIndex && i < startIndex+count {
			newValue = value
		}
		result = append(result, newValue)
	}

	return
}

func SliceFillUint16(slice []uint16, startIndex int, count int, value uint16) (result []uint16) {
	for i := 0; i < len(slice); i++ {
		newValue := slice[i]
		if i >= startIndex && i < startIndex+count {
			newValue = value
		}
		result = append(result, newValue)
	}

	return
}

func SliceFillInt32(slice []int32, startIndex int, count int, value int32) (result []int32) {
	for i := 0; i < len(slice); i++ {
		newValue := slice[i]
		if i >= startIndex && i < startIndex+count {
			newValue = value
		}
		result = append(result, newValue)
	}

	return
}

func SliceFillUint32(slice []uint32, startIndex int, count int, value uint32) (result []uint32) {
	for i := 0; i < len(slice); i++ {
		newValue := slice[i]
		if i >= startIndex && i < startIndex+count {
			newValue = value
		}
		result = append(result, newValue)
	}

	return
}

func SliceFillInt64(slice []int64, startIndex int, count int, value int64) (result []int64) {
	for i := 0; i < len(slice); i++ {
		newValue := slice[i]
		if i >= startIndex && i < startIndex+count {
			newValue = value
		}
		result = append(result, newValue)
	}

	return
}

func SliceFillUint64(slice []uint64, startIndex int, count int, value uint64) (result []uint64) {
	for i := 0; i < len(slice); i++ {
		newValue := slice[i]
		if i >= startIndex && i < startIndex+count {
			newValue = value
		}
		result = append(result, newValue)
	}

	return
}

func SliceFillFloat32(slice []float32, startIndex int, count int, value float32) (result []float32) {
	for i := 0; i < len(slice); i++ {
		newValue := slice[i]
		if i >= startIndex && i < startIndex+count {
			newValue = value
		}
		result = append(result, newValue)
	}

	return
}

func SliceFillFloat64(slice []float64, startIndex int, count int, value float64) (result []float64) {
	for i := 0; i < len(slice); i++ {
		newValue := slice[i]
		if i >= startIndex && i < startIndex+count {
			newValue = value
		}
		result = append(result, newValue)
	}

	return
}
