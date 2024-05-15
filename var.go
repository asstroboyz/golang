package main

import "fmt"

func main() { // redeklarsi tidak boleh
	// var name string
	// name = " Risdandi Ganda gunawan"
	// fmt.Println(name) //string

	name := " Ganda Gunawan"
	fmt.Println(name)

	name = "Risdandi"
	fmt.Println(name)

	var (
		firstName  = "Risdandi"
		middleName = " Ganda"
		lastname   = "gunawan"
	)

	fmt.Println(firstName, middleName, lastname)
}
