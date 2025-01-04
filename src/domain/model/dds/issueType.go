package dds

import (
	"time"

	issueTypeDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	utils "github.com/danielsilveiradevbr/api_jira/src/utils"
)

type Issuetype struct {
	ID           int64 `gorm:"primaryKey;autoIncrement:true"`
	NOME         string
	DATA_CRIACAO time.Time
	DATA_UPDATE  time.Time
}

func NewIssuetype(issuetypedto *issueTypeDto.Issuetype) *Issuetype {
	return &Issuetype{
		NOME:         issuetypedto.Name,
		DATA_CRIACAO: utils.GetADateTimeSaoPaulo(),
	}
}
