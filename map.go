package main

import "fmt"

func main() {

	// var person map[string]string = map[string]string{}
	// person["name"] = "Eko"
	// person["addres"] = "Jl Bandung"
	// person["contact"] = "+621234"

	person := map[string]string{
		"name":   "eko",
		"addres": "subang",
	}

	fmt.Println(person["name"])
	fmt.Println(person["addres"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Ganda"
	book["ups"] = "salah"

	fmt.Println(book)
	delete(book, "ups") //menghapus data ups
	fmt.Println(book)
}
