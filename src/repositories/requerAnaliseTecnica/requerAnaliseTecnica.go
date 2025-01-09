package requerAnaliseTecnica

import (
	requerAnaliseTecnicaDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	requerAnaliseTecnicaModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds"
	"gorm.io/gorm"
)

func SalvaRequerAnaliseTecnica(db *gorm.DB, requerAnaliseTecnicaDTO *requerAnaliseTecnicaDTO.RequerAnaliseTecnica) error {

	requerAnaliseTecnica := requerAnaliseTecnicaModel.NewRequerAnaliseTecnica(requerAnaliseTecnicaDTO)
	result := db.First(&requerAnaliseTecnica, "value = ?", requerAnaliseTecnica.Value)
	if result.RowsAffected == 0 {
		res := db.Create(&requerAnaliseTecnica)
		if res.Error != nil {
			return res.Error
		}
	} else {
		db.Save(&requerAnaliseTecnica)
	}
	return nil
}