package tipoUserModel

import (
	tipoUserDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/tipoUser"
	"gorm.io/gorm"
)

type TipoUser struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey;autoIncrement:true"`
	Nome string `gorm:"unique;not null"`
}

func NewTipoUser(tipoUserdto *tipoUserDto.TipoUser) *TipoUser {
	return &TipoUser{
		Nome: tipoUserdto.Nome,
	}
}
