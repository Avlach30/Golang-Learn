package studentreport

func StudentReport(examValue int) string {
	var report string

	if (examValue >= 90) {
		report = "Selamat! pertahankan ya prestasinya"
	} else if (examValue >= 80) {
		report = "Sudah mulai bagus nih prestasinya"
	} else if (examValue >= 70) {
		report = "Lebih giat lagi ya belajarnya"
	} else if (examValue >= 60) {
		report = "Dikurangi dulu mainnya, lebih fokus belajarnya ya"
	} else {
		report = "Gapapa, coba lagi ya di ujian remedial"
	}

	return report
}