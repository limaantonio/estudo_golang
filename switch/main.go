package main

import "fmt"

func main() {
	a := "Antonio"

	switch a {
		case "Maria":
			fmt.Println("Hey Maria")
		case "Joao":
			fmt.Println("Hey Joao")
		default:
			fmt.Println("Hey")
	}
}