// L = log

package main

import (
	"fmt"
	"math/big"
	"log"
	// "reflect"
)

var MINUS_ONE = big.NewRat(-1, 1)
var ZERO = big.NewRat(0, 1)
var ONE = big.NewRat(1, 1)
var TWO = big.NewRat(2, 1)

func main() {
	a := big.NewRat(2, 1)
	// b := big.NewRat(1, 2)
	// a := big.NewRat(23465735903, 10000000000)
	// c := pow(a, b, 100, false) // uses all methods ~ if this works, high chance of all working
	c := exp_2(a, 100, false)
	// c := exp(a, 10, false)
	// c := log2(a, 50, false)
	// c := ln(a, 10, false)
	// fmt.Println(a.FloatString(10), b.FloatString(10), c.FloatString(10))
	fmt.Println(a.FloatString(10), c.FloatString(10))
}

// a^b = exp(b*ln(a))
// 0 <= a
func pow(a, b *big.Rat, target_precision int, L bool) (c *big.Rat) {

	c = big.NewRat(0, 1)

	// a^0 == 1
	if b_vs_zero := b.Cmp(ZERO); b_vs_zero == 0 {
		c.Add(c, ONE)
		return c
	}

	// 0^b == 0
	if a_vs_zero := a.Cmp(ZERO); a_vs_zero == 0 {
		return c
	} else if a_vs_zero == -1 {
		// a <= 0 fails
		log.Fatal("pow basis cannot be negative")
	}

	l := ln(a, target_precision, L)
	l.Mul(l, b)
	return exp(l, target_precision, L)
}

// approximated using Newton-Raphson on the inverse (ln)
func exp(a *big.Rat, target_precision int, L bool) (b *big.Rat) {

	if L {fmt.Println("exp", a.FloatString(10))}

	b = big.NewRat(1, 1)

	// exp(0) == 1
	if a_vs_zero := a.Cmp(ZERO); a_vs_zero == 0 {
		return b
	}

	precision := 0 // for now, precision is naiive
	for {
		if precision == target_precision {
			break
		}

		if L {fmt.Println("exp, precision", precision)}

		l := ln(b, target_precision, L)
		if L {fmt.Println("exp, l,b 1", l.FloatString(10), b.FloatString(10))}
		l.Neg(l)
		if L {fmt.Println("exp, l,b 2", l.FloatString(10), b.FloatString(10))}
		l.Add(l, a)
		if L {fmt.Println("exp, l,b 3", l.FloatString(10), b.FloatString(10))}
		l.Add(l, ONE)
		if L {fmt.Println("exp, l,b 4", l.FloatString(10), b.FloatString(10))}
		b.Mul(b, l)
		if L {fmt.Println("exp, l,b 5", l.FloatString(10), b.FloatString(10))}

		precision++
	}
	if L {fmt.Println("exp, b end", b.FloatString(10))}

	return b
}

// use taylor expansion
func exp_2(a *big.Rat, target_precision int, L bool) (b *big.Rat) {

	if L {fmt.Println("exp", a.FloatString(10))}

	b = big.NewRat(1, 1)
	
	t := big.NewRat(1, 1)
	f := big.NewRat(1, 1)

	// exp(0) == 1
	if a_vs_zero := a.Cmp(ZERO); a_vs_zero == 0 {
		return b
	}

	precision := 0 // for now, precision is naiive
	for {
		if precision == target_precision {
			break
		}

		if L {fmt.Println("exp, precision", precision)}

		t.Mul(t, a)
		if L {fmt.Println("exp, t", t)}
		ft := big.NewRat(int64(precision + 1), 1)
		if L {fmt.Println("exp, ft", ft)}
		f.Mul(f, ft)
		if L {fmt.Println("exp, f", f)}
		ft.Inv(f)
		if L {fmt.Println("exp, ft", ft)}
		tf := big.NewRat(1, 1)
		if L {fmt.Println("exp, tf", tf)}
		tf.Mul(t, ft)
		if L {fmt.Println("exp, tf", tf)}

		b.Add(b, tf)
		if L {fmt.Println("exp, b", b)}

		precision++
	}
	if L {fmt.Println("exp, b end", b.FloatString(10))}

	return b
}

