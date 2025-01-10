package skuModel

import (
	skuDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/sku"
	"gorm.io/gorm"
)

type Sku struct {
	gorm.Model
	ID   int64  `gorm:"primaryKey;autoIncrement:true"`
	Nome string `gorm:"unique;not null"`
}

func NewSku(skuDTO *skuDto.Sku) *Sku {
	var sku = &Sku{
		Nome: skuDTO.Name,
	}

	return sku
}
