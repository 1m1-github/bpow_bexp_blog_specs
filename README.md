# specs for bpow, bexp, blog for TEAL

## general

a []byte `b` combined with the params `neg` and `radix` is interpreted as a fixed point as follows:

let $L = len(b) - 1$

$$
x = (-1)^{`neg`}\sum_{i=-`radix`}^{L-`radix`} 256^i \cdot b_{L-(i+`radix`)}
$$

`b` is interpreted as a base 256 number with `radix` and polarity (`neg`) given
still represents the current interpretion for opcodes like `b+`. choosing the same interpretation should make interoperability for opcodes easier.

`neg` and `radix` are common `uint`

## bpow

bpow neg radix

Stack: ..., A: []byte, B: []byte → ..., C: []byte

computes $C=A^B$  
A,B,C follow interpretation above

## bexp

bexp neg radix

Stack: ..., A: []byte → ..., C: []byte

computes $C = e^A$  
A,C follow interpretation above

## blog

blog neg radix

Stack: ..., A: []byte → ..., C: []byte

computes $C = ln(A)$  
A,C follow interpretation above

## use case

many applications in Finance, Math, Science need these functions at least over the rationals

## problems

the standard golang math lib  "... does not guarantee bit-identical results across architectures." (https://pkg.go.dev/math) ~ could be a problem if diff nodes get diff results  
potential solutions:  
- use a fixed point lib (downside is adding dependency on non standard libs)  
- restrict range such that bit-identical results are guaranteed
