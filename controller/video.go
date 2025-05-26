package controller

import (
	"github.com/gin-gonic/gin"
	"goutube/service"
)

func CreateVideo(context *gin.Context) {
	createVideoService := service.CreateVideoService{}
	if err := context.ShouldBind(&createVideoService); err == nil {
		res := createVideoService.Create()
		context.JSON(200, res)
	} else {
		context.JSON(200, ErrorResponse(err))
	}
}

func ShowVideo(context *gin.Context) {
	showVideoService := service.ShowVideoService{}
	if err := context.ShouldBind(&showVideoService); err == nil {
		res := showVideoService.Show(context.Param("id"))
		context.JSON(200, res)
	} else {
		context.JSON(200, ErrorResponse(err))
	}
}

func ListVideo(context *gin.Context) {
	listVideoService := service.ListVideoService{}
	if err := context.ShouldBind(&listVideoService); err == nil {
		res := listVideoService.List()
		context.JSON(200, res)
	} else {
		context.JSON(200, ErrorResponse(err))
	}
}

func UpdateVideo(context *gin.Context) {

}

func DeleteVideo(context *gin.Context) {

}
