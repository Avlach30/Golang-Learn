package main

func main() {
	//* TODO
	/*
		* 1. Develop a looping program with limitation of string length, with condition if index is even number, print at console
		* 2. Develop a same looping program with same limitation, but different condition if letter contains vocals only (a, i, u, e, o) 
	*/

	title := "Golang the best language"

	//* TODO 1
	for index, letter := range title {
		if (index % 2 == 0) {
			println(string(letter))
		}
	}

	//* TODO 2
	for _, letter := range title {
		switch string(letter) {
		case "a", "i", "u", "e", "o":
			println(string(letter))
		}
	}
}