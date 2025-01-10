package complexidade

import (
	complexidadeDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/complexidade"
	complexidadeModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/complexidade"
	"gorm.io/gorm"
)

func SalvaComplexidade(db *gorm.DB, complexidadeDTO *complexidadeDto.Complexidade) error {
	complexidade := complexidadeModel.NewComplexidade(complexidadeDTO)
	result := db.First(&complexidade, "codigo = ?", complexidade.Codigo)
	if result.RowsAffected == 0 {
		res := db.Create(&complexidade)
		if res.Error != nil {
			return res.Error
		}
	} else {
		db.Save(&complexidade)
	}
	return nil
}
