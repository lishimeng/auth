package utils

//
//import (
//	"crypto/rand"
//	"crypto/rsa"
//	"crypto/sha256"
//	"testing"
//)
//
//func TestGenerateRSAKey3(t *testing.T) {
//	privateKey, publicKey := GenerateRSAKey()
//
//	h := sha256.New() // sha1.New() or md5.New()
//	var plainText = "123456"
//
//	oaep, err := rsa.EncryptOAEP(h, rand.Reader, publicKey, []byte(plainText), nil)
//	t.Logf("%s", oaep)
//
//	if err != nil {
//		t.Logf("%s", err)
//	}
//	h = sha256.New()
//
//	var bytes []byte
//	bytes, err = rsa.DecryptOAEP(h, rand.Reader, privateKey, oaep, nil)
//
//	t.Logf("%s", bytes)
//}
//
//func TestGenerateRSAKey(t *testing.T) {
//	privateKey, _ := GenerateRSAKey()
//	privPEM := ExportPrivateKeyAsPEM(privateKey)
//	privKeyEncodeString := Base64EncodeString(privPEM)
//	privKeyDecodeString, _ := Base64DecodeString(privKeyEncodeString)
//
//	// 通过密钥获取公钥
//	pubKey := GetPublicKeyFromPriKey(privKeyDecodeString)
//
//	// 公钥加密
//	painText := "hello world"
//	result, _ := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(painText))
//	t.Logf("%s", string(result))
//
//	str, err := DecodingByPrivateKey(privKeyEncodeString, result)
//	if err != nil {
//		t.Logf("%s", err.Error())
//	}
//	t.Logf("%s", string(str))
//}
//
//func TestGenerateRSAKey2(t *testing.T) {
//	privateKey, publicKey := GenerateRSAKey()
//	privPEM := ExportPrivateKeyAsPEM(privateKey)
//	pubPEM := ExportPublicKeyAsPEM(publicKey)
//	t.Logf("%s", privPEM)
//	t.Logf("%s", pubPEM)
//
//	privKeyEncodeString := Base64EncodeString(privPEM)
//	privKeyDecodeString, err := Base64DecodeString(privKeyEncodeString)
//	if err != nil {
//		t.Logf("%s", err)
//	}
//
//	pubKey := GetPublicKeyFromPriKey(privKeyDecodeString)
//	pubPEM = ExportPublicKeyAsPEM(publicKey)
//	t.Logf("%s", pubPEM)
//
//	// 公钥加密
//	painText := "hello world"
//	result, _ := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(painText))
//	t.Logf("%s", string(result))
//
//	str, err := DecodingByPrivateKey(privKeyEncodeString, result)
//	if err != nil {
//		t.Logf("%s", err.Error())
//	}
//	t.Logf("%s", string(str))
//
//	//私钥解密
//	//result, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, result)
//	//if err != nil {
//	//	t.Logf("%s", err.Error())
//	//}
//	//t.Logf("%s", string(result))
//}

//
//func TestRSA(t *testing.T) {
//	// 生成 RSA 密钥对
//	privateKey, publicKey := generateRSAKey()
//
//	// 将公钥、私钥导出为 PEM 格式
//	pubPEM := exportPublicKeyAsPEM(publicKey)
//	privPEM := exportPrivateKeyAsPEM(privateKey)
//
//	// TODO 存储在redis
//	t.Logf("%s", string(pubPEM))
//	t.Logf("%s", string(privPEM))
//
//	// 加载公钥、私钥
//	privKey := loadPrivateKey(privPEM)
//	// 从密钥中直接获取公钥
//	//pubKey := &privKey.PublicKey
//	pubKey := getPublicKeyFromPriKey(privPEM)
//
//	pubPEM = exportPublicKeyAsPEM(pubKey)
//	t.Logf("%s", string(pubPEM))
//	//pubKey := loadPublicKey(pubPEM)
//
//	// 公钥加密
//	painText := "hello world"
//	result, _ := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(painText))
//	t.Logf("%s", string(result))
//
//	//私钥解密
//	result2, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, result)
//	if err != nil {
//		t.Logf("%s", string(result2))
//	}
//	t.Logf("%s", string(result2))
//}
//
//func TestBASE64(t *testing.T) {
//	// 生成 RSA 密钥对
//	privateKey, publicKey := generateRSAKey()
//
//	// 将公钥、私钥导出为 PEM 格式
//	pubPEM := exportPublicKeyAsPEM(publicKey)
//	privPEM := exportPrivateKeyAsPEM(privateKey)
//
//	t.Logf("%s", string(pubPEM))
//	t.Logf("%s", string(privPEM))
//
//	encodeString := best64EncodeString(pubPEM)
//	t.Logf("%s", encodeString)
//
//	pubKey, err := base64DecodeString(encodeString)
//	if err != nil {
//
//	}
//	t.Logf("%s", string(pubKey))
//}
