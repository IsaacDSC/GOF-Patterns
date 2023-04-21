package factory

type User struct {
	name     string
	lastName string
}

func (this *User) User(name string, lastName string) User {
	return User{name: name, lastName: lastName}
}
