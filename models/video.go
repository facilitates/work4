package models

import (
    "gorm.io/gorm"
    "time"
)

// Video 结构体表示视频实体
type Video struct {
    gorm.Model // GORM的Model包含了ID, CreatedAt, UpdatedAt, DeletedAt字段
    Title      string    `gorm:"size:255;not null"` // 视频标题
    Description string   `gorm:"type:text;not null"` // 视频描述
    FilePath   string    `gorm:"not null"` // 视频文件的存储路径
    UserID     uint      `gorm:"not null"` // 上传者的用户ID
    // 可以根据需要添加更多字段，比如播放次数、点赞数、状态等
    ViewCount int       `gorm:"default:0"` // 视频的观看次数
    // 这里可以考虑加上视频的封面，以及其他你觉得需要的信息
    CoverImagePath string // 视频封面图片的路径，可选
    CreatedAt      time.Time // 创建时间，也可以用 gorm.Model 里的 CreatedAt
}