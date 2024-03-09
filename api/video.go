package api

import (
	// "fmt"
	"fmt"
	"work4/models"
	"work4/pkg/utils"
	"work4/service"
	"github.com/gin-gonic/gin"
	// "work4/serializer"
)

func UploadVideo(c *gin.Context) {
	var uploadVideo service.UploadVideoService
    claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err1 := c.ShouldBind(&uploadVideo);err1 == nil{
		file := uploadVideo.Video
		if utils.ParseVideoExt(file.Filename) {
			c.JSON(400, nil)
		}else{
			filepath := "./upload/video/"+claim.UserName+"/"+file.Filename
				if err2 := c.SaveUploadedFile(file, filepath); err2 != nil{
					c.JSON(400, ErrorResponse(err2))
				}else{
					res := uploadVideo.UploadVideo(claim.ID, claim.UserName, filepath)
					c.JSON(200, res)
				}
			}
	}else{
		c.JSON(200, err1)
	}
}

func VideoRank(c *gin.Context){
	key := "videoclick"
	ranklist := models.Redisdb.ZRevRangeWithScores(key, 0, -1)
	count, err := models.Redisdb.ZCard(key).Result()
	if err != nil {
		fmt.Println(err)
    	c.JSON(404, err)
	}
	res := service.RankVideoList(ranklist, int(count))
	c.JSON(200, res)
}

func ShowVideo(c *gin.Context) {
	videoid := c.Param("videoid")
	fmt.Println()
	fmt.Println(videoid)
	fmt.Println()
	res := service.ShowVideoDetail(videoid)
	c.JSON(200, res)
}