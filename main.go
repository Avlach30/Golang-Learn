package main

import "fmt"

func main() {
	//!TODO
	//!TODO slice source => scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	//!TODO 1: Count an average from slice
	//!TODO 2: From same slice, you push each element to new slice with condition each element greater than or equal to 90

	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	//* Assignment for todo 1
	currenttotal := 0
	for _, score := range scores {
		currenttotal += score
	}

	average := float32(currenttotal) / float32(len(scores))
	fmt.Println(average)


	//* Assignment for todo 2
	var goodScores[] int
	for _, score := range scores {
		if (score >= 90) {
			goodScores = append(goodScores, score)
		}
	}
	fmt.Println(goodScores)
}