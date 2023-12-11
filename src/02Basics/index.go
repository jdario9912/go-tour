package basics

import "fmt"

func Main()  {
	favouriteNumber()
	squareNumber()
	piNumber()
	add := add(13, 42)
	fmt.Println(add)
	hola, mundo := swap("Hola", "Mundo")
	fmt.Println(hola, mundo)
	fmt.Println(split(74))
	_, b := split(67)
	fmt.Println(b)
	a, _ := split(67)
	fmt.Println(a)
	localsVariables()
	inicializedVariables()
	shortInicializedVariables()
	types()
	zeroValues()
	typeConvert()
	typeInference()
	constants()
	numericsConstants()
	loopFor()
	loopWhile()
	infinity()
}

// seguir aqui https://go.dev/tour/flowcontrol/5
// falta crear archivo