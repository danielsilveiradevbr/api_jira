package priorityModel

import (
	priorityDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/priority"
	"gorm.io/gorm"
)

type Priority struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;autoIncrement:true"`
	ID_JIRA   string
	DESCRICAO string `gorm:"unique;not null"`
}

func (Priority) TableName() string {
	return "prioritys"
}

func NewPriority(prioritydto *priorityDto.Priority) *Priority {
	return &Priority{
		ID_JIRA:   prioritydto.ID,
		DESCRICAO: prioritydto.Name,
	}
}
