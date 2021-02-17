package endpoints

import (
	"net/http"
	"fmt"
	"os"

	"geoffrey/api"
	"geoffrey/skills"
)

const skillNameQueryParam = "skill"

const rescheduleQueryParam = "reschedule"
const rescheduleWeekQueryParam = "week"

func passiveSkillRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Passive Skill Request recieved!!");

	skillsParam, ok := r.URL.Query()[skillNameQueryParam]

	if !ok || len(skillsParam) < 1 {
		fmt.Printf("Missing %v query param!\n", skillNameQueryParam)
		w.WriteHeader(400)
		return
	}

	// Get the first skill in the array
	skill := skills.GetPassiveSkillByName(skillsParam[0])

	if skill == nil {
		fmt.Printf("Skill (%v) not found!\n", skillsParam[0])
		w.WriteHeader(400)
		return
	}

	fmt.Printf("Running passive skill %v...\n", skillsParam[0])
	skill()

	rescheduleParams, ok := r.URL.Query()[rescheduleQueryParam]

	if ok && len(rescheduleParams) >= 1 {
		if rescheduleParams[0] == rescheduleWeekQueryParam {
			fmt.Println("Rescheduling event for ~1 week")
			fmt.Printf("rescheuling at: %v\n", getServerName() + r.URL.RequestURI())
			api.ScheduleSingleEventInAWeek(getServerName() + r.URL.RequestURI())
		} else {
			fmt.Printf("Unexpected reschedule duration: %t\n", rescheduleParams[0])
			w.WriteHeader(400)
			return
		}
	}

	w.WriteHeader(200)
}

func getServerName() string {
	return os.Getenv("SELF_URL")
}
