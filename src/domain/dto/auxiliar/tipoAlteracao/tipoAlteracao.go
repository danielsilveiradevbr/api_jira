package tipoAlteracaoDto

type TipoAlteracao struct {
	Self     string `json:"self"`
	Value    string `json:"value"`
	Id       string `json:"id"`
	Disabled bool   `json:"disabled"`
}
