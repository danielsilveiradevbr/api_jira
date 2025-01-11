package clienteModel

import (
	clienteDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/cliente"
	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey;autoIncrement:true"`
	Nome string `gorm:"unique;not null"`
}

func NewCliente(clientedto *clienteDto.Cliente) *Cliente {
	var cliente = &Cliente{
		Nome: clientedto.Name,
	}

	return cliente
}
