package api

import (
	// "fmt"
	// "work4/models"
	// "work4/pkg/utils"
	"work4/service"
	"github.com/gin-gonic/gin"
	// "work4/serializer"
)

func Search(c *gin.Context){
	var search service.SearchService
	if err := c.ShouldBind(&search); err != nil {
		c.JSON(400, "参数绑定错误")
	}else{
		username := c.Param("username")
		res := search.SearchAll(username)
		c.JSON(200, res)
	}
}

func SearchHistory(c *gin.Context){
	username := c.Param("username")
	res := service.GetSearchHistory(username)
	c.JSON(200, res)
}