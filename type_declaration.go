package main

import "fmt"

func main() {

	type NoKTP string

	var ktpGanda NoKTP = "111111"

	var contoh string = "22222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpGanda)
	fmt.Println(contohKtp)
}
