package issuetype

import (
	issueTypeDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/issueType"
	issueModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/issueType"
	"gorm.io/gorm"
)

func SalvaIssueType(db *gorm.DB, issueTypeDTO *issueTypeDto.Issuetype) (*issueModel.IssueType, error) {
	issueType := issueModel.NewIssuetype(issueTypeDTO)
	result := db.First(&issueType, "nome = ?", issueType.NOME)
	if result.RowsAffected == 0 {
		res := db.Create(&issueType)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		db.Save(&issueType)
	}
	return issueType, nil
}
