package sprint

import (
	"fmt"

	sprintDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/sprint"
	sprint "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/sprint"
	"gorm.io/gorm"
)

func SalvaSprint(db *gorm.DB, sprintString []string) (*sprint.Sprint, error) {
	// Exemplo de expressão regular para extrair informações básicas
	sprintDto := sprintDto.NewSprintDto(sprintString)
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
