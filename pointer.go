package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1 // Isinya disalin dari address1, namun address1 tidak berubah

	// address2.City = "Bandung"
	// fmt.Println(address1) // Output address1, yang seharusnya tetap sama
	// fmt.Println(address2) // Output address2, yang diubah menjadi "Bandung"

	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}

	// Deklarasi pointer untuk Address
	var address2 *Address = &address1

	// Mengakses dan mengubah nilai melalui pointer
	address2.City = "Bandung"

	// Mencetak nilai dari address1 dan address2
	fmt.Println(address1)
	fmt.Println(address2)
}
