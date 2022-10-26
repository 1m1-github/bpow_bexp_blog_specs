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
	// "runtime"
	// "math"
	"math/big"
	"strings"
	// "log"
	// "reflect"
	// "bpow/int"
)

var ZERO_BIGINT = big.NewInt(0)
var ONE_BIGINT = big.NewInt(1)
var TEN_BIGINT = big.NewInt(10)
// var ONE_DECIMAL = decimal{false, *ONE_BIGINT, 0}

// nomenclature
// n negative
// c coefficient
// q exponent
// params called out are changed in func

// s = n ?? 1 : 0
// (-1)^s * c
// type int struct {
// 	n bool
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
	// L := false

	var a decimal
	out := decimal{false, *ZERO_BIGINT, 0}

	// a := decimal{true, *big.NewInt(75), -2} // -0.75
	// b := decimal{false, *big.NewInt(25), -3} // 0.025
	// c := add(&a, &b, &out, L) // -0.75 + 0.025 = -0.725
	// fmt.Println(a, b, c)

	// a = decimal{true, *big.NewInt(75), -2} // -0.75
	// c = negate(&a, &out, L) // 0.75
	// fmt.Println(a, c)

	// a = decimal{true, *big.NewInt(75), -2} // -0.75
	// b = decimal{false, *big.NewInt(25), -3} // 0.025
	// c = subtract(&a, &b, &out, L) // -0.75 - 0.025 = -0.775
	// fmt.Println(a, b, c)

	// a = decimal{true, *big.NewInt(75), -2} // -0.75
	// b = decimal{false, *big.NewInt(25), -3} // 0.025
	// c = multiply(&a, &b, &out, L) // -0.75 * 0.025 = -0.018750000000000003
	// fmt.Println(a, b, c)

	// a = decimal{true, *big.NewInt(75), -2} // -0.75
	// c = inverse(&a, &out, 5, L) // -4/3
	// fmt.Println(a, c)

	// a = decimal{true, *big.NewInt(75), -2} // -0.75
	// b = decimal{false, *big.NewInt(25), -3} // 0.025
	// c = divide(&a, &b, &out, 5, L) // -0.75 / 0.025 = -0.018750000000000003
	// fmt.Println(a, b, c)

	// as := String(&decimal{false, *big.NewInt(0), 0})
	// bs := "0"
	// fmt.Println(as == bs)
	// as = String(&decimal{true, *big.NewInt(75), -2})
	// bs = "-0.75"
	// fmt.Println(as == bs)
	// as = String(&decimal{false, *big.NewInt(75), 5})
	// bs = "7500000"
	// fmt.Println(as == bs)
	// as := String(&decimal{false, *big.NewInt(75), -1})
	// bs := "7.5"
	// fmt.Println(as)
	// fmt.Println(as == bs)

	// a = decimal{false, *big.NewInt(100), -2} // 1
	// precision := int64(10)
	// normalize(&a, &out, precision, true, true)
	// fmt.Println("a, String(&a), out, String(&out)", a, String(&a), out, String(&out))
	// a = decimal{false, *big.NewInt(5000000000), 10} // 0.5
	// precision = int64(10)
	// normalize(&a, &out, precision, true, true)
	// fmt.Println("a, String(&a), out, String(&out)", a, String(&a), out, String(&out))

	// func round(a, out *decimal, precision int64, normal bool) (*decimal) {
	// a = decimal{false, *big.NewInt(75), -2} // 0.75
	// precision := int64(10)
	// round(&a, &out, precision, true, true)
	// fmt.Println("a, String(&a), out, String(&out)", a, String(&a), out, String(&out))

	// for m := 1 ; m <= 11 ; m++ {
		m := 90
		a = decimal{false, *big.NewInt(1), 0} // 1
		taylor_precision := uint(m)
		precision := int64(100)
		exp_df(&a, &out, taylor_precision, precision, true)
		// fmt.Println(a, out)
		fmt.Println(m, String(&a), String(&out))		
	// }
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

