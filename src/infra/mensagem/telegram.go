package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendMessage(telegramBotToken, destinatario, message string) error {
	// Endpoint da API para enviar mensagem
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramBotToken)

	// Estrutura dos dados que serão enviados
	body, _ := json.Marshal(map[string]interface{}{
		"chat_id": destinatario,
		"text":    message,
	})
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
