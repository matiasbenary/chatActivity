package models

import "github.com/google/uuid"

type User interface {
	GetUUID() uuid.UUID
	GetId() string
	GetName() string
	GetEmail() string
	GetRoleId() string
}

type UserRepository interface {
	AddUser(user User) User
	RemoveUser(user User)
	FindUserById(ID string) User
	FindUserByEmail(Email string) User
}
