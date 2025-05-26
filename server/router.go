package server

import (
	"goutube/controller"
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
		v1.POST("ping", controller.Ping)

		v1.POST("user/register", controller.UserRegister)

		v1.POST("user/login", controller.UserLogin)

		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			auth.GET("user/me", controller.UserMe)
			auth.DELETE("user/logout", controller.UserLogout)
		}

		v1.POST("videos", controller.CreateVideo)
		v1.GET("video/:id", controller.ShowVideo)
		v1.GET("videos", controller.ListVideo)
		v1.PUT("video/:id", controller.UpdateVideo)
		v1.DELETE("video/:id", controller.DeleteVideo)
	}
	return r
}
