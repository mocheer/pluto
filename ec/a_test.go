package ec_test

import (
	"os"
	"testing"

	"github.com/mocheer/pluto/assert"
	"github.com/mocheer/pluto/ec"
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
