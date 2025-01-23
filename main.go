package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/danielsilveiradevbr/api_jira/src/application/usecases/buscaDDS"
	"github.com/danielsilveiradevbr/api_jira/src/application/usecases/cripto"
	dds "github.com/danielsilveiradevbr/api_jira/src/application/usecases/notificaJira/DDS"
	"github.com/danielsilveiradevbr/api_jira/src/application/usecases/recebeDDS"
	"github.com/danielsilveiradevbr/api_jira/src/application/usecases/verificaDDS"
	classificacaoRelevanciaModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/classificacaoRelevancia"
	clienteModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/cliente"
	complexidadeModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/complexidade"
	issueTypeModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/issueType"
	labelModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/label"
	priorityModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/priority"
	progressModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/progress"
	progressTaskModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/progressTask"
	projectModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/project"
	requerAnaliseTecnicaModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/requerAnaliseTecnica"
	requerDocumentacaoModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/requerDocumentacao"
	resolutionModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/resolution"
	skuModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/sku"
	sprintModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/sprint"
	statusModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/status"
	statusTaskModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/statusTask"
	tipoAlteracaoModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/tipoAlteracao"
	userModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/user"
	taskModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/task"
	helper "github.com/danielsilveiradevbr/api_jira/src/helpers"
	b "github.com/danielsilveiradevbr/api_jira/src/infra/banco"
	ddsservice "github.com/danielsilveiradevbr/api_jira/src/service/ddsService"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	r.Use(helper.BasicAuth)
	r.Post("/dds", recebeDDS.RecebeDDS)
	r.Get("/verificadds", verificaDDS.Verificadds)
	r.Post("/cripto", cripto.Cripto)
	r.Post("/notifica", dds.Notifica)

	porta := os.Getenv("PORT_LISTEN")
	if porta == "" {
		porta = "3000"
	}
	porta = ":" + porta
	println(porta)
	go AtualizaDDS()

	http.ListenAndServe(porta, r)
}

func AtualizaDDS() {
	for {
		hora := helper.GetADateTimeSaoPaulo()
		fmt.Println(hora)
		db, err := b.ConnectToPG()
		if err != nil {
			panic(err)
		}
		if hora.Hour() == 16 { //23 && hora.Minute() == 59 && hora.Second() == 0 {
			var sprints []sprintModel.Sprint
			db.Where("STATUS <> ?", "FUTURE").Find(&sprints).Order("id")
			if len(sprints) > 0 {
				for _, v := range sprints {
					jsonJira, err := buscaDDS.BuscaDDS(fmt.Sprintf(" sprint = %s", v.ID_JIRA))
					if err != nil {
						panic(err)
					}
					if jsonJira.Total > 0 {
						if _, err := ddsservice.SalvaDDS(db, jsonJira); err != nil {
							panic(err)
						}
					}

				}
			} else {
				ind := 0
				for {
					ind = ind + 1
					fmt.Printf("ind num %d", ind)
					jsonJira, err := buscaDDS.BuscaDDS(fmt.Sprintf(" sprint = %d", ind))
					if err != nil {
						if !strings.Contains(err.Error(), "does not exist or you do not have permission to view it.") {
							fmt.Println("Deu erro")
							panic(err)
						}
					} else {
						if jsonJira.Total > 0 {
							if _, err := ddsservice.SalvaDDS(db, jsonJira); err != nil {
								panic(err)
							}
						} else {
							break
						}
					}
				}

			}

			jsonJira, err := buscaDDS.BuscaDDS("sprint in (openSprints())")
			if err != nil {
				if !strings.Contains(err.Error(), "does not exist or you do not have permission to view it.") {
					panic(err)
				}
			} else {
				if jsonJira.Total > 0 {
					if _, err := ddsservice.SalvaDDS(db, jsonJira); err != nil {
						panic(err)
					}
				}
			}
		}
		time.Sleep(time.Minute)
	}
}

func init() {
	db, err := b.ConnectToPG()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&priorityModel.Priority{})
	db.AutoMigrate(&issueTypeModel.IssueType{})
	db.AutoMigrate(&projectModel.Project{})
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&sprintModel.Sprint{})
	db.AutoMigrate(&statusModel.Status{})
	db.AutoMigrate(&resolutionModel.Resolution{})
	db.AutoMigrate(&tipoAlteracaoModel.TipoAlteracao{})
	db.AutoMigrate(&classificacaoRelevanciaModel.ClassificacaoRelevancia{})
	db.AutoMigrate(&requerDocumentacaoModel.RequerDocumentacao{})
	db.AutoMigrate(&requerAnaliseTecnicaModel.RequerAnaliseTecnica{})
	db.AutoMigrate(&clienteModel.Cliente{})
	db.AutoMigrate(&skuModel.Sku{})
	db.AutoMigrate(&labelModel.Label{})
	db.AutoMigrate(&complexidadeModel.Complexidade{})
	db.AutoMigrate(&taskModel.Task{})
	db.AutoMigrate(&statusTaskModel.StatusTask{})
	db.AutoMigrate(&progressModel.Progress{})
	db.AutoMigrate(&progressTaskModel.ProgressTask{})
}
