package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var ganda Customer
	fmt.Println(ganda)

	ganda.Name = "Ganda gunawan"
	ganda.Address = "Pekalongan"
	ganda.Age = 21
	fmt.Println(ganda)
	fmt.Println(ganda.Name)
	fmt.Println(ganda.Age)
}
