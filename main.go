package main

import (
	"fmt"
	studentreport "golang-basic-fundamental/student-report"
)

func main() {
	//* Calling a function which is returned multiple value, with define all returned value
	reportAndiCategory, reportAndiDescription := studentreport.StudentReport(87)

	fmt.Println(reportAndiCategory)
	fmt.Println(reportAndiDescription)

	println("===================")

	//* Calling a function which is returned multiple value, but only define a certain returned value
	reportBoobyCategory, _ := studentreport.StudentReport(98)
	fmt.Println(reportBoobyCategory)
}