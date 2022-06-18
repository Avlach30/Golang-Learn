package main

import (
	"fmt"
	studentreport "golang-basic-fundamental/student-report"
)

func main() {
	reportAhmad := studentreport.StudentReport(97)
	
	fmt.Println(reportAhmad)
}