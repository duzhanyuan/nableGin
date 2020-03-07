package helper

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"strings"
)


//hash加密
func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// hash加密对比
func CompareHashAndPassword(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// 获取MD5
func MD5(str string) string {
	_md5 := md5.New()
	_md5.Write([]byte(str))
	return hex.EncodeToString(_md5.Sum([]byte(nil)))
}

// 获取SHA1
func SHA1(str string) string {
	_sha1 := sha1.New()
	_sha1.Write([]byte(str))
	return hex.EncodeToString(_sha1.Sum([]byte(nil)))
}


// 获取HMAC
func HMAC(key, data string) string {
	_hmac := hmac.New(md5.New, []byte(key))
	_hmac.Write([]byte(data))
	return hex.EncodeToString(_hmac.Sum([]byte(nil)))
}

// 合并字符串
func StrJoin(sep string, e... string) string {
	return strings.Join(e, sep)
}

// 去掉左右两边空格字符串
func Strim(s string) string {
	return strings.Trim(s, " ")
}

// 判断值是否为空
func IsEmpty(v interface{}) bool {
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

//字符串截取
func Substr(start, end int, s string) string {
	if start < 0 {
		return s[:end]
	}
	if end < 0 || end > len(s) {
		return s[start:]
	}
	return s[start:end]
}
