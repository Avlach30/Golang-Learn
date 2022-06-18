package calculation

import "errors" //* Must import errors package for customized an error

//!TODO TODO 2 Function

//* Define a function which is returned multiple values, include custom error handling
func Calculation(firstNum int, secondNum int, operator string) (float32, error) {
	var result float32
	var err error //* Define variable with error data type

	switch operator {
	case "+":
		result = float32(firstNum) + float32(secondNum)
	case "-":
		result = float32(firstNum) - float32(secondNum)
	case "*":
		result = float32(firstNum) * float32(secondNum)
	case "/":
		result = float32(firstNum) / float32(secondNum)
	default:
		err = errors.New("unknown operation type") //* Reassign error variable with custom errors
	}

	return result, err
}