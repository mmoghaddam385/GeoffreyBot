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
	factWithoutPunctuation := factFormat[:len(factFormat) - 1]
	factPunctuation := string(factFormat[len(factFormat) - 1])
	
	catFactResult := <-catFactChannel

	if (catFactResult.Err != nil) {
		fmt.Printf("Error getting cat fact for cat fact skill; %v", catFactResult.Err)
		return false
	}

	fact := formatCatFact(catFactResult.Result, factPunctuation)

	finalMessage := greeting + " " + fmt.Sprintf(factWithoutPunctuation, fact)
	api.PostGroupMeMessage(finalMessage)

	return true
}

func formatCatFact(fact string, punctuation string) string {
	// Some cat facts are multiple sentences, find the first punctuation so we can replace it later
	var puncuationIndex = strings.IndexAny(fact, ".?!")
	if puncuationIndex == -1 {
		puncuationIndex = len(fact) - 1
	}

	// Force the first letter to be lower case
	firstLetter := strings.ToLower(fact[:1])

	return firstLetter + fact[1:puncuationIndex] + punctuation + fact[puncuationIndex+1:]
}
