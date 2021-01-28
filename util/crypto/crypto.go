package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

func NewRsaKey() (priKeyByte []byte, pubKeyByte []byte, err error){
	var priKey *rsa.PrivateKey
	priKey, err = rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, nil, err
	}
	priKeyByte = x509.MarshalPKCS1PrivateKey(priKey)
	pubKey := priKey.PublicKey
	pubKeyByte, err = x509.MarshalPKIXPublicKey(&pubKey)
	if err != nil {
		return nil, nil, err
	}
	return
}

func EncodePub(plainText []byte, pubKeyByte []byte) (cipherText []byte, err error) {
	pubKeyInterface, err := x509.ParsePKIXPublicKey(pubKeyByte)
	if err != nil {
		return nil, err
	}
	pubKey := pubKeyInterface.(*rsa.PublicKey)
	cipherText, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey, plainText)
	return
}