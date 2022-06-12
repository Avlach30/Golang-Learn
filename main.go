package main

import (
	"fmt"
	citizen "golang-basic-fundamental/citizen"
)

func main() {
	citizen := citizen.Citizens()
	fmt.Println(citizen)

	println("==================")

	for _, item := range(citizen) {
		fmt.Println(item)
	}
}