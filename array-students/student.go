package arraystudents

//* Define a function with returned an array of strings
func StudentMales() []string {
	//* Define an array with insert elemen undirectly, but must define a length when declaration
	var studentsMale [6]string
	studentsMale[0] = "ahmad"
	studentsMale[1] = "bobby"
	studentsMale[2] = "chandra"
	studentsMale[3] = "doni"
	studentsMale[4] = "eko"
	studentsMale[5] = "faris"

	return studentsMale[:]
}