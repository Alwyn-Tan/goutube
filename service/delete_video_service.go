package service

import (
	"goutube/model"
	"goutube/serializer"
)

type DeleteVideoService struct{}

func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Video not found",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Delete(&video, id).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Delete video failed",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Msg: "Delete video successfully",
	}
}
