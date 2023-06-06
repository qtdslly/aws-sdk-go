// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mediastoreiface provides an interface to enable mocking the AWS Elemental MediaStore service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mediastoreiface

import (
	"github.com/qtdslly/aws-sdk-go/aws"
	"github.com/qtdslly/aws-sdk-go/aws/request"
	"github.com/qtdslly/aws-sdk-go/service/mediastore"
)

// MediaStoreAPI provides an interface to enable mocking the
// mediastore.MediaStore service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Elemental MediaStore.
//	func myFunc(svc mediastoreiface.MediaStoreAPI) bool {
//	    // Make svc.CreateContainer request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := mediastore.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockMediaStoreClient struct {
//	    mediastoreiface.MediaStoreAPI
//	}
//	func (m *mockMediaStoreClient) CreateContainer(input *mediastore.CreateContainerInput) (*mediastore.CreateContainerOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockMediaStoreClient{}
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
type MediaStoreAPI interface {
	CreateContainer(*mediastore.CreateContainerInput) (*mediastore.CreateContainerOutput, error)
	CreateContainerWithContext(aws.Context, *mediastore.CreateContainerInput, ...request.Option) (*mediastore.CreateContainerOutput, error)
	CreateContainerRequest(*mediastore.CreateContainerInput) (*request.Request, *mediastore.CreateContainerOutput)

	DeleteContainer(*mediastore.DeleteContainerInput) (*mediastore.DeleteContainerOutput, error)
	DeleteContainerWithContext(aws.Context, *mediastore.DeleteContainerInput, ...request.Option) (*mediastore.DeleteContainerOutput, error)
	DeleteContainerRequest(*mediastore.DeleteContainerInput) (*request.Request, *mediastore.DeleteContainerOutput)

	DeleteContainerPolicy(*mediastore.DeleteContainerPolicyInput) (*mediastore.DeleteContainerPolicyOutput, error)
	DeleteContainerPolicyWithContext(aws.Context, *mediastore.DeleteContainerPolicyInput, ...request.Option) (*mediastore.DeleteContainerPolicyOutput, error)
	DeleteContainerPolicyRequest(*mediastore.DeleteContainerPolicyInput) (*request.Request, *mediastore.DeleteContainerPolicyOutput)

	DeleteCorsPolicy(*mediastore.DeleteCorsPolicyInput) (*mediastore.DeleteCorsPolicyOutput, error)
	DeleteCorsPolicyWithContext(aws.Context, *mediastore.DeleteCorsPolicyInput, ...request.Option) (*mediastore.DeleteCorsPolicyOutput, error)
	DeleteCorsPolicyRequest(*mediastore.DeleteCorsPolicyInput) (*request.Request, *mediastore.DeleteCorsPolicyOutput)

	DeleteLifecyclePolicy(*mediastore.DeleteLifecyclePolicyInput) (*mediastore.DeleteLifecyclePolicyOutput, error)
	DeleteLifecyclePolicyWithContext(aws.Context, *mediastore.DeleteLifecyclePolicyInput, ...request.Option) (*mediastore.DeleteLifecyclePolicyOutput, error)
	DeleteLifecyclePolicyRequest(*mediastore.DeleteLifecyclePolicyInput) (*request.Request, *mediastore.DeleteLifecyclePolicyOutput)

	DeleteMetricPolicy(*mediastore.DeleteMetricPolicyInput) (*mediastore.DeleteMetricPolicyOutput, error)
	DeleteMetricPolicyWithContext(aws.Context, *mediastore.DeleteMetricPolicyInput, ...request.Option) (*mediastore.DeleteMetricPolicyOutput, error)
	DeleteMetricPolicyRequest(*mediastore.DeleteMetricPolicyInput) (*request.Request, *mediastore.DeleteMetricPolicyOutput)

	DescribeContainer(*mediastore.DescribeContainerInput) (*mediastore.DescribeContainerOutput, error)
	DescribeContainerWithContext(aws.Context, *mediastore.DescribeContainerInput, ...request.Option) (*mediastore.DescribeContainerOutput, error)
	DescribeContainerRequest(*mediastore.DescribeContainerInput) (*request.Request, *mediastore.DescribeContainerOutput)

	GetContainerPolicy(*mediastore.GetContainerPolicyInput) (*mediastore.GetContainerPolicyOutput, error)
	GetContainerPolicyWithContext(aws.Context, *mediastore.GetContainerPolicyInput, ...request.Option) (*mediastore.GetContainerPolicyOutput, error)
	GetContainerPolicyRequest(*mediastore.GetContainerPolicyInput) (*request.Request, *mediastore.GetContainerPolicyOutput)

	GetCorsPolicy(*mediastore.GetCorsPolicyInput) (*mediastore.GetCorsPolicyOutput, error)
	GetCorsPolicyWithContext(aws.Context, *mediastore.GetCorsPolicyInput, ...request.Option) (*mediastore.GetCorsPolicyOutput, error)
	GetCorsPolicyRequest(*mediastore.GetCorsPolicyInput) (*request.Request, *mediastore.GetCorsPolicyOutput)

	GetLifecyclePolicy(*mediastore.GetLifecyclePolicyInput) (*mediastore.GetLifecyclePolicyOutput, error)
	GetLifecyclePolicyWithContext(aws.Context, *mediastore.GetLifecyclePolicyInput, ...request.Option) (*mediastore.GetLifecyclePolicyOutput, error)
	GetLifecyclePolicyRequest(*mediastore.GetLifecyclePolicyInput) (*request.Request, *mediastore.GetLifecyclePolicyOutput)

	GetMetricPolicy(*mediastore.GetMetricPolicyInput) (*mediastore.GetMetricPolicyOutput, error)
	GetMetricPolicyWithContext(aws.Context, *mediastore.GetMetricPolicyInput, ...request.Option) (*mediastore.GetMetricPolicyOutput, error)
	GetMetricPolicyRequest(*mediastore.GetMetricPolicyInput) (*request.Request, *mediastore.GetMetricPolicyOutput)

	ListContainers(*mediastore.ListContainersInput) (*mediastore.ListContainersOutput, error)
	ListContainersWithContext(aws.Context, *mediastore.ListContainersInput, ...request.Option) (*mediastore.ListContainersOutput, error)
	ListContainersRequest(*mediastore.ListContainersInput) (*request.Request, *mediastore.ListContainersOutput)

	ListContainersPages(*mediastore.ListContainersInput, func(*mediastore.ListContainersOutput, bool) bool) error
	ListContainersPagesWithContext(aws.Context, *mediastore.ListContainersInput, func(*mediastore.ListContainersOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*mediastore.ListTagsForResourceInput) (*mediastore.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *mediastore.ListTagsForResourceInput, ...request.Option) (*mediastore.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*mediastore.ListTagsForResourceInput) (*request.Request, *mediastore.ListTagsForResourceOutput)

	PutContainerPolicy(*mediastore.PutContainerPolicyInput) (*mediastore.PutContainerPolicyOutput, error)
	PutContainerPolicyWithContext(aws.Context, *mediastore.PutContainerPolicyInput, ...request.Option) (*mediastore.PutContainerPolicyOutput, error)
	PutContainerPolicyRequest(*mediastore.PutContainerPolicyInput) (*request.Request, *mediastore.PutContainerPolicyOutput)

	PutCorsPolicy(*mediastore.PutCorsPolicyInput) (*mediastore.PutCorsPolicyOutput, error)
	PutCorsPolicyWithContext(aws.Context, *mediastore.PutCorsPolicyInput, ...request.Option) (*mediastore.PutCorsPolicyOutput, error)
	PutCorsPolicyRequest(*mediastore.PutCorsPolicyInput) (*request.Request, *mediastore.PutCorsPolicyOutput)

	PutLifecyclePolicy(*mediastore.PutLifecyclePolicyInput) (*mediastore.PutLifecyclePolicyOutput, error)
	PutLifecyclePolicyWithContext(aws.Context, *mediastore.PutLifecyclePolicyInput, ...request.Option) (*mediastore.PutLifecyclePolicyOutput, error)
	PutLifecyclePolicyRequest(*mediastore.PutLifecyclePolicyInput) (*request.Request, *mediastore.PutLifecyclePolicyOutput)

	PutMetricPolicy(*mediastore.PutMetricPolicyInput) (*mediastore.PutMetricPolicyOutput, error)
	PutMetricPolicyWithContext(aws.Context, *mediastore.PutMetricPolicyInput, ...request.Option) (*mediastore.PutMetricPolicyOutput, error)
	PutMetricPolicyRequest(*mediastore.PutMetricPolicyInput) (*request.Request, *mediastore.PutMetricPolicyOutput)

	StartAccessLogging(*mediastore.StartAccessLoggingInput) (*mediastore.StartAccessLoggingOutput, error)
	StartAccessLoggingWithContext(aws.Context, *mediastore.StartAccessLoggingInput, ...request.Option) (*mediastore.StartAccessLoggingOutput, error)
	StartAccessLoggingRequest(*mediastore.StartAccessLoggingInput) (*request.Request, *mediastore.StartAccessLoggingOutput)

	StopAccessLogging(*mediastore.StopAccessLoggingInput) (*mediastore.StopAccessLoggingOutput, error)
	StopAccessLoggingWithContext(aws.Context, *mediastore.StopAccessLoggingInput, ...request.Option) (*mediastore.StopAccessLoggingOutput, error)
	StopAccessLoggingRequest(*mediastore.StopAccessLoggingInput) (*request.Request, *mediastore.StopAccessLoggingOutput)

	TagResource(*mediastore.TagResourceInput) (*mediastore.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *mediastore.TagResourceInput, ...request.Option) (*mediastore.TagResourceOutput, error)
	TagResourceRequest(*mediastore.TagResourceInput) (*request.Request, *mediastore.TagResourceOutput)

	UntagResource(*mediastore.UntagResourceInput) (*mediastore.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *mediastore.UntagResourceInput, ...request.Option) (*mediastore.UntagResourceOutput, error)
	UntagResourceRequest(*mediastore.UntagResourceInput) (*request.Request, *mediastore.UntagResourceOutput)
}

var _ MediaStoreAPI = (*mediastore.MediaStore)(nil)
