package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type responseData struct {
	Data payload `json:"payload"`
}

type payload struct {
	Url string `json:"url"`
	PictureUrl string `json:"picture_url"`
}

const groupMeImageServiceUrl = "https://image.groupme.com/pictures?access_token=%v"

const groupMeAccessTokenEnvVar = "GM_ACCESS_TOKEN"

// Process Image posts the given image to GroupMe's image service
// and returns the group me url of the image
func ProcessImage(image io.Reader) string {
	token := os.Getenv(groupMeAccessTokenEnvVar)
	if token == "" {
		fmt.Println("Cannot process image; GM_ACCESS_TOKEN env var not set")
		return ""
	}

	url := fmt.Sprintf(groupMeImageServiceUrl, token)

	resp, err := http.Post(url, "image/jpeg", image)

	defer resp.Body.Close()

	if err != nil {
		fmt.Printf("Error from GroupMe image servce: %v", err)
		return ""
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return string(respBody)
	}

	respPayload := responseData{}

	err = json.Unmarshal(respBody, &respPayload)

	if err != nil {
		fmt.Printf("Error converting response body to JSON: %v\n", err)
		return ""
	}

	return respPayload.Data.PictureUrl
}
