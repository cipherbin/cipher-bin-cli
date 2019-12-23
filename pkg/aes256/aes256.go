// Package aes256 is thanks to https://gist.github.com/brettscott/2ac58ab7cb1c66e2b4a32d6c1c3908a7
// Will play nicely with the aes256 encryption created on the client (javascript) side
package aes256

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	"github.com/mergermarket/go-pkcs7"
)

// Encrypt takes a message and a key and encrypts the message with the
// key using the same AES-256 algorithm used on the client side
func Encrypt(message []byte, key string) (string, error) {
	k := []byte(key)

	plainText, err := pkcs7.Pad(message, aes.BlockSize)
	if err != nil {
		return "", fmt.Errorf(`plainText: "%s" has error`, plainText)
	}
	if len(plainText)%aes.BlockSize != 0 {
		err := fmt.Errorf(`plainText: "%s" has the wrong block size`, plainText)
		return "", err
	}

	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return fmt.Sprintf("%x", cipherText), nil
}

// Decrypt takes a message and a key and decrypts the message with the
// key using the same AES-256 algorithm used on the client side
func Decrypt(message string, key string) (string, error) {
	k := []byte(key)
	cipherText, _ := hex.DecodeString(message)

	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	if len(cipherText) < aes.BlockSize {
		return "", errors.New("cipherText too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	if len(cipherText)%aes.BlockSize != 0 {
		return "", errors.New("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, _ = pkcs7.Unpad(cipherText, aes.BlockSize)

	return fmt.Sprintf("%s", cipherText), nil
}
