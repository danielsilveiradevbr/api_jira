package classificacaoRelevancia

import (
	classificacaoRelevanciaDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/classificacaorelevancia"
	classificacaoRelevanciaModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/classificacaoRelevancia"
	"gorm.io/gorm"
)

func SalvaClassificacaoRelevancia(db *gorm.DB, classificacaoRelevanciaDTO *classificacaoRelevanciaDto.ClassificacaoRelevancia) (*classificacaoRelevanciaModel.ClassificacaoRelevancia, error) {

	classificacaoRelevancia := classificacaoRelevanciaModel.NewClassificacaoRelevancia(classificacaoRelevanciaDTO)
	result := db.First(&classificacaoRelevancia, "value = ?", classificacaoRelevancia.Value)
	if result.RowsAffected == 0 {
		res := db.Create(&classificacaoRelevancia)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		db.Save(&classificacaoRelevancia)
	}
	return classificacaoRelevancia, nil
}