// # Convert a decimal to a string
func String(a *decimal) (string) {

	negative := ""
	if a.n {
		negative = "-"
	}

	a_q_int := int(a.q)
	a_c_str := a.c.String()

	if 0 < a.q {
		zeros := strings.Repeat("0", a_q_int)
		return fmt.Sprintf("%v%v%v", negative, a_c_str, zeros)
	} else if a.q < 0 {
		len_c := len(a_c_str)
		shift := a_q_int + len_c
		if 0 < shift {
			return fmt.Sprintf("%v%v%v%v", negative, a_c_str[0:shift], ".", a_c_str[shift:])
		}
		zeros := strings.Repeat("0", -shift)
		return fmt.Sprintf("%v%v%v%v", negative, "0.", zeros, a_c_str)
	}

	return fmt.Sprintf("%v%v", negative, a_c_str)

    // c = string(x.c)
    // negative = (x.s == 1) ? "-" : ""
    // if x.q > 0
    //     print(io, negative, c, repeat("0", x.q))
    // elseif x.q < 0
    //     shift = x.q + length(c)
    //     if shift > 0
    //         print(io, negative, c[1:shift], ".", c[(shift+1):end])
    //     else
    //         print(io, negative, "0.", repeat("0", -shift), c)
    //     end
    // else
    //     print(io, negative, c)
    // end
}

func copy(a *decimal) (*decimal) {
	return &decimal{a.n, a.c, a.q}
}

// c = a + b
func add(a, b, out *decimal, precision int64, L bool) (*decimal) {
	if L {fmt.Println("add", "a", "b", a, String(a), b, String(b))}

	// cx = (-1)^x.s * x.c * 10^max(x.q - y.q, 0)
	// cy = (-1)^y.s * y.c * 10^max(y.q - x.q, 0)

	// aq := a.q
	// bq := b.q
	// aqmbq := sub_int(&aq, &bq)
	aqmbq := a.q - b.q
	if L {fmt.Println("add", "aqmbq", aqmbq)}
	
	aqmbq_abs := aqmbq
	if aqmbq_abs < 0 {
		aqmbq_abs = -aqmbq_abs
	}
	// ten_power := 10^aqmbq.c
	ten_power := big.NewInt(10)
	ten_power.Exp(ten_power, big.NewInt(aqmbq_abs), big.NewInt(0))
	// ten_power :=  big.Int(10)  10^aqmbq
	if L {fmt.Println("add", "ten_power", ten_power, ten_power.String())}

	ca := a.c
	if a.n {
		ca.Neg(&ca)
	}
	if L {fmt.Println("add", "ca", ca, ca.String())}

	cb := b.c
	if b.n {
		cb.Neg(&cb)
	}
	if L {fmt.Println("add", "cb", cb, cb.String())}

	if 0 < aqmbq {
		ca.Mul(&ca, ten_power)
	} else if aqmbq < 0 {
		cb.Mul(&cb, ten_power)
	}
	if L {fmt.Println("add", "ca", ca, ca.String())}
	if L {fmt.Println("add", "cb", cb, cb.String())}
	// cx = (-1)^x.s * x.c * 10^max(x.q - y.q, 0)
	// cy = (-1)^y.s * y.c * 10^max(y.q - x.q, 0)
	

	// s = (abs(cx) > abs(cy)) ? x.s : y.s
	var n bool
	switch ca.CmpAbs(&cb) {
	case 1: n = a.n
	default: n = b.n
	}
	if L {fmt.Println("add", "n", n)}
	// s = (abs(cx) > abs(cy)) ? x.s : y.s

	// c = BigInt(cx) + BigInt(cy)
	var c big.Int
	c.Add(&ca, &cb)
	if L {fmt.Println("add", "c", c, c.String())}
	// c = BigInt(cx) + BigInt(cy)
	
	// min(x.q, y.q)
	q := a.q
	if (b.q < a.q) {
		q = b.q
	}
	if L {fmt.Println("add", "q", q)}
	// min(x.q, y.q)

	out.n = n
	out.c = *c.Abs(&c)
	out.q = q

	// normalize(Decimal(s, abs(c), min(x.q, y.q)))
	if L {fmt.Println("add", "out", out, String(out))}
	normalize(copy(out), out, precision, false, L)
	if L {fmt.Println("add", "out", out, String(out))}
	return out
	// normalize(Decimal(s, abs(c), min(x.q, y.q)))
}

