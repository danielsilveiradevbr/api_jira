package progressTask

import (
	progressTaskDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/progressTask"
	progressTaskModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/progressTask"
	progressModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/progress"
	taskModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/task"
	"gorm.io/gorm"
)

func SalvaProgressTask(db *gorm.DB, progress *progressModel.Progress, task *taskModel.Task) error {
	progressTaskDto := &progressTaskDto.ProgressTask{
		TaskId:    task.ID,
		ProgresId: progress.ID,
	}
	progressTaskModel := progressTaskModel.NewProgressTask(progressTaskDto)
	result := db.First(&progressTaskModel, "taskId = ? and progressId = ?", progressTaskModel.TaskId, progressTaskModel.ProgresId)
	if result.RowsAffected == 0 {
		res := db.Create(&progressTaskModel)
		if res.Error != nil {
			return res.Error
		}
	} else {
		db.Save(&progressTaskModel)
	}
	return nil
}
