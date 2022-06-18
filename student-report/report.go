package studentreport


/*
	* Define a function which is returned more than 1 value
	* Wrapped in parentheses, each returned value can different data type
	* Define a function value depend of order from returned inside function
*/
func StudentReport(examValue int) (resultCategory string, resultDescription string) {
	var category string
	var description string

	if (examValue >= 90) {
		category = "A"
		description = "Selamat ya! dipertahankan prestasinya"
	} else if (examValue >= 80) {
		category = "B"
		description = "Ditingkatkan lagi ya, jangan sampai turun prestasinya"
	} else if (examValue >= 70) {
		category = "C"
		description = "Lebih ditingkatkan lagi intensitas belajarnya ya"
	} else if (examValue >= 60) {
		category = "D"
		description = "Fokus belajar yuk, jangan main terus"
	} else {
		category = "E"
		description = "Gapapa dicoba lagi ya di ujian remedial nanti"
	}


	return category, description
}