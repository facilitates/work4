package service

import (
	"fmt"
	"work4/models"
	"work4/serializer"

	"github.com/go-redis/redis"
)

func RankVideoList(ranklist *redis.ZSliceCmd, count int) serializer.Response {
	VideoRankList, err := ranklist.Result()
	if err != nil {
		fmt.Println(err)
		return serializer.Response{
			Status: 400,
			Msg:    "视频总数获取错误",
		}
	}
	// RankList := make(map[string]int)
	// ScoreList := make(map[string]int)
	// MemberList := make(map[int]string)
	// UrlList := make(map[int]string)
	var videoranks []serializer.Member
	for i, z := range VideoRankList {
    	member := z.Member.(string)
    	score := z.Score
		// RankList[member] = i/2 + 1
		// ScoreList[member] = int(score)
		// MemberList[i/2+1] = member
		var video models.Video
		fmt.Println(member)
		models.DB.First(&video, member)
		fmt.Println(video)
		videorank := serializer.Member{
			Title: video.Title,
			Rank: i/2+1,
			Views: int(score),
			VideoURL: video.FilePath,
		}
		videoranks = append(videoranks, videorank)
	}
	return serializer.BuildListResponse(serializer.BuildRankList(videoranks), uint(count))
}
