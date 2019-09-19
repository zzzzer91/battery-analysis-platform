// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
// 基本思想：
// 生成的随机数有 63 个 bit（还有一个符号位），
// 但我们需要的随机数范围很小，
// 如果只用其中 6 bit 来代表一个随机数，
// 就减少了很多次的随机数生成

package security

import (
	"battery-analysis-platform/pkg/conv"
	"math/rand"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// 6 bits 最多代表 64 个不同索引，所以 letterBytes 长度不能大于 64
	letterIdxBits = 6
	// All 1-bits, as many as letterIdxBits
	letterIdxMask = 1<<letterIdxBits - 1
	// 生成一个 int64,有 63 个 有效 bit
	letterIdxMax = 63 / letterIdxBits
)

var src = rand.NewSource(time.Now().UnixNano())

func GenerateRandomString(n int) string {
	if n <= 0 {
		panic("String length must be positive")
	}
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		// 生成的随机数用完了，重新生成
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		// 因为生成的 idx 大小范围是 0 ～ 63，而 letterBytes 长度可能小于 63，
		// 所以大于这个长度的要忽略
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return conv.Bytes2string(b)
}
