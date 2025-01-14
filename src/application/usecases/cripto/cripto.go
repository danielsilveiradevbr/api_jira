package cripto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	helper "github.com/danielsilveiradevbr/helpercripto/pkg"
	"github.com/joho/godotenv"
)

type TextCripto struct {
	Text   string `json:"text"`
	Action string `json:"action"`
}

func Cripto(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	key := os.Getenv("KEY")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("problemas para %s", "descriptografar")))
		return
	}
	res, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	var texto *TextCripto

	err = json.Unmarshal(res, &texto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	var textoRetorno string
	if texto.Action == "E" {
		textoRetorno = helper.Cripto("E", texto.Text, key)
	} else {
		textoRetorno = helper.Cripto("D", texto.Text, key)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Retorno %s ", textoRetorno)))
}
