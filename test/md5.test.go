package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	// 加密字符串
	fmt.Println(GetGUID())
}

func CryptoStr(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte("guzhongren" + str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// 获取uuid
func GetGUID() string {
	return CryptoStr(uuid.NewV4().String())

}
