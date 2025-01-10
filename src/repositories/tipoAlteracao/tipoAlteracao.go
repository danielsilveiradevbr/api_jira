package tipoalteracao

import (
	tipoAlteracaoDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/tipoAlteracao"
	tipoAlteracaoModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/tipoAlteracao"
	"gorm.io/gorm"
)

func SalvatipoAlteracao(db *gorm.DB, tipoAlteracaoDTO *tipoAlteracaoDto.TipoAlteracao) error {

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
