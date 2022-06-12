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
	
	//* Get index with calling findIndex function
	index := findIndex(name, address, citizen)
	fmt.Println(citizen[index])

	//* Update existing element from slice of maps
	citizen[index]["name"] = "donny"
	citizen[index]["address"] = "merdeka street 17"
	fmt.Println(citizen)

	//* Get all element from slice of maps until index, but exclude index
	fmt.Println(citizen[:index])

	/* 
		* Get all element from slice of maps
		* Started from element which is their index is called index + 1, until last element
	*/
	fmt.Println(citizen[index+1:])

	/* 
		* Delete existing element from slice of maps with index
		* Logic, get all element from slice until called element which is have same index, but exclude element itself
		* After that, separate by comma, then get all element from slice which is started from element have called index + 1, until last element
		* Merge them with three dots operator '...'
	*/
	citizen = append(citizen[:index], citizen[index+1:]...)
	fmt.Println(citizen)
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