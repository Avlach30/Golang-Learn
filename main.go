package main //* must define 'main' package, because go app can be runned and builded

import "fmt"

func main() {
	fmt.Println("hello, i'm learning golang") //* for print data to console/terminal
	sentence := introduction() //* define variables
	fmt.Println(sentence)
}

//! Before develop an app with golang, make sure initialize app directory firstly with command 'go mod init 'name-app''
//* Use command 'go run fileName.go' for running go program
//* Use command 'go build' for build a executable (.exe) go progam
	//* Name of builded program depend of module name has initialized before
	//* Type executable file in terminal for running an builded go program