package dsdk

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func InitLog(debug bool, output string, stdout bool) error {

	log.SetFormatter(&log.JSONFormatter{})

	switch debug {
	default:
		log.SetLevel(log.WarnLevel)
	case true:
		log.SetLevel(log.DebugLevel)
	}

	var f io.Writer
	if _, err := os.Stat(output); output != "" && os.IsNotExist(err) {
		f, err = os.Create(output)
		if err != nil {
			return err
		}
	} else if output != "" {
		f, err = os.OpenFile(output, os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
	}

	if stdout && output != "" {
		log.SetOutput(io.MultiWriter(f, os.Stdout))
	} else if stdout {
		log.SetOutput(os.Stdout)
	} else if output != "" {
		log.SetOutput(f)
	}
	return nil
}
