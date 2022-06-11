package main 

func main() {
	var classify = classifyAge(1999)
	println(classify)
}

func classifyAge(birthYear int) string {
	//* Define a conditional structure with switch-case-default type
	var result string

	switch birthYear {
	case 1990, 1991, 1992, 1993, 1994, 1995: 
	//* Define a multiple case value 
	//* 	but just checking one conditional value because each conditional separated by comma
		result = "Gen millenial humans"
	case 1996, 1997, 1998, 1999, 2000:
		result = "Gen Z humans"
	default: //* Run a default if conditional value not match in defined case value
		result = "Too old or too young to classify"
	} 

	return result
}