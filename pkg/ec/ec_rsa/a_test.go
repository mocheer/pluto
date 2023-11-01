package ec_rsa_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/ec/ec_rsa"
	"github.com/stretchr/testify/assert"
)

// ecc 加密技术测试
func TestRsa(t *testing.T) {
	// 生成私匙和公匙
	ec_rsa.GenPemFiles("testdata", 2048)

	data := "hello world"
	pubKey, _ := ec_rsa.PublicKeyFromFile("testdata/public.pem") // 解密公匙
	encryData, _ := ec_rsa.Encrypt([]byte(data), pubKey)         // 加密数据

	//
	priKey, _ := ec_rsa.PrivateKeyFromFile("testdata/private.pem") // 解密私匙
	decryData, _ := ec_rsa.Decrypt(encryData, priKey)              // 解密数据
	//
	assert.Equal(t, data, string(decryData))

}

// DecodeCipherRSA 解析RSA密文
func TestDecodeCipherRSA(t *testing.T) {
	data := ""      //待解密
	key := []byte{} //密钥
	plainText := ec_rsa.JSEncrypt(data, key)
	t.Log(plainText)
}
