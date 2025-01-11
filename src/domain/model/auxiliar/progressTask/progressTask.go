package progressTaskModel

import (
	progressTaskDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/progressTask"
	"gorm.io/gorm"
)

type ProgressTask struct {
	gorm.Model
	ID         int64 `gorm:"primaryKey;autoIncrement:true"`
	TaskId     int64 `gorm:"primaryKey"`
	ProgressId int64 `gorm:"primaryKey"`
}

func NewProgressTask(progressTaskDto *progressTaskDto.ProgressTask) *ProgressTask {
	return &ProgressTask{
		TaskId:     progressTaskDto.TaskId,
		ProgressId: progressTaskDto.ProgressId,
	}
}
