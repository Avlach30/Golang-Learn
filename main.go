package main

import (
	"fmt"
	slicefruits "golang-basic-fundamental/slice-fruits"
)

func main() {
	tropicalFruits := slicefruits.TropicalFruits();
	fmt.Println(tropicalFruits)

	itemTarget := "Lychee"
	index := findIndexArrayOfStrings(itemTarget, tropicalFruits)
	println(index)
}

//* Define function whis is return an index of array element
func findIndexArrayOfStrings(element string, arrSource []string) int {
	for index, item := range arrSource {
		if (element == item) {
			return index
		}
	}

	return -1
}
