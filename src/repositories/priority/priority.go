package priority

import (
	priorityDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/priority"
	priorityModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/priority"
	"gorm.io/gorm"
)

func SalvaPriority(db *gorm.DB, priorityDTO *priorityDto.Priority) (*priorityModel.Priority, error) {
	priority := priorityModel.NewPriority(priorityDTO)
	result := db.First(&priority, "descricao = ?", priority.DESCRICAO)
	if result.RowsAffected == 0 {
		res := db.Create(&priority)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		db.Save(&priority)
	}
	return priority, nil
}
