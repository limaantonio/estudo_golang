package main

import "fmt"

// func funcName(a int) int {
// 	return a * a
// }

// func namedReturn(a string) (x string) {
// 	x = a 
// 	return
// }

// func moreReturn(a string, b string) (string string) {
// 	return a, b
// }

func variadicFunc(x ...int) int {
	var res int 
	for _, v := range x {
		res += v
	}
	return res
}

func funcInsideFunc() func() int {
	x := 10
	return func() int {
		return x * x
	}
}

func main() {
	fmt.Println(variadicFunc(1, 2, 3, 4))

	z := 0

	add := func() int {
		z += 2
		return z
	}

	fmt.Println(add())
}