package dds

import (
	statusCategoryDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"gorm.io/gorm"
)

type StatusCategory struct {
	gorm.Model
	ID        int64  `gorm:"primaryKey;autoIncrement:true"`
	Nome      string `gorm:"unique;not null"`
	ColorName string
	Key       string
}

func NewStatusCategory(statusCategoryDTO *statusCategoryDTO.StatusCategory) *StatusCategory {
	return &StatusCategory{
		Nome:      statusCategoryDTO.Name,
		ColorName: statusCategoryDTO.ColorName,
		Key:       statusCategoryDTO.Key,
	}
}
