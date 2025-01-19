package requerAnaliseTecnica

import (
	requerAnaliseTecnicaDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/requerAnaliseTecnica"
	requerAnaliseTecnicaModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/requerAnaliseTecnica"
	"gorm.io/gorm"
)

func SalvaRequerAnaliseTecnica(db *gorm.DB, requerAnaliseTecnicaDTO *requerAnaliseTecnicaDto.RequerAnaliseTecnica) (*requerAnaliseTecnicaModel.RequerAnaliseTecnica, error) {
	requerAnaliseTecnica := requerAnaliseTecnicaModel.NewRequerAnaliseTecnica(requerAnaliseTecnicaDTO)
	result := db.First(&requerAnaliseTecnica, "value = ?", requerAnaliseTecnica.Value)
	if result.RowsAffected == 0 {
		res := db.Create(&requerAnaliseTecnica)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		db.Save(&requerAnaliseTecnica)
	}
	return requerAnaliseTecnica, nil
}
