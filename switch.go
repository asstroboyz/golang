package main

import "fmt"

func main() {

	name := "Gunawan"

	switch name {
	case "Ganda":
		fmt.Println("Heloo Ganda")
	case "Gunawan":
		fmt.Println("Heloo Gunawan")
	case "Eko":
		fmt.Println("Heloo Eko")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlau panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlau panjang")
	case length > 5:
		fmt.Println("Nama Lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
