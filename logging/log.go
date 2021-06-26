package logging

import (
	"log"
	"os"
)

var (
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
	Request *log.Logger
)

var initicialized = false

func init() {
	if initicialized {
		return
	}

	file, err := os.OpenFile("flightening.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	Info = log.New(file, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(file, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "[ERRO] ", log.Ldate|log.Ltime|log.Lshortfile)

	file, err = os.OpenFile("requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	Request = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)

	initicialized = true
}
