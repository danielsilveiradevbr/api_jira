package label

import (
	labelDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	labelModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds"
	"gorm.io/gorm"
)

func Salvalabel(db *gorm.DB, labels []string) error {
	for _, value := range labels {
		var labeldto = &labelDTO.Label{
			Name: value,
		}
		var label = labelModel.NewLabel(labeldto)
		result := db.First(&label, "Nome = ?", label.Nome)
		if result.RowsAffected == 0 {
			res := db.Create(&label)
			if res.Error != nil {
				return res.Error
			}
		} else {
			db.Save(&label)
		}
	}

	return nil
}
