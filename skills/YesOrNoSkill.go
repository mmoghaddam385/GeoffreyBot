package skills

import (
	"fmt"
	"geoffrey/api"
	"geoffrey/types"
)

var responseOptionMap = map[api.YesOrNoAnswer] []string {
	api.YES: []string {
		"Yes!",
		"Yea sure why not.",
		"I don't see why not!",
		"Go for it!",
		"Absolutely!",
	},
	api.NO: []string {
		"No!",
		"Nah I don't thinkg so.",
		"I'm gonna go with...no",
		"Ehhh maybe some other time pal.",
		"Not a chance.",
		"Just...no.",
		"Nope.",
	},
	api.MAYBE: []string {
		"Hmm that's a tough one...",
		"idk",
		"Meh...maybe.",
		"Ahhh who cares?",
		"Just do what you want it doesn't really matter",
	},
}

func yesOrNoSkill(message types.GroupMeMessagePost) bool {
	// First check if geoffrey is mentioned
	if (!isGeoffreyMentioned(message.MessageText)) {
		return false
	}

	messageTextWithoutMention := stripGeoffreyMentions(message.MessageText)

	// Next check if it's a yes or no question
	if (!isYesOrNoQuestion(messageTextWithoutMention)) {
		return false
	}

	response, err := api.GetYesOrNo()

	if (err != nil) {
		fmt.Printf("Error getting Yes or No response: %v", err)
		return false
	}

	messageText := pickRandomFromStringArray(responseOptionMap[response.Answer])

	api.PostGroupMeMessageWithPicture(messageText, response.ImageUrl)

	return true
}