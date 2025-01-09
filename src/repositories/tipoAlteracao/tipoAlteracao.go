package tipoalteracao

import (
	tipoAlteracaoDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	tipoAlteracaoModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds"
	"gorm.io/gorm"
)

func SalvatipoAlteracao(db *gorm.DB, tipoAlteracaoDTO *tipoAlteracaoDTO.TipoAlteracao) error {

	tipoAlteracao := tipoAlteracaoModel.NewTipoAlteracao(tipoAlteracaoDTO)
	result := db.First(&tipoAlteracao, "value = ?", tipoAlteracao.Value)
	if result.RowsAffected == 0 {
		res := db.Create(&tipoAlteracao)
		if res.Error != nil {
			return res.Error
		}
	} else {
		db.Save(&tipoAlteracao)
	}
	return nil
}
