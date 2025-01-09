package dds

import (
	statusDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	ID               uint64 `gorm:"primaryKey;autoIncrement:true"`
	ID_JIRA          string
	DESCRICAO        string `gorm:"unique"`
	StatusCategoryId int64
	StatusCategory   StatusCategory
}

func (Status) TableName() string {
	return "status"
}

func NewStatus(status *statusDTO.Status) *Status {
	return &Status{
		ID_JIRA:   status.ID,
		DESCRICAO: status.Description,
	}
}
