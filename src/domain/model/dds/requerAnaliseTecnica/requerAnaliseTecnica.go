package requerAnaliseTecnicaModel

import (
	requerAnaliseTecnicaDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/requerAnaliseTecnica"
	"gorm.io/gorm"
)

type RequerAnaliseTecnica struct {
	gorm.Model
	Id       uint `gorm:"primaryKey;autoIncrement:true"`
	Self     string
	Value    string `gorm:"unique;not null"`
	Disabled bool
}

func NewRequerAnaliseTecnica(requerAnaliseTecnicaDTO *requerAnaliseTecnicaDto.RequerAnaliseTecnica) *RequerAnaliseTecnica {
	return &RequerAnaliseTecnica{
		Self:     requerAnaliseTecnicaDTO.Self,
		Value:    requerAnaliseTecnicaDTO.Value,
		Disabled: requerAnaliseTecnicaDTO.Disabled,
	}
}
