// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"time"

	"github.com/qtdslly/aws-sdk-go/aws"
	"github.com/qtdslly/aws-sdk-go/aws/request"
)

// WaitUntilImageScanComplete uses the Amazon ECR API operation
// DescribeImageScanFindings to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *ECR) WaitUntilImageScanComplete(input *DescribeImageScanFindingsInput) error {
	return c.WaitUntilImageScanCompleteWithContext(aws.BackgroundContext(), input)
}

// WaitUntilImageScanCompleteWithContext is an extended version of WaitUntilImageScanComplete.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECR) WaitUntilImageScanCompleteWithContext(ctx aws.Context, input *DescribeImageScanFindingsInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilImageScanComplete",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "imageScanStatus.status",
				Expected: "COMPLETE",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "imageScanStatus.status",
				Expected: "FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeImageScanFindingsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeImageScanFindingsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilLifecyclePolicyPreviewComplete uses the Amazon ECR API operation
// GetLifecyclePolicyPreview to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *ECR) WaitUntilLifecyclePolicyPreviewComplete(input *GetLifecyclePolicyPreviewInput) error {
	return c.WaitUntilLifecyclePolicyPreviewCompleteWithContext(aws.BackgroundContext(), input)
}

// WaitUntilLifecyclePolicyPreviewCompleteWithContext is an extended version of WaitUntilLifecyclePolicyPreviewComplete.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECR) WaitUntilLifecyclePolicyPreviewCompleteWithContext(ctx aws.Context, input *GetLifecyclePolicyPreviewInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilLifecyclePolicyPreviewComplete",
		MaxAttempts: 20,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "status",
				Expected: "COMPLETE",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "status",
				Expected: "FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetLifecyclePolicyPreviewInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetLifecyclePolicyPreviewRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
