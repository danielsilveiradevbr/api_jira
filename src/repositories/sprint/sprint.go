package sprint

import (
	sprintDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	sprint "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds"
	"gorm.io/gorm"
)

func SalvaSprint(db *gorm.DB, sprintDTO *sprintDto.Sprint) error {
	sprint := sprint.NewSprint(sprintDTO)
	println(sprint)
	// result := db.First(&sprint, "nome = ?", sprint.NOME)
	// if result.RowsAffected == 0 {
	// 	res := db.Create(&sprint)
	// 	if res.Error != nil {
	// 		return res.Error
	// 	}
	// } else {
	// 	db.Save(&sprint)
	// }
	return nil
}
