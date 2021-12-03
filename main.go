package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	token := os.Getenv("BEARER_TOKEN")
	usernames := "usernames=MCChampionship_"
	userFields := "user.fields=description,created_at,verified"
	url := fmt.Sprintf("https://api.twitter.com/2/users/by?%s&%s", usernames, userFields)

	twitterClient := http.Client{
		Timeout: time.Second * 2,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Authorization", "Bearer "+token)

	request.Header.Set("User-Agent", "v2UserLookupGolang")

	result, getErr := twitterClient.Do(request)

	if getErr != nil {
		log.Fatal(getErr)
	}

	if result.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(result.Body)
	}

	test, readErr := ioutil.ReadAll(result.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fmt.Println(test)

	fmt.Println(result.Status)

}
