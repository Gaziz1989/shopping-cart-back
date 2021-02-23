package logger

import (
	"fmt"
	"log"
	"os"
)

var logg *os.File

func OpenLog() error {
	f, err := os.OpenFile("logs/landing-back.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	logg = f
	fmt.Println("Connected to log file")
	return nil
}

func CloseLog() error {
	return logg.Close()
}

func GetLog() *os.File {
	return logg
}

func Logg(text string, logging bool) {
	if logging {
		fmt.Println(text)
	}
	log.New(logg, "landing-back ", log.LstdFlags).Println(text)
}
