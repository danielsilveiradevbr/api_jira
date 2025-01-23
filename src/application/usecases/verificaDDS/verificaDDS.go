package verificaDDS

import (
	"net/http"
	"os"

	b "github.com/danielsilveiradevbr/api_jira/src/infra/banco"
	"github.com/danielsilveiradevbr/api_jira/src/infra/jira"
	ddsservice "github.com/danielsilveiradevbr/api_jira/src/service/ddsService"
	"github.com/joho/godotenv"
)

func Verificadds(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	db, err := b.ConnectToPG()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	if os.Getenv("DEBUGANDO") == "T" {
		println("Verificando")
	}
	jsonJira, err := jira.BuscaDDS("sprint in (openSprints())")
	if err != nil {
		if err.Error() != "Sprint with id 11 does not exist or you do not have permission to view it." {
			panic(err)
		}
	}
	if os.Getenv("DEBUGANDO") == "T" {
		println(jsonJira)
	}
	if jsonJira.Total > 0 {
		if _, err := ddsservice.SalvaDDS(db, jsonJira); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		if os.Getenv("DEBUGANDO") == "T" {
			println("Verificou")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Verificação realizada com sucesso!"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Verificação realizada com sucesso!"))
}
