package assignee

import (
	assigneeDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	assigneeModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds"
	"gorm.io/gorm"
)

func SalvaAssigned(db *gorm.DB, assigneeDTO *assigneeDTO.Assignee) error {
	assignee := assigneeModel.NewAssignee(assigneeDTO)
	result := db.First(&assignee, "email = ?", assignee.Email)
	if result.RowsAffected == 0 {
		res := db.Create(&assignee)
		if res.Error != nil {
			return res.Error
		}
	} else {
		db.Save(&assignee)
	}
	return nil
}
