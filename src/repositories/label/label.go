package label

import (
	labelDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/label"
	labelModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/label"
	"gorm.io/gorm"
)

func Salvalabel(db *gorm.DB, labels []string) error {
	for _, value := range labels {
		var labeldto = &labelDto.Label{
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
