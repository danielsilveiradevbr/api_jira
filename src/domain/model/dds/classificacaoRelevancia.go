package dds

import (
	classificacaoRelevanciaDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"gorm.io/gorm"
)

type ClassificacaoRelevancia struct {
	gorm.Model
	Id       int64 `gorm:"primaryKey;autoIncrement:true"`
	Self     string
	Value    string `gorm:"unique"`
	Disabled bool
}

func NewClassificacaoRelevancia(classificacaoRelevanciaDTO *classificacaoRelevanciaDTO.ClassificacaoRelevancia) *ClassificacaoRelevancia {
	return &ClassificacaoRelevancia{
		Self:     classificacaoRelevanciaDTO.Self,
		Value:    classificacaoRelevanciaDTO.Value,
		Disabled: classificacaoRelevanciaDTO.Disabled,
	}
}
