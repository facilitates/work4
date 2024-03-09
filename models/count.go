package models

import "fmt"

func IncreaseClick(videoID string) error {
	key := videoID
	result, err := Redisdb.ZIncrBy("videoclick", 1, key).Result()
	if err != nil {
		fmt.Println(result, err)
		return err
	}
	fmt.Println(result,"here", err)
	return nil
}