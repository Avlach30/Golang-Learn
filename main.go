package main

import (
	"fmt"
	ratingsmap "golang-basic-fundamental/ratings-map"
)

func main() {
	//* Calling a function which is returned map and assigned it to variables
	mapMovieRatings := ratingsmap.MovieRatings()
	fmt.Println(mapMovieRatings)

	targetKey := mapMovieRatings["kkn"] //* Get a value in map from key
	println(targetKey)

	mapMovieRatings["kkn"] = 9 //* Update existing value in map from key
	fmt.Println(mapMovieRatings)

	delete(mapMovieRatings, "kkn") //* Delete existing key in map
	fmt.Println(mapMovieRatings)

	fmt.Println("==========================")

	programmingLangRatings := ratingsmap.ProgrammingLangRatings()
	fmt.Println(programmingLangRatings)

	programmingLangRatings["java"] = "good for native device" //* Insert a new key-value pair of map
	fmt.Println(programmingLangRatings)

	fmt.Println("==========================")

	length := len(programmingLangRatings) //* Get length of map element's
	println(length)

	fmt.Println("==========================")

	//* Loop over a map with limitation of map length
	for key, rating := range programmingLangRatings {
		fmt.Println("Language:", key, ".", "Slogan:", rating)
	}

	fmt.Println("==========================")

	//* Check the value and check is available or not from called map with key
	value, isAvailable := programmingLangRatings["java"]
	fmt.Println(value)
	fmt.Println(isAvailable)
}