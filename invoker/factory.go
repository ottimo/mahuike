package invoker

func NewInvoker(c *CmdLine) Invoker {
	l.Info("New Invoker")

	i := NewAwsLambdaInvoker()
	return i
}
