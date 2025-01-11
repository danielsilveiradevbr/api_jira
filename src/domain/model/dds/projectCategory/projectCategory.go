package projectCategoryModel

import (
	projectCategoryDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/projectCategory"
	"gorm.io/gorm"
)

type ProjectCategory struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey;autoIncrement:true"`
	Nome      string `gorm:"unique;not null"`
	Descricao string
}

func NewProjectCategory(projectCategoryDTO *projectCategoryDto.ProjectCategory) *ProjectCategory {
	return &ProjectCategory{
		Nome:      projectCategoryDTO.Name,
		Descricao: projectCategoryDTO.Description,
	}
}
