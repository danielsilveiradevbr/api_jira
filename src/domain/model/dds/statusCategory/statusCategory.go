package statusCategoryModel

import (
	statusCategoryDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/statusCategory"
	"gorm.io/gorm"
)

type StatusCategory struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey;autoIncrement:true"`
	Nome      string `gorm:"unique;not null"`
	ColorName string
	Key       string
}

func NewStatusCategory(statusCategoryDTO *statusCategoryDto.StatusCategory) *StatusCategory {
	return &StatusCategory{
		Nome:      statusCategoryDTO.Name,
		ColorName: statusCategoryDTO.ColorName,
		Key:       statusCategoryDTO.Key,
	}
}
