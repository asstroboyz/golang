package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
	//error

}

func main() {
	runApplication()
}
