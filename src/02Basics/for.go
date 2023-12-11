package basics

import "fmt"

// en Go todos los loops se llaman "for".
// dependiendo de los parametros va a ser su comportamiento.

func loopFor() {
	var sum int

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

func loopWhile() {
	sum := 1

	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func infinity() {
	for {

	}
}
