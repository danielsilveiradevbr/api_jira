package projectcategory

import (
	projectCategoryDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	projectCategoryModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/projectCategory"
	"gorm.io/gorm"
)

func SalvaProjectCategory(db *gorm.DB, projectCategoryDTO *projectCategoryDTO.ProjectCategory) (*projectCategoryModel.ProjectCategory, error) {
	projectCategory := projectCategoryModel.NewProjectCategory(projectCategoryDTO)
	result := db.First(&projectCategory, "nome = ?", projectCategory.Nome)
	if result.RowsAffected == 0 {
		res := db.Create(&projectCategory)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		db.Save(&projectCategory)
	}
	return projectCategory, nil
}
