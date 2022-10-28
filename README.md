# specs for bpow, bexp, bln, blog2, blog10 for TEAL

## general

a []byte `b` combined with the params `neg` and `width` is interpreted as a `decimal` `x` as follows:
`
let $L = len(b) - 1$

$$`x`^{'} = (-1)^{`neg`}\sum_{i=0}^{L-1} 256^i \cdot b_{L-i-1}$$

$$`x` = `x`^{'} * 10^{-`width`}$$

`b` is interpreted as a base 256 number first (bid endian unsigned integer from byte array) and then `width` applies a shift to the decimal point and `neg` defines its polarity via its parity.    
this represents the current interpretion for opcodes like `b+`. choosing the same interpretation should make interoperability for opcodes easier.

`neg` and `width` are common `uint64`

## bpow widthB widthD widthY

- Opcode: 0x??
- Stack: ..., A: uint, B: []byte, C: uint, D: []byte &rarr; ..., X: uint, Y: []byte
- Decimal to the power of a `decimal`. B and D are interpreted as big-endian unsigned integers.
- **Cost**: base cost plus cost dependent on widthY
- Availability: v9

[(-1)^A * B * 10^widthB] ^ [(-1)^C * D * 10^widthD] = (-1)^X * Y * 10^widthY. widthB, widthD and widthY are uint64. widthY output precision is guaranteed.
Calculation is based on the following identity: q^w = exp(ln(q)/w).

## bexp widthB widthY

- Opcode: 0x??
- Stack: ..., A: uint, B: []byte &rarr; ..., Y: []byte
- e to the power of a `decimal`. B is interpreted as big-endian unsigned integers.
- **Cost**: base cost plus cost dependent on widthY
- Availability: v9

exp[(-1)^A * B * 10^widthB] = Y * 10^widthY. widthB and widthY are uint64. widthY output precision is guaranteed.
Calculation is based on the Taylor expansion of e^q.

## bln widthA widthY

- Opcode: 0x??
- Stack: ..., A: []byte &rarr; ..., X: uint, Y: []byte
- Natural logarithm of a `decimal`. A is interpreted as big-endian unsigned integers.
- **Cost**: base cost plus cost dependent on widthY
- Availability: v9

ln[A * 10^widthA] = (-1)^X * Y * 10^widthY. widthA and widthY are uint64. widthY output precision is guaranteed.
Calculation is based on the following identity: ln(q) = log2(q) / log2(e)

## blog2 widthA widthY

- Opcode: 0x??
- Stack: ..., A: []byte &rarr; ..., X: uint, Y: []byte
- Logarithm base 2 of a `decimal`. A is interpreted as big-endian unsigned integers.
- **Cost**: base cost plus cost dependent on widthY
- Availability: v9

log2[A * 10^widthA] = (-1)^X * Y * 10^widthY. widthA and widthY are uint64. widthY output precision is guaranteed.
Calculation is based on the on a binary appromixation that allows arbitrary precision.

## blog10 widthA widthY

- Opcode: 0x??
- Stack: ..., A: []byte &rarr; ..., X: uint, Y: []byte
- Logarithm base 10 of a `decimal`. A is interpreted as big-endian unsigned integers.
- **Cost**: base cost plus cost dependent on widthY
- Availability: v9

log10[A * 10^widthA] = (-1)^X * Y * 10^widthY. widthA and widthY are uint64. widthY output precision is guaranteed.
Calculation is based on the following identity: log10(q) = log2(q) / log2(10)


## use case

many applications in Finance, Math, Science need these functions at least over the rationals

## implementation details

### lib

no outside libs will be used ~ the code will only rely on `big.Int`from the `golang` `math` lib

### precision

the representation and algorithms allow for arbitrary precision ~ let the user choose the precision they need and let the opcode cost increase with the target precision
there will still be a max, naturally because `[]byte` has a max ~ currently, `b+` is limited to 512 bit inputs

### error handling

all range violations will result in immediate panic

### var types

the following `decimal` type allows exact representation of all decimal values within a range

```
// s = n ?? 1 : 0
// (-1)^s * c * 10^q
type decimal struct {
	n bool
	c big.Int // >= 0
	q int64
}
```

### bpow

`pow` can be reduced to using `exp` and `log` as follows

$$ a^b = e^{b \cdot ln(a)} $$

### bexp

use Taylor expansion:

$$ e^x = \sum_{i} \frac{x^i}{i!} $$

### bln

$$ ln(x) = \frac{log2(x)}{log2(e)} $$

### blog2

`log2` will be approximated using the algorithm described here: http://www.claysturner.com/dsp/BinaryLogarithm.pdf

that algorithm uses very basic operations only, like binary shifts and integer arithmetic and works upto arbitrary precision

short description of the algorithm:

- we want to calc $log_2(a) = b$
- normalise $a$ such that $1 \le a < 2 <=> 0 \le b < 1$ by halfing $a$ as long as $2 \le a$ and doubling $a$ as long as $a < 1$
- $2 \le a^2$ is equivalent to the next bit in $b$ being set
- in a loop, meeting comparing $a^2$ to 2 to keep revealing bits of $b$ until the desired precision is achieved
- after the comparison in the loop, $a$ is reduced depending on the revealed bit of $b$

### blog10

$$ log10(x) = \frac{log2(x)}{log2(10)} $$
- should users provide output width as a param? using the same width param for inputs and outputs does not really make sense as the output will always have a larger error vs the input ~ if a user requires 5 decimal output precision, than the input should be more precise

## streams of dev
https://youtu.be/Lwqmu2p-bsQ
