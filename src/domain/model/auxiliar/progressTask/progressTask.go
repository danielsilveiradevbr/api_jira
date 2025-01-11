package progressTaskModel

import (
	progressTaskDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/progressTask"
	"gorm.io/gorm"
)

type ProgressTask struct {
	gorm.Model
	ID         uint `gorm:"primaryKey;autoIncrement:true"`
	TaskId     uint `gorm:"primaryKey"`
	ProgressId uint `gorm:"primaryKey"`
}

func NewProgressTask(progressTaskDto *progressTaskDto.ProgressTask) *ProgressTask {
	return &ProgressTask{
		TaskId:     progressTaskDto.TaskId,
		ProgressId: progressTaskDto.ProgressId,
	}
}
