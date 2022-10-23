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
var MAX_ERROR = big.NewRat(1, 10000000)

func main() {
	a := big.NewRat(1, 1)
	b := big.NewRat(1, 2)
	// a := big.NewRat(23465735903, 10000000000)

	for i := 1; i < 20; i++ {
		// c := ln(a, i, false)
		c := exp(a, i, false)
		// c := pow(a, b, i, false) // uses all methods ~ if this works, high chance of all working
		fmt.Println(i, a.FloatString(10), b.FloatString(10), c.FloatString(10))
	}
	// c := log2(a, 5, false)
	
	// c := ln_2(a, 10, false)
	// fmt.Println(a.FloatString(10), b.FloatString(10), c.FloatString(10))
	// c = ln(a, 10, false)
	// fmt.Println(a.FloatString(10), c.FloatString(10))
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
func exp_slow(a *big.Rat, target_precision int, L bool) (b *big.Rat) {

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

		// b *= 1 + a - ln(a)
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
		// b *= 1 + a - ln(a)

		precision++
	}
	if L {fmt.Println("exp, b end", b.FloatString(10))}

	return b
}

// use taylor expansion
func exp(a *big.Rat, target_precision int, L bool) (b *big.Rat) {

	if L {fmt.Println("exp", a.FloatString(10))}

	b = big.NewRat(1, 1)
	
	a_power := big.NewRat(1, 1)
	factorial := big.NewRat(1, 1)

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

		// a^n
		a_power.Mul(a_power, a)
		if L {fmt.Println("exp, a_power", a_power)}
		// a^n

		// n!
		factorial_next := big.NewRat(int64(precision + 1), 1)
		factorial.Mul(factorial, factorial_next)
		if L {fmt.Println("exp, factorial, factorial_next", factorial, factorial_next)}
		// n!
		
		// 1/n!
		factorial_inv := big.NewRat(1, 1)
		factorial_inv.Inv(factorial)
		if L {fmt.Println("exp, factorial_inv", factorial_inv)}
		// 1/n!

		// a^n/n!
		taylor_term := big.NewRat(1, 1)
		taylor_term.Set(a_power)
		taylor_term.Mul(taylor_term, factorial_inv)
		if L {fmt.Println("exp, taylor_term", taylor_term)}
		// a^n/n!

		b.Add(b, taylor_term)
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

// approximated using Newton-Raphson on the inverse (ln)
// 0 < a
func ln_2(_a *big.Rat, target_precision int, L bool) (b *big.Rat) {

	b = big.NewRat(0, 1)
	a := big.NewRat(0, 1)
	a.Set(_a)

	if L {fmt.Println("ln_2", a.FloatString(10))}
	if L {fmt.Println("ln_2, a.Num().Int64()", a.Num().Int64())}
	if L {fmt.Println("ln_2, a.Denom().Int64()", a.Denom().Int64())}
	
	// range
	if a_vs_zero := a.Cmp(ZERO); a_vs_zero <= 0 {
		log.Fatal("log2 not defined for values <= 0");
	}

	// edge
	if a_vs_one := a.Cmp(ONE); a_vs_one == 0 {
		return b
	}

	precision := 0 // for now, precision is naiive
	for {
		if precision == target_precision {
			break
		}

		if L {fmt.Println("ln_2, precision", precision)}

		// b += 2 * (a - exp(b)) / (a + exp(b))
		e := exp(b, target_precision, L)
		if L {fmt.Println("ln_2, l,b 1", e.FloatString(10), b.FloatString(10))}

		l1 := big.NewRat(1, 1)
		l1.Set(e)
		l1.Neg(l1)
		l1.Add(l1, a)
		if L {fmt.Println("ln_2, l1", l1.FloatString(10))}

		l2 := big.NewRat(1, 1)
		l2.Set(e)
		l2.Add(l2, a)
		if L {fmt.Println("ln_2, l2", l2.FloatString(10))}

		l := big.NewRat(1, 1)
		l.Set(l2)
		l.Inv(l)
		l.Mul(l, l1)
		if L {fmt.Println("ln_2, l", l.FloatString(10))}
		
		l.Mul(l, TWO)

		abs_l := big.NewRat(1, 1)
		abs_l.Abs(l)
		if abs_l_vs_max_error := abs_l.Cmp(MAX_ERROR) ; abs_l_vs_max_error != 1 {
			break
		}
		
		b.Add(b, l)
		if L {fmt.Println("ln_2, l, b", l.FloatString(10), b.FloatString(10))}
		// b += 2 * (a - exp(b)) / (a + exp(b))

		precision++
	}
	if L {fmt.Println("ln_2, b end", b.FloatString(10))}

	return b
}