
// Decimal(s, c, q) = (-1)^s * c * 10^q
// Decimal(0, 1, 1) = (-1)^0 * 1 * 10^1 = 10
// Decimal(0, 10, -1) = 10


// Decimal(0, 1000, 1) = Decimal(0, 1, 3)



package main

import (
	"fmt"
	"math/big"
	// "log"
	// "reflect"
)

// (-1)^s * c
type int struct {
	s uint
	c uint
}

// (-1)^s * c * 10^q and q = (-1)^qs * qc
type decimal struct {
	s uint
	c big.Int // >= 0
	qs uint
	qc uint
}

func main() {
	one := decimal{0, *big.NewInt(1), 0, 0}
	answer := add(&one, &one, 2, true)
	fmt.Println(answer)
	a:=int{1,2}
	fmt.Println(a)
}

func add(a, b *decimal, target_precision uint, L bool) (c decimal) {

	// aq := *a

	// cx = (-1)^x.s * x.c * 10^max(x.q - y.q, 0)
    // cy = (-1)^y.s * y.c * 10^max(y.q - x.q, 0)
    // s = (abs(cx) > abs(cy)) ? x.s : y.s
    // c = BigInt(cx) + BigInt(cy)
    // normalize(Decimal(s, abs(c), min(x.q, y.q)))

	return *a
}