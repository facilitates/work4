package service

import (
	"fmt"
	"work4/models"
	// "work4/pkg/utils"
	"work4/serializer"
	// "work4/service"
)

type CommentService struct {
	CommentText string 	`form:"comment" json:"comment" binding:"required"`
	VideoId 	uint 	`form:"video_id" json:"video_id" binding:"required"`
	ParentId 	uint 	`form:"parent_id" json:"parent_id" binding:"-"`
}

func (service *CommentService) Comment(userid uint, username string) serializer.Response {
	var video models.Video
	fmt.Printf("%T", service.VideoId)
	// fmt.Println(service.ParentId)
	models.DB.First(&video, service.VideoId)
	newComment := models.Comment{
    	// Video    	:	video,
		Uid 		:	video.ID,
		ParentId	:	service.ParentId,
		UserId 		:	userid,
		UserName 	: 	username,
		Content 	:	service.CommentText,
	}
	// fmt.Println(newComment.ParentId)
	models.DB.Create(&newComment)
	return serializer.Response{
		Status : 200,
		Data: newComment,
		Msg : "评论成功",
	}
}