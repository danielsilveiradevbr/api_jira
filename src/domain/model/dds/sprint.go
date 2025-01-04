package dds

import (
	"time"

	sprintdto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	utils "github.com/danielsilveiradevbr/api_jira/src/utils"
)

type Sprint struct {
	ID            int64 `gorm:"primaryKey;autoIncrement:true"`
	ID_JIRA       string
	STATUS        string
	NOME          string
	COMPLETE_DATE time.Time
	START_DATE    time.Time
	END_DATE      time.Time
	DATA_CRIACAO  time.Time
	DATA_UPDATE   time.Time
}

func NewSprint(sprintdto *sprintdto.Sprint) *Sprint {

	return &Sprint{
		ID_JIRA:       sprintdto.ID_JIRA,
		STATUS:        sprintdto.State,
		NOME:          sprintdto.Name,
		COMPLETE_DATE: sprintdto.CompleteDate,
		START_DATE:    sprintdto.StartDate,
		END_DATE:      sprintdto.EndDate,
		DATA_CRIACAO:  utils.GetADateTimeSaoPaulo(),
	}

}
