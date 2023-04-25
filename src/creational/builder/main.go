package main

import "fmt"

type Person struct {
	name, address, pin string

	workAddress, company, position string
	salary                         int
}

// ------------------ PersonBuilder ---------------------------
type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (a *PersonAddressBuilder) At(address string) *PersonAddressBuilder {
	a.person.address = address
	return a
}

func (a *PersonAddressBuilder) WithPostalCode(pin string) *PersonAddressBuilder {
	a.person.pin = pin
	return a
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (j *PersonJobBuilder) As(position string) *PersonJobBuilder {
	j.person.position = position
	return j
}

func (j *PersonJobBuilder) For(company string) *PersonJobBuilder {
	j.person.company = company
	return j
}

func (j *PersonJobBuilder) In(companyAddress string) *PersonJobBuilder {
	j.person.workAddress = companyAddress
	return j
}

func (j *PersonJobBuilder) WithSalary(salary int) *PersonJobBuilder {
	j.person.salary = salary
	return j
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func RunBuilderFacet() {
	pb := NewPersonBuilder()
	pb.Lives().
		At("Bangalore").
		WithPostalCode("560102").
		Works().
		As("Software Engineer").
		For("IBM").
		In("Bangalore").
		WithSalary(150000)

	person := pb.Build()

	fmt.Println(person)
}

func main() {
	RunBuilderFacet()
}
