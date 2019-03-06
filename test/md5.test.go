package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte("guzhongren" + "000000"))
	cipherStr := md5Ctx.Sum(nil)
	fmt.Print(cipherStr)
	fmt.Print("\n")
	fmt.Print(hex.EncodeToString(cipherStr))
}
