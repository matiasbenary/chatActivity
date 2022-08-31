package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/matiasbenary/chatActivity/models"
)

type User struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	RoleId string    `json:"role_id"`
}

func (user *User) GetId() string {
	return user.ID.String()
}

func (user *User) GetName() string {
	return user.Name
}

func (user *User) GetEmail() string {
	return user.Email
}

func (user *User) GetRoleId() string {
	return user.Name
}

type UserRepository struct {
	Db *sql.DB
}

func (repo *UserRepository) AddUser(user models.User) {
	stmt, err := repo.Db.Prepare("INSERT INTO user(id, name,email,role_id) values(?,?,?,?)")
	checkErr(err)

	_, err = stmt.Exec(user.GetId(), user.GetName(), user.GetEmail(), user.GetRoleId())
	checkErr(err)
}

func (repo *UserRepository) RemoveUser(user models.User) {
	stmt, err := repo.Db.Prepare("DELETE FROM user WHERE id = ?")
	checkErr(err)

	_, err = stmt.Exec(user.GetId())
	checkErr(err)
}

func (repo *UserRepository) FindUserById(id string) models.User {
	println("FindUserById")
	row := repo.Db.QueryRow("SELECT id, name FROM user where id = ? LIMIT 1", id)

	var user User

	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.RoleId); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return &user

}

func (repo *UserRepository) FindUserByEmail(email string) models.User {
	println("FindUserByEmail")
	row := repo.Db.QueryRow("SELECT id, name FROM user where email = ? LIMIT 1", email)

	var user User

	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.RoleId); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return &user

}
