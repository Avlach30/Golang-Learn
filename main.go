package main

import (
	"fmt"
	entity"golang-basic-fundamental/management-city/entity"
)

func main() {	
	ahmad := entity.Citizen{
		Name: "Ahmad mujahidin",
		Address: "Diponegoro 43",
		NumberPhone: 628123457890,
		IsMarried: true,
	}
	fmt.Println(ahmad)

	println("============")

	
	diana := entity.Citizen{}
	diana.Name = "Diana Anjani"
	diana.Address = "Ahmad yani 7"
	diana.NumberPhone = 628754398765
	diana.IsMarried = false

	fmt.Println(diana)
	fmt.Println(diana.Describe())

	println("============")

	chandra := entity.Citizen{
		Name: "Chandra wijaya",
		Address: "Pattimura 77",
		NumberPhone: 628575438965,
		IsMarried: true,
	}

	surabaya := entity.City{
		Name: "Surabaya",
		Major: chandra,
		Citizens: []entity.Citizen{ahmad, diana, chandra},
	}
	surabaya.Describe()
}
