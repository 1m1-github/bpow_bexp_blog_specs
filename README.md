# specs for bpow, bexp, blog for TEAL

## general

a []byte `b` combined with the params `neg` and `width` is interpreted as a fixed point `x` as follows:
`
let $L = len(b) - 1$

$$`x`^{'} = (-1)^{`neg`}\sum_{i=0}^{L-1} 256^i \cdot b_{L-i-1}$$

$$`x` = `x`^{'} * 10^{-`width`}$$

`b` is interpreted as a base 256 number first and then `width` applies a shift to the decimal point and `neg` defines its polaity via its parity.    
this represents the current interpretion for opcodes like `b+`. choosing the same interpretation should make interoperability for opcodes easier.

`neg` and `width` are common `uint`

## bpow

bpow widthA widthB

Stack: ..., negA: uint, A: []byte, negB: uint, B: []byte → ..., negC: uint, C: []byte

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

## implementation details

### lib

no outside libs will be used ~ the code will only rely on `big.Int` and `big.Rat` from the `golang` `math` lib

### error handling

all range violations will result in immediate panic

### var types

`[]byte` inputs are interpreted as decimal fixed point values as described above ~ these will be represented as `big.Rat` rational numbers in `golang` ~ the return type of all the internal methods will also be `big.Rat` which can eventually be converted into a `[]byte` to return to the `opcode` caller

the `[]byte` to decimal method above is a bijection and covers all decimal values upto a specific width

### bpow

`pow` can be reduced to using `exp` and `log` as follows

$$ a^b = e^{b \cdot ln(a)} $$

### bexp

`exp` can be approximated using the Newton-Raphson method on the inverse (`log`) as follows

$$ e^a = b <=> a = ln(b) $$

$$ b_{n+1} = b_n - {f(b_n) \over f'(b_n)} = b_n - {ln(b_n) - a \over {1 \over b_n}} = b_n \cdot (1 - ln(b_n) + a)$$

### blog

`log` will be approximated using the algorithm described here: http://www.claysturner.com/dsp/BinaryLogarithm.pdf

that algorithm uses very basic operations only, like binary shifts and integer arithmetic and works upto arbitrary precision

## discussion

- should users provide output width as a param? using the same width param for inputs and outputs does not really make sense as the output will always have a larger error vs the input ~ if a user requires 5 decimal output precision, than the input should be more precise
