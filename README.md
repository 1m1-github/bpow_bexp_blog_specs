# specs for bpow, bexp, blog for TEAL

## general

a []byte `b` combined with the params `neg` and `width` is interpreted as a fixed point as follows:

let $L = len(b) - 1$

$$
x^' = (-1)^{`neg`}\sum_{i=0}^{L} 256^i \cdot b_{L-i}
x = x^' * 10^{-'width'}
$$

`b` is interpreted as a base 256 number first and the with `width` is applies a shift to the decimal point and `neg` defines its polaity via its parity.    
this represents the current interpretion for opcodes like `b+`. choosing the same interpretation should make interoperability for opcodes easier.

`neg` and `width` are common `uint`

## bpow

bpow widthA widthB

Stack: ..., negA: uint, A: []byte, negA: uint, B: []byte → ..., negC: uint, C: []byte

computes $C=A^B$  
A,B,C follow interpretation above

## bexp

bexp width

Stack: ..., neg: uint, A: []byte → ..., C: []byte

computes $C = e^A$  
A,C follow interpretation above

## blog

blog width

Stack: ..., A: []byte → ..., neg: uint, C: []byte

computes $C = ln(A)$  
`neg` represents polarity of C
A,C follow interpretation above


## use case

many applications in Finance, Math, Science need these functions at least over the rationals

## problems

the standard golang math lib  "... does not guarantee bit-identical results across architectures." (https://pkg.go.dev/math) ~ could be a problem if diff nodes get diff results  
potential solutions:  
- use a fixed point lib (downside is adding dependency on non standard libs)  
- restrict range such that bit-identical results are guaranteed
