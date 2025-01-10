package ddsservice

import (
	jsonDdsDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/jsonDDS"
	taskModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/task"
	b "github.com/danielsilveiradevbr/api_jira/src/infra/banco"
	taskRep "github.com/danielsilveiradevbr/api_jira/src/repositories/task"
)

func SalvaDDS(DDSJson *jsonDdsDto.JsonDDS) (*taskModel.Task, error) {
	db, err := b.ConnectToPG()
	if err != nil {
		return nil, err
	}
	var task = &taskModel.Task{}
	for _, issue := range DDSJson.Issues {
		task, err = taskRep.SalvaTask(db, &issue)
		if err != nil {
			return nil, err
		}
	}

	return task, nil
}
