package ec

import (
	"crypto/md5"
	"encoding/hex"
)

// EncryptToString 获取data的md5值
func EncryptToString(data []byte) string {
	hash := md5.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}
