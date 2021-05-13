package ecc

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io"
	"os"
	"path"

	"github.com/mocheer/pluto/fs"
)

// RSA_DecodeJSEncrypt 用于用户名、密码解密
func RSA_DecodeJSEncrypt(data string, privatePemPath string) string {
	// JSEncrypt 生成的编码本身会再加上base64编码
	b, _ := base64.StdEncoding.DecodeString(data)
	//
	plain, err := RSA_DecryptByPem(b, privatePemPath) // 解密私匙
	if err != nil {
		panic(nil)
	}
	return string(plain)
}

// RSA_GenPemFiles
func RSA_GenPemFiles(dir string, bits int) error {
	privatePemPath := path.Join(dir, "private.pem")
	publicPemPath := path.Join(dir, "public.pem")
	pri := fs.MustCreate(privatePemPath)
	pub := fs.MustCreate(publicPemPath)
	defer pri.Close()
	defer pub.Close()
	return RSA_GenPems(pri, pub, bits)
}

func RSA_GenKeys(bits int) (privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, err error) {
	// 生成私匙，提供一个随机数和私匙的长度，目前主流的长度为1024、2048、3072、4096，
	// 但1024已经不在推荐使用了。
	privateKey, err = rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	publicKey = &privateKey.PublicKey
	return
}

// RSA_GenPems 产生公钥私钥对应的pem文件
func RSA_GenPems(privateKeyWriter, publicKeyWriter io.Writer, bits int) error {
	priKey, pubKey, err := RSA_GenKeys(bits)
	if err != nil {
		return err
	}
	// 将rsa私钥序列化为ASN.1 PKCS#1 DER编码
	data := x509.MarshalPKCS1PrivateKey(priKey)
	block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: data,
	}
	// 将私匙做pem数据编码，然后写入文件
	err = pem.Encode(privateKeyWriter, &block)
	if err != nil {
		return err
	}

	// 将rsa公钥序列化为ASN.1 PKCS#1 DER编码
	// pubKeyData := x509.MarshalPKCS1PublicKey(&pubKey)
	// 这里改成这种算法，跟 JSEncrypt 配合
	pubKeyData, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return err
	}
	// 将公匙做pem数据编码，然后写入文件
	err = pem.Encode(publicKeyWriter, &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyData,
	})
	if err != nil {
		return err
	}

	return nil
}

// 解析公匙
func RSA_PublicKeyFromFile(file string) (*rsa.PublicKey, error) {
	// 读取公匙文件
	pubByte, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	// pem解码
	b, _ := pem.Decode(pubByte)
	if b == nil {
		return nil, errors.New("error public key")
	}
	// der解码，最终返回一个公匙对象
	// pubKey, err := x509.ParsePKCS1PublicKey(b.Bytes)
	pubKey, err := x509.ParsePKIXPublicKey(b.Bytes)
	if err != nil {
		return nil, err
	}
	return pubKey.(*rsa.PublicKey), nil
}

// 解析私匙
func RSA_PrivateKeyFromFile(file string) (*rsa.PrivateKey, error) {
	// 读取私匙
	priByte, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	// pem解码
	block, _ := pem.Decode(priByte)
	if block == nil {
		return nil, errors.New("error private key")
	}
	// der加密，返回一个私匙对象
	prikey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return prikey, nil
}

func RSA_DecryptByPem(src []byte, file string) ([]byte, error) {
	//
	privateKey, err := RSA_PrivateKeyFromFile(file) // 解密私匙
	if err != nil {
		return nil, err
	}
	// 解密
	res, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//  rsa公匙加密
func RSA_Encrypt(src []byte, publickey *rsa.PublicKey) ([]byte, error) {
	// 使用公匙加密数据，需要一个随机数生成器和公匙和需要加密的数据
	data, err := rsa.EncryptPKCS1v15(rand.Reader, publickey, src)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// rsa私匙解密
func RSA_Decrypt(src []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	// 使用私匙解密数据，需要一个随机数生成器和私匙和需要解密的数据
	data, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)
	if err != nil {
		return nil, err
	}
	return data, nil
}
