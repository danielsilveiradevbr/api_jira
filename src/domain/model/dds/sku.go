package dds

import (
	skuDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"gorm.io/gorm"
)

type Sku struct {
	gorm.Model
	ID   int64  `gorm:"primaryKey;autoIncrement:true"`
	Nome string `gorm:"unique"`
}

func NewSku(skuDTO *skuDTO.Sku) *Sku {
	var sku = &Sku{
		Nome: skuDTO.Name,
	}

	return sku
}
