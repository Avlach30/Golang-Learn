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

	println("=============")

	target := citizen[0]["address"] //* Get only address value from index 0 in citizen slice
	println(target)

	name := "doddy"
	address := "asia afrika, 7"
	

	index := findIndex(name, address, citizen)
	fmt.Println(citizen[index])
}

func findIndex(targetNameValue string, targetAddressValue string, arrSouce []map[string]string) int {
	targetIdx := -1

	for idx, item := range(arrSouce) {
		if (item["name"] == targetNameValue && item["address"] == targetAddressValue) {
			targetIdx = idx
		}
	}

	return targetIdx
}