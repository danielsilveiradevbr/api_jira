package resolution

import (
	resolutionDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/resolution"
	resolutionModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/resolution"
	"gorm.io/gorm"
)

func SalvaResolution(db *gorm.DB, resolutionDTO *resolutionDto.Resolution) (*resolutionModel.Resolution, error) {

	resolution := resolutionModel.NewResolution(resolutionDTO)
	result := db.First(&resolution, "name = ?", resolution.Name)
	if result.RowsAffected == 0 {
		res := db.Create(&resolution)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		db.Save(&resolution)
	}
	return resolution, nil
}
