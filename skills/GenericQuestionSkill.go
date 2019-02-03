package skills

import (
	"geoffrey/types"
	"geoffrey/api"
	"fmt"
	"strings"
	"net/url"
)

var answerOptions = []string {
	"Uhh I don't really know how to answer that...",
	"Hey listen idk alright leave me out of this.",
	"Believe me if I knew I'd tell you!",
	"I'm not even going to justify that question with an answer...",
	"I'm just a guinea pig I don't know these things!",
	"Ugh why do I have to do everything around here... http://lmgtfy.com/?q=%v",
	"Listen that's not my problem buddy good luck figuring it out though.",
}

func genericQuestionSkill(message types.GroupMeMessagePost) bool {

	// First check if geoffrey is mentioned
	if (!isGeoffreyMentioned(message.MessageText)) {
		return false
	}

	messageTextWithoutMention := stripGeoffreyMentions(message.MessageText)

	// Next check if it's a question
	if (!isQuestion(messageTextWithoutMention)) {
		return false
	}

	response := pickRandomFromStringArray(answerOptions)

	// If it's the lmgtfy response, throw in the query string
	if (strings.Contains(response, "lmgtfy")) {
		response = fmt.Sprintf(response, url.QueryEscape(messageTextWithoutMention))
	}

	api.PostGroupMeMessage(response)
	return true
}