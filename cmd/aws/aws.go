package aws

import (
	"mahuike/invoker"
	"mahuike/logger"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	AwsCmd = &cobra.Command{
		Use:  "aws SERVICE [OPTIONS]",
		Args: cobra.MinimumNArgs(1),
	}
)

var l = logger.GetLogger().WithFields(log.Fields{"pkg": "aws"})
var c invoker.CmdLine

func init() {
	AwsCmd.PersistentFlags().StringVarP(&c.Region, "region", "r", "", "Region where the resource is located")
	AwsCmd.AddCommand(lambdaCmd)
}
