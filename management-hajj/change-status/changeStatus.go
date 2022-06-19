package managementHajj

import managementHajj"golang-basic-fundamental/management-hajj/entity"


/* 
	* Define a method for change Person name globally with pointer 
	* accept parameter from struct
	* parameter type must dereferencing for reassign value globally
*/
func ChangeHajjStatus (person *managementHajj.Person) {
	person.Name = "Hj. " + person.Name
}