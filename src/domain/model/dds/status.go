package dds

import "github.com/danielzinhors/api_jira/src/domain/model/dds"

type Status struct {
	Self           string              `json:"self"`
	Description    string              `json:"description"`
	IconURL        string              `json:"iconUrl"`
	Name           string              `json:"name"`
	ID             string              `json:"id"`
	StatusCategory *dds.StatusCategory `json:"statusCategory"`
}
