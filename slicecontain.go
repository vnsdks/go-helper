package go_helper

func SliceContain(slice interface{}, needle interface{}) bool {
	_, result := SliceSearch(slice, needle)

	return result
}

func SliceContainString(slice []string, needle string) bool {
	_, result := SliceSearchString(slice, needle)

	return result
}

func SliceContainStringInsensitive(slice []string, needle string) bool {
	_, result := SliceSearchStringInsensitive(slice, needle)

	return result
}

func SliceContainInt(slice []int, needle int) bool {
	_, result := SliceSearchInt(slice, needle)

	return result
}

func SliceContainUint(slice []uint, needle uint) bool {
	_, result := SliceSearchUint(slice, needle)

	return result
}

func SliceContainInt8(slice []int8, needle int8) bool {
	_, result := SliceSearchInt8(slice, needle)

	return result
}

func SliceContainUint8(slice []uint8, needle uint8) bool {
	_, result := SliceSearchUint8(slice, needle)

	return result
}

func SliceContainInt16(slice []int16, needle int16) bool {
	_, result := SliceSearchInt16(slice, needle)

	return result
}

func SliceContainUint16(slice []uint16, needle uint16) bool {
	_, result := SliceSearchUint16(slice, needle)

	return result
}

func SliceContainInt32(slice []int32, needle int32) bool {
	_, result := SliceSearchInt32(slice, needle)

	return result
}

func SliceContainUint32(slice []uint32, needle uint32) bool {
	_, result := SliceSearchUint32(slice, needle)

	return result
}

func SliceContainInt64(slice []int64, needle int64) bool {
	_, result := SliceSearchInt64(slice, needle)

	return result
}

func SliceContainUint64(slice []uint64, needle uint64) bool {
	_, result := SliceSearchUint64(slice, needle)

	return result
}

func SliceContainFloat32(slice []float32, needle float32) bool {
	_, result := SliceSearchFloat32(slice, needle)

	return result
}

func SliceContainFloat64(slice []float64, needle float64) bool {
	_, result := SliceSearchFloat64(slice, needle)

	return result
}
