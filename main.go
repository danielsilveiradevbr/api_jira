package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	controllers "github.com/danielsilveiradevbr/api_jira/src/application/controllers"
	service "github.com/danielsilveiradevbr/api_jira/src/service/ddsService"
	u "github.com/danielsilveiradevbr/api_jira/src/utils"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	r.Post("/dds", controllers.RecebeDDS)

	porta := os.Getenv("PORT")
	if porta == "" {
		porta = "3000"
	}
	porta = ":" + porta
	println(porta)
	go AtualizaDDS()

	http.ListenAndServe(porta, r)
}

func AtualizaDDS() {
	for {
		hora := u.GetADateTimeSaoPaulo()
		fmt.Println(hora)
		// if hora.Hour() == 23 && hora.Minute() == 59 && hora.Second() == 0 {
		jsonJira, err := controllers.AtualizaDDS()
		if err != nil {
			panic(err)
		}
		service.SalvaDDS(jsonJira)
		// }
		time.Sleep(time.Minute)
	}
}
