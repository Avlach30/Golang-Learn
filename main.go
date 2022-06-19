package main

import "fmt"

func main() {
	number := 11
	fmt.Printf("Value before processing: %d", number)
	println("")
	changeValue(&number, 100)
	//* Call function with referencing technique for get address memory where number value stored
	println("")
	fmt.Printf("Value after processing: %d", number)
}

//* Define a functions which is changed value globally with implement of pointer
func changeValue(oldNumber *int, newNumber int) {
	//* Process for reassign reference value with dereferencing technique (get real value from address memory)
	*oldNumber = newNumber
	fmt.Printf("New value: %d", *oldNumber)
}