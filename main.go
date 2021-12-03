package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type AutoGenerated struct {
	Data []struct {
		Name        string    `json:"name"`
		Verified    bool      `json:"verified"`
		Description string    `json:"description"`
		Username    string    `json:"username"`
		CreatedAt   time.Time `json:"created_at"`
		ID          string    `json:"id"`
	} `json:"data"`
}

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

	testByte, readErr := ioutil.ReadAll(result.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var data AutoGenerated

	err = json.Unmarshal(testByte, &data)
	if err != nil {
		return
	}
	fmt.Println(data)
	fmt.Println(string(testByte))
	fmt.Println(result.Status)

}
