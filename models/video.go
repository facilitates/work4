package models

import (
    "github.com/jinzhu/gorm"
)

// Video 结构体表示视频实体
type Video struct {
    gorm.Model // GORM的Model包含了ID, CreatedAt, UpdatedAt, DeletedAt字段
    Title      string    `gorm:"size:255;not null"` // 视频标题
    Description string   `gorm:"type:text;not null"` // 视频描述
    FilePath   string    `gorm:"not null"` // 视频文件的存储路径
    // 可以根据需要添加更多字段，比如播放次数、点赞数、状态等
    // ViewCount int       `gorm:"default:0"` // 视频的观看次数
    // CoverImagePath string 
    Type    string      
    Uid uint `gorm:"not null"`
    User User `gorm:"ForeignKey:Uid"`
}