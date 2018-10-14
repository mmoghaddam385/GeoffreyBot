package endpoints

import (
	"geoffrey/skills"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
)

const postBodyMessageId = "id"
const postBodyGroupId = "group_id"
const postBodySenderName = "name"
const postBodySenderType = "sender_type"

// messageRecieved handles POST requests from GroupMe that get sent for each message sent
// in the group that the bot is in
func messageRecieved(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Message Recieved!!\n")

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if (err != nil) {
		fmt.Printf("Error reading request body: %v\n", err)
		return
	}

	fmt.Printf("Super raw: %v", string(body))

	bodyMap := make(map[string] string)
	json.Unmarshal(body, &bodyMap)

	logMessage(bodyMap)

	// Only process messages from humans
	if (bodyMap[postBodySenderType] != "user") {
		return
	}

	for _, skill := range skills.GetActiveSkills() {
		if (skill(bodyMap)) {
			break
		}
	}
}

func logMessage(body map[string] string) {
	fmt.Printf("raw: %v", body)
	fmt.Println("Message: " + body["text"])
	fmt.Printf("Message (id: %v) from %v in group %v\n", body[postBodyMessageId], body[postBodySenderName], body[postBodyGroupId])
}
