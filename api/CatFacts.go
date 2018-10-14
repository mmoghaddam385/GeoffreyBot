package api

import (
	"net/http"
	"errors"
)

const catFactsBaseUrl = "https://catfact.ninja"

const catFactEndpoint = catFactsBaseUrl + "/fact"
const catFactFactKey = "fact"

func GetCatFact() (string, error) {
	resp, err := http.Get(catFactEndpoint)

	if (err != nil) {
		return "", err
	}

	defer resp.Body.Close()

	if (resp.StatusCode != 200) {
		return "", buildStatusError("Error recieved from cat facts endpoint", resp)
	}

	body, err := getBodyJson(resp)

	if (err != nil) {
		return "", err
	}

	if fact, exists := body[catFactFactKey]; exists {
		return fact, nil
	} else {
		return "", errors.New("Could not parse fact out of response body!")
	}
}

func GetCatFactAsync(channel chan StringResult) {
	fact, err := GetCatFact()
	channel <- StringResult{fact, err}
}