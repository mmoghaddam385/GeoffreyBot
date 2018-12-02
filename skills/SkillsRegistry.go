package skills

import (
	"geoffrey/types"
)

// A Skill is a function that takes a map containing the POST body of a GroupMe message POST
// and may or may not perform an action based on it

// ActiveSkills return true/false depending on whether or not they consumed the event
type ActiveSkill func(types.GroupMeMessagePost) bool

// Passive skills just do their thing man
type PassiveSkill func()

var activeSkills = []ActiveSkill {
	yesOrNoSkill,
	genericQuestionSkill,
	catFactSkill,
}

// GetActiveSkills returns all registered active  skills in order or priority
func GetActiveSkills() []ActiveSkill {
	return activeSkills
}

var passiveSkills = map[string] PassiveSkill {
	"summon-skill": summonSkill,
}

func GetPassiveSkillByName(name string) PassiveSkill {
	return passiveSkills[name]
}