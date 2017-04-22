package week5

import (
	"math/big"
	"log"
)

// DLog computes the discrete log modulo a prime p using Man-in-The-Middle attack.
// That is, computes z such that x^z = y (mod m)
func DLog(x, y, m *big.Int, bitLen int) *big.Int {
	one := big.NewInt(1)

	var middleNumber *big.Int = new(big.Int)
	middleNumber.Lsh(one, uint(bitLen / 2))

	log.Println("Building Hash Table...")
	hashTable := make(map[string]*big.Int, middleNumber.Int64())

	xInverse := new(big.Int)
	xInverse.ModInverse(x, m)

	power := new(big.Int)
	power.Set(y)
	hashTable[power.String()] = big.NewInt(0)

	for i := big.NewInt(1); i.Cmp(middleNumber) == -1; i.Add(i, one) {
		power.Mul(power, xInverse)
		power.Mod(power, m)
		hashTable[power.String()] = new(big.Int).Set(i)
	}

	log.Println("Searching Hash Table...")

	xPowerMiddleNumber := new(big.Int)
	xPowerMiddleNumber.Exp(x, middleNumber, m)

	bigPower := big.NewInt(1)
	for i := big.NewInt(0); i.Cmp(middleNumber) == -1; i.Add(i, one) {
		if needle, ok := hashTable[bigPower.String()]; ok {
			result := new(big.Int);
			result.Mul(i, middleNumber)
			result.Mod(result, m)
			result.Add(result, needle)
			result.Mod(result, m)
			return result;
		}
		bigPower.Mul(bigPower, xPowerMiddleNumber)
		bigPower.Mod(bigPower, m)
	}

	return nil;
}
