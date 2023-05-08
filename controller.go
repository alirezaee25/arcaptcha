package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
)

const dataListKey = "data_list"

type SampleData struct {
	IP string `json:"IP"`
	CR string `json:"CR"`
}

func main() {
	jobs := make(chan SampleData, 15)
	for w := 1; w <= 8; w++ {
		go worker(w, jobs)
	}
	fmt.Println("Go Redis Tutorial")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "fc935b29-b386-4701-b3ae-61439f69cc18.hsvc.ir:30220",
		Password: "m5oW2lYnUgWSm6hihybRidIC4MXqLee4",
		DB:       0,
	})
	for range time.Tick(time.Second * 1) {
		result, err := rdb.ZPopMin(ctx, dataListKey, 15).Result()
		if err != nil {
			panic(err)
		}
		for _, v := range result {
			sample := SampleData{}
			json.Unmarshal([]byte(v.Member.(string)), &sample)
			jobs <- sample
			fmt.Println(v, sample)
		}
		time.Sleep(time.Second * 5)
}
	if len(jobs) == 0 {
		close(jobs)
	}

}
