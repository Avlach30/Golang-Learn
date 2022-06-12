package ratingsmap

//* Map: Like array and slice, but the key can customized

//* Define a function which is returned map with key of string and value of integers 
func MovieRatings() map[string]int {

	//* Declarating a map which is insert an element indirectly
	var ratings map[string]int
	ratings = map[string]int{} //* Must initialize firstly before insert a new element

	ratings["spiderman"] = 8
	ratings["dr-strange"] = 9
	ratings["kkn"] = 10
	ratings["x-men"] = 8

	return ratings
}