package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	dds "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/jsonDDS"
	cripto "github.com/danielsilveiradevbr/helpercripto/pkg"
	"github.com/joho/godotenv"
)

func BuscaDDS(sprintFiltro string) (*dds.JsonDDS, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	method := "POST"
	endpoint := cripto.Cripto("D", os.Getenv("ENDPOINT_JIRA"), os.Getenv("KEY"))
	if os.Getenv("DEBUGANDO") == "T" {
		println(endpoint)
	}

	body, _ := json.Marshal(map[string]string{
		"jql": "project = \"Desenvolvimento de Software\" AND " + sprintFiltro,
	})

	payload := bytes.NewBuffer(body)

	req, err := http.NewRequest(method, endpoint, payload)
	if err != nil {
		return nil, err
	}

	// Adicione o cabeçalho de autenticação básica
	req.SetBasicAuth(cripto.Cripto("D", os.Getenv("USER_JIRA"), os.Getenv("KEY")), cripto.Cripto("D", os.Getenv("PASS_JIRA"), os.Getenv("KEY")))
	req.Header.Set("Content-Type", "application/json")
	// Faça a solicitação HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if os.Getenv("DEBUGANDO") == "T" {
		fmt.Println("Response status:", resp.Status)
	}
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
	if strings.Contains(string(res), "does not exist or you do not have permission to view it.") {
		return nil, fmt.Errorf("does not exist or you do not have permission to view it.")
	}
	if os.Getenv("DEBUGANDO") == "T" {
		println(jsonDDS.Total)
	}
	return jsonDDS, nil
}
