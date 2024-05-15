package main

import "fmt"

// Fungsi untuk menghitung faktorial menggunakan loop
func factorialLoop(value int) int {
	result := 1

	// Loop dari value ke 1, mengalikan result dengan nilai i setiap iterasi
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// Fungsi untuk menghitung faktorial menggunakan rekursi
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	// Menghitung faktorial 10 secara manual untuk memastikan hasilnya benar
	result := 1
	for i := 10; i > 0; i-- {
		result *= i
	}
	fmt.Println(result) // Output manual

	// Menghitung faktorial 10 menggunakan fungsi factorialLoop
	fmt.Println(factorialLoop(10)) // Output menggunakan loop

	// Menghitung faktorial 10 menggunakan fungsi factorialRecursive
	fmt.Println(factorialRecursive(10)) // Output menggunakan rekursi
}
