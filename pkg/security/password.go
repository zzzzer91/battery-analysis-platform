package security

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"strconv"
	"strings"
)

const (
	defaultPbkdf2Iterations = 150000
)

//var (
//	ErrInvaildPasswordHash = errors.New("加密密码格式有误")
//	ErrPasswordDoesntMatch = errors.New("密码错误")
//)

// 生成加密后密码，源自 Python 的 werkzeug 库
func GeneratePasswordHash(password, method string, saltLength int) (string, error) {
	salt := GenerateSalt(saltLength)
	h, actualMethod, err := hashInternal(method, salt, password)
	return fmt.Sprintf("%s$%s$%s", actualMethod, salt, h), err
}

func CheckPasswordHash(pwhash, password string) error {
	if strings.Count(pwhash, "$") < 2 {
		return errors.New("加密密码格式有误")
	}
	// 把字符串分成三部分
	args := strings.SplitN(pwhash, "$", 3)
	tmp, _, err := hashInternal(args[0], args[1], password)
	if err != nil {
		return err
	}
	if args[2] != tmp {
		return errors.New("密码错误")
	}
	return nil
}

func hashInternal(method, salt, password string) (string, string, error) {
	if method == "plain" {
		return password, method, nil
	}
	if strings.HasPrefix(method, "pbkdf2:") {
		args := strings.Split(method[7:], ":")
		method = args[0]
		var iterations int
		// 注意：error 要在这里定义，不然下面执行 strconv.Atoi 函数时，
		// 使用 := 赋值符会导致 iterations 作用域在当前大括号中
		var err error
		if len(args) == 1 {
			iterations = defaultPbkdf2Iterations
		} else if len(args) == 2 {
			iterations, err = strconv.Atoi(args[1])
			if err != nil || iterations <= 0 {
				iterations = defaultPbkdf2Iterations
			}
		} else {
			return "", "", errors.New("invalid number of arguments for PBKDF2")
		}
		actualMethod := fmt.Sprintf("pbkdf2:%s:%d", method, iterations)
		rv := pbkdf2.Key([]byte(password), []byte(salt), iterations, 32, sha256.New)
		// 将 []byte 转换成十六进制小写字符串
		encodedStr := hex.EncodeToString(rv)
		return encodedStr, actualMethod, nil
	} else {
		return "", "", errors.New("don't support other method")
	}
}
