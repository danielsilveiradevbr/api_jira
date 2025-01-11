package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	jsonDdsDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/jsonDDS"
	b "github.com/danielsilveiradevbr/api_jira/src/infra/banco"
	service "github.com/danielsilveiradevbr/api_jira/src/service/ddsService"
	"github.com/joho/godotenv"
)

func RecebeDDS(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	db, err := b.ConnectToPG()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	res, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	var jsonDDS *jsonDdsDto.JsonDDS

	err = json.Unmarshal(res, &jsonDDS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	// println(jsonDDS.)
	task, err := service.SalvaDDS(db, jsonDDS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Task criada com sucesso id %d ", task.ID)))
}
