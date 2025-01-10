package sku

import (
	skuDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/sku"
	skuModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/sku"
	"gorm.io/gorm"
)

func SalvaSku(db *gorm.DB, skus []string) (*skuModel.Sku, error) {
	var sku = &skuModel.Sku{}
	for _, value := range skus {
		var skudto = &skuDto.Sku{
			Name: value,
		}
		sku = skuModel.NewSku(skudto)
		result := db.First(&sku, "Nome = ?", sku.Nome)
		if result.RowsAffected == 0 {
			res := db.Create(&sku)
			if res.Error != nil {
				return nil, res.Error
			}
		} else {
			db.Save(&sku)
		}
	}

	return sku, nil
}
