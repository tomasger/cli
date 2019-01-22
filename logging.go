package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func SetupLogging() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(file)
	switch options.Logging {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	}
}
