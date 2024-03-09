package service

import (
	"encoding/json"
	"fmt"
	"time"
	"work4/models"
	"work4/serializer"
	// "github.com/go-redis/redis"
)

type SearchService struct {
	Target      string `form:"target" binding:"required" json:"Target"`
	Name		string `form:"name" binding:"-" json:"Name"`
	Category    string `form:"category" binding:"-" json:"Category"`
	PublishTime string `form:"publishtime" binding:"-" json:"Publishtime"`
}

func (service *SearchService) SearchAll(username string) serializer.Response {
	if service.Target == "user" && service.Name != "" {
		var users []models.User
		var count int
		if err := models.DB.Model(&models.User{}).
		Select("id, user_name, avatar_file_path").
		Where("user_name LIKE ? OR id LIKE?", "%"+service.Name+"%","%"+service.Name+"%").
		Find(&users).Count(&count).Error; err != nil {
			return serializer.Response {
				Status: 200,
				Msg: "没找到相关内容",
			}
		}else{
			searchHistoryString, _ := json.Marshal(service)
			timestamp := time.Now().Unix()
			models.Redisdb.Do("ZADD", username, timestamp, searchHistoryString)
			return serializer.BuildListResponse(serializer.BuildSearchUsersResponse(users), uint(count))
		}
	} else if service.Target == "video" && (service.Name != "" || service.Category != "" || service.PublishTime != ""){
		var videos []models.Video
		var count int
		query := models.DB.Model(&models.Video{})
		if service.Name != "" {
			query = query.Where("title = ?", service.Name)
		}
		if service.PublishTime != "" {
			query = query.Where("YEAR(created_at) = ?", service.PublishTime)
		}
		if service.Category != "" {
			query = query.Where("type = ?", service.Category)
		}
		if err := query.Find(&videos).Count(&count).Error; err != nil {
			return serializer.Response {
				Status: 200,
				Msg: "没找到相关内容",
			}
		}else{
			searchHistoryString, _ := json.Marshal(service)
			timestamp := time.Now().Unix()
			models.Redisdb.Do("ZADD", username, timestamp, searchHistoryString)
			return serializer.BuildListResponse(serializer.BuildSearchVideosResponse(videos), uint(count))
		}
	}else{
		return serializer.Response{
			Status: 400,
			Msg: "请输入更多搜索条件",
		}
	}
}

func GetSearchHistory(username string) serializer.Response{
	searchRecords := models.Redisdb.Do("ZRANGEBYSCORE", username, "0", fmt.Sprintf("%d", time.Now().Unix()))
	result, _ := searchRecords.Result()
	fmt.Println(result)
	list, _ := result.([]interface{})
	var searchrecords []SearchService
	// jsonString := fmt.Sprintf(`%s`, result)
	// err := json.Unmarshal([]byte(jsonString), &searchrecords)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	for _, item := range list {
		fmt.Println(item)
		// detail, err := item.(map[string]interface{})
		// fmt.Println(detail, err)
		// searchrecord := SearchService{
		// 	Target: detail["Target"],
		// }

		var searchrecord SearchService
		jsonString := fmt.Sprintf("%s", item)
		err := json.Unmarshal([]byte(jsonString), &searchrecord)
		if err != nil {
			fmt.Println(err)
		}
		searchrecords = append(searchrecords, searchrecord)
	}
	return serializer.BuildListResponse(searchrecords , uint(len(searchrecords)))
}