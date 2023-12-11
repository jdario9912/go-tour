package basics

import "fmt"

const Pi = 3.14

func constants() {
	const World = "Mundo"

	fmt.Println(Pi)
	fmt.Println(World)
}

const (
	Big   = 1 << 62
	Small = Big >> 99
)

func numericsConstants() {

	fmt.Println(Big)
	fmt.Println(Small)
}
