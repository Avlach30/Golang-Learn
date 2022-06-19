package main

import "fmt"

/* 
	* Interface is container of defined method, but only definition. Has no content when define it
	* Interface only can be use if it already has content inside
	* To populate method of interface itself, make sure that defined method is assign to method in Struct
	* For assign method of interface to Struct, make sure name, params, and data type when returned it must same
*/
type BangunDatar interface {
	Luas() int
}

type Persegi struct {
	Sisi int
}

//* Populate Luas() method from interface to Persegi struct
func (persegi Persegi) Luas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar int
}

//* Populate Luas() method from interface to PersegiPanjang struct
func (persegiPanjang PersegiPanjang) Luas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

//! Luas() method already defined in 2 struct
/* 
	* Defining function which is accept parameter from interface 
	* So Luas() method which is already define in interface can passing 
			* to different Struct which is have same defined method from Interface
*/
func HitungLuas(bangunDatar BangunDatar) int {
	return bangunDatar.Luas()
}

//* Calling HitungLuas() method in main function
func main()  {
	persegiPanjang := PersegiPanjang{Panjang: 12, Lebar: 4}
	luas := HitungLuas(persegiPanjang)
	fmt.Printf("Luas persegi panjang : %d", luas)
	println("")
	persegi := Persegi{Sisi: 8}
	luas = HitungLuas(persegi)
	fmt.Printf("Luas persegi : %d", luas)
}
