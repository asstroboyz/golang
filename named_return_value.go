package main
//
import "fmt"

func getCompleteName() (firstName, middleName, lastName string){
	firstName = "Risdandi"
	middleName = "Ganda"
	lastName = "Gunawan"

	return firstName, middleName,lastName
}
	
func main() {
 a,b,c := getCompleteName()
	fmt.Println(a,b,c)

}