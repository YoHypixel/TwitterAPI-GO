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

type TweetStructure struct {
	Data []struct {
		AuthorId        string `json:"author_id"`
		Id              string `json:"id"`
		InReplyToUserId string `json:"in_reply_to_user_id,omitempty"`
		Text            string `json:"text"`
		Lang            string `json:"lang"`
		Attachments     struct {
			MediaKeys []string `json:"media_keys"`
		} `json:"attachments,omitempty"`
	} `json:"data"`
}

type tweetMedia struct {
	Data []struct {
		Attachments struct {
			MediaKeys []string `json:"media_keys"`
		} `json:"attachments"`
		Id   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
	Includes struct {
		Media []struct {
			MediaKey string `json:"media_key"`
			Type     string `json:"type"`
			Url      string `json:"url,omitempty"`
		} `json:"media"`
	} `json:"includes"`
}

func main() {
	getAuthor()
}
