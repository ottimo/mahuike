package invoker

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"

	"fmt"
	"os"
)

type getItemsRequest struct {
	SortBy     string
	SortOrder  string
	ItemsToGet int
}

type getItemsResponseError struct {
	Message string `json:"message"`
}

type getItemsResponseData struct {
	Item string `json:"item"`
}

type getItemsResponseBody struct {
	Result string                 `json:"result"`
	Data   []getItemsResponseData `json:"data"`
	Error  getItemsResponseError  `json:"error"`
}

type getItemsResponseHeaders struct {
	ContentType string `json:"Content-Type"`
}

type getItemsResponse struct {
	StatusCode int                     `json:"statusCode"`
	Headers    getItemsResponseHeaders `json:"headers"`
	Body       getItemsResponseBody    `json:"body"`
}

type AwsLambdaInvoker struct {
}

func NewAwsLambdaInvoker() *AwsLambdaInvoker {
	inv := new(AwsLambdaInvoker)

	return inv
}

func (i *AwsLambdaInvoker) setDefaults(c *CmdLine) {
	if len(c.Region) == 0 {
		c.Region = os.Getenv("AWS_DEFAULT_REGION")
	}
}

func (i *AwsLambdaInvoker) Run(c CmdLine) {

	i.setDefaults(&c)

	l.Infof("Invoking %s", c.ServiceId)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := lambda.New(sess, &aws.Config{Region: aws.String(c.Region)})

	request := c.Data
	payload := []byte(request)
	l.Infof("Sending: %s", payload)

	result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String(c.ServiceId), Payload: payload})
	if err != nil {
		fmt.Println("Error calling MyGetItemsFunction")
		os.Exit(0)
	}

	fmt.Printf("%s", result.Payload)
}
