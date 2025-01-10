package requerDocumentacao

import (
	requerDocumentacaoDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	requerDocumentacaoModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/requerDocumentacao"
	"gorm.io/gorm"
)

func SalvarequerDocumentacao(db *gorm.DB, requerDocumentacaoDTO *requerDocumentacaoDTO.RequerDocumentacao) error {

	requerDocumentacao := requerDocumentacaoModel.NewRequerDocumentacao(requerDocumentacaoDTO)
	result := db.First(&requerDocumentacao, "value = ?", requerDocumentacao.Value)
	if result.RowsAffected == 0 {
		res := db.Create(&requerDocumentacao)
		if res.Error != nil {
			return res.Error
		}
	} else {
		db.Save(&requerDocumentacao)
	}
	return nil
}
