package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encrypter struct {
	Key []byte
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("KEY is not set")
	}
	return &Encrypter{
		Key: []byte(key),
	}
}

func (enc *Encrypter) Encrypt(text []byte) []byte {
	block, err := aes.NewCipher(enc.Key)
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	return aesGCM.Seal(nonce, nonce, text, nil)
}

func (enc *Encrypter) Decrypt(text []byte) []byte {
	block, err := aes.NewCipher(enc.Key)
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce, cipherText := text[:aesGCM.NonceSize()], text[aesGCM.NonceSize():]
	plainText, err := aesGCM.Open(nil, []byte(nonce), []byte(cipherText), nil)
	if err != nil {
		panic(err.Error())
	}
	return plainText
}
