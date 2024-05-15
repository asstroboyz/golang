package main 

import "fmt"

// Mendefinisikan alias tipe untuk fungsi blacklist
type Blacklist func(string) bool

// Fungsi untuk mendaftarkan pengguna dan memeriksa blacklist
func registerUser(name string, blacklist Blacklist) {
    if blacklist(name) {
        fmt.Println("You are blocked", name)
    } else {
        fmt.Println("Welcome", name)
    }
}

func main() {
    // Mendefinisikan fungsi blacklist yang memblokir nama "anjing"
    blacklist := func(name string) bool {
        return name == "anjing"
    }

    // Mendaftarkan pengguna "eko" dan memeriksa blacklist
    registerUser("eko", blacklist)

    // Mendaftarkan pengguna "anjing" dengan fungsi blacklist inline
    registerUser("anjing", func(name string) bool {
        return name == "anjing"
    })
}
