package models

type User interface {
	GetId() string
	GetName() string
	GetEmail() string
	GetRoleId() string
}

type UserRepository interface {
	AddUser(user User)
	RemoveUser(user User)
	FindUserById(ID string) User
	FindUserByEmail(Email string) User
}
