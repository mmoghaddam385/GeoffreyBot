package endpoints

import (
	"geoffrey/skills"
	"geoffrey/types"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
)

const postBodyMessageId = "id"
const postBodyMessageText = "text"
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

	bodyMap := make(map[string] interface{})
	err = json.Unmarshal(body, &bodyMap)

	if (err != nil) {
		fmt.Printf("Error unmarshalling body into map; %v", err)
		return
	}

	logMessage(bodyMap)
	message := buildMessageStruct(bodyMap)

	// Only process messages from humans
	if (message.SenderType != "user") {
		return
	}

	for _, skill := range skills.GetActiveSkills() {
		if (skill(message)) {
			break
		}
	}
}

func buildMessageStruct(bodyMap map[string] interface{}) types.GroupMeMessagePost {
	var msg types.GroupMeMessagePost
	msg.Id = bodyMap[postBodyMessageId].(string)
	msg.GroupId = bodyMap[postBodyGroupId].(string)
	msg.Sender = bodyMap[postBodySenderName].(string)
	msg.SenderType = bodyMap[postBodySenderType].(string)
	msg.MessageText = bodyMap[postBodyMessageText].(string)

	return msg
}

func logMessage(body map[string] interface{}) {
	fmt.Printf("Message (id: %v) from %v in group %v\n", body[postBodyMessageId], body[postBodySenderName], body[postBodyGroupId])
}
