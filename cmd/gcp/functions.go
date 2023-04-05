package gcp

import (
	"bufio"
	"mahuike/invoker"
	"os"

	"github.com/spf13/cobra"
)

var (
	functionCmd = &cobra.Command{
		Use:              "functions",
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	functionCmd.Flags().StringVarP(&c.ServiceId, "functionName", "f", "", "Function to invoke")
	functionCmd.Flags().StringVarP(&c.Project, "projectName", "p", "", "Project where the function is deployed")

	functionCmd.MarkFlagsRequiredTogether("functionName", "projectName")
}

func run() {
	l.Info("Start")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		c.Data = scanner.Text()
	}

	if scanner.Err() != nil {
		l.Error(scanner.Err().Error())
	}

	inv := invoker.NewGcpFunctionInvoker()
	inv.Run(c)
}
