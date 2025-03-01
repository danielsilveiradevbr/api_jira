package issueTypeModel

import (
	issueTypeDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/issueType"
	"gorm.io/gorm"
)

type IssueType struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey;autoIncrement:true"`
	NOME string `gorm:"unique;not null"`
}

func NewIssuetype(issueTypedto *issueTypeDto.Issuetype) *IssueType {
	return &IssueType{
		NOME: issueTypedto.Name,
	}
}
