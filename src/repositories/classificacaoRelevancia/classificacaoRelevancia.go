package classificacaoRelevancia

import (
	classificacaoRelevanciaDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	classificacaoRelevanciaModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds"
	"gorm.io/gorm"
)

func SalvaclassificacaoRelevancia(db *gorm.DB, classificacaoRelevanciaDTO *classificacaoRelevanciaDTO.ClassificacaoRelevancia) error {

	classificacaoRelevancia := classificacaoRelevanciaModel.NewClassificacaoRelevancia(classificacaoRelevanciaDTO)
	result := db.First(&classificacaoRelevancia, "value = ?", classificacaoRelevancia.Value)
	if result.RowsAffected == 0 {
		res := db.Create(&classificacaoRelevancia)
		if res.Error != nil {
			return res.Error
		}
	} else {
		db.Save(&classificacaoRelevancia)
	}
	return nil
}
