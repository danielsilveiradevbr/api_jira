package dds

import (
	"strings"
	"time"

	sprint "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
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

func NewSprint(sprintdto *sprint.Sprint) *Sprint {
	sprint := strings.Split(sprintdto.NOME, ",")
	println(sprint)
	return &Sprint{}
}
