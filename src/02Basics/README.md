# Types in Go

## Basics

bool

string

int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
// represents a Unicode code point

float32 float64

complex64 complex128

## Type conversions

Some numeric conversions:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

Or, put more simply:

i := 42
f := float64(i)
u := uint(f)

In Go assignment between items of different type requires an explicit conversion.

```golang
func typeConvert() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
```

## Type inference

When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:

```golang
var i int
j := i // j is an int
```

But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:

```golang
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```
