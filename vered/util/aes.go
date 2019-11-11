package util

import (
	"crypto/aes"
	"crypto/sha256"
	"encoding/base64"
	"errors"

	"github.com/afterpay/sdk/vered/util/ebc"
)

func fixLengthKeyHash(bytes []byte) []byte {
	h := sha256.New()
	h.Write(bytes)
	return h.Sum(nil)
}

func FixAesKeyLength(data []byte) []byte {
	key := fixLengthKeyHash(data)
	return key[:aes.BlockSize]
}

func AESEncrypt(key []byte, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plaintext = PKCS5Padding(plaintext, blockSize)

	ciphertext := make([]byte, len(plaintext))
	blockMode := ebc.NewECBEncrypter(block)
	blockMode.CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

// AESEncrypt encrypts string to base64 crypto using AES
func AESEncryptBase64(key []byte, plaintext []byte) (string, error) {
	ciphertext, err := AESEncrypt(key, plaintext)
	if err != nil {
		return "", err
	}
	// convert to base64
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AESDecrypt from base64 to decrypted string
func AESDecryptBase64(key []byte, cryptoText string) ([]byte, error) {
	ciphertext, _ := base64.StdEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	if len(ciphertext) < blockSize {
		return nil, errors.New("ciphertext too short")
	}
	origData := make([]byte, len(ciphertext))
	blockMode := ebc.NewECBDecrypter(block)
	blockMode.CryptBlocks(origData, ciphertext)
	origData = PKCS5UnPadding(origData)

	return origData, nil
}
