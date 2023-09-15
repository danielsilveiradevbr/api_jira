package controllers

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"encoding/base64"

	"github.com/joho/godotenv"
)

func AtualizaDDS() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	endpoint := os.Getenv("ENDPOINT_JIRA")
	jsonVar := bytes.NewBuffer([]byte(`{"jql":"project = \"Desenvolvimento de Software\" AND Sprint = 76"}`))
	auth := os.Getenv("USER_JIRA") + ":" + os.Getenv("PASS_JIRA")
	base64Auth := base64.StdEncoding.EncodeToString([]byte(auth))
	req, err := http.NewRequest("POST", endpoint, jsonVar)
	if err != nil {
		return err
	}

	// Adicione o cabeçalho de autenticação básica
	req.Header.Set("Authorization", "Basic "+base64Auth)

	// Faça a solicitação HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	return nil
}
