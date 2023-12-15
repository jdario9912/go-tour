package basics

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// condicional tradicional
	if x < 0 {
		return sqrt(x) + "i"
	}

	return fmt.Sprint(math.Sqrt((x)))
}

func pow(x, n, lim float64) float64 {
	// condicional que crea un valor y despues lo compara
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}
