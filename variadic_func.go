package main

import "fmt"

// Fungsi sumAll menerima sejumlah argumen integer yang tidak terbatas (variadic function)
func sumAll(numbers ...int) int {
    total := 0 // Inisialisasi variabel total dengan nilai 0

    // Loop melalui setiap angka dalam slice numbers
    for _, number := range numbers {
        total += number // Tambahkan setiap angka ke total
    }
    return total // Kembalikan nilai total
}

func main() {
    // Cetak hasil dari pemanggilan fungsi sumAll dengan tiga argumen: 10, 101, dan 10
    fmt.Println(sumAll(10, 101, 10))

	 // Inisialisasi slice dengan beberapa angka
    numbers := []int{10, 10, 10, 10}
    // Cetak hasil dari pemanggilan fungsi sumAll dengan slice numbers yang diurai menjadi argumen variadic
    fmt.Println(sumAll(numbers...))
	// menggunakan titik 3 agar slice bisa terkirimkan kedalam nubers yang berupa int
}
