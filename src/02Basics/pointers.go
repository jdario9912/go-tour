package basics

import "fmt"

func pointers() {
	i := 42
	j := 1207

	p := &i
	fmt.Println(p)
	*p = 21
	fmt.Println(i)

	p = &j
	fmt.Println(j)
}