package project

import (
	projectDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/project"
	projectModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/project"
	projectCategoryRep "github.com/danielsilveiradevbr/api_jira/src/repositories/projectCategory"
	"gorm.io/gorm"
)

func SalvaProject(db *gorm.DB, projetoDTO *projectDto.Project) (*project.Project, error) {
	projectCategory, err := projectCategoryRep.SalvaProjectCategory(db, &projetoDTO.ProjectCategory)
	if err != nil {
		return nil, err
	}
	projeto := projectModel.NewProject(projetoDTO)
	projeto.ProjectCategoryID = projectCategory.ID
	result := db.First(&projeto, "descricao = ?", projeto.DESCRICAO)
	if result.RowsAffected == 0 {
		res := db.Create(&projeto)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		db.Save(&projeto)
	}
	return projeto, nil
}
