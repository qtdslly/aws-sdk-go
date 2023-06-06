// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sqsiface provides an interface to enable mocking the Amazon Simple Queue Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package sqsiface

import (
	"github.com/qtdslly/aws-sdk-go/aws"
	"github.com/qtdslly/aws-sdk-go/aws/request"
	"github.com/qtdslly/aws-sdk-go/service/sqs"
)

// SQSAPI provides an interface to enable mocking the
// sqs.SQS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Simple Queue Service.
//	func myFunc(svc sqsiface.SQSAPI) bool {
//	    // Make svc.AddPermission request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := sqs.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockSQSClient struct {
//	    sqsiface.SQSAPI
//	}
//	func (m *mockSQSClient) AddPermission(input *sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockSQSClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type SQSAPI interface {
	AddPermission(*sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error)
	AddPermissionWithContext(aws.Context, *sqs.AddPermissionInput, ...request.Option) (*sqs.AddPermissionOutput, error)
	AddPermissionRequest(*sqs.AddPermissionInput) (*request.Request, *sqs.AddPermissionOutput)

	ChangeMessageVisibility(*sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error)
	ChangeMessageVisibilityWithContext(aws.Context, *sqs.ChangeMessageVisibilityInput, ...request.Option) (*sqs.ChangeMessageVisibilityOutput, error)
	ChangeMessageVisibilityRequest(*sqs.ChangeMessageVisibilityInput) (*request.Request, *sqs.ChangeMessageVisibilityOutput)

	ChangeMessageVisibilityBatch(*sqs.ChangeMessageVisibilityBatchInput) (*sqs.ChangeMessageVisibilityBatchOutput, error)
	ChangeMessageVisibilityBatchWithContext(aws.Context, *sqs.ChangeMessageVisibilityBatchInput, ...request.Option) (*sqs.ChangeMessageVisibilityBatchOutput, error)
	ChangeMessageVisibilityBatchRequest(*sqs.ChangeMessageVisibilityBatchInput) (*request.Request, *sqs.ChangeMessageVisibilityBatchOutput)

	CreateQueue(*sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error)
	CreateQueueWithContext(aws.Context, *sqs.CreateQueueInput, ...request.Option) (*sqs.CreateQueueOutput, error)
	CreateQueueRequest(*sqs.CreateQueueInput) (*request.Request, *sqs.CreateQueueOutput)

	DeleteMessage(*sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error)
	DeleteMessageWithContext(aws.Context, *sqs.DeleteMessageInput, ...request.Option) (*sqs.DeleteMessageOutput, error)
	DeleteMessageRequest(*sqs.DeleteMessageInput) (*request.Request, *sqs.DeleteMessageOutput)

	DeleteMessageBatch(*sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error)
	DeleteMessageBatchWithContext(aws.Context, *sqs.DeleteMessageBatchInput, ...request.Option) (*sqs.DeleteMessageBatchOutput, error)
	DeleteMessageBatchRequest(*sqs.DeleteMessageBatchInput) (*request.Request, *sqs.DeleteMessageBatchOutput)

	DeleteQueue(*sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, error)
	DeleteQueueWithContext(aws.Context, *sqs.DeleteQueueInput, ...request.Option) (*sqs.DeleteQueueOutput, error)
	DeleteQueueRequest(*sqs.DeleteQueueInput) (*request.Request, *sqs.DeleteQueueOutput)

	GetQueueAttributes(*sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error)
	GetQueueAttributesWithContext(aws.Context, *sqs.GetQueueAttributesInput, ...request.Option) (*sqs.GetQueueAttributesOutput, error)
	GetQueueAttributesRequest(*sqs.GetQueueAttributesInput) (*request.Request, *sqs.GetQueueAttributesOutput)

	GetQueueUrl(*sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error)
	GetQueueUrlWithContext(aws.Context, *sqs.GetQueueUrlInput, ...request.Option) (*sqs.GetQueueUrlOutput, error)
	GetQueueUrlRequest(*sqs.GetQueueUrlInput) (*request.Request, *sqs.GetQueueUrlOutput)

	ListDeadLetterSourceQueues(*sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, error)
	ListDeadLetterSourceQueuesWithContext(aws.Context, *sqs.ListDeadLetterSourceQueuesInput, ...request.Option) (*sqs.ListDeadLetterSourceQueuesOutput, error)
	ListDeadLetterSourceQueuesRequest(*sqs.ListDeadLetterSourceQueuesInput) (*request.Request, *sqs.ListDeadLetterSourceQueuesOutput)

	ListDeadLetterSourceQueuesPages(*sqs.ListDeadLetterSourceQueuesInput, func(*sqs.ListDeadLetterSourceQueuesOutput, bool) bool) error
	ListDeadLetterSourceQueuesPagesWithContext(aws.Context, *sqs.ListDeadLetterSourceQueuesInput, func(*sqs.ListDeadLetterSourceQueuesOutput, bool) bool, ...request.Option) error

	ListQueueTags(*sqs.ListQueueTagsInput) (*sqs.ListQueueTagsOutput, error)
	ListQueueTagsWithContext(aws.Context, *sqs.ListQueueTagsInput, ...request.Option) (*sqs.ListQueueTagsOutput, error)
	ListQueueTagsRequest(*sqs.ListQueueTagsInput) (*request.Request, *sqs.ListQueueTagsOutput)

	ListQueues(*sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error)
	ListQueuesWithContext(aws.Context, *sqs.ListQueuesInput, ...request.Option) (*sqs.ListQueuesOutput, error)
	ListQueuesRequest(*sqs.ListQueuesInput) (*request.Request, *sqs.ListQueuesOutput)

	ListQueuesPages(*sqs.ListQueuesInput, func(*sqs.ListQueuesOutput, bool) bool) error
	ListQueuesPagesWithContext(aws.Context, *sqs.ListQueuesInput, func(*sqs.ListQueuesOutput, bool) bool, ...request.Option) error

	PurgeQueue(*sqs.PurgeQueueInput) (*sqs.PurgeQueueOutput, error)
	PurgeQueueWithContext(aws.Context, *sqs.PurgeQueueInput, ...request.Option) (*sqs.PurgeQueueOutput, error)
	PurgeQueueRequest(*sqs.PurgeQueueInput) (*request.Request, *sqs.PurgeQueueOutput)

	ReceiveMessage(*sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error)
	ReceiveMessageWithContext(aws.Context, *sqs.ReceiveMessageInput, ...request.Option) (*sqs.ReceiveMessageOutput, error)
	ReceiveMessageRequest(*sqs.ReceiveMessageInput) (*request.Request, *sqs.ReceiveMessageOutput)

	RemovePermission(*sqs.RemovePermissionInput) (*sqs.RemovePermissionOutput, error)
	RemovePermissionWithContext(aws.Context, *sqs.RemovePermissionInput, ...request.Option) (*sqs.RemovePermissionOutput, error)
	RemovePermissionRequest(*sqs.RemovePermissionInput) (*request.Request, *sqs.RemovePermissionOutput)

	SendMessage(*sqs.SendMessageInput) (*sqs.SendMessageOutput, error)
	SendMessageWithContext(aws.Context, *sqs.SendMessageInput, ...request.Option) (*sqs.SendMessageOutput, error)
	SendMessageRequest(*sqs.SendMessageInput) (*request.Request, *sqs.SendMessageOutput)

	SendMessageBatch(*sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, error)
	SendMessageBatchWithContext(aws.Context, *sqs.SendMessageBatchInput, ...request.Option) (*sqs.SendMessageBatchOutput, error)
	SendMessageBatchRequest(*sqs.SendMessageBatchInput) (*request.Request, *sqs.SendMessageBatchOutput)

	SetQueueAttributes(*sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error)
	SetQueueAttributesWithContext(aws.Context, *sqs.SetQueueAttributesInput, ...request.Option) (*sqs.SetQueueAttributesOutput, error)
	SetQueueAttributesRequest(*sqs.SetQueueAttributesInput) (*request.Request, *sqs.SetQueueAttributesOutput)

	TagQueue(*sqs.TagQueueInput) (*sqs.TagQueueOutput, error)
	TagQueueWithContext(aws.Context, *sqs.TagQueueInput, ...request.Option) (*sqs.TagQueueOutput, error)
	TagQueueRequest(*sqs.TagQueueInput) (*request.Request, *sqs.TagQueueOutput)

	UntagQueue(*sqs.UntagQueueInput) (*sqs.UntagQueueOutput, error)
	UntagQueueWithContext(aws.Context, *sqs.UntagQueueInput, ...request.Option) (*sqs.UntagQueueOutput, error)
	UntagQueueRequest(*sqs.UntagQueueInput) (*request.Request, *sqs.UntagQueueOutput)
}

var _ SQSAPI = (*sqs.SQS)(nil)
