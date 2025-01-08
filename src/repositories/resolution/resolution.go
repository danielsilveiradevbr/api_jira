package resolution

import (
	resolutionDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	resolutionModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds"
	"gorm.io/gorm"
)

func SalvaResolution(db *gorm.DB, resolutionDTO *resolutionDTO.Resolution) error {

	resolution := resolutionModel.NewResolution(resolutionDTO)
	result := db.First(&resolution, "name = ?", resolution.Name)
	if result.RowsAffected == 0 {
		res := db.Create(&resolution)
		if res.Error != nil {
			return res.Error
		}
	} else {
		db.Save(&resolution)
	}
	return nil
}
