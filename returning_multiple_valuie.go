//mengembalikan multiple value
package main
import "fmt"

func getFullName()(string,string){
	return "Ganda", "Gunawan"
}
func main(){
	// firstName, lastName := getFullName()
	// fmt.Println(firstName,lastName) //return sesuai tipe data (multiple value)

	firstName, _ := getFullName()
	fmt.Println(firstName)
}