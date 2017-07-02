package week2

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

// UnPKCS7Padding undo the PKCS7 Padding
func UnPKCS7Padding(data []byte) []byte {
	length := len(data)
	padding := int(data[length-1])
	return data[:(length - padding)]
}

// CBCDecrypt calls built-in CBC-mode decryption on keyString and cipher text string
func CBCDecrypt(keyStr string, ciphertextStr string) []byte {
	key, _ := hex.DecodeString(keyStr)
	ciphertext, _ := hex.DecodeString(ciphertextStr)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(ciphertext, ciphertext)

	return UnPKCS7Padding(ciphertext)
}

// CTRDecrypt calls built-in CTR-mode decryption on keyString and cipher text string
func CTRDecrypt(keyStr string, ciphertextStr string) []byte {
	key, _ := hex.DecodeString(keyStr)
	ciphertext, _ := hex.DecodeString(ciphertextStr)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCTR(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext
}
