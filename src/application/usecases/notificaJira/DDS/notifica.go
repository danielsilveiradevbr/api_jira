package dds

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	notificaJiraDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/notificaJira/DDS"
	telegram "github.com/danielsilveiradevbr/api_jira/src/infra/mensagem"
)

func Notifica(w http.ResponseWriter, r *http.Request) {
	res, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if !strings.Contains(string(res), "changelog") {
		w.WriteHeader(http.StatusOK)
		return
	}

	var jsonNotifica *notificaJiraDto.JsonNotifica
	err = json.Unmarshal(res, &jsonNotifica)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	tipoNoticacao := SetTipoNotificacao(jsonNotifica)
	if tipoNoticacao == "" {
		w.WriteHeader(http.StatusOK)
		return
	}

	key_jira := jsonNotifica.Issue.Key
	err = EnviaNotificacao(tipoNoticacao, key_jira)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Mensagem enviada com sucesso"))
	}
}

func SetTipoNotificacao(jsonNotifica *notificaJiraDto.JsonNotifica) string {
	var tipoNotifica = ""
	var items = jsonNotifica.Changelog.Items
	for _, item := range items {
		if item.Field == "status" {
			tipoNotifica = item.ToString
			break
		}
	}

	return strings.ToUpper(tipoNotifica)
}

func EnviaNotificacao(tipoNotificacao, key_jira string) error {
	var users []string
	if tipoNotificacao == "EM ANÁLISE TÉCNICA TESTE" ||
		tipoNotificacao == "REVISÃO NEGÓCIO" {
		users = strings.Split(os.Getenv("USERS_ENVIO_MSG_TELEGRAM_EM_ANALISE"), "|")
	} else if tipoNotificacao == "EM DESENVOLVIMENTO" ||
		tipoNotificacao == "AGUARDANDO PULL REQUEST" ||
		tipoNotificacao == "REVISÃO TÉCNICA" ||
		tipoNotificacao == "AGUARDANDO DEPLOY" {
		users = strings.Split(os.Getenv("USERS_ENVIO_MSG_TELEGRAM_EM_DEV"), "|")
	}

	if len(users) > 0 {
		telegramBotToken := os.Getenv("TELEGRAM_BOT")
		for _, user := range users {
			err := telegram.SendMessage(telegramBotToken, user, "O "+key_jira+" entrou no status "+tipoNotificacao+" e aguarda sua atenção")
			if err != nil {
				return err
			}
		}
	}
	return nil
}
