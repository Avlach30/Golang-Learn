package main

import "fmt"

//* Define a struct in go
//* 	Struct like a class, so developer can make a blueprint of object and can implement a oop too
type Citizen struct {
	Name string
	Address string
	NumberPhone int
	IsMarried bool 
}

//* Define a function which is accept parameter from struct
func describeCitizen (citizen Citizen) string {
	return fmt.Sprintf(
		"Hello!, my name is %s, i live at %s, i'm %t married. You can contact me at %d", 
		citizen.Name, citizen.Address, citizen.IsMarried, citizen.NumberPhone)
}

func main() {

	//* Make istance from struct with assign value directly
	ahmad := Citizen{
		Name: "Ahmad mujahidin",
		Address: "Diponegoro 43",
		NumberPhone: 628123457890,
		IsMarried: true,
	}
	fmt.Println(ahmad)
	describeAhmad := describeCitizen(ahmad) //* Calling a function with parameter of istance from struct
	fmt.Println(describeAhmad)

	//* Make istance from struct with assign value undirectly
	diana := Citizen{}
	diana.Name = "Diana Anjani"
	diana.Address = "Ahmad yani 7"
	diana.NumberPhone = 628754398765
	diana.IsMarried = false

	fmt.Println(diana)
}
