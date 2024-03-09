package middleware

import (
	// "context"
	"fmt"
	// "net/http"
	// "time"
	"work4/models"
	"github.com/gin-gonic/gin"
	// "github.com/go-redis/redis/v8"
)

func COUNT() gin.HandlerFunc{
		return func(c *gin.Context){
		videoId := c.Param("videoid")
		err := models.IncreaseClick(videoId)
		if err != nil {
			fmt.Println(err)
		}
		c.Next()
	}
}