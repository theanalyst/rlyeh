package main

import (
	log "github.com/Sirupsen/logrus"
	"io"
)

func LogInit(verbose bool, logfile io.Writer) {

	log.SetFormatter(&log.TextFormatter{})
	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	log.SetOutput(logfile)
}
