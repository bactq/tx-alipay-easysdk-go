package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var iv [aes.BlockSize]byte

func Encrypt(plainText, key []byte) ([]byte, error) {
	// key base64 decode
	key, err := base64.StdEncoding.DecodeString(string(key))
	if err != nil {
		return nil, err
	}
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plainText = aesPadding(plainText)
	dest := make([]byte, len(plainText))
	cipher.NewCBCEncrypter(aesBlock, iv[:]).CryptBlocks(dest, plainText)
	// base64 encode
	return []byte(base64.StdEncoding.EncodeToString(dest)), nil
}

func Decrypt(cipherText, key []byte) ([]byte, error) {
	// key base64 decode
	key, err := base64.StdEncoding.DecodeString(string(key))
	if err != nil {
		return nil, err
	}
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// base64 decode
	ds, err := base64.StdEncoding.DecodeString(string(cipherText))
	if err != nil {
		return nil, err
	}
	cipherText = aesUnPadding(ds)
	dest := make([]byte, len(cipherText))
	cipher.NewCBCDecrypter(aesBlock, iv[:]).CryptBlocks(dest, cipherText)
	return dest, nil
}

// PKCS5Padding
func aesPadding(data []byte) []byte {
	dl := len(data)
	pl := aes.BlockSize - (dl % aes.BlockSize)
	for i := dl; i < dl+pl; i++ {
		data = append(data, byte(pl))
	}
	return data
}

// un PKCS5Padding
func aesUnPadding(data []byte) []byte {
	dl := len(data)
	lastByte := data[dl-1]
	padValue := int(lastByte & 0x0ff)
	if padValue < 0x01 || padValue > aes.BlockSize {
		return data
	}
	for i := dl - padValue; i < dl; i++ {
		if data[i] != lastByte {
			return data
		}
	}
	return data[:dl-padValue]
}
