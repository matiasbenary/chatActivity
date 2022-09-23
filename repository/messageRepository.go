package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/matiasbenary/chatActivity/models"
)

type Message struct {
	ID     uuid.UUID `json:"id"`
	Value  string    `json:"value"`
	SendAt string    `json:"send_at"`
	UserId string    `json:"user_id"`
	RoomId string    `json:"room_id"`
}

func (message *Message) GetId() string {
	return message.ID.String()
}

func (message *Message) GetValue() string {
	return message.Value
}

func (message *Message) GetSendAt() string {
	return message.SendAt
}

func (message *Message) GetUserId() string {
	return message.UserId
}

func (message *Message) GetRoomId() string {
	return message.RoomId
}

// m Message

type MessageRepository struct {
	Db *sql.DB
}

//repo *MessageRepository

func (repo *MessageRepository) AddMessage(message models.Message) {
	stmt, err := repo.Db.Prepare("INSERT INTO message(id, value ,user_id , room_id,send_at) values(?,?,?,?,?)")
	checkErr(err)
	_, err = stmt.Exec(message.GetId(), message.GetValue(), message.GetUserId(), message.GetRoomId(), message.GetSendAt())
	checkErr(err)
}

func (repo *MessageRepository) FindRoomByID(id string) []models.Message {

	rows, err := repo.Db.Query("SELECT id, value ,user_id , room_id,send_at  FROM message where room_id = ? order by send_at desc LIMIT 10", id)

	if err != nil {
		log.Println(err)
		return nil
	}

	var msjs []models.Message

	for rows.Next() {
		var msj Message
		if err := rows.Scan(&msj.ID, &msj.Value, &msj.UserId, &msj.RoomId, &msj.SendAt); err != nil {
			log.Println(err)
			if err == sql.ErrNoRows {
				return nil
			}
			panic(err)
		}
		msjs = append(msjs, &msj)

	}

	return msjs
}

func NewMessage(value string, userId string, roomId string) *Message {
	return &Message{
		ID:     uuid.New(),
		Value:  value,
		UserId: userId,
		RoomId: roomId,
		SendAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}
