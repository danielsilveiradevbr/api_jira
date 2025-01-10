package sku

import (
	skuDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	skuModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/sku"
	"gorm.io/gorm"
)

func SalvaSku(db *gorm.DB, skus []string) error {
	for _, value := range skus {
		var skudto = &skuDTO.Sku{
			Name: value,
		}
		var sku = skuModel.NewSku(skudto)
		result := db.First(&sku, "Nome = ?", sku.Nome)
		if result.RowsAffected == 0 {
			res := db.Create(&sku)
			if res.Error != nil {
				return res.Error
			}
		} else {
			db.Save(&sku)
		}
	}

	return nil
}
