package basics

import "fmt"

// estas variables estaran disponibles dentro de este archivo
// "false" es el valor por defecto para variables declaradas como "bool"
var c, python, java bool

func localsVariables()  {
	// estas variables estaran disponibles solo dentro de la funcion
	// "0" es el valor por defecto para variables declaradas como "int"
	var i int

	fmt.Println(c, python, java, i)
}

// variables definidas como "string" e inicializadas
var typescript, php string = "Typescript", "PHP"

func inicializedVariables()  {
	// estas variables tomaran el tipo de dato de su valor
	var c, python, java = true, false, "Java"

	fmt.Println(typescript, php, c, python, java)
}

func shortInicializedVariables()  {
	// cuando una variable tiene su valor definido se puede omitir la palabra "var"
	// solo funciona para variables declaradas dentro de una funcion
	k := 3

	fmt.Println(k)
}

