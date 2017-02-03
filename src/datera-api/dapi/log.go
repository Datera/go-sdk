package dapi

import (
	log "github.com/Sirupsen/logrus"
	"io"
	"os"
)

const (
	LogFile = "datera-go-api.log"
)

func InitLog(debug bool, output string) error {

	customFormatter := new(log.TextFormatter)
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
	var o string
	switch output {
	default:
		o = output
	case "":
		o = LogFile
	}

	switch debug {
	default:
		log.SetLevel(log.WarnLevel)
	case true:
		log.SetLevel(log.DebugLevel)
	}

	var f io.Writer
	if _, err := os.Stat(o); os.IsNotExist(err) {
		f, err = os.Create(o)
		if err != nil {
			return err
		}
	} else {
		f, err = os.OpenFile(o, os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
	}

	log.SetOutput(f)
	return nil
}
