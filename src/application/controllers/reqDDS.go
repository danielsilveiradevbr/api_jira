package controllers

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"encoding/base64"
	"encoding/json"

	dds "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"github.com/joho/godotenv"
)

func AtualizaDDS() (*dds.DDS, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	endpoint := os.Getenv("ENDPOINT_JIRA")

	jsonVar := bytes.NewBuffer([]byte(`{"jql":"project = \"Desenvolvimento de Software\" AND Sprint = 75"}`))
	auth := os.Getenv("USER_JIRA") + ":" + os.Getenv("PASS_JIRA")

	base64Auth := base64.StdEncoding.EncodeToString([]byte(auth))
	req, err := http.NewRequest("GET", endpoint, jsonVar)
	if err != nil {
		return nil, err
	}
	// Adicione o cabeçalho de autenticação básica
	req.Header.Add("Authorization", "Basic "+base64Auth)

	// Faça a solicitação HTTP
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// fmt.Println("Response status:", resp.Status)
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jsondds *dds.DDS

	err = json.Unmarshal(res, &jsondds)
	if err != nil {
		return nil, err
	}

	return jsondds, nil
}
