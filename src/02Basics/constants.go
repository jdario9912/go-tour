package basics

import "fmt"

const Pi = 3.14

func Constants() {
	const World = "Mundo"

	fmt.Println(Pi)
	fmt.Println(World)
}

const (
	Big   = 1 << 62
	Small = Big >> 99
)

func NumericsConstants() {

	fmt.Println(Big)
	fmt.Println(Small)
}
