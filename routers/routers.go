package routers

import (
	"work4/middleware"
	"work4/api"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouters() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("it's_a_secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("v1/")
	 {
		video := v1.Group("videos")
		video.Use(middleware.COUNT())
		video.GET("/:videoid", api.ShowVideo)
		v1.POST("register", api.UserRegister)
		v1.POST("login", api.Userlogin)
		v1.GET("rank", api.VideoRank)
		authed := v1.Group(":username")
		authed.Use(middleware.JWT())
		{	
			search := authed.Group("/")
			{
				search.GET("search", api.Search)
				search.GET("search/history", api.SearchHistory)
			}
			chat := authed.Group("/chat")
			{
				chat.GET("/:receivername", api.UserChat)
				chat.POST("/history", api.ChatHistory)
			}
			authed.POST("/avatar", api.UploadAvatar)
			video := authed.Group("/video/")
			{
				video.POST(":videoid", api.ShowVideo)
				video.POST("upload", api.UploadVideo)
				video.POST("comment", api.AddComment)
			}
		}
	 }
	 return r
}