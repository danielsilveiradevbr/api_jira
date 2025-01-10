package progressDto

type Progress struct {
	Progress int     `json:"progress"`
	Total    int     `json:"total"`
	Percent  float32 `json:"percent"`
}
