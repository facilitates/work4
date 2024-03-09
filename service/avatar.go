package service

import (
	"work4/models"
	"work4/serializer"
	//"time"
	//"github.com/gin-gonic/gin"
	//"path/filepath"
	"mime/multipart"
	
)

type UploadAvatarService struct {
    Avatar   *multipart.FileHeader `form:"avatar" binding:"required"`
}

func (service *UploadAvatarService) UploadAvatar(id uint, filepath string) serializer.Response {
	var user models.User
	models.DB.First(&user, id)
	user.AvatarFilePath = filepath
	models.DB.Save(&user)
	return serializer.Response{
		Status: 200,
		Data: serializer.BuildUser(user),
		Msg: "头像更新完成",
	}
}