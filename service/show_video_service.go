package service

import (
	"goutube/model"
	"goutube/serializer"
)

type ShowVideoService struct{}

func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Show video failed",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: video,
	}
}
