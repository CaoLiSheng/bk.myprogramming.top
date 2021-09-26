package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
	"unicode"

	"github.com/mozillazg/go-pinyin"
)

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func Pinyin(str string) string {
	args := pinyin.NewArgs()
	args.Style = pinyin.Tone
	py := make([]string, len(str))
	for i, r := range str {
		if unicode.Is(unicode.Han, r) {
			py[i] = pinyin.SinglePinyin(r, args)[0]
		} else {
			py[i] = string(r)
		}
	}
	return strings.Join(py, "")
}
