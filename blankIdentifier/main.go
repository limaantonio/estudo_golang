package main

import (
	"fmt"
	"net/http" 
	"io/ioutil"
)

func main() {
	res, _ := http.Get("https://google.com")
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", body) 


}