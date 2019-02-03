package skills

import (
	"os"
	"geoffrey/api"
	"geoffrey/types"
)

// summons the beast...
func summonSkill() {
	var messageText = "@Kaie Westmaas"
	var mention = types.GroupMeMessageMention {
		UserId: os.Getenv("SUMMON_USER_ID"),
		StartIndex: 0,
		Length: len(messageText),
	}

	api.PostGroupMeMessageWithMentions(messageText, mention)
}
