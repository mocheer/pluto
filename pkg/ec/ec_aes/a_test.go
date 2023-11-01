package ec_aes_test

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/mocheer/pluto/pkg/ec/ec_aes"
)

func TestAes(t *testing.T) {
	origData := []byte("Hello World") // 待加密的数据
	key := []byte("ABCDEFGHIJKLMNOP") // 加密的密钥，必须是16、24、32
	t.Log("原文：", string(origData))

	t.Log("------------------ CBC模式 --------------------")
	encrypted := ec_aes.EncryptCBC(origData, key)
	t.Log("密文(hex)：", hex.EncodeToString(encrypted))
	t.Log("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := ec_aes.DecryptCBC(encrypted, key)
	t.Log("解密结果：", string(decrypted))

	// 电码本模式
	t.Log("------------------ ECB模式 --------------------")
	encrypted = ec_aes.EncryptECB(origData, key)
	t.Log("密文(hex)：", hex.EncodeToString(encrypted))
	t.Log("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = ec_aes.DecryptECB(encrypted, key)
	t.Log("解密结果：", string(decrypted))

	// 密码反馈模式
	t.Log("------------------ CFB模式 --------------------")
	encrypted = ec_aes.EncryptCFB(origData, key)
	t.Log("密文(hex)：", hex.EncodeToString(encrypted))
	t.Log("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = ec_aes.DecryptCFB(encrypted, key)
	t.Log("解密结果：", string(decrypted))
}
