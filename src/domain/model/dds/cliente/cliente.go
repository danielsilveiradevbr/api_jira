package clienteModel

import (
	clienteDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	ID   int64  `gorm:"primaryKey;autoIncrement:true"`
	Nome string `gorm:"unique;not null"`
}

func NewCliente(clientedto *clienteDTO.Cliente) *Cliente {
	var cliente = &Cliente{
		Nome: clientedto.Name,
	}

	return cliente
}
