package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hellow", name, "my name is", customer.Name)
}

func main() {
	var ganda Customer
	fmt.Println(ganda)

	ganda.Name = "Ganda Gunawan"
	ganda.Address = "Pekalongan"
	ganda.Age = 21
	fmt.Println(ganda)
	fmt.Println(ganda.Name)
	fmt.Println(ganda.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Pekalongan",
		Age:     12,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "jawa", 21}
	fmt.Println(budi)

	budi.sayHello("Agus")
	ganda.sayHello("Agus")
}
