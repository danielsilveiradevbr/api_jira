package creator

import (
	creatorDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	creatorModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds"
	"gorm.io/gorm"
)

func SalvaCreator(db *gorm.DB, creatorDTO *creatorDTO.Creator) error {
	creator := creatorModel.NewCreator(creatorDTO)
	result := db.First(&creator, "email = ?", creator.Email)
	if result.RowsAffected == 0 {
		res := db.Create(&creator)
		if res.Error != nil {
			return res.Error
		}
	} else {
		db.Save(&creator)
	}
	return nil
}
