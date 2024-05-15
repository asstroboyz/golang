package main

import "fmt"

func main() {
	// Mendeklarasikan variabel counter dengan nilai awal 0
	counter := 0

	// Mendefinisikan fungsi anonim increment
	increment := func() {
		// Mencetak "Increment" ke konsol
		fmt.Println("Increment")

		// Menambahkan nilai counter dengan 1
		counter++
	}

	// Memanggil fungsi increment beberapa kali
	increment() // Pertama kali
	increment() // Kedua kali
	increment() // Ketiga kali

	// Mencetak nilai counter setelah beberapa kali pemanggilan fungsi increment
	fmt.Println(counter)
}
