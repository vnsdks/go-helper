package go_helper

import (
	"math/rand"
	"reflect"
	"time"
)

func SliceRand(slice interface{}) interface{} {
	arrV := reflect.ValueOf(slice)
	if arrV.Kind() != reflect.Slice {
		panic("parameter should be a slice")
	}
	rand.Seed(time.Now().Unix())

	return arrV.Index(rand.Intn(arrV.Len())).Interface()
}

func SliceRandString(slice []string) string {
	rand.Seed(time.Now().Unix())

	return slice[rand.Intn(len(slice))]
}

func SliceRandInt(slice []int) int {
	rand.Seed(time.Now().Unix())

	return slice[rand.Intn(len(slice))]
}

func SliceRandUint(slice []uint) uint {
	rand.Seed(time.Now().Unix())

	return slice[rand.Intn(len(slice))]
}

func SliceRandInt8(slice []int8) int8 {
	rand.Seed(time.Now().Unix())

	return slice[rand.Intn(len(slice))]
}

func SliceRandUint8(slice []uint8) uint8 {
	rand.Seed(time.Now().Unix())

	return slice[rand.Intn(len(slice))]
}

func SliceRandInt16(slice []int16) int16 {
	rand.Seed(time.Now().Unix())

	return slice[rand.Intn(len(slice))]
}

func SliceRandUint16(slice []uint16) uint16 {
	rand.Seed(time.Now().Unix())

	return slice[rand.Intn(len(slice))]
}

func SliceRandInt32(slice []int32) int32 {
	rand.Seed(time.Now().Unix())

	return slice[rand.Intn(len(slice))]
}

func SliceRandUint32(slice []uint32) uint32 {
	rand.Seed(time.Now().Unix())

	return slice[rand.Intn(len(slice))]
}

func SliceRandInt64(slice []int64) int64 {
	rand.Seed(time.Now().Unix())

	return slice[rand.Intn(len(slice))]
}

func SliceRandUint64(slice []uint64) uint64 {
	rand.Seed(time.Now().Unix())

	return slice[rand.Intn(len(slice))]
}

func SliceRandFloat32(slice []float32) float32 {
	rand.Seed(time.Now().Unix())

	return slice[rand.Intn(len(slice))]
}

func SliceRandFloat64(slice []float64) float64 {
	rand.Seed(time.Now().Unix())

	return slice[rand.Intn(len(slice))]
}
