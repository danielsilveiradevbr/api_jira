package sprintModel

import (
	"time"

	sprintdto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"gorm.io/gorm"
)

type Sprint struct {
	gorm.Model
	ID            int64 `gorm:"primaryKey;autoIncrement:true"`
	ID_JIRA       string
	STATUS        string
	NOME          string `gorm:"unique;not null"`
	COMPLETE_DATE time.Time
	START_DATE    time.Time
	END_DATE      time.Time
}

func NewSprint(sprintdto *sprintdto.Sprint) *Sprint {

	return &Sprint{
		ID_JIRA:       sprintdto.ID_JIRA,
		STATUS:        sprintdto.State,
		NOME:          sprintdto.Name,
		COMPLETE_DATE: sprintdto.CompleteDate,
		START_DATE:    sprintdto.StartDate,
		END_DATE:      sprintdto.EndDate,
	}

}
