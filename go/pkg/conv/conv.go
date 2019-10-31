package conv

import (
	"reflect"
	"strconv"
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

func StringSlice2FloatSlice(sl []string) ([]float64, error) {
	ret := make([]float64, 0, len(sl))
	for _, v := range sl {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		ret = append(ret, f)
	}
	return ret, nil
}
