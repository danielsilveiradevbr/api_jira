package issuetype

import (
	issueTypeDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	issueModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/issueType"
	"gorm.io/gorm"
)

func SalvaIssueType(db *gorm.DB, issueTypeDTO *issueTypeDTO.Issuetype) error {
	issueType := issueModel.NewIssuetype(issueTypeDTO)
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
