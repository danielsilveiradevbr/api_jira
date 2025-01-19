package statuscategory

import (
	statusCategoryDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/statusCategory"
	statusCategoryModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/statusCategory"
	"gorm.io/gorm"
)

func SalvaStatusCategory(db *gorm.DB, statusCategoryDTO *statusCategoryDto.StatusCategory) (*statusCategoryModel.StatusCategory, error) {
	statusCategory := statusCategoryModel.NewStatusCategory(statusCategoryDTO)
	result := db.First(&statusCategory, "nome = ?", statusCategory.Nome)
	if result.RowsAffected == 0 {
		res := db.Create(&statusCategory)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		db.Save(&statusCategory)
	}
	return statusCategory, nil
}
