package skills

import (
	"math/rand"
	"time"
	"strings"

	"geoffrey/api"
	"geoffrey/types"
)

var aliases = []string {
	"geoffrey",
	"geoff",
}

const postBodyMessageText = "text"
const postBodySenderName = "sender"

// isGeoffreyMentioned checks the given messageText for any instances of geoffrey's aliases
// preceeded by an '@'
func isGeoffreyMentioned(messageText string) bool {
	// Case doesn't matter!
	var lowerCaseMessage = strings.ToLower(messageText)
	
	// Loop over each of Geoffrey's aliases to see if he's been mentioned in the message
	for _, alias := range aliases {
		if strings.Contains(lowerCaseMessage, "@" + alias) {
			return true
		}
	}

	return false
}

// stripGeoffreyMentions removes all @aliases from the given string
// Note the returned string also is all lower case
func stripGeoffreyMentions(message string) string {
	var lowerCaseMessage = strings.ToLower(message)

	for _, alias := range aliases {
		lowerCaseMessage = strings.Replace(lowerCaseMessage, "@" + alias, "", -1)
	}

	return strings.TrimSpace(lowerCaseMessage)
}

// isQuestion checks if the given message text is a question or not
func isQuestion(message string) bool {
	// Naively just check if there's a question mark in the string :\
	return strings.Contains(message, "?")
}

// isYesOrNoQuestion checks if the given message text is a yes or no question
func isYesOrNoQuestion(message string) bool {
	// First confirm it's a generic question
	if (!isQuestion(message)) {
		return false
	}

	// Case does't matter!
	message = strings.ToLower(message)

	// Now check if it starts with a yes/no question starter
	for _, starter := range []string { "do", "did", "should", "will", "am", "is", "are" } {
		if (strings.Contains(message, starter)) {
			return true
		}
	}

	return false
}

// pickRandomFromArray returns a random string from the given array
func pickRandomFromStringArray(arr []string) string {
	rand.Seed(time.Now().Unix())
	return arr[rand.Int() % len(arr)]
}

var errorMessageOptions = []string {
	"I don't feel so good...",
	"Maybe try checking my logs every once in a while :/",
	"I've got a cyber stomach ache pls help",
}

// postErrorMessage sends a message to the groupme indicating something went wrong
func postErrorMessage() {
	mention := types.GroupMeMessageMention {
		UserId: "20626795",
		StartIndex: 0,
		Length: len("@Michael Moghaddam"),
	}

	message := "@Michael Moghaddam " + pickRandomFromStringArray(errorMessageOptions)

	api.PostGroupMeMessageWithMentions(message, mention)
}
