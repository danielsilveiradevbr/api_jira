package helper

import (
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func NewLog(level int, message string) {
	if strings.Trim(message, "") == "" {
		return
	}

	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo de log: %v", err)
	}
	defer logFile.Close()

	handler := slog.NewJSONHandler(logFile, nil)
	l := slog.New(handler)
	slog.SetDefault(l)
	err = godotenv.Load()
	if err != nil {
		slog.Error(err.Error())
	}
	if os.Getenv("DEBUGANDO") == "T" {
		level := new(slog.LevelVar)
		handler = slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: level})
		l = slog.New(handler)
		slog.SetDefault(l)
		level.Set(slog.LevelDebug)
		slog.Debug(message)
	}
	if level == 1 {
		slog.Info(message)
	} else if level == 2 {
		slog.Error(message)
	} else if level == 3 {
		slog.Warn(message)
	}

}
