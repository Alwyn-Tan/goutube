package service

import (
	"goutube/model"
	"goutube/serializer"
)

type ListVideoService struct{}

func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Status: 200,
			Msg:    "List video failed",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
