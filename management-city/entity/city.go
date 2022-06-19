package entity 

import "fmt"

type City struct {
	Name string
	Major Citizen //* Imported from Citizen struct, because same package so not defined at import keyword
	Citizens []Citizen
}

func (city City) Describe() {
	citizenTotal := len(city.Citizens)

	fmt.Printf(
		"Welcome to %s city, this city is led by %s. Total population in this city is %d person",
		city.Name, city.Major.Name, citizenTotal,
	)

	println("")

	fmt.Println("Citizen list: ")
	for _, citizen := range city.Citizens {
		fmt.Printf(
			"Name: %s, Address: %s", 
			citizen.Name, citizen.Address,
		)
		println("")
	}
}