package week1

import (
	"encoding/hex"
	"log"
	"unicode/utf8"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

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

type ManyTimePaddingInteractiveCracker struct {
	target           string
	targetText       []byte
	cipherTextString []string
	XORedText        [][]byte
}

// Initialize preprocess the Cracker by computing the XORedText of every cipherText xor targetText
func (c *ManyTimePaddingInteractiveCracker) Initialize() {
	targetText, err := hex.DecodeString(c.target)
	checkError(err)
	c.targetText = targetText

	var XORedText [][]byte
	for i := 0; i < len(c.cipherTextString); i++ {
		cipherText, err := hex.DecodeString(c.cipherTextString[i])
		checkError(err)
		dst := make([]byte, len(targetText))
		safeXORBytes(dst, cipherText, targetText)
		XORedText = append(XORedText, dst)
	}

	c.XORedText = XORedText
}

// FindFirstHeuristic iterates on targetText to find given heuristic, typically we use ` the ` as fragment
func (c *ManyTimePaddingInteractiveCracker) FindFirstHeuristic(fragment string) {
	for index := 0; index < len(c.targetText)-len(fragment); index++ {
		c.ValidateHeuristic(index, fragment)
	}
}

// ExpandHeuristic takes expanded heuristic from given cipher text started from given index and outputs the expanded target plaintext fragment
func (c *ManyTimePaddingInteractiveCracker) ExpandHeuristic(cipherTextIndex int, index int, ptFragmentString string) []byte {
	result := expandHeuristic(c.XORedText[cipherTextIndex], index, ptFragmentString)
	log.Println("Expanded:")
	log.Printf("target: index = %d, text = [%s]\n", index, result)
	return result
}

// ValidateHeuristic computes the plaintext of given cipher text corresponding to the fragment of target plaintext
func (c *ManyTimePaddingInteractiveCracker) ValidateHeuristic(index int, fragment string) string {
	result, test := validateHeuristic(c.XORedText, index, fragment)
	if test == true {
		log.Println("Validated: ")
		for i := 0; i < len(c.cipherTextString); i++ {
			log.Printf("A%d: index = %d, text = [%s]\n", i, index, result[i])
		}
	}
	return fragment
}

func isRuneAllowable(rune rune) bool {
	return rune <= 122 && rune >= 97 ||
		rune <= 90 && rune >= 65 ||
		rune <= 57 && rune >= 48 ||
		rune <= 63 && rune >= 32 ||
		rune >= 64256 && rune <= 64262 || // latin small ligature
		rune == 0
}

func validateHeuristic(XORedText [][]byte, index int, fragmentString string) ([][]byte, bool) {
	fragment := []byte(fragmentString)
	fragmentLength := len(fragment)
	var guessResult [][]byte
	positive := true

	for i := 0; i < len(XORedText); i++ {
		dst := make([]byte, fragmentLength)

		var fragmentText []byte
		if index+fragmentLength >= len(XORedText[i]) {
			fragmentText = XORedText[i][index:]
		} else {
			fragmentText = XORedText[i][index : index+fragmentLength]
		}
		safeXORBytes(dst, fragmentText, fragment)
		guessResult = append(guessResult, dst)

		// now we test for positiveness
		for j, w := 0, 0; j < fragmentLength; j += w {
			runeValue, width := utf8.DecodeRune(dst[j:])
			if isRuneAllowable(runeValue) == false {
				log.Printf("Unacceptable %#U found in A%d\n", runeValue, i)
				positive = false
				break
			}
			w = width
		}
	}
	return guessResult, positive
}

func expandHeuristic(SomeXORedText []byte, index int, plaintextFragmentString string) []byte {
	ptFragment := []byte(plaintextFragmentString)
	ptFragmentLength := len(ptFragment)
	dst := make([]byte, ptFragmentLength)
	var fragmentText []byte
	if index+ptFragmentLength >= len(SomeXORedText) {
		fragmentText = SomeXORedText[index:]
	} else {
		fragmentText = SomeXORedText[index : index+ptFragmentLength]
	}
	safeXORBytes(dst, fragmentText, ptFragment)

	return dst
}
