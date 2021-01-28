package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"log"
)

type Cryptor struct {
	pubKeyByte []byte
	priKey *rsa.PrivateKey
}

func NewCryptor() *Cryptor {
	tmp := &Cryptor{}
	var (
		err error
		priKeyByte []byte
	)
	priKeyByte, tmp.pubKeyByte, err = NewRsaKey()
	if err != nil {
		log.Println(err)
	}
	tmp.priKey, err = x509.ParsePKCS1PrivateKey(priKeyByte)
	if err != nil {
		log.Println(err)
	}
	return tmp
}

func (c *Cryptor) GetPubKeyByte() []byte {
	if c == nil {
		return nil
	}
	return c.pubKeyByte
}

func (c *Cryptor) Decode(cipherText []byte) ([]byte, error) {
	var (
		plainText []byte
		err error
	)
	plainText, err = rsa.DecryptPKCS1v15(rand.Reader, c.priKey, cipherText)
	return plainText, err
}