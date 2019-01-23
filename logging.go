package main

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func SetupLogging(loggingLevel string) {
	logrus.SetFormatter(&logrus.TextFormatter{})
	// Logging should be setup even if logging isn't enabled so we would not get logs to our command line
	// As such we either discard logging output or set it to the log file
	if loggingLevel == "" {
		logrus.SetOutput(ioutil.Discard)
	} else {
		file, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		logrus.SetOutput(file)
		if err != nil {
			logrus.Fatal(err)
		}
		switch loggingLevel {
		case "debug":
			logrus.SetLevel(logrus.DebugLevel)
		case "warn":
			logrus.SetLevel(logrus.WarnLevel)
		}
	}

}
