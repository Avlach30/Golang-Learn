package main

import (
	"fmt"
	studentreport "golang-basic-fundamental/student-report"
	calculation "golang-basic-fundamental/calculation"
)

func main() {
	reportAhmad := studentreport.StudentReport(97)
	
	fmt.Println(reportAhmad)

	calculationResult := calculation.Calculation(22, 4, "/")

	fmt.Println(calculationResult)
}