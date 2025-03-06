package verificaAPI

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	helper "github.com/danielsilveiradevbr/api_jira/src/helpers"
	telegram "github.com/danielsilveiradevbr/api_jira/src/infra/mensagem"
	//cripto "github.com/danielsilveiradevbr/helpercripto/pkg"
)

func VerificaAPI() {
	apisTestar := strings.Split(os.Getenv("APIS_TESTAR"), "|")
	helper.NewLog(1, "Iniciou o teste de apis")
	if len(apisTestar) > 0 {
		for {
			for _, apiTestar := range apisTestar {
				api := strings.Split(apiTestar, ",")

				if len(api) > 0 {
					if strings.Trim(api[0], "") != "" {
						for i := 0; i < 3; i++ {
							c := http.Client{Timeout: 5 * time.Second}
							resp, err := c.Get(api[0])
							if err != nil {
								if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
									helper.NewLog(2, "Tempo limite de conexão")
									EnviaNotificacao(fmt.Sprintf("API %s NO ENDERECO  %s TEVE UM ERRO DE TIMEOUT DE CONEXAO", api[1], api[0]))
								} else if os.IsTimeout(err) {
									helper.NewLog(2, "Tempo limite de operação")
									EnviaNotificacao(fmt.Sprintf("API %s NO ENDERECO  %s TEVE UM ERRO DE TIMEOUT DE OPERACAO", api[1], api[0]))
								} else {
									helper.NewLog(2, "Erro de conexão: "+err.Error())
								}
								continue
							}
							defer resp.Body.Close()
							body, err := io.ReadAll(resp.Body)
							if err != nil {
								helper.NewLog(2, err.Error())
							}

							if string(body) != "ONLINE" {
								helper.NewLog(2, "erro na api "+api[0])
								EnviaNotificacao(fmt.Sprintf("API %s NÃO RESPONDEU NO ENDEREÇO %s", api[1], api[0]))
							} else {
								helper.NewLog(2, "ok com api "+api[0])
							}
							break
						}
					}
				}
			}
			time.Sleep(time.Minute * 5)
		}

	}
}

func EnviaNotificacao(prMensagem string) {
	users := strings.Split(os.Getenv("USERS_ENVIO_MSG_TELEGRAM_EM_DEV"), "|") //cripto.Cripto("D", os.Getenv("USERS_ENVIO_MSG_TELEGRAM_EM_DEV"), os.Getenv("KEY")), "|")
	telegramBotToken := os.Getenv("TELEGRAM_BOT")                             //cripto.Cripto("D", os.Getenv("TELEGRAM_BOT"), os.Getenv("KEY"))
	for _, user := range users {
		err := telegram.SendMessage(telegramBotToken, user, prMensagem)
		if err != nil {
			break
		}
	}
}
