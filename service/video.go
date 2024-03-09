package service

import (
	"fmt"
	"work4/models"
	"work4/serializer"
	"strconv"
	//"time"
	//"github.com/gin-gonic/gin"
	//"path/filepath"
	"mime/multipart"
)

type UploadVideoService struct {
    Video   *multipart.FileHeader `form:"video" binding:"required"`
	Title      string    `form:"title"` // 视频标题
    Description string   `form:"description"` // 视频描述
	Type		string   `form:"type"`
}

func (service *UploadVideoService) UploadVideo(id uint, username string, filepath string) serializer.Response {
	var user models.User
	models.DB.First(&user, id)
	video := models.Video{
		User : user,
		Uid : user.ID,
		Title : service.Title,  
    	Description : service.Description,
    	FilePath : filepath,
		Type : service.Type,
	}
	err := models.DB.Create(&video).Error
	fmt.Println(err)
	code := 200
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg: "视频上传失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg: "视频上传成功",
	}
}

func ShowVideoDetail(videoid string) serializer.Response {
	videoId, _ := strconv.ParseUint(videoid, 10, 64)
	var video models.Video
	models.DB.First(&video, videoId)
	return serializer.Response{
		Status: 200,
		Data: serializer.BuildVideo(video),
	}
}