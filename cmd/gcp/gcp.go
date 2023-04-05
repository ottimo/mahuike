package gcp

import (
	"mahuike/invoker"
	"mahuike/logger"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	GcpCmd = &cobra.Command{
		Use:  "gcp SERVICE [OPTIONS]",
		Args: cobra.MinimumNArgs(1),
	}
)

var l = logger.GetLogger().WithFields(log.Fields{"pkg": "gcp"})
var c invoker.CmdLine

func init() {
	GcpCmd.PersistentFlags().StringVarP(&c.Region, "region", "r", "", "Region where the resource is located")
	GcpCmd.MarkPersistentFlagRequired("region")
	GcpCmd.AddCommand(functionCmd)
}
