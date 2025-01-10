package priorityModel

import (
	priority "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"gorm.io/gorm"
)

type Priority struct {
	gorm.Model
	ID        int64 `gorm:"primaryKey;autoIncrement:true"`
	ID_JIRA   string
	DESCRICAO string `gorm:"unique;not null"`
}

func NewPriority(prioritydto *priority.Priority) *Priority {
	return &Priority{
		ID_JIRA:   prioritydto.ID,
		DESCRICAO: prioritydto.Name,
	}
}
