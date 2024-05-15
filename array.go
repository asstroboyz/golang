package main

import "fmt"

func main() {

	var names [3]string //jumlahnya maksimal tidak bisa mnegisi lebih dari ini

	names[0] = "Risdandi"
	names[1] = "Ganda"
	names[2] = "Gunawan"

	//isinya tidak bisa melebihi 3 karena dihitung 3 dari 0

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		90,
		80,
		97,
		120,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values[3])
	fmt.Println(len(values))
}
