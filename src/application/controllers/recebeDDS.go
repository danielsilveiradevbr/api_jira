package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	dds "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	service "github.com/danielsilveiradevbr/api_jira/src/service/ddsService"
	"github.com/joho/godotenv"
)

func RecebeDDS(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	res, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	var jsonDDS *dds.JsonDDS

	err = json.Unmarshal(res, &jsonDDS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	// println(jsonDDS.)
	err = service.SalvaDDS(jsonDDS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Recebido com sucesso"))
}
