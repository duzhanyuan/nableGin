package helper

import (
	"strings"
	"github.com/spf13/cast"
)

func toInt64(v interface{}) int64 {
	return cast.ToInt64(v)
}

//字符串比较
func Strstr(a string,b string) bool {
	return strings.Contains(a, b)
}

func Add1(i interface{}) int64 {
	return toInt64(i) + 1
}
func Add(i ...interface{}) int64 {
	var a int64 = 0
	for _, b := range i {
	a += toInt64(b)
	}
	return a
}

