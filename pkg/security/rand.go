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
	// url safe 的 base64 支持的字符，64 个
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	// 6 bits 代表 64 个不同索引，所以 letterBytes 长度必须为 64
	letterIdxBits = 6
	// 先位移，然后减 1
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
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; i-- {
		if remain == 0 {
			// 生成的随机数用完了，重新生成
			cache, remain = src.Int63(), letterIdxMax
		}
		b[i] = letterBytes[cache&letterIdxMask]
		cache >>= letterIdxBits
		remain--
	}
	return conv.Bytes2string(b)
}
