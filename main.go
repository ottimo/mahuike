package main

import (
	"mahuike/cmd"
	"mahuike/invoker"
	"mahuike/logger"

	log "github.com/sirupsen/logrus"
)

var c invoker.CmdLine
var l = logger.GetLogger().WithFields(log.Fields{"pkg": "main"})

func main() {
	cmd.Execute()

}
