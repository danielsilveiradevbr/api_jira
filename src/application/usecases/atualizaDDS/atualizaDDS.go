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
		time.Sleep(time.Minute)
		hora := time.Now() //helper.GetADateTimeSaoPaulo()
		helper.NewLog(1, "Hora da consulta "+hora.GoString())

		//if (hora.Hour() >= 23 && hora.Minute() == 58 && hora.Second() == 0) && (hora.Hour() <= 23 && hora.Minute() == 59 && hora.Second() == 59) {

		db, err := banco.ConnectToPG()
		if err != nil {
			helper.NewLog(2, err.Error())
		}
		//defer db.Close()
		var sprints []sprintModel.Sprint
		db.Find(&sprints).Order("id")
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
			//defer db.Close()
			for {
				ind = ind + 1
				helper.NewLog(1, fmt.Sprintf("IND NUM %d", ind))
				time.Sleep(time.Second * 15)
				jsonJira, err := buscaDDS.BuscaDDS(fmt.Sprintf("project = \"Desenvolvimento de Software\" AND sprint in (%d) ", ind))
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

		jsonJira, err := buscaDDS.BuscaDDS("project = \"Desenvolvimento de Software\" AND sprint in (openSprints())")
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
		//		}
	}
}
