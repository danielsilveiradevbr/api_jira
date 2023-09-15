package model

import "github.com/danielzinhors/api_jira/src/model"

type Status struct {
	Self           string                `json:"self"`
	Description    string                `json:"description"`
	IconURL        string                `json:"iconUrl"`
	Name           string                `json:"name"`
	ID             string                `json:"id"`
	StatusCategory *model.StatusCategory `json:"statusCategory"`
}
