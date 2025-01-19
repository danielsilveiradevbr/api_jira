package progressModel

import (
	progressDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/progress"
	"gorm.io/gorm"
)

type Progress struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement:true"`
	Progress int
	Total    int
	Percent  float32
}

func (Progress) TableName() string {
	return "progress"
}

func NewProgress(progressDto *progressDto.Progress) *Progress {
	return &Progress{
		Progress: progressDto.Progress,
		Total:    progressDto.Total,
		Percent:  progressDto.Percent,
	}
}
