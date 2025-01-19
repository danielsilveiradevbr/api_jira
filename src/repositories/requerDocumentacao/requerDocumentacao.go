package requerDocumentacao

import (
	requerDocumentacaoDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/requerDocumentacao"
	requerDocumentacaoModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/requerDocumentacao"
	"gorm.io/gorm"
)

func SalvaRequerDocumentacao(db *gorm.DB, requerDocumentacaoDTO *requerDocumentacaoDto.RequerDocumentacao) (*requerDocumentacaoModel.RequerDocumentacao, error) {
	requerDocumentacao := requerDocumentacaoModel.NewRequerDocumentacao(requerDocumentacaoDTO)
	result := db.First(&requerDocumentacao, "value = ?", requerDocumentacao.Value)
	if result.RowsAffected == 0 {
		res := db.Create(&requerDocumentacao)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		db.Save(&requerDocumentacao)
	}
	return requerDocumentacao, nil
}
