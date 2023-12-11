package basics

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func types() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func zeroValues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func typeConvert() {
	var x, y int = 3, 4
	// para que "f" pueda tomar el valor de la ecuacion entre "x" e "y"
	// debe existir una conversion explicita
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func typeInference()  {
	// num1 es de tipo numero
	// sera "int", "float64" o "complex128" dependiendo de la presicion del valor
	// num1 := 4
	// num1 := 3.14
	num1 := 0.867 + 0.5i


	fmt.Printf("num1 is type %T and your value is %v.\n", num1, num1)
}