package service

import (
	"fmt"
	"log"
	"work4/models"
	"work4/serializer"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
)

type SearchChatHistoryService struct{
	Receiver  string	`form:"receiver" binding:"required"`	
	StartTime string	`form:"starttime" binding:"required"`
	EndTime	  string	`form:"endtime" binding:"required"`
}

func SendMessageToRabbitMQ(sendername, receivername, message string) (uint, serializer.Response) {
	ch, err := models.Conn.Channel()
	if err != nil {
		log.Println(err)
		return 0, serializer.Response{
			Status: 400,
			Msg: "通道创建失败",
		}
	}
	defer ch.Close()
	var chatlist models.ChatList
	result := models.DB.Where("(user1 = ? AND user2 = ?) OR (user1 = ? AND user2 = ?)", sendername, receivername, receivername, sendername).First(&chatlist)
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error){
			makeChatList := models.ChatList{
				User1: sendername,
				User2: receivername,
			}
			models.DB.Create(&makeChatList)
		}else{
			fmt.Println(result.Error)
			return 0, serializer.Response{
				Status: 400,
				Msg: "查询数据库失败",
			}
		}
	}
	// 发送消息到 exchange
	err = ch.Publish(
		"exchange",    // exchange 名称
		fmt.Sprintf("%d", chatlist.ID),          // routing key
		false,       // 强制发送
		false,       // 立即发送
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("%s %s", receivername, message)),
		},
	)
	if err != nil {
		log.Println(err)
		return 0, serializer.Response{
			Status: 400,
			Msg: "发送失败",
		}
	}
	return chatlist.ID, serializer.Response{
		Status: 200,
		Msg: "发送成功",
	}
}

func SaveChatHistory(uid uint, sendername string, receivername string, message string) serializer.Response{
	chathistory := models.ChatHistory{
		Uid : uid,
		Context: message,
		Sender: sendername,
		Receiver: receivername,
	}
	err := models.DB.Create(&chathistory).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg: "聊天记录保存失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg: "聊天记录保存成功",
	}
}

func (service *SearchChatHistoryService) SearchChatHistory(sendername string) serializer.Response{
	receivername := service.Receiver
	var chatlist models.ChatList
	var count int
	result := models.DB.Where("(user1 = ? AND user2 = ?) OR (user1 = ? AND user2 = ?)", sendername, receivername, receivername, sendername).First(&chatlist)
	if result.Error != nil {
		return serializer.Response{
			Status: 400,
			Msg: "查询数据库失败",
		}
	}
	var chathistory []models.ChatHistory
	result = models.DB.Where("uid = ? AND created_at BETWEEN ? AND ?", chatlist.ID, service.StartTime, service.EndTime).Find(&chathistory).Count(&count)
	fmt.Println(chathistory)
	if result.Error != nil {
		return serializer.Response{
			Status: 400,
			Msg: "查询数据库错误",
		}
	}
	return serializer.BuildListResponse(serializer.BuildMessagesList(chathistory), uint(count))
}