package managementHajj

type Person struct {
	Name string
	Age int
}

/* 
	* Define a method for change Person name globally with pointer 
	* accept parameter from struct
	* parameter type must dereferencing for reassign value globally
*/
func (person *Person) ChangeHajjStatus() {
	person.Name = "Hj. " + person.Name
}