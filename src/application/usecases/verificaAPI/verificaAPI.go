package verificaAPI

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	telegram "github.com/danielsilveiradevbr/api_jira/src/infra/mensagem"
)

func VerificaAPI() {
	apisTestar := strings.Split(os.Getenv("APIS_TESTAR"), "|")
	fmt.Println(len(apisTestar))
	if len(apisTestar) > 0 {
		for {
			for _, apiTestar := range apisTestar {
				api := strings.Split(apiTestar, ",")
				fmt.Println(api[0])
				if len(api) > 0 {
					if strings.Trim(api[0], "") != "" {
						c := http.Client{Timeout: time.Second}
						resp, err := c.Get(api[0])
						if err != nil {
							panic(err)
						}
						defer resp.Body.Close()
						body, err := io.ReadAll(resp.Body)
						if err != nil {
							panic(err)
						}

						if string(body) != "ONLINE" {
							users := strings.Split(os.Getenv("USERS_ENVIO_MSG_TELEGRAM_EM_DEV"), "|")
							telegramBotToken := os.Getenv("TELEGRAM_BOT")
							for _, user := range users {
								err := telegram.SendMessage(telegramBotToken, user, fmt.Sprintf("API %s NÃO RESPONDEU NO ENDEREÇO %s", api[1], api[0]))
								if err != nil {
									break
								}
							}
						}
					}
				}
			}
			time.Sleep(time.Minute * 5)
		}

	}
}
