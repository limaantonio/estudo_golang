package scope

import "fmt"

var y int = 20//declarando uma variavel int

//go run *.go roda toda a aplicacao (todos os pacotes)

func main() {
	x := 10
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}


func printY() {
	fmt.Println(y)
}