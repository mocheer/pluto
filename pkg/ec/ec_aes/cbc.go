package ec_aes

import (
	"crypto/aes"
	"crypto/cipher"
)

// =================== CBC ======================
// EncryptCBC aes-密码分组链接模式
func EncryptCBC(origData []byte, key []byte) (encrypted []byte) {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                 // 获取秘钥块的长度
	iv := key[:blockSize]                          // 这里直接将key作为向量，IV是一个随机生成的值，或者由应用程序提供，以确保加密结果不同于在同一密钥下加密相同数据块的其他结果。
	origData = PKCS7Padding(origData, blockSize)   // 补全码
	blockMode := cipher.NewCBCEncrypter(block, iv) // 加密模式
	encrypted = make([]byte, len(origData))        // 创建数组
	blockMode.CryptBlocks(encrypted, origData)     // 加密
	return encrypted
}

// DecryptCBC
func DecryptCBC(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)                 // 分组秘钥
	blockSize := block.BlockSize()                 // 获取秘钥块的长度
	iv := key[:blockSize]                          // 这里直接将key作为向量，IV是一个随机生成的值，或者由应用程序提供，以确保加密结果不同于在同一密钥下加密相同数据块的其他结果。
	blockMode := cipher.NewCBCDecrypter(block, iv) // 加密模式
	decrypted = make([]byte, len(encrypted))       // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)    // 解密
	decrypted = PKCS7UnPadding(decrypted)          // 去除补全码
	return decrypted
}
