package main

import "fmt"
import "math/big"

func factorial(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}
	if (n==0) {
		return big.NewInt(1)
	}

	bigN := big.NewInt(n)
	return bigN.Mul(bigN, factorial(n-1))
}

func main() {
	n := factorial(100)
	r, s, zero := big.NewInt(0), big.NewInt(0), big.NewInt(0)

	for n.Cmp(zero) > 0 {
		n.QuoRem(n, big.NewInt(10), r)
		s.Add(s, r)
	}

	fmt.Println(s)
}