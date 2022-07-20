package ec_test

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"os"
	"testing"

	"github.com/mocheer/pluto/pkg/ec"
	"github.com/stretchr/testify/assert"
)

// ecc 加密技术测试
func TestRsa(t *testing.T) {
	// 生成私匙和公匙
	ec.RSA_GenPemFiles("test", 2048)

	data := "hello world"
	pubKey, _ := ec.RSA_PublicKeyFromFile("test/public.pem") // 解密公匙
	encryData, _ := ec.RSA_Encrypt([]byte(data), pubKey)     // 加密数据

	//
	priKey, _ := ec.RSA_PrivateKeyFromFile("test/private.pem") // 解密私匙
	decryData, _ := ec.RSA_Decrypt(encryData, priKey)          // 解密数据
	//
	os.RemoveAll("test")

	assert.Equal(t, data, string(decryData))

}

func TestAes(t *testing.T) {
	origData := []byte("Hello World") // 待加密的数据
	key := []byte("ABCDEFGHIJKLMNOP") // 加密的密钥，必须是16、24、32
	log.Println("原文：", string(origData))

	log.Println("------------------ CBC模式 --------------------")
	encrypted := ec.AesEncryptCBC(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := ec.AesDecryptCBC(encrypted, key)
	log.Println("解密结果：", string(decrypted))

	// 电码本模式
	log.Println("------------------ ECB模式 --------------------")
	encrypted = ec.AesEncryptECB(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = ec.AesDecryptECB(encrypted, key)
	log.Println("解密结果：", string(decrypted))

	// 密码反馈模式
	log.Println("------------------ CFB模式 --------------------")
	encrypted = ec.AesEncryptCFB(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = ec.AesDecryptCFB(encrypted, key)
	log.Println("解密结果：", string(decrypted))
}
