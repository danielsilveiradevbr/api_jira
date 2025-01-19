package complexidade

import (
	complexidadeDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/complexidade"
	complexidadeModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/complexidade"
	"gorm.io/gorm"
)

func SalvaComplexidade(db *gorm.DB, complexidadeDTO *complexidadeDto.Complexidade) (*complexidadeModel.Complexidade, error) {
	complexidade := complexidadeModel.NewComplexidade(complexidadeDTO)
	result := db.First(&complexidade, "codigo = ?", complexidade.Codigo)
	if result.RowsAffected == 0 {
		res := db.Create(&complexidade)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		db.Save(&complexidade)
	}
	return complexidade, nil
}
