package projectCategoryModel

import (
	projectCategoryDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"gorm.io/gorm"
)

type ProjectCategory struct {
	gorm.Model
	ID        int64  `gorm:"primaryKey;autoIncrement:true"`
	Nome      string `gorm:"unique;not null"`
	Descricao string
}

func NewProjectCategory(projectCategoryDTO *projectCategoryDTO.ProjectCategory) *ProjectCategory {
	return &ProjectCategory{
		Nome:      projectCategoryDTO.Name,
		Descricao: projectCategoryDTO.Description,
	}
}
