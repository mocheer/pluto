package ec_sha1

import (
	"crypto/sha1"
	"encoding/hex"
)

func EncryptToString(data []byte) string {
	r := sha1.New()
	r.Write(data)
	return hex.EncodeToString(r.Sum([]byte("")))
}
