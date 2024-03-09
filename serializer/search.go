package serializer

import (
	"work4/models"
)

type SearchHistory struct {

}

func BuildSearchVideosResponse(items []models.Video) []Video {
	var videos []Video
	for _, item := range items{
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}

func BuildSearchUsersResponse(items []models.User) []User {
	var users []User
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}
