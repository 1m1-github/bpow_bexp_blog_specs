package main

import (
	"fmt"
	"math/big"
	"log"
)

func main() {
	a := big.NewRat(9, 1)
	b := logE(a, 10)
	fmt.Println("a", a.FloatString(10), b.FloatString(10))
}

func exp(a *big.Rat) (b *big.Rat) {

}

func logE(a *big.Rat, target_precision int) (b *big.Rat) {
	b = log2(a, target_precision)
	C := big.NewRat(69314718056, 100000000000)
	b.Mul(b, C)
	return b
}

func log2(_a *big.Rat, target_precision int) (b *big.Rat) {

	// fmt.Println("log2", _a.FloatString(10))

	a := big.NewRat(_a.Num().Int64(), _a.Denom().Int64())

	ZERO := big.NewRat(0, 1)
	if a_vs_zero := a.Cmp(ZERO); a_vs_zero <= 0 {
		log.Fatal("log2 not defined for values <= 0");
	}


	ONE := big.NewRat(1, 1)
	if a_vs_one := a.Cmp(ONE); a_vs_one == 0 {
		return ZERO;
	}
	
	MINUS_ONE := big.NewRat(-1, 1)
	TWO := big.NewRat(2, 1)

	b = big.NewRat(0, 1)

	// double a until 1 <= a
	for {

		if a_vs_one := a.Cmp(ONE); a_vs_one != -1 {
			break
		}

		a.Num().Lsh(a.Num(), 1) // double
		b.Add(b, MINUS_ONE)
	}
	// fmt.Println("doubled", a.FloatString(10), b.FloatString(10))

	// half a until a < 2
	for {

		if a_vs_two := a.Cmp(TWO); a_vs_two == -1 {
			break
		}

		a.Denom().Lsh(a.Denom(), 1) // half
		b.Add(b, ONE)
	}
	// fmt.Println("halved", a.FloatString(10), b.FloatString(10))

	// from here: 1 <= a < 2 <=> 0 <= b < 1

	// compare a^2 to 2 to reveal b bit-by-bit
	precision_counter := 0
	v := big.NewRat(1, 2)
	for {
		if target_precision == precision_counter {
			break
		}

		// fmt.Println("precision_counter", precision_counter)
		// fmt.Println("v", v.FloatString((10)))
		// fmt.Println("a", a.FloatString((10)))
		// fmt.Println("b", b.FloatString((10)))

		a.Mul(a, a)

		// fmt.Println("a^2", a.FloatString((10)))

		if a2_vs_two := a.Cmp(TWO); a2_vs_two != -1 {
			
			// fmt.Println("2 <= a^2", a.FloatString(10))

			a.Denom().Lsh(a.Denom(), 1) // half
			b.Add(b, v)
		} 
		// else {
		// 	fmt.Println("a^2 < 2")
		// }

		v.Denom().Lsh(v.Denom(), 1) // half

		precision_counter++
	}

	// fmt.Println("b", b.FloatString((10)))

	return b;
}