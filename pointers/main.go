package main

import "fmt"

func main() {
	x := 10

	fmt.Println(x)

	y := &x

	fmt.Println(y)//valor da memoria
	fmt.Println(*y)//10

	*y = 20
	fmt.Println(x)

	var z *int = &x
}