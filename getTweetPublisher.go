package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func getAuthor() {
	token := os.Getenv("BEARER_TOKEN")
	x := getTweets()

	ids := "ids="

	for y := 0; y < len(x); y++ {
		if ids == "ids=" {
			ids = ids + x[y]
		} else {
			ids = ids + "," + x[y]
		}

	}

	tweetFields := "tweet.fields=lang,author_id,in_reply_to_user_id,attachments"
	url := fmt.Sprintf("https://api.twitter.com/2/tweets?%s&%s", ids, tweetFields)

	twitterClient := http.Client{
		Timeout: time.Second * 2,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Println("test")
		log.Fatal(err)
	}

	request.Header.Set("Authorization", "Bearer "+token)

	request.Header.Set("User-Agent", "v2UserLookupGolang")

	result, getErr := twitterClient.Do(request)

	if getErr != nil || result.StatusCode != 200 {
		fmt.Println("test1")
		fmt.Println(result.Status)

		log.Fatal(result.StatusCode)
	}

	testByte, readErr := ioutil.ReadAll(result.Body)
	if readErr != nil {
		fmt.Println("test2")

		log.Fatal(readErr)
	}

	var data TweetStructure

	parseErr := json.Unmarshal(testByte, &data)

	if parseErr != nil {
		log.Fatal(parseErr)
	}

	for x := 0; x < len(data.Data); x++ {
		getMedia(data.Data[x].Id)

	}

}
