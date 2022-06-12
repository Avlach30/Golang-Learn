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
}