package issuetypeModel

import (
	issueTypeDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"gorm.io/gorm"
)

type Issuetype struct {
	gorm.Model
	ID   int64  `gorm:"primaryKey;autoIncrement:true"`
	NOME string `gorm:"unique;not null"`
}

func NewIssuetype(issuetypedto *issueTypeDto.Issuetype) *Issuetype {
	return &Issuetype{
		NOME: issuetypedto.Name,
	}
}
