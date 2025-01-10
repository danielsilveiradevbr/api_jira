package label

import (
	labelDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/label"
	labelModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/label"
	"gorm.io/gorm"
)

func Salvalabel(db *gorm.DB, labels []string) (*labelModel.Label, error) {
	var label = &labelModel.Label{}
	for _, value := range labels {
		var labeldto = &labelDto.Label{
			Name: value,
		}
		label = labelModel.NewLabel(labeldto)
		result := db.First(&label, "Nome = ?", label.Nome)
		if result.RowsAffected == 0 {
			res := db.Create(&label)
			if res.Error != nil {
				return nil, res.Error
			}
		} else {
			db.Save(&label)
		}
	}

	return label, nil
}
