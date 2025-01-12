package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func SendMessage(message string) error {
	telegramBotToken := os.Getenv("TELEGRAM_BOT")
	chatsID := strings.Split(os.Getenv("USERS_ENVIO_MSG_TELEGRAM"), "|")
	// Endpoint da API para enviar mensagem
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramBotToken)
	for _, chatID := range chatsID {
		// Estrutura dos dados que serão enviados
		body, _ := json.Marshal(map[string]interface{}{
			"chat_id": chatID,
			"text":    message,
		})
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
		if err != nil {
			return err
		}
		defer resp.Body.Close()
	}

	return nil
}
