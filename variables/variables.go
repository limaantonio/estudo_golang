package variabels

import "fmt"

func main() {
    a := 10
    b := "hello"
    c := 10.45
    d := false
    e := "W"
    f := `Uouuu`

   fmt.Printf("%v \n", a)//T mostra o tipo, v = valor, 
   fmt.Printf("%T \n", b)
   fmt.Printf("%T \n", c)
   fmt.Printf("%T \n", d)
   fmt.Printf("%T \n", e)
   fmt.Printf("%T \n", f)

   const x int = 13
   const (
       aa = 55
       bb = 34
   )

}