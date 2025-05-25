package main

import (
	"goutube/conf"
	"goutube/server"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.Init()

	gin.SetMode(os.Getenv("GIN_MODE"))
	r := server.NewRouter()
	r.Run(":3000")
}
