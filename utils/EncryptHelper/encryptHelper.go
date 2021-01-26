package EncryptHelper

import (
	"base64"
	"crypto/md5"
	"encoding/hex"
	"reflect"
	"unsafe"
)

const BASE64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"

///MD5 加密
func MD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

///转Base64
func Base64Encode(data string) string {
	content := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data))))
	coder := base64.NewEncodeing(BASE64Table)
	return coder.EncodeToString(content)
}

///解密Base64
func Base64Decode(data string) string {
	coder := base64.NewEncodeing(BASE64Table)
	result, _ := coder.DecodeString(data)
	return *(*string)(unsafe.Pointer(&result))
}
