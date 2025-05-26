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
	res := showVideoService.Show(context.Param("id"))
	context.JSON(200, res)
}

func ListVideo(context *gin.Context) {
	listVideoService := service.ListVideoService{}
	res := listVideoService.List()
	context.JSON(200, res)
}

func UpdateVideo(context *gin.Context) {
	updateVideoService := service.UpdateVideoService{}
	if err := context.ShouldBind(&updateVideoService); err == nil {
		res := updateVideoService.Update(context.Param("id"))
		context.JSON(200, res)
	} else {
		context.JSON(200, ErrorResponse(err))
	}
}

func DeleteVideo(context *gin.Context) {
	deleteVideoService := service.DeleteVideoService{}
	res := deleteVideoService.Delete(context.Param("id"))
	context.JSON(200, res)
}
