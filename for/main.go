package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//while nao exite

	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	for { //loop 
		x++
		if x == 10 {
			continue
		}
		if x == 100 {
			break
		}
	}
}