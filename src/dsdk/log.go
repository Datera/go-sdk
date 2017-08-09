package dsdk

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

const (
	LogFile = "dsdk.log"
)

func InitLog(debug bool, output string, stdout bool) error {

	log.SetFormatter(&log.JSONFormatter{})
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

	if stdout {
		log.SetOutput(io.MultiWriter(f, os.Stdout))
	} else {
		log.SetOutput(f)
	}
	return nil
}
