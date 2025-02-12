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
		hora := time.Now() //helper.GetADateTimeSaoPaulo()
		helper.NewLog(1, "Hora da consulta "+hora.GoString())

		if hora.Hour() == 23 && hora.Minute() == 59 && hora.Second() == 0 {
			db, err := banco.ConnectToPG()
			if err != nil {
				helper.NewLog(2, err.Error())
			}
			//defer db.Close()
			var sprints []sprintModel.Sprint
			db.Where("STATUS <> ?", "FUTURE").Find(&sprints).Order("id")
			if len(sprints) > 0 {
				for _, v := range sprints {
					jsonJira, err := buscaDDS.BuscaDDS(fmt.Sprintf(" sprint = %s", v.ID_JIRA))
					if err != nil {
						helper.NewLog(2, err.Error())
					}
					if jsonJira.Total > 0 {
						if _, err := ddsservice.SalvaDDS(db, jsonJira); err != nil {
							helper.NewLog(2, err.Error())
						}
					}

				}
			} else {
				ind := 0
				db, err := banco.ConnectToPG()
				if err != nil {
					helper.NewLog(2, err.Error())
				}
				//defer db.Close()
				for {
					ind = ind + 1
					fmt.Printf("ind num %d", ind)
					jsonJira, err := buscaDDS.BuscaDDS(fmt.Sprintf(" sprint = %d", ind))
					if err != nil {
						if !strings.Contains(err.Error(), "does not exist or you do not have permission to view it.") {
							helper.NewLog(2, err.Error())
						}
					} else {
						if jsonJira.Total > 0 {
							if _, err := ddsservice.SalvaDDS(db, jsonJira); err != nil {
								helper.NewLog(2, err.Error())
							}
						} else {
							break
						}
					}
				}

			}

			jsonJira, err := buscaDDS.BuscaDDS("sprint in (openSprints())")
			helper.NewLog(1, "iniciou a busca da sprint aberta")
			if err != nil {
				if !strings.Contains(err.Error(), "does not exist or you do not have permission to view it.") {
					helper.NewLog(2, err.Error())
				}
			} else {
				if jsonJira.Total > 0 {
					helper.NewLog(1, fmt.Sprintf("encontrou %d dds", jsonJira.Total))
					if _, err := ddsservice.SalvaDDS(db, jsonJira); err != nil {
						helper.NewLog(2, err.Error())
					}
				}
			}
		}
		time.Sleep(time.Minute)
	}
}
