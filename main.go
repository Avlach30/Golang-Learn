package main

import (
	"fmt"
	sumarrays "golang-basic-fundamental/sum-arrays"
	calculation "golang-basic-fundamental/calculation"
)

func main() {
	//!TODO
	/*
		* 1: Create a function which is accept parameter from array of numbers, return it with sum of each element 		of arrays
		* 2: Create a calculate function which is accept three parameter, such as firstNumber, secondNumber, and 			operator. Add a custom error if operator isn't add, substract, multiply, or divide. So returned value 	is two values. Calculate result and error (if available)
	*/


	//* Assignment result of todo 1
	scores := []int{10, 5, 8, 9, 7}
	total := sumarrays.Sum(scores)
	fmt.Println(total)

	println("=================")

	//* Assignment result of todo 2
	calculateDivide, err := calculation.Calculation(22, 4, "/")
	if (err != nil) { //* If error not null (available)
		fmt.Println(err)
	} else {
		fmt.Println(calculateDivide)
	}

	calculateUnknown, err := calculation.Calculation(35, 8, "]")
	if (err != nil) { //* If error not null (available)
		fmt.Println(err)
	} else {
		fmt.Println(calculateUnknown)
	}
}