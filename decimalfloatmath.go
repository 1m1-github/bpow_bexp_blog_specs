// https://github.com/JuliaMath/Decimals.jl

// this package shadows the type int
// that is ok as this pkg only uses uint and big.Int from go

// Decimal(s, c, q) = (-1)^s * c * 10^q
// Decimal(0, 1, 1) = (-1)^0 * 1 * 10^1 = 10
// Decimal(0, 10, -1) = 10

// Decimal(0, 1000, 1) = Decimal(0, 1, 3)

package main

import (
	"fmt"
	// "math"
	"math/big"
	// "log"
	// "reflect"
	// "bpow/int"
)

// nomenclature
// n negative
// c coefficient
// q exponent

// s = n ?? 1 : 0
// (-1)^s * c
// type int struct {
// 	n nool
// 	c uint
// }

// s = n ?? 1 : 0
// (-1)^s * c * 10^q
type decimal struct {
	n bool
	c big.Int // >= 0
	q int64
}

func main() {

	// a := int{false, 2}
	// b := int{false, 3}
	// c := add_int(&a, &b)
	// fmt.Println(c)

	// a = int{false, 2}
	// b = int{true, 3}
	// c = add_int(&a, &b)
	// fmt.Println(c)

	// a = int{false, 2}
	// b = int{true, 2}
	// c = add_int(&a, &b)
	// fmt.Println(c)

	// a = int{true, 2}
	// b = int{false, 2}
	// c = add_int(&a, &b)
	// fmt.Println(c)

	// a = int{true, 5}
	// b = int{true, 7}
	// c = sub_int(&a, &b)
	// fmt.Println(c)

	// zero := int{false, 0}
	L := false

	one := decimal{true, *big.NewInt(1), 0}
	two := add(&one, &one, L)
	four := add(&two, &two, L)
	six := add(&four, &two, L)
	fmt.Println(one, two, four, six)

	a := decimal{true, *big.NewInt(75), -2} // -0.75
	b := decimal{false, *big.NewInt(25), -3} // 0.025
	c := add(&a, &b, L) // -0.75 + 0.025 = -0.725
	fmt.Println(a, b, c)

	a = decimal{true, *big.NewInt(75), -2} // -0.75
	c = negate(&a, L) // 0.75
	fmt.Println(a, c)

	a = decimal{true, *big.NewInt(75), -2} // -0.75
	b = decimal{false, *big.NewInt(25), -3} // 0.025
	c = subtract(&a, &b, L) // -0.75 - 0.025 = -0.775
	fmt.Println(a, b, c)

	a = decimal{true, *big.NewInt(75), -2} // -0.75
	b = decimal{false, *big.NewInt(25), -3} // 0.025
	c = multiply(&a, &b, L) // -0.75 * 0.025 = -0.018750000000000003
	fmt.Println(a, b, c)

	a = decimal{true, *big.NewInt(75), -2} // -0.75
	c = inverse(&a, 5, L) // -4/3
	fmt.Println(a, c)

	a = decimal{true, *big.NewInt(75), -2} // -0.75
	b = decimal{false, *big.NewInt(25), -3} // 0.025
	c = divide(&a, &b, 5, L) // -0.75 / 0.025 = -0.018750000000000003
	fmt.Println(a, b, c)
}

// func add_int(a, b *int) (int) { // a + b
// 	if (*a).p == (*b).p {return int{(*a).p, (*a).c + (*b).c}}
// 	if (*a).c <= (*b).c {return int{(*b).p, (*b).c - (*a).c}}
// 	return int{(*a).p, (*a).c - (*b).c}
// }
// func negate_int(a *int) (int) { // -a
// 	return int{!(*a).p, (*a).c}
// }
// func sub_int(a, b *int) (int) { // a - b
// 	bn := negate_int(b)
// 	return add_int(a, &bn)
// }

// a + b
func add(a, b *decimal, L bool) (decimal) {
	if L {fmt.Println("add, a, b", a, b)}

	// cx = (-1)^x.s * x.c * 10^max(x.q - y.q, 0)
	// cy = (-1)^y.s * y.c * 10^max(y.q - x.q, 0)

	// aq := a.q
	// bq := b.q
	// aqmbq := sub_int(&aq, &bq)
	aqmbq := a.q - b.q
	if L {fmt.Println("aqmbq", aqmbq)}

	// ten_power := 10^aqmbq.c
	ten_power := big.NewInt(10)
	ten_power.Exp(ten_power, big.NewInt(aqmbq), big.NewInt(0))
	// ten_power :=  big.Int(10)  10^aqmbq
	if L {fmt.Println("ten_power", ten_power)}

	ca := a.c
	if a.n {
		ca.Neg(&ca)
	}
	if L {fmt.Println("ca", ca)}

	cb := b.c
	if b.n {
		cb.Neg(&cb)
	}
	if L {fmt.Println("cb", cb)}

	if 0 < aqmbq {
		ca.Mul(&ca, ten_power)
	} else if aqmbq < 0 {
		cb.Mul(&cb, ten_power)
	}
	if L {fmt.Println("ca", ca)}
	if L {fmt.Println("cb", cb)}
	// cx = (-1)^x.s * x.c * 10^max(x.q - y.q, 0)
	// cy = (-1)^y.s * y.c * 10^max(y.q - x.q, 0)
	

	// s = (abs(cx) > abs(cy)) ? x.s : y.s
	var n bool
	switch ca.CmpAbs(&cb) {
	case 1: n = a.n
	default: n = b.n
	}
	if L {fmt.Println("n", n)}
	// s = (abs(cx) > abs(cy)) ? x.s : y.s

	// c = BigInt(cx) + BigInt(cy)
	var c big.Int
	c.Add(&ca, &cb)
	if L {fmt.Println("c", c)}
	// c = BigInt(cx) + BigInt(cy)
	
	// min(x.q, y.q)
	q := a.q
	if (b.q < a.q) {
		q = b.q
	}
	if L {fmt.Println("q", q)}
	// min(x.q, y.q)

	// normalize(Decimal(s, abs(c), min(x.q, y.q)))
	// todo
	// normalize(Decimal(s, abs(c), min(x.q, y.q)))

	return decimal{n, *c.Abs(&c), q}
}

// -a
func negate(a *decimal, L bool) (decimal) {
	return decimal{!a.n, a.c, a.q}
}

// a - b
func subtract(a, b *decimal, L bool) (decimal) {
	bn := negate(b, L)
	return add(a, &bn, L)
}

// a * b
func multiply(a, b *decimal, L bool) (decimal) {
	n := !(a.n && b.n)
	var c big.Int
	c.Mul(&a.c, &b.c)
	return decimal{n, c, a.q + b.q}
}

// 1 / a
func inverse(a *decimal, precision int64, L bool) (decimal) {
	ten_power := big.NewInt(10)
	ten_power.Exp(ten_power, big.NewInt(-a.q + precision), big.NewInt(0))
	var c big.Int
	c.Div(ten_power, &a.c)
	q := -precision
	return decimal{a.n, c, q}

	// c = round(BigInt(10)^(-x.q + DIGITS) / x.c) # the decimal point of 1/x.c is shifted by -x.q so that the integer part of the result is correct and then it is shifted further by DIGITS to also cover some digits from the fractional part.
    // q = -DIGITS # we only need to remember that there are these digits after the decimal point
    // normalize(Decimal(x.s, c, q))
}

// a / b
func divide(a, b *decimal, precision int64, L bool) (decimal) {
	bi := inverse(b, precision, L)
	return multiply(a, &bi, L)
}