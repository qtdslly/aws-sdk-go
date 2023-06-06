// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"time"

	"github.com/qtdslly/aws-sdk-go/aws"
	"github.com/qtdslly/aws-sdk-go/aws/request"
)

// WaitUntilAlarmExists uses the CloudWatch API operation
// DescribeAlarms to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *CloudWatch) WaitUntilAlarmExists(input *DescribeAlarmsInput) error {
	return c.WaitUntilAlarmExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilAlarmExistsWithContext is an extended version of WaitUntilAlarmExists.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudWatch) WaitUntilAlarmExistsWithContext(ctx aws.Context, input *DescribeAlarmsInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilAlarmExists",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "length(MetricAlarms[]) > `0`",
				Expected: true,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeAlarmsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeAlarmsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilCompositeAlarmExists uses the CloudWatch API operation
// DescribeAlarms to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *CloudWatch) WaitUntilCompositeAlarmExists(input *DescribeAlarmsInput) error {
	return c.WaitUntilCompositeAlarmExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilCompositeAlarmExistsWithContext is an extended version of WaitUntilCompositeAlarmExists.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudWatch) WaitUntilCompositeAlarmExistsWithContext(ctx aws.Context, input *DescribeAlarmsInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilCompositeAlarmExists",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "length(CompositeAlarms[]) > `0`",
				Expected: true,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeAlarmsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeAlarmsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
