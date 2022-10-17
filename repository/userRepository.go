package repository

import (
	"database/sql"
	"log"

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

func (user *User) GetUUID() uuid.UUID {
	return user.ID
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

func (repo *UserRepository) AddUser(user models.User) models.User {
	userFind := repo.FindUserByEmail(user.GetEmail())

	if userFind != nil {
		println(userFind)
		return userFind
	}

	stmt, err := repo.Db.Prepare("INSERT INTO user(id, name,email,role_id) values(?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(user.GetId(), user.GetName(), user.GetEmail(), user.GetRoleId())
	println(res)
	checkErr(err)
	log.Printf("check %s", res)
	userFind = repo.FindUserByEmail(user.GetEmail())
	println(userFind)
	return userFind
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
	println(email)
	row := repo.Db.QueryRow("SELECT id, name,email,role_id FROM user where email = ? LIMIT 1", email)

	var user User

	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.RoleId); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return &user

}
