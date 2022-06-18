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

func main() {

	//* Make istance from struct with assign value directly
	ahmad := Citizen{
		Name: "Ahmad mujahidin",
		Address: "Diponegoro 43",
		NumberPhone: 628123457890,
		IsMarried: true,
	}
	fmt.Println(ahmad)

	//* Make istance from struct with assign value undirectly
	diana := Citizen{}
	diana.Name = "Diana Anjani"
	diana.Address = "Ahmad yani 7"
	diana.NumberPhone = 628754398765
	diana.IsMarried = false

	fmt.Println(diana)
}
