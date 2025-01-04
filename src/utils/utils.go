package utils

import "time"

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
