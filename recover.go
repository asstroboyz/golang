package main

import "fmt"

// menghentikan program dipanggil ketika terjadi panic saat prgram berjalan

func endApp() {
	fmt.Println("end app")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups error")
	}

}
func main() {
	runApp(true)
	fmt.Println("Risdandi Ganda")
}
