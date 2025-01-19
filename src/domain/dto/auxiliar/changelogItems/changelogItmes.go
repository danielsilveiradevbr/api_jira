package changelogItemsDto

type Items struct {
	Field      string `json:"field"`
	Fieldtype  string `json:"fieldtype"`
	From       string `json:"from"`
	FromString string `json:"fromString"`
	To         string `json:"to"`
}
