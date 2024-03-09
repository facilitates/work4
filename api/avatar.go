package api

import (
	"fmt"
	"work4/pkg/utils"
	"work4/service"
	"github.com/gin-gonic/gin"
)

func UploadAvatar(c *gin.Context) {
	var uploadAvatar service.UploadAvatarService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err1 := c.ShouldBind(&uploadAvatar);err1 == nil{
		file := uploadAvatar.Avatar
		if utils.ParseAvatarExt(file.Filename){
			c.JSON(400, nil)
		}else{
			filepath := "./upload/avatar/"+claim.UserName+"/"+file.Filename
			if err2 := c.SaveUploadedFile(file, filepath); err2 != nil {
				fmt.Println(err2)
				c.JSON(400, err2)
			}else{
				res := uploadAvatar.UploadAvatar(claim.ID, filepath)
				c.JSON(200, res)
			}
		}
	}else{
		fmt.Println(err1)
		c.JSON(200, err1)
	}
}
