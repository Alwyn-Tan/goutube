package service

import (
	"goutube/model"
	"goutube/serializer"
)

type CreateVideoService struct {
	Title       string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Description string `form:"description" json:"description" binding:"max=100"`
}

func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title:       service.Title,
		Description: service.Description,
	}
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "Create video failed",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
