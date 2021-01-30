package crypto

import (
	"log"
	"testing"
)

func TestAESCrypto(t *testing.T) {
	crypter := &AESCrypto{
		key: []byte("acim12320099qw21"),
	}
	var s = "我呃呃呃呃呃呃呃呃呃呃呃呃呃恶心死了"
	cipherText, err := crypter.EncodeCBC([]byte(s))
	if err != nil {
		log.Println(err)
	}
	log.Println(len(cipherText))
	plainText, err := crypter.DecodeCBC(cipherText)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(plainText))
	if string(plainText) != s {
		log.Fatal("NO EQUAL")
	}
}