package main

import "fmt"

type Employee struct {
	Name    string
	Gender  string
	Address *Address
}
type Address struct {
	StreetName string
	City       string
}

func (add *Address) DeepCopy() *Address {
	return &Address{
		StreetName: add.StreetName,
		City:       add.City,
	}
}

func Prototype_Copy() {

	employee1 := Employee{
		Name:   "John",
		Gender: "Male",
		Address: &Address{
			StreetName: "Baker Street",
			City:       "London",
		},
	}

	employee2 := employee1
	employee2.Name = "Surya"
	employee2.Address = employee1.Address.DeepCopy()
	employee2.Address.StreetName = "Marine Drive"
	employee2.Address.City = "Mumbai"
	fmt.Println(employee1, employee1.Address)
	fmt.Println(employee2, employee2.Address)

}

func main() {
	Prototype_Copy()
}
