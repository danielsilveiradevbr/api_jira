package project

import (
	projectDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	project "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds"
	"gorm.io/gorm"
)

func SalvaProject(db *gorm.DB, projetoDTO *projectDTO.Project) error {
	projeto := project.NewProject(projetoDTO)
	result := db.First(&projeto, "descricao = ?", projeto.DESCRICAO)
	if result.RowsAffected == 0 {
		res := db.Create(&projeto)
		if res.Error != nil {
			return res.Error
		}
	} else {
		db.Save(&projeto)
	}
	return nil
}
