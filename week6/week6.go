package week6

import (
	"math/big"
	"errors"
)

/* FactorCloselyFactorSemiPrime finds p, q such that
	N = p*q when |p - q| < 2N^{1/4} and p <= q
 */

func FactorCloselyFactorSemiPrime(N *big.Int) (*big.Int, *big.Int, error) {
	// A = \lceil\sqrt{N}\rceil
	NSqrt := new(big.Int).Sqrt(N)
	NSqrt.Add(NSqrt, big.NewInt(1))

	// \lfloor x\rfloor = \sqrt{A^2 - N}\rfloor
	Difference := new(big.Int).Mul(NSqrt, NSqrt)
	Difference.Sub(Difference, N)

	SqrtDifference := new(big.Int).Sqrt(Difference)

	SquaredSqrtDifference := new(big.Int).Mul(SqrtDifference, SqrtDifference);

	// check if x == \lfloor x \rfloor
	if SquaredSqrtDifference.Cmp(Difference) == 0 {
		p := new(big.Int).Sub(NSqrt, SqrtDifference)
		q := new(big.Int).Add(NSqrt, SqrtDifference)
		return p, q, nil
	}

	return nil, nil, errors.New("The factor is not closely enough for efficient factoring")
}

