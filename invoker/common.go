package invoker

import (
	"mahuike/logger"

	log "github.com/sirupsen/logrus"
)

type Invoker interface {
	Run(CmdLine)
}

type CmdLine struct {
	Service   string
	ServiceId string
	Data      string
	Region    string
}

var l = logger.GetLogger().WithFields(log.Fields{"pkg": "invoker"})
