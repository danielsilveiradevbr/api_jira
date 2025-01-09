package dds

import (
	requerAnaliseTecnicaDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"gorm.io/gorm"
)

type RequerAnaliseTecnica struct {
	gorm.Model
	Id       int64 `gorm:"primaryKey;autoIncrement:true"`
	Self     string
	Value    string `gorm:"unique"`
	Disabled bool
}

func NewRequerAnaliseTecnica(requerAnaliseTecnicaDTO *requerAnaliseTecnicaDTO.RequerAnaliseTecnica) *RequerAnaliseTecnica {
	return &RequerAnaliseTecnica{
		Self:     requerAnaliseTecnicaDTO.Self,
		Value:    requerAnaliseTecnicaDTO.Value,
		Disabled: requerAnaliseTecnicaDTO.Disabled,
	}
}
