package controllers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"encoding/base64"
	"encoding/json"

	"github.com/danielzinhors/api_jira/src/domain/model/dds"
	"github.com/joho/godotenv"
)

func AtualizaDDS() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	endpoint := os.Getenv("ENDPOINT_JIRA")
	jsonVar := bytes.NewBuffer([]byte(`{"jql":"project = \"Desenvolvimento de Software\" AND Sprint = 75"}`))
	auth := os.Getenv("USER_JIRA") + ":" + os.Getenv("PASS_JIRA")
	base64Auth := base64.StdEncoding.EncodeToString([]byte(auth))
	req, err := http.NewRequest("POST", endpoint, jsonVar)
	if err != nil {
		return err
	}

	// Adicione o cabeçalho de autenticação básica
	req.Header.Set("Authorization", "Basic "+base64Auth)
	fmt.Println(base64Auth)
	// Faça a solicitação HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Println("passou")
	defer resp.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	var dds *dds.DDS
	err = json.Unmarshal(res, &dds)
	if err != nil {
		return err
	}

	println(&dds.Total)
	return nil
}
