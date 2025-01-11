package tipoAlteracaoModel

import (
	tipoAlteracaoDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/tipoAlteracao"
	"gorm.io/gorm"
)

type TipoAlteracao struct {
	gorm.Model
	Id       uint `gorm:"primaryKey;autoIncrement:true"`
	Self     string
	Value    string `gorm:"unique;not null"`
	Disabled bool
}

func (TipoAlteracao) TableName() string {
	return "tipo_alteracoes"
}

func NewTipoAlteracao(tipoAlteracaoDTO *tipoAlteracaoDto.TipoAlteracao) *TipoAlteracao {
	return &TipoAlteracao{
		Self:     tipoAlteracaoDTO.Self,
		Value:    tipoAlteracaoDTO.Value,
		Disabled: tipoAlteracaoDTO.Disabled,
	}
}
