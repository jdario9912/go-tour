package basics

import "fmt"

type Vertex struct {
	X uint
	Y uint
}

func vertex() {
	vertexes := Vertex{1, 5874}
	fmt.Println(vertexes)
	fmt.Println(vertexes.X)
	fmt.Println(vertexes.Y)
}

func pointersToStructs() {
	println("v se inicia en {21, 48}")
	v := Vertex{21, 48}
	println("")

	println("Se define una variable p := v")
	p := v

	
	p.Y = 56

	// "b" apunta al espacio en memoria que esta guardado en "v"
	b := &v
	b.Y = 67

	fmt.Printf("v: %v\n", v)
	println("")

	println("p es una instancia de v y se inicia con sus valores")
	println("Cuando p redefine algun valor, los valores de v se mantienen iguales")
	fmt.Printf("p.X: %v\n", p.X)
	fmt.Printf("p.Y: %v\n", p.Y)

	println("")
	println("b apunta al espacio en memoria que esta guardado en v y se inicia con sus valores")
	println("Cuando b redefine algun valor, los valores de v tambien cambian")
	fmt.Printf("b.X: %v\n", b.X)
	fmt.Printf("b.Y: %v\n", b.Y)

	fmt.Printf("v: %v\n", v)
}
