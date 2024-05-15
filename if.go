package main

// if statement
import "fmt"

func main() {
	name := "Risdandi"

	if name == "Ganda" {
		fmt.Println("Hello Ganda")
	} else if name == "Budi" {
		fmt.Println("Hi, Budi")
	} else if name == "Joko" {
		fmt.Println("Hai Jokondo")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama Sudah benar")
	}
}
