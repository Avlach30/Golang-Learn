package studentreport

func CalculationReport(value int) string {
	//* Define conditional structure with if-else type
	var result string

	if (value >= 90) {
		result = "Selamat atas prestasinya! dipertahankan ya"
	} else if (value >= 80) {
		result = "Prestasi yang bagus, lebih bagus lagi kalau ditingkatkan belajarnya"
	} else if (value >= 70) {
		if (value == 75) { //* Nested conditional (if inside if)
			result = "Kamu beruntung, bisa diambang batas minimal"
		} else {
			result = "Lebih giat belajar lagi yaa"
		}
	} else if (value >= 60) {
		result = "Ditunda dulu kesenangannya"
	} else {
		result = "Siap-siap remedial yaa"
	}

	return result
}