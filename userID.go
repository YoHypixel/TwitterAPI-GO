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

func getID() string {
	// basic setup

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

	// does request

	result, getErr := twitterClient.Do(request)

	// error checking and reading od data from request

	if getErr != nil || result.StatusCode != 200 {
		log.Fatal(result.StatusCode)
	}
	testByte, readErr := ioutil.ReadAll(result.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// sorts data and return ID

	var data TwitterUser
	err = json.Unmarshal(testByte, &data)
	FinalData := data.Data[0]
	trueData := FinalData.Id

	return trueData
}
