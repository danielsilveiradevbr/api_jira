package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	controllers "github.com/danielsilveiradevbr/api_jira/src/application/controllers"
	progressTaskModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/progressTask"
	statusTaskModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/statusTask"
	classificacaoRelevanciaModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/classificacaoRelevancia"
	clienteModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/cliente"
	complexidadeModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/complexidade"
	issueTypeModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/issueType"
	labelModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/label"
	priorityModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/priority"
	progressModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/progress"
	projectModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/project"
	requerAnaliseTecnicaModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/requerAnaliseTecnica"
	requerDocumentacaoModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/requerDocumentacao"
	resolutionModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/resolution"
	skuModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/sku"
	sprintModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/sprint"
	statusModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/status"
	taskModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/task"
	tipoAlteracaoModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/tipoAlteracao"
	userModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/user"
	b "github.com/danielsilveiradevbr/api_jira/src/infra/banco"
	service "github.com/danielsilveiradevbr/api_jira/src/service/ddsService"
	u "github.com/danielsilveiradevbr/api_jira/src/utils"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	r.Post("/dds", controllers.RecebeDDS)

	porta := os.Getenv("PORT")
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
		hora := u.GetADateTimeSaoPaulo()
		fmt.Println(hora)
		if hora.Hour() == 23 && hora.Minute() == 59 && hora.Second() == 0 {
			jsonJira, err := controllers.AtualizaDDS()
			if err != nil {
				panic(err)
			}
			service.SalvaDDS(jsonJira)
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
