package api

import (
	"io/ioutil"
	"bytes"
	"net/http"
	"fmt"
	"os"
)

const groupMeBaseUrl = "https://api.groupme.com/v3"
const botPostMessageUrl = groupMeBaseUrl + "/bots/post"

const jsonContentType = "application/json"

const botIdEnvironmentVar = "BOT_ID"

var botId string = ""

func PostGroupMeMessage(message string) {
	id := getBotId()
	if (id == "") {
		fmt.Printf("Cannot send message; $%v environment variable not set\n", botIdEnvironmentVar)
		return
	}
	
	body := fmt.Sprintf(`{
		"bot_id": "%v",
		"text": "%v"
	}`, id, message)

	postGroupMeMessage(body)
}

func PostGroupMeMessageWithPicture(message string, imageUrl string) {
	id := getBotId()
	if (id == "") {
		fmt.Printf("Cannot send message; $%v environment variable not set\n", botIdEnvironmentVar)
		return
	}

	body := fmt.Sprintf(`{
		"bot_id": "%v",
		"text": "%v",
		"picture_url": "%v"
	}`, id, message, imageUrl)

	postGroupMeMessage(body)
}

// postGroupMeMessage is the internal function that both public post functions delegate to
func postGroupMeMessage(postBody string) {
	resp, err := http.Post(botPostMessageUrl, jsonContentType, bytes.NewBufferString(postBody))

	if (err != nil) {
		fmt.Printf("Error posting message; %v", err)
		return
	}

	if (resp.StatusCode != 202) {
		respBody, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Post failed; response code: %v: %v;\n\tbody: %v\n", resp.StatusCode, resp.Status, string(respBody))
	}

	resp.Body.Close()
}

func getBotId() string {
	if botId == "" {
		botId = os.Getenv(botIdEnvironmentVar)
	}

	return botId
}