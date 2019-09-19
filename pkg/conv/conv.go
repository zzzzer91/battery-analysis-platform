package conv

import (
	"reflect"
	"unsafe"
)

func String2bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
