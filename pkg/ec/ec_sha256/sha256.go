package ec_sha256

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncryptToString(data []byte) string {
	r := sha256.New()
	r.Write(data)
	return hex.EncodeToString(r.Sum([]byte("")))
}
