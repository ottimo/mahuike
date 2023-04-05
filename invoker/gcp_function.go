package invoker

import (
	"context"
	"fmt"
	"os"

	functions "cloud.google.com/go/functions/apiv1"
	functionspb "cloud.google.com/go/functions/apiv1/functionspb"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

//https://cloud.google.com/functions/docs/running/direct

type GcpFunctionInvoker struct {
}

func NewGcpFunctionInvoker() *GcpFunctionInvoker {
	inv := new(GcpFunctionInvoker)

	return inv
}

func (i *GcpFunctionInvoker) Run(c CmdLine) {

	ctx := context.Background()
	credentials, err := google.FindDefaultCredentials(ctx)
	if err != nil {
		fmt.Printf("cannot get credentials: %v", err)
		os.Exit(1)
	}

	client, err := functions.NewCloudFunctionsClient(ctx,
		option.WithCredentials(credentials),
	)
	if err != nil {
		fmt.Printf("cannot create the client: %v", err)
		os.Exit(1)
	}
	defer client.Close()

	req := &functionspb.CallFunctionRequest{
		// Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/functions/apiv1/functionspb#CallFunctionRequest.
	}
	req.Name = "projects/" + c.Project + "/locations/" + c.Region + "/functions/" + c.ServiceId
	req.Data = c.Data
	resp, err := client.CallFunction(ctx, req)
	if err != nil {
		fmt.Printf("cannot fetch result: %v", err)
		os.Exit(1)
	}
	fmt.Printf("%v", resp.GetResult())
}
