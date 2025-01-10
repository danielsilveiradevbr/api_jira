package creator

import (
	creatorDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/creator"
	creatorModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/user"
	"gorm.io/gorm"
)

func SalvaCreator(db *gorm.DB, creatorDTO *creatorDto.Creator) error {
	creator := creatorModel.NewCreator(db, creatorDTO)
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
