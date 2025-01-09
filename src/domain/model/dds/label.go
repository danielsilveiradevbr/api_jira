package dds

import (
	labelDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"gorm.io/gorm"
)

type Label struct {
	gorm.Model
	ID   int64  `gorm:"primaryKey;autoIncrement:true"`
	Nome string `gorm:"unique"`
}

func NewLabel(labeldto *labelDTO.Label) *Label {
	var label = &Label{
		Nome: labeldto.Name,
	}

	return label
}
