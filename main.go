package main

import (
	"fmt"
	ratingsmap "golang-basic-fundamental/ratings-map"
)

func main() {
	//* Calling a function which is returned map and assigned it to variables
	mapMovieRatings := ratingsmap.MovieRatings()
	fmt.Println(mapMovieRatings)

	
}