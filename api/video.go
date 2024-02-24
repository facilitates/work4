package api

import (
	"mime/multipart"
	"net/http"
	"path/filepath"
	"work4/models"
	"work4/service"

	"github.com/gin-gonic/gin"

	// "gorm.io/gorm"
	"fmt"
	"io"
	"log"
)

func UploadVideo(c *gin.Context){
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(400, ErrorResponse(err))
    }
    err, url := UploadAvatarService(file)

    // 验证视频类型
    ext := filepath.Ext(header.Filename)
    if ext != ".mp4" && ext != ".avi" { // 添加或修改以支持更多视频格式
        c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported video format"})
        return
    }

    // 定义视频保存的路径
    filePath := fmt.Sprintf("./upload/video/%s", header.Filename)

    // 保存文件
    if err := c.SaveUploadedFile(header, filePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the video"})
        log.Println("Failed to save the video:", err)
        return
    }

    // 将视频信息保存到数据库
    video := models.Video{
        Title:      "Example Video Title", 
        FilePath:   filePath,
    }
    if result := models.DB.Create(&video); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save video info"})
        return
    }

    // 上传成功
    c.JSON(http.StatusCreated, gin.H{"message": "视频上传成功", "videoId": video.ID})
}