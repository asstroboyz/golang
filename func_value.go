package main

import "fmt"

// Fungsi getGoodBye menerima sebuah string 'name' dan mengembalikan sebuah string.
func getGoodBye(name string) string {
    return "Good Bye " + name // Menambahkan spasi setelah "Good Bye" untuk pemisahan yang benar
}

func main() {
    contoh := getGoodBye // Mendeklarasikan variabel 'contoh' yang mengacu pada fungsi getGoodBye
    misal := getGoodBye  // Mendeklarasikan variabel 'misal' yang mengacu pada fungsi getGoodBye

    // Memanggil fungsi 'contoh' dengan argumen "Ganda" dan mencetak hasilnya
    fmt.Println(contoh("Ganda"))

    // Memanggil fungsi 'misal' dengan argumen "Ganda" dan mencetak hasilnya
    fmt.Println(misal("Ganda"))

	//function as a value yang artnya function juga bisa digunakan sebagai value
}
