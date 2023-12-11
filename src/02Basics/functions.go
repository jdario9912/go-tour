package basics

// cuando los argumentos comparten el tipo se declara despues del ultimo
func add(x, y int) int {
	return x + y
}

// las funciones pueden retornar mas de un valor
func swap(x, y string) (string, string) {
	return x, y
}

// las funciones pueden retornar valores nombrados
// es de utilidad para la documentacion del significado de los valores
func split(sum int)(x, y int)  {
	x = sum * 4 / 9
	y = sum - x
	return
}