package sherbet

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// PKCS7Padding encrypt bytes
func PKCS7Padding(text []byte, size int) []byte {
	var padding = size - len(text)%size
	var padtext = bytes.Repeat([]byte{byte(padding)}, padding)

	return append(text, padtext...)
}

// PKCS7UnPadding decrypt bytes
func PKCS7UnPadding(data []byte) []byte {
	var length = len(data)
	var unpadding = int(data[length-1])

	return data[:(length - unpadding)]
}

// AesEncrypt encrypt data
func AesEncrypt(data, key []byte) ([]byte, error) {
	var crypted []byte
	var err error

	if block, e := aes.NewCipher(key); err == nil {
		var size = block.BlockSize()
		var mode = cipher.NewCBCEncrypter(block, key[:size])
		data = PKCS7Padding(data, size)
		crypted = make([]byte, len(data))
		mode.CryptBlocks(crypted, data)
	} else {
		err = e
	}

	return crypted, err
}

// AesDecrypt decrypt data
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	var data []byte
	var err error

	if block, e := aes.NewCipher(key); err == nil {
		var size = block.BlockSize()
		var mode = cipher.NewCBCDecrypter(block, key[:size])

		data = make([]byte, len(crypted))
		mode.CryptBlocks(data, crypted)
		data = PKCS7UnPadding(data)
	} else {
		err = e
	}

	return data, err
}
