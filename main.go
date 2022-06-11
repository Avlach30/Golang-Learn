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
	
	tropicalFruits[index] = "Rambutans"
	fmt.Println(tropicalFruits)

	tropicalFruits = append(tropicalFruits[:index], tropicalFruits[index+1:]...)
	/*
		* Delete a slice element from index
	  * 	With copying executable slice, but excluce index element which is want to delete it
	*/
	fmt.Println(tropicalFruits)
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
