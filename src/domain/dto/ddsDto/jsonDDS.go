package ddsDto

type JsonDDS struct {
	Expand     string   `json:"expand"`
	StartAt    int      `json:"startAt"`
	MaxResults int      `json:"maxResults"`
	Total      int      `json:"total"`
	Issues     []Issues `json:"issues"`
}