package week4

import "fmt"

func ExampleGenerateCipherTextOnShortAES() {
	plaintext := "Pay Bob 100$"
	target := "Pay Bob 500$"
	ciphertext := "20814804c1767293b99f1d9cab3bc3e7ac1e37bfb15599e5f40eef805488281d"
	newCipherText := GenerateCipherTextOnShortAES(plaintext, ciphertext, target)
	fmt.Println(newCipherText)
	// Output: 20814804c1767293bd9f1d9cab3bc3e7ac1e37bfb15599e5f40eef805488281d
}

func ExamplePaddingOracle() {
	var er = "f20bdba6ff29eed7b046d1df9fb7000058b1ffb4210a580f748b4ac714c001bd4a61044426fb515dad3f21f18aa577c0bdf302936266926ff37dbf7035d5eeb4"
	fmt.Println(PaddingOracle(er))
	// Output: The Magic Words are Squeamish Ossifrage
}