// -a
func negate(a, out *decimal, L bool) (*decimal) {
	out.n = !a.n
	out.c = a.c
	out.q = a.q
	return out
}

// a - b
func subtract(a, b, out *decimal, precision int64, L bool) (*decimal) {
	negate(b, out, L)
	add(a, out, out, precision, L)
	return out
}

// a * b
func multiply(a, b, out *decimal, precision int64, L bool) (*decimal) {
	// if L {fmt.Println("multiply", "a", String(a), "b", String(b), "precision", precision)}
	// if L {fmt.Println("multiply", "a", a, "b", b)}
	out.n = a.n != b.n
	// if L {fmt.Println("multiply", "out.n", out.n)}
	out.c.Mul(&a.c, &b.c)
	// if L {fmt.Println("multiply", "out.c", out.c)}
	out.q = a.q + b.q
	// if L {fmt.Println("multiply", "out.q", out.q)}
	return normalize(copy(out), out, precision, false, L)
}

// 1 / a
func inverse(a, out *decimal, precision int64, L bool) (*decimal) {
	if L {fmt.Println("inverse", "a", String(a), "precision", precision)}

	out.n = a.n

	if L {fmt.Println("inverse", "out.n", out.n)}

	ten_power := big.NewInt(10)
	ten_power.Exp(ten_power, big.NewInt(-a.q + precision), big.NewInt(0))
	out.c.Div(ten_power, &a.c)

	if L {fmt.Println("inverse", "out.c", out.c)}
	
	out.q = -precision

	if L {fmt.Println("inverse", "out.q", out.q)}
	if L {fmt.Println("inverse", "out", out, String(out))}
	
	norm := normalize(copy(out), out, precision, false, L)
	if L {fmt.Println("inverse", "norm", norm, String(norm))}
	return norm
	
	// c = round(BigInt(10)^(-x.q + DIGITS) / x.c) # the decimal point of 1/x.c is shifted by -x.q so that the integer part of the result is correct and then it is shifted further by DIGITS to also cover some digits from the fractional part.
    // q = -DIGITS # we only need to remember that there are these digits after the decimal point
    // normalize(Decimal(x.s, c, q))
}


// utils
func iszero(a *decimal) (bool) {
	return a.c.Cmp(ZERO_BIGINT) == 0
}
// utils

// a / b
func divide(a, b, out *decimal, precision int64, L bool) (*decimal) {
	inverse(b, out, precision, L)
	multiply(a, copy(out), out, precision, L)
	return out
}

// e^a
// df decimal float
func exp_df(a, out *decimal, taylor_precision uint, precision int64, L bool) (*decimal) {

	if L {fmt.Println("a", String(a), "taylor_precision", taylor_precision, "precision", precision)}

	if iszero(a) {
		out.n = false
		out.c = *ONE_BIGINT // possible problem
		out.q = 0
		return out
	}

	ONE := decimal{false, *ONE_BIGINT, 0} // 1
	a_power := decimal{false, *ONE_BIGINT, 0} // 1
	factorial := decimal{false, *ONE_BIGINT, 0} // 1
	factorial_next := decimal{false, *ZERO_BIGINT, 0} // 0
	factorial_inv := decimal{false, *ONE_BIGINT, 0} // 1
	
	// out = 1
	out.n = false
	out.c = *ONE_BIGINT
	out.q = 0

	if L {fmt.Println("out", String(out))}

	for i := uint(0) ; i < taylor_precision ; i++ {
		if L {fmt.Println("i", i)}

		if L {fmt.Println("a", String(a), a)}
		if L {fmt.Println("a_power", String(&a_power), a_power)}
		multiply(copy(&a_power), a, &a_power, precision, false) // a^i
		if L {fmt.Println("a_power", String(&a_power), a_power)}

		if L {fmt.Println("ONE", String(&ONE), ONE)}
		if L {fmt.Println("factorial_next", String(&factorial_next) ,factorial_next)}
		add(copy(&factorial_next), &ONE, &factorial_next, precision, false) // i + 1
		if L {fmt.Println("factorial_next", String(&factorial_next), factorial_next)}
		
		if L {fmt.Println("factorial", String(&factorial), factorial)}
		multiply(copy(&factorial), &factorial_next, &factorial, precision, false) // i!
		if L {fmt.Println("factorial", String(&factorial), factorial)}
		
		if L {fmt.Println("factorial_inv", String(&factorial_inv), factorial_inv)}
		inverse(&factorial, &factorial_inv, precision, false) // 1 / i!
		if L {fmt.Println("factorial_inv", String(&factorial_inv), factorial_inv)}

		multiply(&a_power, copy(&factorial_inv), &factorial_inv, precision, false) // store in factorial_inv as not needed anymore
		if L {fmt.Println("factorial_inv", String(&factorial_inv), factorial_inv)}

		if L {fmt.Println("out", String(out), out)}
		add(copy(out), &factorial_inv, out, precision, true)
		if L {fmt.Println("out", String(out), out)}
	}

	if L {fmt.Println("out", String(out))}

	return out
}

