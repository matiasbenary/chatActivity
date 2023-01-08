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
	Email  string    `json:"email"`
	Name   string    `json:"name"`
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

func (message *Message) GetEmail() string {
	return message.Email
}

func (message *Message) GetName() string {
	return message.Name
}

type Activity struct {
	Id   string `json:"id"`
	Cant string `json:"cant"`
	Name string `json:"name"`
}

func (activity *Activity) GetId() string {
	return activity.Id
}

func (activity *Activity) GetCant() string {
	return activity.Cant
}

func (activity *Activity) GetName() string {
	return activity.Name
}

// m Message

type MessageRepository struct {
	Db *sql.DB
}

//repo *MessageRepository

func (repo *MessageRepository) AddMessage(message models.Message) {
	stmt, err := repo.Db.Prepare("INSERT INTO message(id, value ,user_id , room_id,send_at) values(?,?,?,?,?)")
	checkErr(err)
	_, err = stmt.Exec(message.GetId(), message.GetValue(), message.GetUserId(), message.GetRoomId(), time.Now().Format("2006-01-02 15:04:05"))
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

func (repo *MessageRepository) MoreMessage(roomId string) []models.Message {

	rows, err := repo.Db.Query(`SELECT p.id, p.value,p.send_at,p.user_id,p.room_id,u.name,u.email FROM (SELECT m.id, m.value,m.send_at,m.user_id,m.room_id
								FROM chat.message m
								inner join (select * FROM  message where id = ?) c
								on m.send_at < c.send_at or (m.send_at = c.send_at and c.id != m.id) 
								where m.room_id = c.room_id
								order by send_at desc limit 10) p
								join user u
								on u.id = p.user_id
								order by send_at asc `, roomId)

	if err != nil {
		log.Println(err)
		return nil
	}

	var msjs []models.Message

	for rows.Next() {
		var msj Message
		if err := rows.Scan(&msj.ID, &msj.Value, &msj.SendAt, &msj.UserId, &msj.RoomId, &msj.Name, &msj.Email); err != nil {
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

func (repo *MessageRepository) LastMessage(id string) []models.Message {

	rows, err := repo.Db.Query(`SELECT p.id, p.value,p.send_at,p.user_id,p.room_id,u.name,u.email FROM (SELECT m.id, m.value,m.send_at,m.user_id,m.room_id
								FROM chat.message m
								inner join (select * FROM  message ) c
								on c.id = m.id
								where m.room_id = ?
								order by send_at desc limit 10) p
								join user u
								on u.id = p.user_id
								order by send_at asc `, id)

	if err != nil {
		log.Println(err)
		return nil
	}

	var msjs []models.Message

	for rows.Next() {
		var msj Message
		if err := rows.Scan(&msj.ID, &msj.Value, &msj.SendAt, &msj.UserId, &msj.RoomId, &msj.Name, &msj.Email); err != nil {
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

func (repo *MessageRepository) DeleteMessage(id string) string {

	_, err := repo.Db.Query(`DELETE FROM chat.message
	WHERE id=?;
`, id)

	if err != nil {
		log.Println(err)
		return "fail"
	}

	return "ok"
}

func (repo *MessageRepository) LastActivity() []models.Activity {

	rows, err := repo.Db.Query(`SELECT name,room.id,COUNT(*) 
FROM chat.message
join chat.room
on room.id = message.room_id 
where send_at >= ?
group by name ;`, time.Now().Add(-24*time.Hour))

	if err != nil {
		log.Println(err)
		return nil
	}

	var msjs []models.Activity

	for rows.Next() {
		var msj Activity
		if err := rows.Scan(&msj.Name, &msj.Id, &msj.Cant); err != nil {
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
		SendAt: time.Now().Format(time.RFC3339),
	}
}
