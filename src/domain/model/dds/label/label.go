package labelModel

import (
	labelDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/label"
	"gorm.io/gorm"
)

type Label struct {
	gorm.Model
	ID   int64  `gorm:"primaryKey;autoIncrement:true"`
	Nome string `gorm:"unique;not null"`
}

func NewLabel(labeldto *labelDto.Label) *Label {
	var label = &Label{
		Nome: labeldto.Name,
	}

	return label
}
