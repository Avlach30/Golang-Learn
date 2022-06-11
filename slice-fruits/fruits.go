package slicefruits

func TropicalFruits() []string {
	//* Slice like an array, but without define a length when declaration, so flexible

	//* Insert an element in slice directly
	tropicalFruits := []string{"Pineapple", "Mango", "Mangosteen", "Jackfruit", "Lychee", "Papaya"}

	return tropicalFruits
}

func WesternFruits() []string {

	//* Define a slice, but insert an element indirectly
	var westernFruits []string

	westernFruits = append(westernFruits, "Apple")
	westernFruits = append(westernFruits, "Berry")
	westernFruits = append(westernFruits, "Lemon")
	westernFruits = append(westernFruits, "Citrus")
	westernFruits = append(westernFruits, "Grape")

	return westernFruits
}