package models

type Message interface {
	GetId() string
	GetValue() string
	GetSendAt() string
	GetUserId() string
	GetRoomId() string
}

type MessageRepository interface {
	AddMessage(message Message)
	FindRoomByID(id string) []Message
	MoreMessage(id string) []Message
	LastMessage(roomId string) []Message
}
