package dds

import (
	complexidadeDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"gorm.io/gorm"
)

type Complexidade struct {
	gorm.Model
	ID     int64  `gorm:"primaryKey;autoIncrement:true"`
	Nome   string `gorm:"unique;not null;not null"`
	Codigo string
}

func NewComplexidade(complexidadeDTO *complexidadeDTO.Complexidade) *Complexidade {
	var nome = "Baixa"
	if complexidadeDTO.Codigo == "2" {
		nome = "Média"
	} else if complexidadeDTO.Codigo == "3" {
		nome = "Alta"
	}
	return &Complexidade{
		Nome:   nome,
		Codigo: complexidadeDTO.Codigo,
	}
}