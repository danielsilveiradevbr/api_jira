package ddsservice

import (
	"fmt"

	jsonDdsDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/jsonDDS"
	taskModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/task"
	taskRep "github.com/danielsilveiradevbr/api_jira/src/repositories/task"
	"gorm.io/gorm"
)

func SalvaDDS(db *gorm.DB, DDSJson *jsonDdsDto.JsonDDS) (*taskModel.Task, error) {
	var task = &taskModel.Task{}
	var err = fmt.Errorf("")
	for _, issue := range DDSJson.Issues {
		task, err = taskRep.SalvaTask(db, &issue)
		if err != nil {
			return nil, err
		}
	}

	return task, nil
}
