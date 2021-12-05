package main

import (
	"time"
)

type TwitterUser struct {
	Data []struct {
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		Username    string    `json:"username"`
		Name        string    `json:"name"`
		Id          string    `json:"id"`
		Verified    bool      `json:"verified"`
	} `json:"data"`
}

type Tweets struct {
	Data []struct {
		Id   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
	Meta struct {
		OldestId    string `json:"oldest_id"`
		NewestId    string `json:"newest_id"`
		ResultCount int    `json:"result_count"`
		NextToken   string `json:"next_token"`
	} `json:"meta"`
}

func main() {
	getTweets()
}
