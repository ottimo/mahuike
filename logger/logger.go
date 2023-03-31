package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func GetLogger() *log.Entry {
	log.SetOutput(os.Stderr)
	return log.WithFields(
		log.Fields{"mod": "mahuike"},
	)
}
