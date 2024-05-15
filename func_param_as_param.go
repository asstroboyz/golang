package main

import "fmt"

type Filter func(string)string

func sayHellowithFilter(name string, filter Filter){
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string{
	if name == "Kasar"{
		return "..." //jika name berisi kasar maka akan diganti menjadi ...
	} else {
		return name
	}
}

func main(){
	sayHellowithFilter("Musa",spamFilter)

	filter := spamFilter
	sayHellowithFilter("Kasar", filter)
}