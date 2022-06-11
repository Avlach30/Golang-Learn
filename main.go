package main

func main() {
	result := "Saya sedang mencari hadiah"

	//* Define a plain looping with limit of number
	for index:= 1; index < 20; index++ {
		if (index == 8) { //* Add if conditional with certain index
			result = "Selamat! anda menemukan hadiahnya"
			println(result)
		}

		result = "Silahkan tunggu"
		println(result)
	}
}