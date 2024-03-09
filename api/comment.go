package api

import (
	"fmt"
	"work4/pkg/utils"
	"work4/service"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	var comment service.CommentService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&comment); err != nil {
		fmt.Println(err)
		c.JSON(400, "参数绑定错误")
	}else{
		res := comment.Comment(claim.ID, claim.UserName)
		c.JSON(200, res)
	}
}
