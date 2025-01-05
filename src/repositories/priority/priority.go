package priority

import (
	priorityDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	priorityModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds"
	"gorm.io/gorm"
)

func SalvaPriority(db *gorm.DB, priorityDTO *priorityDTO.Priority) error {
	priority := priorityModel.NewPriority(priorityDTO)
	result := db.First(&priority, "descricao = ?", priority.DESCRICAO)
	if result.RowsAffected == 0 {
		res := db.Create(&priority)
		if res.Error != nil {
			return res.Error
		}
	} else {
		db.Save(&priority)
	}
	return nil
}