// # Rounding
// function round(x::Decimal; digits::Int=0, normal::Bool=false)
//     shift = BigInt(digits) + x.q
//     if shift > BigInt(0) || shift < x.q
//         (normal) ? x : normalize(x, rounded=true)
//     else
//         c = Base.round(x.c / BigInt(10)^(-shift))
//         d = Decimal(x.s, BigInt(c), x.q - shift)
//         (normal) ? d : normalize(d, rounded=true)
//     end
// end
func round(a, out *decimal, precision int64, normal bool, L bool) (*decimal) {

	shift := big.NewInt(precision)
	q := big.NewInt(a.q)
	shift.Add(shift, q)

	out.n = a.n
	out.c = a.c
	out.q = a.q

	if shift.Cmp(ZERO_BIGINT) == 1 || shift.Cmp(q) == -1 {
		if normal {
			return out
		}
		return normalize(out, out, precision, true, L)
	}

	shift.Neg(shift) // shift *= -1
	var ten_power big.Int
	ten_power.Exp(TEN_BIGINT, shift, ZERO_BIGINT) // 10^shift
	out.c.Div(&out.c, &ten_power)
	out.q += shift.Int64()

	if normal {
		return out
	}

	return normalize(copy(out), out, precision, true, L)
}


// # Normalization: remove trailing zeros in coefficient
// function normalize(x::Decimal; rounded::Bool=false)
//     p = 0
//     if x.c != 0
//         while x.c % 10^(p+1) == 0
//             p += 1
//         end
//     end
//     c = BigInt(x.c / 10^p)
//     q = (c == 0 && x.s == 0) ? 0 : x.q + p
//     if rounded
//         Decimal(x.s, abs(c), q)
//     else
//         round(Decimal(x.s, abs(c), q), digits=DIGITS, normal=true)
//     end
// end

func normalize(a, out *decimal, precision int64, rounded bool, L bool) (*decimal) {
	out.n = a.n
	
	p := int64(0)
	ten_power := big.NewInt(10) // 10^(p+1)
	if a.c.Cmp(ZERO_BIGINT) != 0 { // if a.c != 0
		for {
			var t big.Int
			if t.Mod(&a.c, ten_power).Cmp(ZERO_BIGINT) != 0 { // if a.c % 10^(p+1) != 0
				break
			}
			p++ // p = p + 1
			ten_power.Mul(ten_power, TEN_BIGINT) // 10^(p+1)
		}
	}

	ten_power.Div(ten_power, TEN_BIGINT) // 10^p
	out.c.Div(&a.c, ten_power) // out.c = a.c / 10^p
	out.c.Abs(&out.c) // out.c = abs(out.c)

	out.q = 0
	if !(out.c.Cmp(ZERO_BIGINT) == 0 && !a.n) { // if out.c == 0
		out.q = a.q + p
	}
	
	if rounded {
		return out
	}

	return round(copy(out), out, precision, true, L)
}
