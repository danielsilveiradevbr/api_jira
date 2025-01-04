package issuetype

import (
	issueTypetDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	issue "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds"
	"gorm.io/gorm"
)

func SalvaIssueType(db *gorm.DB, issueDTO *issueTypetDTO.Issuetype) error {
	issueType := issue.NewIssuetype(issueDTO)
	result := db.First(&issueType, "nome = ?", issueType.NOME)
	if result.RowsAffected == 0 {
		res := db.Create(&issueType)
		if res.Error != nil {
			return res.Error
		}
	} else {
		db.Save(&issueType)
	}
	return nil
}
