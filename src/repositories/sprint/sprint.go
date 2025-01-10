package sprint

import (
	"fmt"

	dtoSprint "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	sprint "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds"
	"gorm.io/gorm"
)

func SalvaSprint(db *gorm.DB, sprintString []string) (*sprint.Sprint, error) {
	// Exemplo de expressão regular para extrair informações básicas
	sprintDto := dtoSprint.NewSprintDto(sprintString)
	if sprintDto == nil {
		return nil, fmt.Errorf("nao foi possivel criar a sprint")
	}
	sprint := sprint.NewSprint(sprintDto)
	result := db.First(&sprint, "nome = ?", sprint.NOME)
	if result.RowsAffected == 0 {
		res := db.Create(&sprint)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		db.Save(&sprint)
	}
	return sprint, nil
}
