package main

import "fmt"

func main() {
	implementPointer(5)
	println("=======")
	implementPointerWithVar(6)
}

func implementPointer(initNumber int) {
	//* Implementing of pointer
	secondNum := &initNumber 
	//* Get a memory address where value of initNumber stored, this is called referencing
	thirdNum := *secondNum 
	//* Get a real value from pointer variable, this is called dereferencing

	fmt.Println(secondNum)
	fmt.Println(thirdNum)

	*secondNum = initNumber * 3 
	//* Reassign real value from pointer, so affected pointer source (initNumber) value is changed too
	fmt.Println(*secondNum)

	fmt.Println(initNumber)
}

//* Implement a pointer with var keyword
func implementPointerWithVar(initNumber int) {
	var secondNum *int = &initNumber //* Referencing process
	var thirdNum = *secondNum //* Dereferencing process

	fmt.Println(secondNum)
	fmt.Println(thirdNum)

	*secondNum = initNumber * 3 //* Reassign real value from referencing result
	fmt.Println(*secondNum)

	fmt.Println(initNumber)
}

