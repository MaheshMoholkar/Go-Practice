package util

import "fmt"

var (
	_float64    float64 = 1.5
	_float32    float32 = 1.5
	_string     string  = "Gopher"
	_bool       bool    = true
	_int        int     = 1
	_int8       int8    = 1
	_int16      int16   = 1
	_int32      int32   = 1
	_int64      int64   = 1
	_uint       uint    = 1
	_uint8      uint8   = 1
	_uint16     uint16  = 1
	_uint32     uint32  = 1
	_uint64     uint64  = 1
	_byte       byte    = 1
	_rune       rune    = 1
	_complex64  complex64
	_complex128 complex128
	_error      error
	// _uintptr     uintptr
)

func slices() {
	numbers := []int{1, 2, 3, 4, 5}
	otherNumbers := make([]int, 2)
	numbers = append(numbers, 6, 7, 8, 9, 10)
	fmt.Println(numbers)
	fmt.Println(otherNumbers)
}
