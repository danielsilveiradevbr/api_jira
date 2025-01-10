package project

import (
	projectDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	project "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds"
	projectCategoryRep "github.com/danielsilveiradevbr/api_jira/src/repositories/projectCategory"
	"gorm.io/gorm"
)

func SalvaProject(db *gorm.DB, projetoDTO *projectDTO.Project) (*project.Project, error) {

	projectCategory, err := projectCategoryRep.SalvaProjectCategory(db, &projetoDTO.ProjectCategory)
	if err != nil {
		return nil, err
	}

	projeto := project.NewProject(projetoDTO)
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
