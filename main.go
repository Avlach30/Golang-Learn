package main

import (
	"fmt"
	arraystudents "golang-basic-fundamental/array-students"
)

func main() {
	studentMales := arraystudents.StudentMales() //* Call a function which is returned an array
	studentMalesLen := len(studentMales) //* Get a length of an array

	fmt.Println(studentMales)
	fmt.Println(studentMalesLen)
}