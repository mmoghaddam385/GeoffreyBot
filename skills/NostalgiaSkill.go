package skills

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"os"

	"geoffrey/api"
)

const nostalgiaPicFilePath = "/pics/%v.jpg"
const numNostalgiaPicsEnvVar = "NUM_NOSTALGIA_PICS"

var passiveGreetingOptions = []string {
	"Let's go on a stroll down memory lane...",
	"Hahahaha remember this?? (I don't...ya know cuz you guys kicked me out and all...)",
	"Woah check out this blast from the past:",
	"Ahhh those were the days...",
	"How come no one ever invites me to cool stuff like this?",
	"Woah when did this happen?",
	"You guys sure had a good time without me...#FOMO",
	"What's the story behind this beauty??",
}

func nostalgiaPassiveSkill() {
	pictureUrl := getRandomNostalgiaPicUrl()
	
	if pictureUrl != "" {
		fmt.Printf("Random nostalgia pic url: %v\n", pictureUrl)

		messageText := pickRandomFromStringArray(passiveGreetingOptions)

		api.PostGroupMeMessageWithPicture(messageText, pictureUrl)
	} else {
		postErrorMessage()
	}
}

func getRandomNostalgiaPicUrl() string {
	numNostalgiaPicsStr := os.Getenv(numNostalgiaPicsEnvVar)
	if numNostalgiaPicsStr == "" {
		fmt.Printf("Could not retrieve random nostalgia pic; %v env var not set!\n", numNostalgiaPicsEnvVar)
		return ""
	}

	numNostalgiaPics, _ := strconv.Atoi(numNostalgiaPicsStr)

	rand.Seed(time.Now().Unix())
	// I started the pics at 1 instead of 0, sue me
	randIndex := 1 + (rand.Int() % numNostalgiaPics)
	filePath := fmt.Sprintf(nostalgiaPicFilePath, randIndex)

	fmt.Printf("Random nostalgia pic index: %v\n", randIndex)
	fmt.Printf("Downloading dropbox file %v...\n", filePath)

	file, err := api.DownloadFile(filePath)

	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
		return ""
	}

	defer file.Close()

	fmt.Printf("Processing %v in GroupMe Image service...\n", filePath)

	processedUrl := api.ProcessImage(file)

	if processedUrl == "" {
		fmt.Printf("Error processing image with GroupMe Image service :(\n")
	}

	return processedUrl
}
