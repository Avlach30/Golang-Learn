package entity

import "fmt"

type Citizen struct {
	Name string
	Address string
	NumberPhone int
	IsMarried bool 
}

func (citizen Citizen) Describe() string {
	return fmt.Sprintf(
		"Hello!, my name is %s, i live at %s, i'm %t married. You can contact me at %d", 
		citizen.Name, citizen.Address, citizen.IsMarried, citizen.NumberPhone)
}