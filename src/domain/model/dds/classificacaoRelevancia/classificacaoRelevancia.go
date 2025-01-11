package classificacaoRelevanciaModel

import (
	classificacaoRelevanciaDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/classificacaorelevancia"
	"gorm.io/gorm"
)

type ClassificacaoRelevancia struct {
	gorm.Model
	Id       uint `gorm:"primaryKey;autoIncrement:true"`
	Self     string
	Value    string `gorm:"unique;not null"`
	Disabled bool
}

func NewClassificacaoRelevancia(classificacaoRelevanciaDTO *classificacaoRelevanciaDTO.ClassificacaoRelevancia) *ClassificacaoRelevancia {
	return &ClassificacaoRelevancia{
		Self:     classificacaoRelevanciaDTO.Self,
		Value:    classificacaoRelevanciaDTO.Value,
		Disabled: classificacaoRelevanciaDTO.Disabled,
	}
}
