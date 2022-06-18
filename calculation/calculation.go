package calculation

func Calculation(firstNumber int, secondNumber int, operator string) float32 {
	var result float32

	switch operator {
	case "+":
		result = float32(firstNumber) + float32(secondNumber)
	case "-":
		result = float32(firstNumber) - float32(secondNumber)
	case "*":
		result = float32(firstNumber) * float32(secondNumber)
	case "/":
		result = float32(firstNumber) / float32(secondNumber)
	}

	return result
}