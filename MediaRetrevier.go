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

func getMedia(TweetId string) {

	// General housekeeping to get Twitter API key and sort through Ids given by the getMedia() method

	token := os.Getenv("BEARER_TOKEN")

	ids := "ids=" + TweetId

	url := fmt.Sprintf("https://api.twitter.com/2/tweets?%s&expansions=attachments.media_keys&media.fields=preview_image_url,url",
		ids)

	twitterClient := http.Client{
		Timeout: time.Second * 2,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Authorization", "Bearer "+token)

	request.Header.Set("User-Agent", "v2UserLookupGolang")

	// Does the request and sorts it

	result, getErr := twitterClient.Do(request)

	if getErr != nil || result.StatusCode != 200 {
		fmt.Println(result.Status)

		log.Fatal(result.StatusCode)
	}

	testByte, readErr := ioutil.ReadAll(result.Body)
	if readErr != nil {

		log.Fatal(readErr)
	}

	var data tweetMedia

	parseErr := json.Unmarshal(testByte, &data)

	if parseErr != nil {
		log.Fatal(parseErr)
	}

	// sorts the sorted Data and outputs it to the console

	if len(data.Includes.Media) <= 0 {

	} else {
		if data.Includes.Media[0].Type != "video" {
			fmt.Println(data.Includes.Media[0].Url)
			fmt.Println(data.Includes.Media[0].Type)
		}

	}
}
