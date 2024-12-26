package utils

import "time"

func GetADateTimeSaoPaulo() time.Time {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	return time.Now().In(loc)
}
