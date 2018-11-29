package skills

import (
	"geoffrey/api"
	"geoffrey/types"
)

func summonSkill(message types.GroupMeMessagePost) bool {

	var messageText = "@Kaie Westmaas"
	var mention = types.GroupMeMessageMention {
		UserId: "1234",
		StartIndex: 0,
		Length: len(messageText),
	}

	api.PostGroupMeMessageWithMentions(messageText, mention)

	return true;
}
