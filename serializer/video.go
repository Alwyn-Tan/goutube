package serializer

import "goutube/model"

type Video struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}

func BuildVideo(item model.Video) Video {
	return Video{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		CreatedAt:   item.CreatedAt.Unix(),
	}
}

func BuildVideos(item []model.Video) []Video {
	var videos []Video
	for _, item := range item {
		videos = append(videos, BuildVideo(item))
	}
	return videos
}
