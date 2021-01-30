package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// len(key) must be 16 or 32.
type AESCrypto struct {
	key []byte
}

func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText) % blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

func (A *AESCrypto) EncodeCBC(plainText []byte) (cipherText []byte, err error) {
	block, err := aes.NewCipher(A.key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	originData := PKCS5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, A.key[:blockSize])
	cipherText = make([]byte, len(originData))
	blockMode.CryptBlocks(cipherText, originData)
	return cipherText, nil
}

func (A *AESCrypto) DecodeCBC(cipherText []byte) (plainText []byte, err error) {
	block, err := aes.NewCipher(A.key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, A.key[:blockSize])
	plainText = make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	plainText = PKCS5UnPadding(plainText)
	return plainText, nil
}