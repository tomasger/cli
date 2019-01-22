package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func SetupLogging() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.SetFormatter(&log.TextFormatter{})
	switch options.Logging {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	}
}
