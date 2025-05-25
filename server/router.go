package server

import (
	"golang.org/x/net/html/atom"
	"goutube/api"
	"goutube/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		v1.POST("user/register", api.UserRegister)

		v1.POST("user/login", api.UserLogin)

		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}

		v1.POST("video", api.CreateVideo)
		v1.GET("video/:id", api.ShowVideo)
		v1.GET("video", api.ListVideo)
		v1.PUT("video/:id", api.UpdateVideo)
		v1.DELETE("video/:id", api.DeleteVideo)
	}
	return r
}
