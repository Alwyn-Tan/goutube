package service

import (
	"goutube/model"
	"goutube/serializer"
)

type UpdateVideoService struct {
	Title       string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Description string `form:"description" json:"description" binding:"max=100"`
}

func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Video not found",
			Error:  err.Error(),
		}
	}

	video.Title = service.Title
	video.Description = service.Description
	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Update video failed",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
		Msg:  "Update video successfully",
	}
}
