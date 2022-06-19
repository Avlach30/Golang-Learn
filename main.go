package main

import (
	"fmt"
	hajj"golang-basic-fundamental/management-hajj/entity"
)

func main() {
	muzakki := hajj.Person{Name: "Ahmad Muzakki", Age: 67}
	fmt.Printf("Name before hajj: %s", muzakki.Name)
	println("")
	muzakki.ChangeHajjStatus() //* Calling a method with referencing technique for change struct value globally
	fmt.Printf("Name after hajj: %s", muzakki.Name)
}