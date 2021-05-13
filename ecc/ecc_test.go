package ecc_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/mocheer/pluto/ecc"
)

// ecc 加密技术测试
func TestRsa(t *testing.T) {
	// 生成私匙和公匙
	pri, _ := os.Create("private.pem")
	pub, _ := os.Create("public.pem")
	defer pri.Close()
	defer pub.Close()
	ecc.RSA_GenPems(pri, pub, 2048)

	pubKey, _ := ecc.RSA_PublicKeyFromFile("public.pem") // 解密公匙
	fmt.Printf("1 => %v\n", pubKey)
	encryData, _ := ecc.RSA_Encrypt([]byte("hello world"), pubKey) // 加密数据
	fmt.Printf("2 => %x\n", encryData)

	priKey, _ := ecc.RSA_PrivateKeyFromFile("private.pem") // 解密私匙
	decryData, _ := ecc.RSA_Decrypt(encryData, priKey)     // 解密数据
	fmt.Printf("3 => %s\n", decryData)

	t.Error("4")
}
