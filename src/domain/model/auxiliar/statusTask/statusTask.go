package statusTaskModel

import (
	statusTaskDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/statusTask"
	"gorm.io/gorm"
)

type StatusTask struct {
	gorm.Model
	ID       int64 `gorm:"primaryKey;autoIncrement:true"`
	TaskId   int64 `gorm:"primaryKey"`
	StatusId int64 `gorm:"primaryKey"`
}

func NewStatusTask(statusTaskDto *statusTaskDto.StatusTask) *StatusTask {
	return &StatusTask{
		TaskId:   statusTaskDto.TaskId,
		StatusId: statusTaskDto.StatusId,
	}
}
