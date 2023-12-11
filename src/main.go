package main

import (
	"fmt"
	welcome "go-tour/src/01Welcome"
	basics "go-tour/src/02Basics"
)

func main() {
	// welcome
  welcome.Hello()
	welcome.Time()

	// basics
	basics.FavouriteNumber()
	basics.SquareNumber()
	basics.PiNumber()
	add := basics.Add(13, 42)
	fmt.Println(add)
	hola, mundo := basics.Swap("Hola", "Mundo")
	fmt.Println(hola, mundo)
	fmt.Println(basics.Split(74))
	_, b := basics.Split(67)
	fmt.Println(b)
	a, _ := basics.Split(67)
	fmt.Println(a)
	basics.LocalsVariables()
	basics.InicializedVariables()
	basics.ShortInicializedVariables()
	basics.Types()
	basics.ZeroValues()
	basics.TypeConvert()
	//  seguir aca https://go.dev/tour/basics/14
}
