package sumarrays

//!TODO TODO 1 Function

//* Define a function with parameter array of numbers, return a value sum of element from arrays parameter
func Sum(scores []int) int {
	currentResult := 0
	for _, score := range scores {
		currentResult += score
	}

	return currentResult
}