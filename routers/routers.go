package routers

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"work4/api"
	"work4/middleware"
)

func NewRouters() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("it's_a_secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("/v1")//进行路由器分组
	 {
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.Userlogin)
		authed := v1.Group("/user")//路由分组
		authed.Use(middleware.JWT())//中间件函数添加到链中。
		{
			video := authed.Group("/video")
			{
				video.POST("upload", api.VideoUpload)
				video.POST("comment", api.Comment)
				video.GET("rank", api.Rank)
			}
			chat := authed.Group("/chat")
			{
				chat.POST("chating", api.UserChat)
				chat.GET("search", api.HistorySearch)
			}
			search := authed.Group("/search")
			{
				search.GET("searchall", api.Searchall)
			}
		}
	 }
	 return r
}