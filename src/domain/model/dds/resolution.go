package dds

import (
	resolutionDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"gorm.io/gorm"
)

type Resolution struct {
	gorm.Model
	ID        int64 `gorm:"primaryKey;autoIncrement:true"`
	Url       string
	Descricao string
	Name      string `gorm:"unique"`
}

func NewResolution(resolutionDTO *resolutionDTO.Resolution) *Resolution {
	return &Resolution{
		Url:       resolutionDTO.Self,
		Descricao: resolutionDTO.Description,
		Name:      resolutionDTO.Name,
	}
}
