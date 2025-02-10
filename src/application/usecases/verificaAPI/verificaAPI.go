package verificaAPI

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	helper "github.com/danielsilveiradevbr/api_jira/src/helpers"
	telegram "github.com/danielsilveiradevbr/api_jira/src/infra/mensagem"
	cripto "github.com/danielsilveiradevbr/helpercripto/pkg"
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
						c := http.Client{Timeout: time.Second}
						resp, err := c.Get(api[0])
						if err != nil {
							helper.NewLog(2, err.Error())
						}
						defer resp.Body.Close()
						body, err := io.ReadAll(resp.Body)
						if err != nil {
							helper.NewLog(2, err.Error())
						}

						if string(body) != "ONLINE" {
							helper.NewLog(2, "erro na api "+api[0])
							users := strings.Split(cripto.Cripto("D", os.Getenv("USERS_ENVIO_MSG_TELEGRAM_EM_DEV"), os.Getenv("KEY")), "|")
							telegramBotToken := cripto.Cripto("D", os.Getenv("TELEGRAM_BOT"), os.Getenv("KEY"))
							for _, user := range users {
								err := telegram.SendMessage(telegramBotToken, user, fmt.Sprintf("API %s NÃO RESPONDEU NO ENDEREÇO %s", api[1], api[0]))
								if err != nil {
									break
								}
							}
						} else {
							helper.NewLog(2, "ok com api "+api[0])
						}
					}
				}
			}
			time.Sleep(time.Minute * 5)
		}

	}
}
