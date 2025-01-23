package atualizaDDS

import (
	"fmt"
	"strings"
	"time"

	"github.com/danielsilveiradevbr/api_jira/src/application/usecases/buscaDDS"
	sprintModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/sprint"
	helper "github.com/danielsilveiradevbr/api_jira/src/helpers"
	"github.com/danielsilveiradevbr/api_jira/src/infra/banco"
	ddsservice "github.com/danielsilveiradevbr/api_jira/src/service/ddsService"
)

func AtualizaDDS() {
	for {
		hora := helper.GetADateTimeSaoPaulo()
		fmt.Println(hora)
		db, err := banco.ConnectToPG()
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
