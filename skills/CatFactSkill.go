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

var passiveGreetingSubjectOptions = []string {
	"y'all",
	"guys",
	"everyone",
	"fam",
	"friends",
	"losers",
}

var factPrefixOptions = []string {
	"Did you know %v?",
	"Have you heard that %v?",
	"Can you believe that %v?",
	"Wanna hear a fun cat fact?? Here ya go: %v.",
	"Guess what! Ahh you'll never guess I'll just tell you: %v!",
}

// catFactActiveSkill sends the group a cat fact if geoffrey is 
// mentioned in a message
func catFactActiveSkill(message types.GroupMeMessagePost) bool {
	if (!isGeoffreyMentioned(message.MessageText)) {
		return false
	}

	// Get only the first name of the sender
	sender := message.Sender
	indexOfSpace := strings.Index(sender, " ")
	if indexOfSpace > 0 {
		sender = sender[:indexOfSpace]
	}

	finalMessage, err := getCatFact(sender)
	if err != nil {
		fmt.Printf("Error getting cat fact for cat fact skill; %v", err)
		return false
	}
	
	api.PostGroupMeMessage(finalMessage)

	return true
}

// catFactPassiveSkill sends a random cat fact to the group un prompted!
func catFactPassiveSkill() {
	finalMessage, err := getCatFact(pickRandomFromStringArray(passiveGreetingSubjectOptions))
	if err != nil {
		fmt.Printf("Error getting cat fact for cat fact skill; %v", err)
		return
	}
	
	api.PostGroupMeMessage(finalMessage)
}

// getCatFact gets the cat fact from the API and formats the message while its waiting for a response
func getCatFact(name string) (string, error) {
	// Go get the cat fact while we figure out who to @
	catFactChannel := make(chan api.StringResult)
	go api.GetCatFactAsync(catFactChannel)

	
	greeting := fmt.Sprintf(pickRandomFromStringArray(greetingOptions), name)
	factFormat := string(pickRandomFromStringArray(factPrefixOptions))
	factWithoutPunctuation := factFormat[:len(factFormat) - 1]
	factPunctuation := string(factFormat[len(factFormat) - 1])
	
	catFactResult := <-catFactChannel

	if (catFactResult.Err != nil) {
		return "", catFactResult.Err
	}

	fact := formatCatFact(catFactResult.Result, factPunctuation)

	finalMessage := greeting + " " + fmt.Sprintf(factWithoutPunctuation, fact)

	return finalMessage, nil
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
