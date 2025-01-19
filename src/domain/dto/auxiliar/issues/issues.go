package issuesDto

import fieldsDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/fields"

type Issues struct {
	Expand string           `json:"expand"`
	ID     string           `json:"id"`
	Self   string           `json:"self"`
	Key    string           `json:"key"`
	Fields fieldsDto.Fields `json:"fields"`
}
