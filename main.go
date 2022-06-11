package main 

func main() {
	//* List option of define variable methods

	var firstNum int = 23 //* First option of define variables, with assign value directly

	var secondNum int //* Second option of define variables, can assign value undirectly or later
	secondNum = 15

	//* Third option of define variables without keyword var, but must assign value directly
	totalCalculation := firstNum * secondNum 
	println(totalCalculation)

	totalCalculation = 55 //* Reassign existing variable with new value
	println(totalCalculation)
}