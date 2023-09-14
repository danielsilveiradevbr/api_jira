package model

import "github.com/danielzinhors/api_jira/src/model"

type DDS struct {
	Expand     string          `json:"expand"`
	StartAt    int             `json:"startAt"`
	MaxResults int             `json:"maxResults"`
	Total      int             `json:"total"`
	Issues     []*model.Issues `json:"issues"`
}
