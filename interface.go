package main

import "fmt"

// Definisikan interface HasName
type HasName interface {
	GetName() string
}

// Fungsi SayHello menerima parameter yang memiliki interface HasName
func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

// Definisikan struct Person
type Person struct {
	Name string
}

// Definisikan struct Animal
type Animal struct {
	Name string
}

// Implementasikan metode GetName untuk struct Animal
func (animal Animal) GetName() string {
	return animal.Name
}

// Implementasikan metode GetName untuk struct Person
func (person Person) GetName() string {
	return person.Name
}

func main() {
	// Inisialisasi variabel person sebagai struct Person
	person := Person{Name: "Eko"}
	// Panggil fungsi SayHello dengan parameter person
	SayHello(person)

	animal := Animal{Name: "Kucing"}
	SayHello(animal)
}
