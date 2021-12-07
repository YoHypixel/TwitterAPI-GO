package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func getDescription() (string, string) {

	// general housekeeping

	token := os.Getenv("BEARER_TOKEN")
	usernames := "usernames=MCChampionship_"
	userFields := "user.fields=description,created_at,verified"
	url := fmt.Sprintf("https://api.twitter.com/2/users/by?%s&%s", usernames, userFields)

	// setting up request client

	twitterClient := http.Client{
		Timeout: time.Second * 2,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Authorization", "Bearer "+token)

	request.Header.Set("User-Agent", "v2UserLookupGolang")

	// doing the request

	result, getErr := twitterClient.Do(request)

	// error checking of the request

	if getErr != nil || result.StatusCode != 200 {
		log.Fatal(result.StatusCode)
	}
	testByte, readErr := ioutil.ReadAll(result.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// sorting and reading of the data

	var data TwitterUser

	err = json.Unmarshal(testByte, &data)
	FinalData := data.Data[0]
	trueData := FinalData.Description

	split := strings.Split(trueData, "\n")
	x := split[2]

	split1 := strings.Split(x, "Next event: ")
	split2 := strings.Split(split1[1], " @ ")
	date := split2[0]
	t := strings.Split(split2[1], " GMT")[0]
	return date, t

}
