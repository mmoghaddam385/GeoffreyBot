package skills

import (
	"fmt"
	"strings"
	"geoffrey/api"
	"geoffrey/types"
)

var greetingOptions = []string {
	"Hi %v!",
	"Hi there %v!",
	"Hey %v!",
	"What's crackin' %v?!",
	"Hello there %v.",
}

var factPrefixOptions = []string {
	"Did you know %v?",
	"Have you heard that %v?",
	"Can you believe that %v?",
	"Wanna hear a fun cat fact?? Here ya go: %v.",
	"Guess what! Ahh you'll never guess I'll just tell you: %v!",
}

// catFactSkill sends the group a cat fact if geoffrey is 
// mentioned in a message
func catFactSkill(message types.GroupMeMessagePost) bool {
	if (!isGeoffreyMentioned(message.MessageText)) {
		return false
	}

	// Go get the cat fact while we figure out who to @
	catFactChannel := make(chan api.StringResult)
	go api.GetCatFactAsync(catFactChannel)

	// Get only the first name of the sender
	sender := message.Sender
	indexOfSpace := strings.Index(sender, " ")
	if indexOfSpace > 0 {
		sender = sender[:indexOfSpace]
	}

	greeting := fmt.Sprintf(pickRandomFromStringArray(greetingOptions), sender)
	factFormat := string(pickRandomFromStringArray(factPrefixOptions))
	
	catFactResult := <-catFactChannel

	if (catFactResult.Err != nil) {
		fmt.Printf("Error getting cat fact for cat fact skill; %v", catFactResult.Err)
		return false
	}

	fact := formatCatFact(catFactResult.Result)

	finalMessage := greeting + " " + fmt.Sprintf(factFormat, fact)
	api.PostGroupMeMessage(finalMessage)

	return true
}

func formatCatFact(fact string) string {
	firstLetter := strings.ToLower(fact[:1])
	lastLetter := fact[len(fact) - 1]

	// Chop off punctuation if it's there
	if (lastLetter == '.' || lastLetter == '?' || lastLetter == '!') {
		return firstLetter + fact[1:len(fact) - 1]
	} else {
		return firstLetter + fact[1:]
	}
}