// logT(x) = log2(x) / log(T)
// ln = logE
// 0 < a
func ln(a *big.Rat, target_precision int, L bool) (b *big.Rat) {
	if L {fmt.Println("ln", a.FloatString(10))}

	// exp(0) == 1
	if a_vs_zero := a.Cmp(ZERO); a_vs_zero == 0 {
		return b
	}
	
	b = log2(a, target_precision, L)
	if L {fmt.Println("ln, a,b", a.FloatString(10), b.FloatString(10))}
	
	C := big.NewRat(69314718056, 100000000000)
	if L {fmt.Println("ln, C", C.FloatString(10))}
	
	b.Mul(b, C)
	if L {fmt.Println("ln, a,b,C", a.FloatString(10), b.FloatString(10), C.FloatString(10))}
	
	return b
}

// http://www.claysturner.com/dsp/BinaryLogarithm.pdf
// 0 < a
func log2(_a *big.Rat, target_precision int, L bool) (b *big.Rat) {

	b = big.NewRat(0, 1)
	a := big.NewRat(0, 1)
	a.Set(_a)

	if L {fmt.Println("log2", a.FloatString(10))}
	if L {fmt.Println("log2, a.Num().Int64()", a.Num().Int64())}
	if L {fmt.Println("log2, a.Denom().Int64()", a.Denom().Int64())}
	
	if a_vs_zero := a.Cmp(ZERO); a_vs_zero <= 0 {
		log.Fatal("log2 not defined for values <= 0");
	}

	if a_vs_one := a.Cmp(ONE); a_vs_one == 0 {
		return b
	}
	
	// double a until 1 <= a
	for {

		if a_vs_one := a.Cmp(ONE); a_vs_one != -1 {
			break
		}

		a.Num().Lsh(a.Num(), 1) // double
		b.Add(b, MINUS_ONE)
	}
	if L {fmt.Println("log2 doubled", a.FloatString(10), b.FloatString(10))}

	// half a until a < 2
	for {

		if a_vs_two := a.Cmp(TWO); a_vs_two == -1 {
			break
		}

		a.Denom().Lsh(a.Denom(), 1) // half
		b.Add(b, ONE)
	}
	if L {fmt.Println("log2 halved", a.FloatString(10), b.FloatString(10))}

	// from here: 1 <= a < 2 <=> 0 <= b < 1

	// compare a^2 to 2 to reveal b bit-by-bit
	precision_counter := 0 // for now, precision is naiive
	v := big.NewRat(1, 2)
	for {
		if target_precision == precision_counter {
			break
		}

		if L {
			fmt.Println("log2 precision_counter", precision_counter)
			fmt.Println("log2 v", v.FloatString(10))
			fmt.Println("log2 a", a.FloatString(10))
			fmt.Println("log2 b", b.FloatString(10))
		}

		a.Mul(a, a) // THIS IS SLOW
		// a = big.NewRat(a.Num().Int64()*a.Num().Int64(), a.Denom().Int64()*a.Denom().Int64())

		if L {fmt.Println("log2 a^2", a.FloatString(10))}

		if a2_vs_two := a.Cmp(TWO); a2_vs_two != -1 {
			
			if L {fmt.Println("log2 2 <= a^2", a.FloatString(10))}

			a.Denom().Lsh(a.Denom(), 1) // half
			b.Add(b, v)
		} else {
			if L {fmt.Println("log2 a^2 < 2")}
		}

		v.Denom().Lsh(v.Denom(), 1) // half

		precision_counter++
	}

	if L {fmt.Println("log2 b", b.FloatString((10)))}

	return b;
}