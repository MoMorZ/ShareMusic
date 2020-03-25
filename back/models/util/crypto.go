package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"math/big"
	mrand "math/rand"
	"strings"
)

const iv = "0102030405060708"
const presetKey = "0CoJUm6Qyw8W8jud"
const linuxapiKey = "rFgB&h#%2?^eDg:Q"
const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const publicKey = "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgtQn2JZ34ZC28NWYpAUd98iZ37BUrX/aKzmFbt7clFSs6sXqHauqKWqdtLkF2KexO40H1YTX8z2lSgBBOAxLsvaklV8k4cBFK9snQXE9/DDaFt6Rr7iVZMldczhC0JNgTz+SHXT6CBHuX3e9SdB1Ua44oncaTWz7OBGLbCiK45wIDAQAB\n-----END PUBLIC KEY-----"
const eapiKey = "e82ckenh8dichen8"

func WeApi(object map[string]interface{}) map[string]interface{} {
	text, _ := json.Marshal(object)
	secretKey := createSecretKey(16)
	params := base64.StdEncoding.EncodeToString(
		aesEncrypt([]byte(
			base64.StdEncoding.EncodeToString(
				aesEncrypt(text, "cbc", []byte(presetKey), []byte(iv)))), "cbc", secretKey, []byte(iv)))
	encSecKey := hex.EncodeToString(rsaEncrypt(reverse(secretKey), []byte(publicKey)))
	return map[string]interface{}{
		"params":    params,
		"encSecKey": encSecKey,
	}
}

func LinuxApi(object map[string]interface{}) map[string]interface{} {
	text, _ := json.Marshal(object)
	eparams := strings.ToUpper(hex.EncodeToString(aesEncrypt(text, "ecb", []byte(linuxapiKey), []byte{})))
	return map[string]interface{}{
		"eparams": eparams,
	}
}

// 翻转字符串
func reverse(s []byte) []byte {
	n := len(s)
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		ret[n-i-1] = s[i]
	}
	return ret
}

// 生成一个大小为size的随机密钥
// size应默认为16
func createSecretKey(size int) []byte {
	key := make([]byte, size)
	for i := 0; i < size; i++ {
		num := mrand.Intn(62)
		key[i] = base62[num]
	}
	return key
}

// aes 加密
func aesEncrypt(buffer []byte, mode string, key []byte, iv []byte) []byte {
	if mode == "cbc" {
		return AesEncryptCBC(buffer, key, iv)
	} else if mode == "ecb" {
		return AesEncryptECB(buffer, key)
	}
	return []byte{}
}

// rsa 加密
func rsaEncrypt(buffer []byte, key []byte) []byte {
	data := make([]byte, 128-len(buffer))
	buffer = append(data, buffer...)
	ret, _ := RsaEncrypt(buffer, key)
	return ret
}

// 下面的 aes和rsa 模块
// 主要参考自
// http://www.361way.com/golang-rsa-aes/5828.html
// https://blog.csdn.net/mirage003/article/details/87868999

// aes 模块

// =================== CBC ======================
func AesEncryptCBC(origData []byte, key []byte, iv []byte) (encrypted []byte) {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                 // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)   // 补全码
	blockMode := cipher.NewCBCEncrypter(block, iv) // 加密模式
	encrypted = make([]byte, len(origData))        // 创建数组
	blockMode.CryptBlocks(encrypted, origData)     // 加密
	return encrypted
}
func AesDecryptCBC(encrypted []byte, key []byte, iv []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key) // 分组秘钥
	//blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, iv) // 加密模式
	decrypted = make([]byte, len(encrypted))       // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)    // 解密
	decrypted = pkcs5UnPadding(decrypted)          // 去除补全码
	return decrypted
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// =================== ECB ======================
func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}
func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// rsa 模块

// 加密
func RsaEncrypt(origData []byte, key []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	encrypted := new(big.Int)
	e := big.NewInt(int64(pub.E))
	payload := new(big.Int).SetBytes(origData)
	encrypted.Exp(payload, e, pub.N)
	return encrypted.Bytes(), nil
}
