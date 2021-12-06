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
	token := os.Getenv("BEARER_TOKEN")

	ids := "ids=" + TweetId

	// tweetFields := "preview_image_url,width,height"
	url := fmt.Sprintf("https://api.twitter.com/2/tweets?%s&expansions=attachments.media_keys&media.fields=preview_image_url,url",
		ids)

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

	var data tweetMedia

	parseErr := json.Unmarshal(testByte, &data)

	if parseErr != nil {
		log.Fatal(parseErr)
	}

	if len(data.Includes.Media) <= 0 {

	} else {
		fmt.Println(data.Includes.Media[0].Url)
		fmt.Println(data.Includes.Media[0].Type)
	}
}
