package api

import (
	"time"
	"math/rand"
	"net/http"
)

type YesOrNoAnswer int
type YesOrNoResponse struct {
	Answer YesOrNoAnswer
	ImageUrl string
}

const (
	YES YesOrNoAnswer  = iota
	NO
	MAYBE
)

const yesOrNoUrl = "https://yesno.wtf/api"
const forceMaybeUrl = yesOrNoUrl + "/?force=maybe"

func GetYesOrNo() (YesOrNoResponse, error) {
	var retval YesOrNoResponse
	url := yesOrNoUrl

	if shouldForceMaybe() {
		url = forceMaybeUrl
	}

	resp, err := http.Get(url)

	if (err != nil) {
		return retval, err
	}

	defer resp.Body.Close()

	if (resp.StatusCode != 200) {
		return retval, buildStatusError("Error recieved from cat facts endpoint", resp)
	}

	// Get the body and parse it into a map
	body, err := getBodyJson(resp)

	if (err != nil) {
		return retval, err
	}

	// Everything has gone well, read the response now
	switch body["answer"] {
		case "yes": retval.Answer = YES
		case "no":  retval.Answer = NO
		default:    retval.Answer = MAYBE
	}

	retval.ImageUrl = body["image"]

	return retval, nil
}

// Force maybe ~1/15 times to make things more interesting
func shouldForceMaybe() bool {
	rand.Seed(time.Now().Unix())
	return rand.Intn(15) == 1
}