package models

import (
    "time"
)

// Message 表示聊天消息的结构体。
type Message struct {
    ID        uint      `gorm:"primaryKey"` // 使用gorm标签为ORM框架提供指令
    SenderID  uint      `gorm:"index"` // 发送者ID
    ReceiverID uint     `gorm:"index"` // 接收者ID
    Content   string    // 消息内容
    CreatedAt time.Time // 发送时间
}

type ChatSession struct {
    ID          uint      `gorm:"primaryKey"`
    Participant1 uint     `gorm:"index"` // 参与者1ID
    Participant2 uint     `gorm:"index"` // 参与者2ID
    CreatedAt    time.Time
    LastMessage  string    // 存储最后一条消息的摘要
    UpdatedAt    time.Time // 最后一条消息的时间
}