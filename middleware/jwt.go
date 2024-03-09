package middleware

import (
	"fmt"
	"time"
	"work4/pkg/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context){
		code := 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		}else{
			claim, err := utils.ParseToken(token)
			username := c.Param("username")
			fmt.Println(username)
			if err != nil {
				code = 403 //无权限吗，token是无权限的，是假的
			}else if time.Now().Unix() > claim.ExpiresAt {
				code = 401 // token无效
			}else if claim.UserName != username {
				// fmt.Println(claim.UserName)
				code = 403 //无权限
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"status" : code,
				"msg" : "token解析错误",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}