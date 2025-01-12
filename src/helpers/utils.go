package helper

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func GetADateTimeSaoPaulo() time.Time {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	return time.Now().In(loc)
}

func StrToTimeTime(dateStr string) time.Time {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return GetADateTimeSaoPaulo()
	}
	return t
}

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Pego as credenciais da req
		// Extrair a query string da URL
		auth := r.Header.Get("Authorization")
		if auth == "" {

			queryParams := r.URL.Query()

			// Acessar parâmetros específicos
			auth = queryParams.Get("Authorization")
		}

		//auth := r.URL.Query("Authorization") //auth := r.Header.Get("Authorization")
		if auth == "" {
			unauthorized(w)
			return
		}

		// Verifico se as credenciais são válidas
		valid := checkAuth(auth)
		if !valid {
			unauthorized(w)
			return
		}

		// Passo para o próximo passo
		next.ServeHTTP(w, r)
	})
}

func checkAuth(auth string) bool {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// Extraio as credenciais
	creds := auth[len("Basic "):]
	decoded, err := base64.StdEncoding.DecodeString(creds)
	if err != nil {
		return false
	}
	// Verifico as  username e password
	validUsername := os.Getenv("AUTH_USER")
	validPassword := os.Getenv("AUTH_PASSWORD")
	return string(decoded) == fmt.Sprintf("%s:%s", validUsername, validPassword)
}

func unauthorized(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("401 Unauthorized\n"))
}
