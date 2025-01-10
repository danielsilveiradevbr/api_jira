package statusModel

import (
	statusDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/status"
	statusCategoryModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/statusCategory"
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	ID               int64 `gorm:"primaryKey;autoIncrement:true"`
	ID_JIRA          string
	DESCRICAO        string `gorm:"unique;not null"`
	StatusCategoryId int64
	StatusCategory   statusCategoryModel.StatusCategory
}

func (Status) TableName() string {
	return "status"
}

func NewStatus(status *statusDto.Status) *Status {
	return &Status{
		ID_JIRA:   status.ID,
		DESCRICAO: status.Description,
	}
}
