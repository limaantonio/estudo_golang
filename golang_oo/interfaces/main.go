package main

import "fmt"

type vehicle interface {
	start() string
}

type car struct {
	name string
}

func (c car) start() string {
	return "The car started"
}

type motorcycle struct {
	name string
}

func (m motorcycle) start() string {
	return "The motorcycle started"
}

func startVehicle(v vehicle) string {
	return v.start()
}

func main() {
	car := car{"Gol"}
	m := motorcycle{"Titan"}
	fmt.Println(startVehicle(car))
	fmt.Println(startVehicle(m))

}
