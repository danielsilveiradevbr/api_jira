package sprintDto

import (
	"strings"
	"time"

	helper "github.com/danielsilveiradevbr/api_jira/src/helpers"
)

type Sprint struct {
	ID_JIRA       string
	State         string
	Name          string
	StartDate     time.Time
	EndDate       time.Time
	CompleteDate  time.Time
	ActivatedDate time.Time
}

func NewSprintDto(sprintdto []string) *Sprint {
	data := sprintdto[0]
	indice := strings.Index(data, "[")
	data = data[indice+1 : len(data)-1]
	// data = strings.ReplaceAll(data, "=", ":")
	dados := strings.Split(data, ",")
	var sprint Sprint
	for key, dado := range dados {
		keyValue := strings.Split(dado, "=")

		if keyValue[0] == "sequence" {
			break
		}
		if key == 0 {
			sprint.ID_JIRA = keyValue[1]
		} else if key == 2 {
			sprint.State = keyValue[1]
		} else if key == 3 {
			sprint.Name = keyValue[1]
		} else if key == 4 {
			sprint.StartDate = helper.StrToTimeTime(keyValue[1])
		} else if key == 5 {
			sprint.EndDate = helper.StrToTimeTime(keyValue[1])
		} else if key == 6 {
			sprint.CompleteDate = helper.StrToTimeTime(keyValue[1])
		} else if key == 7 {
			sprint.ActivatedDate = helper.StrToTimeTime(keyValue[1])
		}

	}
	return &sprint
}
