package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func getTweets() []string {
	token := os.Getenv("BEARER_TOKEN")
	ID := getID()
	UrlAPI := fmt.Sprintf("https://api.twitter.com/2/users/%s/tweets", ID)

	twitterClient := http.Client{
		Timeout: time.Second * 2,
	}

	request, err := http.NewRequest(http.MethodGet, UrlAPI, nil)

	// I don't know why it has to be different, but it is

	param := url.Values{}
	param.Add("tweet.fields", "created_at")

	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Authorization", "Bearer "+token)

	request.Header.Set("User-Agent", "v2UserLookupGolang")

	// sends the request

	result, getErr := twitterClient.Do(request)

	if getErr != nil || result.StatusCode != 200 {
		log.Fatal(result.StatusCode)
	}

	// reads the request in a not human-readable way

	testByte, readErr := ioutil.ReadAll(result.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var data Tweets

	// sorts data and it is human-readable too

	parseErr := json.Unmarshal(testByte, &data)

	if parseErr != nil {
		log.Fatal(parseErr)
	}

	// gets the tweet ID and returns it

	// Very Important for other functions

	s := make([]string, 0)

	for x := 0; x < len(data.Data); x++ {
		s = append(s, data.Data[x].Id)
	}
	return s

}
