package ec_test

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/mocheer/pluto/pkg/ec"
	"github.com/stretchr/testify/assert"
)

// ecc 加密技术测试
func TestRsa(t *testing.T) {
	// 生成私匙和公匙
	ec.RSA_GenPemFiles("testdata", 2048)

	data := "hello world"
	pubKey, _ := ec.RSA_PublicKeyFromFile("testdata/public.pem") // 解密公匙
	encryData, _ := ec.RSA_Encrypt([]byte(data), pubKey)         // 加密数据

	//
	priKey, _ := ec.RSA_PrivateKeyFromFile("testdata/private.pem") // 解密私匙
	decryData, _ := ec.RSA_Decrypt(encryData, priKey)              // 解密数据
	//
	assert.Equal(t, data, string(decryData))

}

func TestAes(t *testing.T) {
	origData := []byte("Hello World") // 待加密的数据
	key := []byte("ABCDEFGHIJKLMNOP") // 加密的密钥，必须是16、24、32
	t.Log("原文：", string(origData))

	t.Log("------------------ CBC模式 --------------------")
	encrypted := ec.AesEncryptCBC(origData, key)
	t.Log("密文(hex)：", hex.EncodeToString(encrypted))
	t.Log("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := ec.AesDecryptCBC(encrypted, key)
	t.Log("解密结果：", string(decrypted))

	// 电码本模式
	t.Log("------------------ ECB模式 --------------------")
	encrypted = ec.AesEncryptECB(origData, key)
	t.Log("密文(hex)：", hex.EncodeToString(encrypted))
	t.Log("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = ec.AesDecryptECB(encrypted, key)
	t.Log("解密结果：", string(decrypted))

	// 密码反馈模式
	t.Log("------------------ CFB模式 --------------------")
	encrypted = ec.AesEncryptCFB(origData, key)
	t.Log("密文(hex)：", hex.EncodeToString(encrypted))
	t.Log("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = ec.AesDecryptCFB(encrypted, key)
	t.Log("解密结果：", string(decrypted))
}

// DecodeCipherRSA 解析RSA密文
func TestDecodeCipherRSA(t *testing.T) {
	data := ""      //待解密
	key := []byte{} //密钥
	plainText := ec.RSA_JSEncrypt(data, key)
	t.Log(plainText)
}
