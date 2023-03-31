package main

import (
	"bufio"
	"flag"
	"mahuike/invoker"
	"mahuike/logger"
	"os"

	log "github.com/sirupsen/logrus"
)

var c invoker.CmdLine
var l = logger.GetLogger().WithFields(log.Fields{"pkg": "main"})

func parseArgs() {
	flag.StringVar(&c.ServiceId, "f", "", "Function to invoke")
	flag.StringVar(&c.Region, "r", "", "Region where the resource is located")
	flag.StringVar(&c.Service, "s", "aws-lambda", "Service to call")

	flag.Parse()
}

func main() {
	parseArgs()
	l.Info("Start")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		c.Data = scanner.Text()
	}

	if scanner.Err() != nil {
		l.Error(scanner.Err().Error())
	}

	inv := invoker.NewInvoker(&c)
	inv.Run(c)
}
