package week6

import (
	"math/big"
	"errors"
)

func isSquareNumber(n *big.Int) (bool, *big.Int) {
	sqrtFloor := new(big.Int).Sqrt(n)
	squareSqrtFloor := new(big.Int).Mul(sqrtFloor, sqrtFloor)
	return squareSqrtFloor.Cmp(n) == 0, sqrtFloor
}

/* FactorCloselyFactorSemiPrime finds p, q such that
	N = p*q when |p - q| < 2N^{1/4} and p <= q
 */
func FactorCloselyFactorSemiPrime(N *big.Int) (*big.Int, *big.Int, error) {
	return FactorNearlyFactorSemiPrime(N, 0)
}

/* FactorCloselyFactorSemiPrime finds p, q such that
	N = p*q when |p - q| < 2^{magnitude + 1} N^{1/4} and p <= q
	Notice that when magnitude = 1, it reduces to `FactorCloselyFactorSemiPrime`
 */

func FactorNearlyFactorSemiPrime(N *big.Int, magnitude uint) (*big.Int, *big.Int, error) {
	// A = \lceil\sqrt{N}\rceil
	NSqrt := new(big.Int).Sqrt(N)
	AvgGuess := new(big.Int).Set(NSqrt)
	AvgGuessSqaure := new(big.Int).Mul(AvgGuess, AvgGuess)
	Difference := new(big.Int).Sub(AvgGuessSqaure, N)
	one := big.NewInt(1)

	for i := 0; i < (1 << ( magnitude + magnitude)); i++ {
		// increment Difference by 2*A_guess + 1
		Difference.Add(Difference, AvgGuess)
		Difference.Add(Difference, AvgGuess)
		Difference.Add(Difference, one)

		// increment A_guess by 1
		AvgGuess.Add(AvgGuess, one)

		// check if x^2 is square number
		if isSquare, SqrtDifference := isSquareNumber(Difference); isSquare {
			p := new(big.Int).Sub(AvgGuess, SqrtDifference)
			q := new(big.Int).Add(AvgGuess, SqrtDifference)
			return p, q, nil
		}
	}

	return nil, nil, errors.New("The factor is not closely enough for efficient factoring")
}
