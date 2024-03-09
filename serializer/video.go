package serializer

import(
	"work4/models"
)

type Video struct{
	Title 		string 		`json:"title"`
	ID			uint		`json:"id"`
	Author  	string 		`json:"author"`
	FilePath  	string  	`json:"filepath"`
	Type		string 		`json:"type"`
	Description string		`json:"description"`
}

func BuildVideo(item models.Video) Video{
	return Video{
		Title: item.Title,
		ID: item.ID,
		Author: item.User.UserName,
		FilePath: item.FilePath,
		Type: item.Type,
		Description: item.Description,
	}
}