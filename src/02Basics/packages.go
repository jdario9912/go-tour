package basics

// el nombre del paquete siempre responde a la ultima palabra de la importacion
import (
	"fmt"
	"math/rand"
)

func favouriteNumber() {
	fmt.Println("My favoutite number is", rand.Intn(200))
}
