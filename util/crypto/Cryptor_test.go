package crypto

import (
	"log"
	"testing"
)


func TestNewCryptor(t *testing.T) {
	c := NewCryptor()
	test := "我是啦啦啦啦啦啦时代峰峻未婚夫dfsdf带防辐射服ewfwfwegggggggg"
	cipherText, err := EncodePub([]byte(test), c.GetPubKeyByte())
	if err != nil {
		log.Println(err)
	}
	plainText, err := c.Decode(cipherText)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(plainText))
	if string(plainText) != test {
		t.Fatal("decoding failed")
	}
}