package statusTask

import (
	statusTaskDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/statusTask"
	statusTaskModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/statusTask"
	statusModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/status"
	taskModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/task"
	"gorm.io/gorm"
)

func SalvaStatusTask(db *gorm.DB, status *statusModel.Status, task *taskModel.Task) error {
	statusTaskDto := &statusTaskDto.StatusTask{
		TaskId:   task.ID,
		StatusId: status.ID,
	}
	statusTaskModel := statusTaskModel.NewStatusTask(statusTaskDto)
	result := db.First(&statusTaskModel, "task_id = ? and status_id = ?", statusTaskModel.TaskId, statusTaskModel.StatusId)
	if result.RowsAffected == 0 {
		res := db.Create(&statusTaskModel)
		if res.Error != nil {
			return res.Error
		}
	} else {
		db.Save(&statusTaskModel)
	}
	return nil
}
