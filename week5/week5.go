package week5

import (
	"math/big"
	"log"
)

// DLog computes the discrete log modulo a prime p using Man-in-The-Middle attack.
// That is, computes z such that x^z = y (mod m)
func DLog(x, y, m *big.Int, bitLen uint) (int64, error) {
	middleNumber := int64(1 << (bitLen / 2))

	log.Println("Building Hash Table...")
	hashTable := make(map[string]int64, middleNumber)

	xInverse := new(big.Int).ModInverse(x, m)

	power := new(big.Int).Set(y)
	hashTable[power.String()] =0

	for i := int64(1); i < middleNumber; i++ {
		power.Mul(power, xInverse)
		power.Mod(power, m)
		hashTable[power.String()] = i
	}

	log.Println("Searching Hash Table...")

	middleNumberBigInt := new(big.Int).SetInt64(middleNumber)
	xPowerMiddleNumber := new(big.Int).Exp(x, middleNumberBigInt, m)

	bigPower := big.NewInt(1)
	for i := int64(0); i < middleNumber; i++ {
		if needle, ok := hashTable[bigPower.String()]; ok {
			return i * middleNumber + needle, nil;
		}
		bigPower.Mul(bigPower, xPowerMiddleNumber)
		bigPower.Mod(bigPower, m)
	}

	return 0, nil;
}
