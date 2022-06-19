package main

import "fmt"

type Gamer struct {
	Name string
	Games []string
}

//!TODO: Create a method with istance from Gamer, which is accept 1 parameter is string. For append to Games
func (gamer *Gamer) AddGame(game string) {
	//* Implement referencing to struct for reassign value globally
	gamer.Games = append(gamer.Games, game) //* Append parameter to Games array of string
}

func main() {
	gamerAhmad := Gamer{Name: "Ahmad"}
	fmt.Printf("List game before addedd: %v", gamerAhmad.Games)
	println("")
	gamerAhmad.AddGame("Minecraft")
	gamerAhmad.AddGame("Fifa 2021") //* Call method for append to Games globally
	fmt.Printf("List game after addedd: %v", gamerAhmad.Games)
}