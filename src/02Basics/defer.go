package basics

import "fmt"

func statementDefer()  {
	defer fmt.Println("world")

	fmt.Println("Hello")
}

func stakingDefer()  {
	fmt.Println("counting")

	for i := 0; i < 10; i ++ {
		
		if i == 4 {
			continue
		}
		defer fmt.Println(i)
	}

	fmt.Println("done")
}