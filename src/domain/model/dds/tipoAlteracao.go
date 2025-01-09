package dds

import (
	tipoAlteracaoDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"gorm.io/gorm"
)

type TipoAlteracao struct {
	gorm.Model
	Id       int64 `gorm:"primaryKey;autoIncrement:true"`
	Self     string
	Value    string `gorm:"unique"`
	Disabled bool
}

func NewTipoAlteracao(tipoAlteracaoDTO *tipoAlteracaoDTO.TipoAlteracao) *TipoAlteracao {
	return &TipoAlteracao{
		Self:     tipoAlteracaoDTO.Self,
		Value:    tipoAlteracaoDTO.Value,
		Disabled: tipoAlteracaoDTO.Disabled,
	}
}
