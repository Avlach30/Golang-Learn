package main

import (
	"fmt"
	citizen "golang-basic-fundamental/citizen"
)

func main() {
	citizen := citizen.Citizens()
	fmt.Println(citizen)

	println("==================")

		//* Insert a new map with multiple key-value pair in existing slice
		citizen = append(citizen, map[string]string{
			"name": "frans", "address": "pahlawan street 4",
		})
	
		fmt.Println(citizen)
	
		println("===================")

	for _, item := range(citizen) {
		fmt.Println(item)
	}
}