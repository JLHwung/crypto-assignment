package week4

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"log"
	"net/http"
	"time"
	"sync"
)

func safeXORBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return n
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generatePKCS7Padding(paddingCount int) []byte {
	return bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
}

// AddPkcs7 pads given byte array using pkcs7 padding schema till it has blockSize length in bytes
func addPkcs7(data []byte, blockSize int) []byte {

	var paddingCount int

	if paddingCount = blockSize - (len(data) % blockSize); paddingCount == 0 {
		paddingCount = blockSize
	}

	return append(data, generatePKCS7Padding(paddingCount)...)
}

func GenerateCipherTextOnShortAES(plaintext, ciphertext, target string) string {
	cipherTextBytes, err := hex.DecodeString(ciphertext)
	check(err)
	plaintextBytes := []byte(plaintext)
	targetBytes := []byte(target)

	iv := cipherTextBytes[:aes.BlockSize]
	ciphertextWithoutIV := cipherTextBytes[aes.BlockSize:]
	newIv := make([]byte, aes.BlockSize)

	paddedPlaintext := addPkcs7(plaintextBytes, aes.BlockSize)
	paddedTarget := addPkcs7(targetBytes, aes.BlockSize)
	diff := make([]byte, len(paddedPlaintext))
	safeXORBytes(diff, paddedPlaintext, paddedTarget)
	safeXORBytes(newIv, iv, diff)
	var newIvBuf []byte = newIv[:]
	newCipherText := hex.EncodeToString(append(newIvBuf, ciphertextWithoutIV[:]...))
	return newCipherText

}

func PaddingOracle(victim string) string {
	victimBytes, err := hex.DecodeString(victim)
	check(err)
	blockNumber := len(victimBytes) / aes.BlockSize
	plaintext := make([]byte, (blockNumber-1)*aes.BlockSize)
	mu := sync.Mutex{}

	for blockIndex := blockNumber - 2; blockIndex >= 0; blockIndex-- {
		for byteIndex := aes.BlockSize - 1; byteIndex >= 0; byteIndex-- {
			paddingCount := aes.BlockSize - byteIndex
			injectIndex := blockIndex*aes.BlockSize + byteIndex
			log.Printf("processing paddingCount = %d", paddingCount)

			ch := make(chan byte)

			for guess := 0x00; guess < 0x100; guess++ {
				guessByte := byte(guess)

				var testBytes []byte = make([]byte, (blockIndex+2)*aes.BlockSize)
				copy(testBytes, victimBytes)
				testBytes[injectIndex] = victimBytes[injectIndex] ^ guessByte ^ byte(paddingCount)
				for index := injectIndex + 1; index < (blockIndex+1)*aes.BlockSize; index++ {
					testBytes[index] = victimBytes[index] ^ plaintext[index] ^ byte(paddingCount)
				}

				go func(erByte []byte, guessByte byte, injectIndex int, paddingCount int, ch chan byte) {
					er := hex.EncodeToString(erByte)
				try:
					resp, err := http.Get("https://crypto-class.appspot.com/po?er=" + er)
					if err != nil {
						log.Println(err)
						time.Sleep(time.Duration(300) * time.Millisecond)
						goto try
					}
					defer resp.Body.Close()
					// if the StatusCode is 404, the padding is good but the format is malformed
					// if the StatusCode is 200 and the paddingCount is as specified,
					// it is exactly the orignal ciphertext so we shall include this situation
					mu.Lock()
					defer mu.Unlock()
					if resp.StatusCode == 404 || (resp.StatusCode == 200 && paddingCount == int(plaintext[len(plaintext)-1])) {
						log.Println(guessByte, injectIndex, resp.StatusCode)
						ch <- guessByte
					}

				}(testBytes, guessByte, injectIndex, paddingCount, ch)
			}

			plaintext[injectIndex] = <-ch
			close(ch)
		}
	}

	return string(plaintext[:])
}
