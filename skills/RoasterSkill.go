package skills

import (
	"time"
	"math/rand"
	"geoffrey/api"
	"geoffrey/types"
)

type Roast struct {
	text string
	mention types.GroupMeMessageMention
}

// map of string (user id) to potential roasts
// user id of "" can apply to anyone
var roasts = map[string] []Roast {
	"21004947": { // Johnny
		{
			text: "@Johnny Bollash when're you gonna drop the whole 'Plutarch' thing and admit your middle name is Francis?",
			mention: types.GroupMeMessageMention {
				UserId: "21004947",
				StartIndex: 0,
				Length: len("@Johnny Bollash"),
			},
		},
		{
			text: "@Johnny Bollash Connecticut is small. Boom roasted",
			mention: types.GroupMeMessageMention {
				UserId: "21004947",
				StartIndex: 0,
				Length: len("@Johnny Bollash"),
			},
		},
	},
	"20596690": { // Anokhi 
		{
			text: "@Anohki Patel I was going through some of the old messages I missed while I was gone and...I'm a f**king guinea pig alright??",
			mention: types.GroupMeMessageMention {
				UserId: "20596690",
				StartIndex: 0,
				Length: len("@Anohki Patel"),
			},
		},
	},
	"18172472": { // Apurva 
		{
			text: "@Apurva Kasam you live in Missouri. Boom roasted",
			mention: types.GroupMeMessageMention {
				UserId: "18172472",
				StartIndex: 0,
				Length: len("@Apurva Kasam"),
			},
		},
	},
	"20626795": { // Michael 
		{
			text: "@Michael Moghaddam you write shitty bots",
			mention: types.GroupMeMessageMention {
				UserId: "20626795",
				StartIndex: 0,
				Length: len("@Michael Moghaddam"),
			},
		},
		{
			text: "@Michael Moghaddam D.C. isn't part of Maryland. Get over yourself.",
			mention: types.GroupMeMessageMention {
				UserId: "20626795",
				StartIndex: 0,
				Length: len("@Michael Moghaddam"),
			},
		},
	},
	"17123786": { // Heman
		{
			text: "@Hemanth Koralla...do I know you?",
			mention: types.GroupMeMessageMention {
				UserId: "17123786",
				StartIndex: 0,
				Length: len("@Hemanth Koralla"),
			},
		},
	},
	"20868132": { // David 
		{
			text: "@David morrison why didn't you capitalize the 'm' in your last name?",
			mention: types.GroupMeMessageMention {
				UserId: "20868132",
				StartIndex: 0,
				Length: len("@David morrison"),
			},
		},
	},
	"22602314": { // Kaie
		{
			text: "@Kaie Westmaas have you moved to Japan yet?",
			mention: types.GroupMeMessageMention {
				UserId: "22602314",
				StartIndex: 0,
				Length: len("@Kaie Westmaas"),
			},
		},
	},
	"21405378": { // Ryan
		{
			text: "@Ryan Miller You dress like a dad. Boom roasted.",
			mention: types.GroupMeMessageMention {
				UserId: "21405378",
				StartIndex: 0,
				Length: len("@Ryan Miller"),
			},
		},
	},
	"21498740": { // Vicki 
		{
			text: "@Victoria Kravets got any new stalkers recently?",
			mention: types.GroupMeMessageMention {
				UserId: "21498740",
				StartIndex: 0,
				Length: len("@Victoria Kravets"),
			},
		},
	},
}

func roasterPassiveSkill() {
	roast := getRandomRoast()

	api.PostGroupMeMessageWithMentions(roast.text, roast.mention)
}

func getRandomRoast() Roast {
	var allRoasts []Roast
	for _, roastList := range roasts {
		allRoasts = append(allRoasts, roastList...)
	}

	rand.Seed(time.Now().Unix())
	return allRoasts[rand.Intn(len(allRoasts))]
}