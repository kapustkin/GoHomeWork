package main	

import (
	"fmt"
	"time"	
	"github.com/beevik/ntp"
)

func main() {
	p := fmt.Printf

	ntpTime, err := ntp.Time("time.apple.com")
	if err != nil {
		p("Ошибка при получении времени %s \n", err)
	} else {
		p("Точное время: %s \n", ntpTime.Format(time.RFC3339))
	}
	p("Текущее время: %s \n", time.Now().Format(time.RFC3339))
}