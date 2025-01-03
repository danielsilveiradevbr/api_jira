package controllers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"encoding/json"

	dds "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"github.com/joho/godotenv"
)

func AtualizaDDS() (*dds.JsonDDS, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	method := "POST"
	endpoint := os.Getenv("ENDPOINT_JIRA")
	if os.Getenv("DEBUGANDO") == "T" {
		println(endpoint)
	}

	body, _ := json.Marshal(map[string]string{
		"jql": "project = \"Desenvolvimento de Software\" AND Sprint = 75",
	})

	payload := bytes.NewBuffer(body)

	req, err := http.NewRequest(method, endpoint, payload)
	if err != nil {
		return nil, err
	}

	// Adicione o cabeçalho de autenticação básica
	req.SetBasicAuth(os.Getenv("USER_JIRA"), os.Getenv("PASS_JIRA"))
	req.Header.Set("Content-Type", "application/json")
	// Faça a solicitação HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println("Response status:", resp.Status)
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jsonDDS *dds.JsonDDS
	if os.Getenv("DEBUGANDO") == "T" {
		println(string(res))
	}
	err = json.Unmarshal(res, &jsonDDS)
	if err != nil {
		return nil, err
	}
	if os.Getenv("DEBUGANDO") == "T" {
		println(jsonDDS.Total)
	}
	return jsonDDS, nil
}
