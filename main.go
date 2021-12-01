package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type DATA struct {
	Data []string `json:"data"`
}

func main() {
	// token := "AAAAAAAAAAAAAAAAAAAAANE9WQEAAAAAZIRfcpuGktyArDC4ppNKH7YefsU%3Do9xRlDeOyxsrMBJL0mqHS2vZz14vaa5LgPmqorpzSFVImh0vL1"
	usernames := "usernames=MCChampionship_"
	userFields := "user.fields=description,created_at,verified"
	url := fmt.Sprintf("https://api.twitter.com/2/users/by?%s&%s", usernames, userFields)

	twitterClient := http.Client{
		Timeout: time.Second *2,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Authorization", "AAAAAAAAAAAAAAAAAAAAANE9WQEAAAAAZIRfcpuGktyArDC4ppNKH7YefsU%3Do9xRlDeOyxsrMBJL0mqHS2vZz14vaa5LgPmqorpzSFVImh0vL1")

	request.Header.Set("User-Agent", "v2UserLookupGolang")

	result, getErr := twitterClient.Do(request)

	if getErr != nil {
		log.Fatal(getErr)
	}

	if result.Body != nil {
		defer result.Body.Close()
	}

	test, readErr := ioutil.ReadAll(result.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fmt.Println(test)

	fmt.Println(result.Status)

}
