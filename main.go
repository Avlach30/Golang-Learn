package main

import (
	"fmt"
	personHajj"golang-basic-fundamental/management-hajj/entity"
	changeHajj"golang-basic-fundamental/management-hajj/change-status"
)

func main() {
	muzakki := personHajj.Person{Name: "Ahmad Muzakki", Age: 67}
	fmt.Printf("Name before hajj: %s", muzakki.Name)
	println("")
	changeHajj.ChangeHajjStatus(&muzakki) 
	//* Calling a method with referencing technique for applying change object value globally
	fmt.Printf("Name after hajj: %s", muzakki.Name)
}