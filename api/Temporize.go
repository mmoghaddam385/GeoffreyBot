package api

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"fmt"
	"time"
	"math/rand"
	"os"
)

const temporizeURLEnvironmentVar = "TEMPORIZE_URL"

var temporizeUrl = ""

// ScheduleSingleEventInAWeek will schedule an event for approximately a week (4-7 days) from today
// at a random time around noon 
func ScheduleSingleEventInAWeek(callbackUrl string) {
	now := time.Now()
	rand.Seed(now.Unix())

	later := now.AddDate(0, 0, 4 + rand.Intn(3))
	later = time.Date(later.Year(), 
					later.Month(),
					later.Day(),
					rand.Intn(24),
					rand.Intn(60),
					0, 0, time.UTC)

	ScheduleSingleEvent(later, callbackUrl)
}

func ScheduleSingleEvent(t time.Time, callbackUrl string) {
	baseUrl := getTemporizeUrl()
	if baseUrl == "" {
		fmt.Printf("Cannot schedule event; $%v environment variable not set\n", temporizeURLEnvironmentVar)
		return
	}

	url := fmt.Sprintf("%v/v1/events/%v/%v", baseUrl, t.Format(time.RFC3339), url.QueryEscape(callbackUrl))

	resp, err := http.Post(url, "text/plain", nil)

	if (err != nil) {
		fmt.Printf("Error scheduling event; %v", err)
		return
	}

	if (resp.StatusCode != 200) {
		respBody, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Error scheduling event; Post failed; response code: %v: %v;\n\tbody: %v\n", resp.StatusCode, resp.Status, string(respBody))
	}

	resp.Body.Close()
}

func getTemporizeUrl() string {
	if temporizeUrl == "" {
		temporizeUrl = os.Getenv(temporizeURLEnvironmentVar)
	}

	return temporizeUrl
}