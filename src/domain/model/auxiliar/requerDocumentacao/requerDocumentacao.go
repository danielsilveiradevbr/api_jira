package requerDocumentacaoModel

import (
	requerDocumentacaoDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/requerDocumentacao"
	"gorm.io/gorm"
)

type RequerDocumentacao struct {
	gorm.Model
	Id       uint `gorm:"primaryKey;autoIncrement:true"`
	Self     string
	Value    string `gorm:"unique;not null"`
	Disabled bool
}

func (RequerDocumentacao) TableName() string {
	return "requer_documentacoes"
}

func NewRequerDocumentacao(requerDocumentacaoDTO *requerDocumentacaoDto.RequerDocumentacao) *RequerDocumentacao {
	return &RequerDocumentacao{
		Self:     requerDocumentacaoDTO.Self,
		Value:    requerDocumentacaoDTO.Value,
		Disabled: requerDocumentacaoDTO.Disabled,
	}
}
