package week2

import (
	"crypto/aes"
	"encoding/hex"
	"crypto/cipher"
	"fmt"
)

func CBC_decrypt (keyStr string, ciphertextStr string) {
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

	fmt.Printf("%s\n", ciphertext)
}

func CTR_decrypt (keyStr string, ciphertextStr string) {
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

	fmt.Printf("%s\n", ciphertext)
}