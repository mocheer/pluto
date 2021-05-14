package ecc_test

import (
	"os"
	"testing"

	"github.com/mocheer/pluto/ecc"
)

// ecc 加密技术测试
func TestRsa(t *testing.T) {
	// 生成私匙和公匙
	ecc.RSA_GenPemFiles("test", 2048)

	data := "hello world"
	pubKey, _ := ecc.RSA_PublicKeyFromFile("test/public.pem") // 解密公匙
	encryData, _ := ecc.RSA_Encrypt([]byte(data), pubKey)     // 加密数据

	priKey, _ := ecc.RSA_PrivateKeyFromFile("test/private.pem") // 解密私匙
	decryData, _ := ecc.RSA_Decrypt(encryData, priKey)          // 解密数据
	if data != string(decryData) {
		t.Error("错误", data, "!=", decryData)
	}
	os.RemoveAll("test")
}
