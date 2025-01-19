package sprintModel

import (
	"time"

	sprintDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/sprint"
	"gorm.io/gorm"
)

type Sprint struct {
	gorm.Model
	ID            uint `gorm:"primaryKey;autoIncrement:true"`
	ID_JIRA       string
	STATUS        string
	NOME          string `gorm:"unique;not null"`
	COMPLETE_DATE time.Time
	START_DATE    time.Time
	END_DATE      time.Time
}

func NewSprint(sprintdto *sprintDto.Sprint) *Sprint {
	return &Sprint{
		ID_JIRA:       sprintdto.ID_JIRA,
		STATUS:        sprintdto.State,
		NOME:          sprintdto.Name,
		COMPLETE_DATE: sprintdto.CompleteDate,
		START_DATE:    sprintdto.StartDate,
		END_DATE:      sprintdto.EndDate,
	}

}
