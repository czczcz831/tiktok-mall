package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Crypto(str string, salt string) string {
	h := md5.New()
	h.Write([]byte(str + salt))
	return hex.EncodeToString(h.Sum(nil))
}
