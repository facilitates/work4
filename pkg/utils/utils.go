package utils

import (
	"path/filepath"
	"time"
	"github.com/streadway/amqp"
	//"github.com/dgrijalva/jwt-go"
	"fmt"
	"os"
	"work4/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var JWTsecret = []byte("ABAB")

type Claims struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(id uint, username string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		ID:       id,
		UserName: username,
		PassWord: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "work4",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JWTsecret)
	return token, err
}

func ParseToken(token string)(*Claims, error){
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token)(interface{}, error){
		return JWTsecret, nil
	})
	// fmt.Println(err)
	if tokenClaims != nil{
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid{
			return claims, nil
		}
	}
	return nil, err
}

func ParseAvatarExt(filename string) bool {
	fileExt := filepath.Ext(filename)
	if fileExt != ".jpeg" && fileExt != ".jpg" && fileExt != ".png" {
        return true
    }
	return false
}

func ParseVideoExt(filename string) bool {
	fileExt := filepath.Ext(filename)
	if fileExt != ".mp4" {
        return true
    }
	return false
}

func CreateFolder(username string) error {
	avatarfilepath := "./upload/avatar/"+username
	err1 := os.Mkdir(avatarfilepath, 0755)
	if err1 != nil {
		fmt.Println(err1)
		return err1
	}
	videofilepath := "./upload/video/"+username
	err2 := os.Mkdir(videofilepath, 0755)
	if err2 != nil {
		fmt.Println(err2)
		return err2
	}
	return nil
}

func ParseURL(c *gin.Context) string {
	username := c.Param("id")
	return username
}

func MakeQueue(queuename string) (error, *amqp.Channel) {
	ch, err := models.Conn.Channel()
    if err != nil {
		fmt.Println(err,"here")
        return err, nil
    }
    _, err = ch.QueueDeclare(
        queuename, // 队列名称
        false,           // 持久化
        false,           // 自动删除
        false,           // 独占
        false,           // 等待
        nil,              // 其他参数
    )
	if err != nil {
		fmt.Println(err)
        return err, nil
    }
	err = ch.ExchangeDeclare(
        "exchange", // 交换机名称
        "direct",         // 交换机类型
        true,             // 持久化
        false,            // 自动删除
        false,				// 内部交换机
		true,            
        amqp.Table{},             // 其他参数
    )
    if err != nil {
		fmt.Println(err,"here1")
        return err, nil
    }

    // 绑定队列到交换机
    err = ch.QueueBind(
        queuename, // 队列名称
        queuename,        // 路由键
        "exchange", // 交换机名称
        false,
        nil,
    )
    if err != nil {
		fmt.Println(err,"here2")
        return err, nil
    }
    return nil, ch
}

func SendMessage(message string, ch *amqp.Channel, queuename string) error {
	err := ch.Publish(
        "exchange",     // 交换机名称
        queuename,     // 路由键
        false,      // 是否等待
		false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(message),
        })
    if err != nil {
        return err
    }
	return nil
}

func ReceiveMessage(queuename string, ch *amqp.Channel) error {
	msgs, err := ch.Consume(
        queuename, // 队列名称
        "my_consumer", // 消费者标签
        true,    // 自动应答
        false,   //是否本地消费者
        false,   // 是否独占
		false,		// 是否等待
        nil,     // 额外参数
    )
	go func() {
        for d := range msgs {
            // 处理接收到的消息
            fmt.Printf("Received a message: %s", d.Body)
            // 手动确认消息
            if err := ch.Ack(d.DeliveryTag, false); err != nil {
                fmt.Printf("Failed to acknowledge message: %s", err)
            }
        }
    }()
    if err != nil {
       	return err
    }
	return nil
}