package dds

import "github.com/danielzinhors/api_jira/src/domain/model/dds"

type DDS struct {
	Expand     string        `json:"expand"`
	StartAt    int           `json:"startAt"`
	MaxResults int           `json:"maxResults"`
	Total      int           `json:"total"`
	Issues     []*dds.Issues `json:"issues"`
}
