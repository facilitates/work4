package api

import (
	//"mime/multipart"
	//"net/http"
	//"path/filepath"
	//"work4/models"
	"work4/pkg/utils"
	"work4/serializer"
	"work4/service"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	//"fmt"
	//"io"
	//"log"
)

func UserRegister(c *gin.Context){
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister);err == nil{
		res := userRegister.Register()
		c.JSON(200, res)
	}else{
		c.JSON(400, ErrorResponse(err))
	}
}

func Userlogin(c *gin.Context){
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin);err == nil{
		res := userLogin.Login()
		c.JSON(200, res)
	}else{
		c.JSON(400, ErrorResponse(err))
	}
}

func UserUploadAvatar(c *gin.Context){
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, ErrorResponse(err))
	}
	if utils.ParseAvatarExt(file.Filename) {
		c.JSON(400, serializer.Response{
			Status:400,
			Msg:"文件类型错误",
		})
	}
	filepath := "./upload/avatar"+file.Filename
	if err := c.SaveUploadedFile(file, filepath); err != nil{
		c.JSON(400, gin.H{
			"error":"保存头像失败",
		})
	}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	res := service.UploadAvatar(claim.ID)
	
	c.JSON(200, res)
}
// func UploadAmatar(c *gin.Context){
//     file, err := c.FormFile("file")
//     if err != nil {
//         c.JSON(400, ErrorResponse(err))
//     }
//     res := service.CheckAvatar(file)
// 	filePath := fmt.Sprintf("./upload/avatar", file.Filename)
//     if err := c.SaveUploadedFile(file, filePath); err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "保存头像失败"
// 		})
//         log.Println("保存头像失败", err)
//         return
//     }
//     video := models.Video{
//         Title:      "Example Video Title", 
//         FilePath:   filePath,
//     }
//     if result := models.DB.Create(&video); result.Error != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save video info"})
//         return
//     }
//     c.JSON(http.StatusCreated, gin.H{"message": "视频上传成功", "videoId": video.ID})
// }

// func UpdateAmatar(c *gin.Context){
//     file, err := c.FormFile("file")
//     if err != nil {
//         c.JSON(400, ErrorResponse(err))
//     }
//     res := service.CheckAvatar(file)
// 	filePath := fmt.Sprintf("./upload/avatar", file.Filename)
//     if err := c.SaveUploadedFile(file, filePath); err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "保存头像失败"
// 		})
//         log.Println("保存头像失败", err)
//         return
//     }
//     video := models.Video{
//         Title:      "Example Video Title", 
//         FilePath:   filePath,
//     }
//     if result := models.DB.Create(&video); result.Error != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save video info"})
//         return
//     }
//     c.JSON(http.StatusCreated, gin.H{"message": "视频上传成功", "videoId": video.ID})
// }