package api

import (
	"net/http"
	"time"
	"work4/conf"
	"work4/service"
    "fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	// "github.com/streadway/amqp"
	"log"
	"sync"
)

func UserChat(c *gin.Context){
	username := c.Param("username")
	var (
		upgrader   = websocket.Upgrader{
			ReadBufferSize: 1024,
			WriteBufferSize: 1024,
			HandshakeTimeout: 5 * time.Second,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		lock       sync.Mutex
	)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// 加锁，确保并发安全
	lock.Lock()
	conf.UserConns[username] = conn
	lock.Unlock()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		var receivername, message string
		if _, err := fmt.Sscanf(string(msg), "%s %s", &receivername, &message); err != nil {
			continue
		}
		// 将消息发送到 RabbitMQ
		uid, res := service.SendMessageToRabbitMQ(username, receivername, message)
		c.JSON(200, res)
		// 根据接收方用户名直接发送消息到目标 WebSocket 连接
		lock.Lock()
		recvConn, ok := conf.UserConns[receivername]
		lock.Unlock()
		if !ok {
			log.Printf("User %s is not online\n", receivername)
			continue
		}
		if err := recvConn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Println(err)
		}
		service.SaveChatHistory(uid, username, receivername, message)
	}
	lock.Lock()
	delete(conf.UserConns, username)
	lock.Unlock()
}

func ChatHistory(c *gin.Context){
	var searchChatHistory service.SearchChatHistoryService
	if err := c.ShouldBind(&searchChatHistory); err != nil {
		log.Println(err)
		c.JSON(400, "参数绑定错误")
	}else{
		username := c.Param("username")
		res := searchChatHistory.SearchChatHistory(username)
		c.JSON(200, res)
	}
}
