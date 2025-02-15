package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func MD5Crypto(str string, salt string) string {
	fmt.Println("md5" + str + salt)
	h := md5.New()
	h.Write([]byte(str + salt))
	return hex.EncodeToString(h.Sum(nil))
}
