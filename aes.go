package gorest

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

func test() {
	key := "a very very very" // 16 bytes
	plaintext := "some really really really long plaintext"
	fmt.Println(plaintext)

	ciphertext, err := EncryptStr(key, plaintext)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ciphertext)

	result, err := DecryptStr(key, ciphertext)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func EncryptStr(key, text string) (string, error) {
	cipherBin, err := Encrypt([]byte(key), []byte(text))
	return hex.EncodeToString(cipherBin), err
}
func DecryptStr(key, text string) (string, error) {
	ciphyBin, err := hex.DecodeString(text)
	if err != nil {
		return "", err
	}
	plainBin, err := Decrypt([]byte(key), ciphyBin)
	return string(plainBin), err

}

func Encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func Decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
