package main //* configure package in main, so this file can be executable in main file apps (modularization)

import "fmt"

//* define function returned a string value
func introduction() string {
	return fmt.Sprintf("Hello, my name is %s, and im %d years old", "Ahmad muzakki", 21)
}