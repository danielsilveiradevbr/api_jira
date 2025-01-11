package statusModel

import (
	statusDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/status"
	statusCategoryModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/statusCategory"
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	ID               uint `gorm:"primaryKey;autoIncrement:true"`
	ID_JIRA          string
	DESCRICAO        string
	Nome             string `gorm:"unique;not null"`
	StatusCategoryId *uint
	StatusCategory   statusCategoryModel.StatusCategory
}

func (Status) TableName() string {
	return "status"
}

func (s *Status) BeforeCreate(db *gorm.DB) (err error) {
	if *s.StatusCategoryId == 0 {
		s.StatusCategoryId = nil
	}
	return nil
}

func NewStatus(status *statusDto.Status) *Status {
	return &Status{
		ID_JIRA:          status.ID,
		DESCRICAO:        status.Description,
		Nome:             status.Name,
		StatusCategoryId: nil,
	}
}
