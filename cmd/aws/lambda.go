package aws

import (
	"bufio"
	"mahuike/invoker"
	"os"

	"github.com/spf13/cobra"
)

var (
	lambdaCmd = &cobra.Command{
		Use: "lambda [OPTIONS]",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	lambdaCmd.Flags().StringVarP(&c.ServiceId, "functionArn", "f", "", "Function to invoke")
	lambdaCmd.MarkFlagRequired("functionArn")
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

	inv := invoker.NewAwsLambdaInvoker()
	inv.Run(c)
}
