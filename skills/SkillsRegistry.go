package skills

import (
	"fmt"
	"time"
	"math/rand"
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
	catFactActiveSkill,
}

// GetActiveSkills returns all registered active  skills in order or priority
func GetActiveSkills() []ActiveSkill {
	return activeSkills
}

var passiveSkills = map[string] PassiveSkill {
	"summon": summonSkill,
	"cat-fact": catFactPassiveSkill,
	"roast": roasterPassiveSkill,
	"nostalgia": nostalgiaPassiveSkill,
	"nostalgia-rand-bump": nostalgiaPassiveSkill, // Bump up the likelihood of this skill getting picked randomly
}

func GetPassiveSkillByName(name string) PassiveSkill {
	if (name == "random") {
		rand.Seed(time.Now().Unix())

		// Calculate the random index to use
		var target = rand.Intn(len(passiveSkills))
		var idx = 0
		// Iterate over the map until we hit target, then return the skill
		// Side note: I can't believe there isn't a better way to do this :\ 
		for name, skill := range passiveSkills {
			if (idx == target) {
				fmt.Printf("Getting random skill...%v\n", name)
				return skill
			}

			idx++
		}
	}
	
	return passiveSkills[name]
}