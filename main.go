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

//* Define embedded/nested struct (One of struct data type is from another struct)
type City struct {
	Name string
	Major Citizen //* Define data type is from another struct
	Citizens []Citizen //* Define data type is array or slice of struct
}

//* Define a method for Citizen struct
func (citizen Citizen) Describe() string {
	return fmt.Sprintf(
		"Hello!, my name is %s, i live at %s, i'm %t married. You can contact me at %d", 
		citizen.Name, citizen.Address, citizen.IsMarried, citizen.NumberPhone)
}

//* Define a function which is accept parameter from struct and returned string
func describeCity (city City) string {
	citizenTotal := len(city.Citizens)

	describeCity := fmt.Sprintf(
		"Welcome to %s city, this city is led by %s. Total population in this city is %d person",
		city.Name, city.Major.Name, citizenTotal,
	)

	println("")

	fmt.Println("Citizen list: ")
	for _, citizen := range city.Citizens {
		fmt.Println(citizen.Name)
	}

	return describeCity
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

	println("============")

	//* Make istance from struct with assign value undirectly
	diana := Citizen{}
	diana.Name = "Diana Anjani"
	diana.Address = "Ahmad yani 7"
	diana.NumberPhone = 628754398765
	diana.IsMarried = false

	fmt.Println(diana)
	fmt.Println(diana.Describe()) //* Calling a method

	println("============")

	chandra := Citizen{
		Name: "Chandra wijaya",
		Address: "Pattimura 77",
		NumberPhone: 628575438965,
		IsMarried: true,
	}

	surabaya := City{
		Name: "Surabaya",
		Major: chandra,
		Citizens: []Citizen{ahmad, diana, chandra}, //*Append or insert struct to array or slice
	}

	fmt.Println(describeCity(surabaya))
}
