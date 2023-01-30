package util

import (
	"bytes"
	"encoding/json"
	"reflect"
	"unsafe"
)

func MustMarshal(t interface{}) string {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(t); err != nil {
		return ""
	} else {
		b := buffer.Bytes()
		// trim new line
		return ByteSliceToString(b[:len(b)-1])
	}
}

func MustMarshalIndent(t interface{}) string {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(t); err != nil {
		return ""
	} else {
		b := buffer.Bytes()
		// trim new line
		return ByteSliceToString(b[:len(b)-1])
	}
}

// ByteSliceToString converts a byte slice to a string.
//
// This is a shallow copy, means that the returned string reuse the
// underlying array in byte slice, it's your responsibility to keep
// the input byte slice survive until you don't access the string anymore.
//
func ByteSliceToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

// StringToByteSlice converts a string to a byte slice.
//
// This is a shallow copy, means that the returned byte slice reuse
// the underlying array in string, so you can't change the returned
// byte slice in any situations.
//
func StringToByteSlice(s string) []byte {
	var bs []byte
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return bs
}
