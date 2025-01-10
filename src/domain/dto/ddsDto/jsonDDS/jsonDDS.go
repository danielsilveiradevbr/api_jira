package jsonDdsDto

import issuesDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/issues"

type JsonDDS struct {
	Expand     string             `json:"expand"`
	StartAt    int                `json:"startAt"`
	MaxResults int                `json:"maxResults"`
	Total      int                `json:"total"`
	Issues     []issuesDto.Issues `json:"issues"`
}
