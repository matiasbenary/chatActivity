package models

type Message interface {
	GetId() string
	GetValue() string
	GetSendAt() string
	GetUserId() string
	GetRoomId() string
}

type Activity interface {
	GetId() string
	GetName() string
	GetCant() string
}

type MessageRepository interface {
	AddMessage(message Message)
	FindRoomByID(id string) []Message
	MoreMessage(id string) []Message
	LastMessage(roomId string) []Message
	DeleteMessage(id string) string
	LastActivity() []Activity
}
