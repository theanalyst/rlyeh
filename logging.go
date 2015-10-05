package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func LogInit(verbose bool, logfile string) {
	log.SetFormatter(&log.TextFormatter{})
	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	log.SetOutput(os.Stderr)
}
