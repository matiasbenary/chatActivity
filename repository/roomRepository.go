package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/matiasbenary/chatActivity/models"
)

type Room struct {
	ID      uuid.UUID
	Name    string
	Private bool
}

func (room *Room) GetId() string {
	return room.ID.String()
}

func (room *Room) GetName() string {
	return room.Name
}

func (room *Room) GetPrivate() bool {
	return room.Private
}

type RoomRepository struct {
	Db *sql.DB
}

func (repo *RoomRepository) AddRoom(room models.Room) {
	println("AddRoom")
	stmt, err := repo.Db.Prepare("INSERT INTO room(id, name, private) values(?,?,?)")
	checkErr(err)

	_, err = stmt.Exec(room.GetId(), room.GetName(), room.GetPrivate())
	checkErr(err)
}

func (repo *RoomRepository) FindRoomByName(name string) models.Room {
	println(name)
	println("FindRoomByName")
	row := repo.Db.QueryRow("SELECT id, name, private FROM room where name=? LIMIT 1", name)
	println("qtest")
	var room Room

	if err := row.Scan(&room.ID, &room.Name, &room.Private); err != nil {

		println("Asd")
		println(err)
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return &room

}

func checkErr(err error) {
	if err != nil {
		println("123asd")
		println(err)
		panic(err)
	}
}
