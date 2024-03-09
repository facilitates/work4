package serializer

import (
	"time"
	"work4/models"
)

type UserChatHistory struct {
	SendTime 	time.Time 	`json:"time"`
	Content  	string		`json:"content"`
	Receiver	string		`json:"receiver"`
}

func BuildMessagesList(items []models.ChatHistory) (chathistorys []UserChatHistory){
	for _, item := range items {
		chathistory := BuildMessage(item)
		chathistorys = append(chathistorys, chathistory)
	}
	return chathistorys
}

func BuildMessage(item models.ChatHistory) UserChatHistory{
	return UserChatHistory {
		SendTime: item.CreatedAt,
		Content: item.Context,
		Receiver: item.Receiver,
	}
}