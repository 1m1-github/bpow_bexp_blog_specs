# specs for bpow, bexp, blog for TEAL

## general

a []byte `b` combined with the params `neg` and `radix` is interpreted as a fixed point as follows:

let $L = len(b) - 1$

$$
x = (-1)^{`neg`}\sum_{i=-`radix`}^{L-`radix`} 256^i \cdot b_{L-(i+`radix`)}
$$

`b` is interpreted as a base 256 number with `radix` and polarity (`neg`) given  
this represents the current interpretion for opcodes like `b+`. choosing the same interpretation should make interoperability for opcodes easier.

`neg` and `radix` are common `uint`

## bpow

bpow negA radixA negB radixB

Stack: ..., A: []byte, B: []byte → ..., C: []byte

computes $C=A^B$  
A,B,C follow interpretation above

## bexp

bexp neg radix

Stack: ..., A: []byte → ..., C: []byte

computes $C = e^A$  
A,C follow interpretation above

## blog

blog radix

Stack: ..., A: []byte → ..., C: []byte

computes $C = ln(A)$  
A,C follow interpretation above

## decimal vs binary

the above is a binary (n-ary with $n=2^m$) representation of fixed point values ~ an alternative is to use decimal representations  

### pros binary

- homogeneity: same representation as other opcodes (`b+` etc.)

- size: if each digit is a byte, then a 256 base is optimal ~ using base 10 wastes a lot of space ~ could interpret a byte as 2 digits in the decimal expansion, which would waste less space, but gets arguably convoluted ~ on the other hand, []byte has max 4096 length, allowing a large range even in decimal even wasting space using a byte per digit ~ to represent a 512 bit binary value in decimal, we would need ca. 155 decimal digits

### pros decimal

- famously, 0.1 cannot be represented in binary fixed point finitely ~ this is not the case vice versa, i.e. no finite binary representation has an infinite decimal representation

### question

- if adopting the decimal representation, would it make sense to change the existing opcodes to the same to allow for composition of e.g. $bpow \circ b+ $


## use case

many applications in Finance, Math, Science need these functions at least over the rationals

## problems

the standard golang math lib  "... does not guarantee bit-identical results across architectures." (https://pkg.go.dev/math) ~ could be a problem if diff nodes get diff results  
potential solutions:  
- use a fixed point lib (downside is adding dependency on non standard libs)  
- restrict range such that bit-identical results are guaranteed
