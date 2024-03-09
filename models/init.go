package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"github.com/streadway/amqp"
	"log"
)

var DB *gorm.DB

var Redisdb *redis.Client

var Conn *amqp.Connection

func DbInit(connstring string, RedisAddr string, RedisPW string, RedisDbName int, RabbitMqPath string) {
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		fmt.Println(err)
		panic("Mysql数据库连接错误")
	}
	fmt.Println("数据库连接成功")
	db.LogMode(true) //启动模式
	if gin.Mode() == "release" {
		db.LogMode(false)
	} //如果是发行版本
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(20) // 设置连接池
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPW,
		DB:       RedisDbName,
	})
	_, err = Redisdb.Ping().Result()
	if err != nil {
		fmt.Println(err)
		panic("Redis数据库连接错误")
	}
	conn, err := amqp.Dial(RabbitMqPath)
    if err != nil {
		fmt.Println(err)
        panic("rabbitMQ连接失败")
    }
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	err = ch.ExchangeDeclare(
		"exchange",     // exchange 名称
		amqp.ExchangeDirect, // exchange 类型
		false,         // 不持久化
		false,         // 不自动删除
		false,         // 不等待服务器响应
		false,         // 无内部队列
		nil,           // 参数
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}
	Conn = conn
	migration()
}
