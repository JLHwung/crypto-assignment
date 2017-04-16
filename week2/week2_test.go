package week2

import (
	"fmt"
)

func ExampleCBC_decrypt() {
	const keyStr string = "140b41b22a29beb4061bda66b6747e14"
	const ciphertextStr string = "5b68629feb8606f9a6667670b75b38a5b4832d0f26e1ab7da33249de7d4afc48e713ac646ace36e872ad5fb8a512428a6e21364b0c374df45503473c5242a253"
	fmt.Printf("%s", CBC_decrypt(keyStr, ciphertextStr))
	// Output: Our implementation uses rand. IV
}

func ExampleCTR_decrypt() {
	const keyStr string = "36f18357be4dbd77f050515c73fcf9f2"
	const ciphertextStr string = "69dda8455c7dd4254bf353b773304eec0ec7702330098ce7f7520d1cbbb20fc388d1b0adb5054dbd7370849dbf0b88d393f252e764f1f5f7ad97ef79d59ce29f5f51eeca32eabedd9afa9329"
	fmt.Printf("%s", CTR_decrypt(keyStr, ciphertextStr))
	// Output: CTR mode lets you build a stream cipher from a block cipher.
}
