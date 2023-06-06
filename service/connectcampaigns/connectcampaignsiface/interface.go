// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package connectcampaignsiface provides an interface to enable mocking the AmazonConnectCampaignService service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package connectcampaignsiface

import (
	"github.com/qtdslly/aws-sdk-go/aws"
	"github.com/qtdslly/aws-sdk-go/aws/request"
	"github.com/qtdslly/aws-sdk-go/service/connectcampaigns"
)

// ConnectCampaignsAPI provides an interface to enable mocking the
// connectcampaigns.ConnectCampaigns service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AmazonConnectCampaignService.
//	func myFunc(svc connectcampaignsiface.ConnectCampaignsAPI) bool {
//	    // Make svc.CreateCampaign request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := connectcampaigns.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockConnectCampaignsClient struct {
//	    connectcampaignsiface.ConnectCampaignsAPI
//	}
//	func (m *mockConnectCampaignsClient) CreateCampaign(input *connectcampaigns.CreateCampaignInput) (*connectcampaigns.CreateCampaignOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockConnectCampaignsClient{}
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
type ConnectCampaignsAPI interface {
	CreateCampaign(*connectcampaigns.CreateCampaignInput) (*connectcampaigns.CreateCampaignOutput, error)
	CreateCampaignWithContext(aws.Context, *connectcampaigns.CreateCampaignInput, ...request.Option) (*connectcampaigns.CreateCampaignOutput, error)
	CreateCampaignRequest(*connectcampaigns.CreateCampaignInput) (*request.Request, *connectcampaigns.CreateCampaignOutput)

	DeleteCampaign(*connectcampaigns.DeleteCampaignInput) (*connectcampaigns.DeleteCampaignOutput, error)
	DeleteCampaignWithContext(aws.Context, *connectcampaigns.DeleteCampaignInput, ...request.Option) (*connectcampaigns.DeleteCampaignOutput, error)
	DeleteCampaignRequest(*connectcampaigns.DeleteCampaignInput) (*request.Request, *connectcampaigns.DeleteCampaignOutput)

	DeleteConnectInstanceConfig(*connectcampaigns.DeleteConnectInstanceConfigInput) (*connectcampaigns.DeleteConnectInstanceConfigOutput, error)
	DeleteConnectInstanceConfigWithContext(aws.Context, *connectcampaigns.DeleteConnectInstanceConfigInput, ...request.Option) (*connectcampaigns.DeleteConnectInstanceConfigOutput, error)
	DeleteConnectInstanceConfigRequest(*connectcampaigns.DeleteConnectInstanceConfigInput) (*request.Request, *connectcampaigns.DeleteConnectInstanceConfigOutput)

	DeleteInstanceOnboardingJob(*connectcampaigns.DeleteInstanceOnboardingJobInput) (*connectcampaigns.DeleteInstanceOnboardingJobOutput, error)
	DeleteInstanceOnboardingJobWithContext(aws.Context, *connectcampaigns.DeleteInstanceOnboardingJobInput, ...request.Option) (*connectcampaigns.DeleteInstanceOnboardingJobOutput, error)
	DeleteInstanceOnboardingJobRequest(*connectcampaigns.DeleteInstanceOnboardingJobInput) (*request.Request, *connectcampaigns.DeleteInstanceOnboardingJobOutput)

	DescribeCampaign(*connectcampaigns.DescribeCampaignInput) (*connectcampaigns.DescribeCampaignOutput, error)
	DescribeCampaignWithContext(aws.Context, *connectcampaigns.DescribeCampaignInput, ...request.Option) (*connectcampaigns.DescribeCampaignOutput, error)
	DescribeCampaignRequest(*connectcampaigns.DescribeCampaignInput) (*request.Request, *connectcampaigns.DescribeCampaignOutput)

	GetCampaignState(*connectcampaigns.GetCampaignStateInput) (*connectcampaigns.GetCampaignStateOutput, error)
	GetCampaignStateWithContext(aws.Context, *connectcampaigns.GetCampaignStateInput, ...request.Option) (*connectcampaigns.GetCampaignStateOutput, error)
	GetCampaignStateRequest(*connectcampaigns.GetCampaignStateInput) (*request.Request, *connectcampaigns.GetCampaignStateOutput)

	GetCampaignStateBatch(*connectcampaigns.GetCampaignStateBatchInput) (*connectcampaigns.GetCampaignStateBatchOutput, error)
	GetCampaignStateBatchWithContext(aws.Context, *connectcampaigns.GetCampaignStateBatchInput, ...request.Option) (*connectcampaigns.GetCampaignStateBatchOutput, error)
	GetCampaignStateBatchRequest(*connectcampaigns.GetCampaignStateBatchInput) (*request.Request, *connectcampaigns.GetCampaignStateBatchOutput)

	GetConnectInstanceConfig(*connectcampaigns.GetConnectInstanceConfigInput) (*connectcampaigns.GetConnectInstanceConfigOutput, error)
	GetConnectInstanceConfigWithContext(aws.Context, *connectcampaigns.GetConnectInstanceConfigInput, ...request.Option) (*connectcampaigns.GetConnectInstanceConfigOutput, error)
	GetConnectInstanceConfigRequest(*connectcampaigns.GetConnectInstanceConfigInput) (*request.Request, *connectcampaigns.GetConnectInstanceConfigOutput)

	GetInstanceOnboardingJobStatus(*connectcampaigns.GetInstanceOnboardingJobStatusInput) (*connectcampaigns.GetInstanceOnboardingJobStatusOutput, error)
	GetInstanceOnboardingJobStatusWithContext(aws.Context, *connectcampaigns.GetInstanceOnboardingJobStatusInput, ...request.Option) (*connectcampaigns.GetInstanceOnboardingJobStatusOutput, error)
	GetInstanceOnboardingJobStatusRequest(*connectcampaigns.GetInstanceOnboardingJobStatusInput) (*request.Request, *connectcampaigns.GetInstanceOnboardingJobStatusOutput)

	ListCampaigns(*connectcampaigns.ListCampaignsInput) (*connectcampaigns.ListCampaignsOutput, error)
	ListCampaignsWithContext(aws.Context, *connectcampaigns.ListCampaignsInput, ...request.Option) (*connectcampaigns.ListCampaignsOutput, error)
	ListCampaignsRequest(*connectcampaigns.ListCampaignsInput) (*request.Request, *connectcampaigns.ListCampaignsOutput)

	ListCampaignsPages(*connectcampaigns.ListCampaignsInput, func(*connectcampaigns.ListCampaignsOutput, bool) bool) error
	ListCampaignsPagesWithContext(aws.Context, *connectcampaigns.ListCampaignsInput, func(*connectcampaigns.ListCampaignsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*connectcampaigns.ListTagsForResourceInput) (*connectcampaigns.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *connectcampaigns.ListTagsForResourceInput, ...request.Option) (*connectcampaigns.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*connectcampaigns.ListTagsForResourceInput) (*request.Request, *connectcampaigns.ListTagsForResourceOutput)

	PauseCampaign(*connectcampaigns.PauseCampaignInput) (*connectcampaigns.PauseCampaignOutput, error)
	PauseCampaignWithContext(aws.Context, *connectcampaigns.PauseCampaignInput, ...request.Option) (*connectcampaigns.PauseCampaignOutput, error)
	PauseCampaignRequest(*connectcampaigns.PauseCampaignInput) (*request.Request, *connectcampaigns.PauseCampaignOutput)

	PutDialRequestBatch(*connectcampaigns.PutDialRequestBatchInput) (*connectcampaigns.PutDialRequestBatchOutput, error)
	PutDialRequestBatchWithContext(aws.Context, *connectcampaigns.PutDialRequestBatchInput, ...request.Option) (*connectcampaigns.PutDialRequestBatchOutput, error)
	PutDialRequestBatchRequest(*connectcampaigns.PutDialRequestBatchInput) (*request.Request, *connectcampaigns.PutDialRequestBatchOutput)

	ResumeCampaign(*connectcampaigns.ResumeCampaignInput) (*connectcampaigns.ResumeCampaignOutput, error)
	ResumeCampaignWithContext(aws.Context, *connectcampaigns.ResumeCampaignInput, ...request.Option) (*connectcampaigns.ResumeCampaignOutput, error)
	ResumeCampaignRequest(*connectcampaigns.ResumeCampaignInput) (*request.Request, *connectcampaigns.ResumeCampaignOutput)

	StartCampaign(*connectcampaigns.StartCampaignInput) (*connectcampaigns.StartCampaignOutput, error)
	StartCampaignWithContext(aws.Context, *connectcampaigns.StartCampaignInput, ...request.Option) (*connectcampaigns.StartCampaignOutput, error)
	StartCampaignRequest(*connectcampaigns.StartCampaignInput) (*request.Request, *connectcampaigns.StartCampaignOutput)

	StartInstanceOnboardingJob(*connectcampaigns.StartInstanceOnboardingJobInput) (*connectcampaigns.StartInstanceOnboardingJobOutput, error)
	StartInstanceOnboardingJobWithContext(aws.Context, *connectcampaigns.StartInstanceOnboardingJobInput, ...request.Option) (*connectcampaigns.StartInstanceOnboardingJobOutput, error)
	StartInstanceOnboardingJobRequest(*connectcampaigns.StartInstanceOnboardingJobInput) (*request.Request, *connectcampaigns.StartInstanceOnboardingJobOutput)

	StopCampaign(*connectcampaigns.StopCampaignInput) (*connectcampaigns.StopCampaignOutput, error)
	StopCampaignWithContext(aws.Context, *connectcampaigns.StopCampaignInput, ...request.Option) (*connectcampaigns.StopCampaignOutput, error)
	StopCampaignRequest(*connectcampaigns.StopCampaignInput) (*request.Request, *connectcampaigns.StopCampaignOutput)

	TagResource(*connectcampaigns.TagResourceInput) (*connectcampaigns.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *connectcampaigns.TagResourceInput, ...request.Option) (*connectcampaigns.TagResourceOutput, error)
	TagResourceRequest(*connectcampaigns.TagResourceInput) (*request.Request, *connectcampaigns.TagResourceOutput)

	UntagResource(*connectcampaigns.UntagResourceInput) (*connectcampaigns.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *connectcampaigns.UntagResourceInput, ...request.Option) (*connectcampaigns.UntagResourceOutput, error)
	UntagResourceRequest(*connectcampaigns.UntagResourceInput) (*request.Request, *connectcampaigns.UntagResourceOutput)

	UpdateCampaignDialerConfig(*connectcampaigns.UpdateCampaignDialerConfigInput) (*connectcampaigns.UpdateCampaignDialerConfigOutput, error)
	UpdateCampaignDialerConfigWithContext(aws.Context, *connectcampaigns.UpdateCampaignDialerConfigInput, ...request.Option) (*connectcampaigns.UpdateCampaignDialerConfigOutput, error)
	UpdateCampaignDialerConfigRequest(*connectcampaigns.UpdateCampaignDialerConfigInput) (*request.Request, *connectcampaigns.UpdateCampaignDialerConfigOutput)

	UpdateCampaignName(*connectcampaigns.UpdateCampaignNameInput) (*connectcampaigns.UpdateCampaignNameOutput, error)
	UpdateCampaignNameWithContext(aws.Context, *connectcampaigns.UpdateCampaignNameInput, ...request.Option) (*connectcampaigns.UpdateCampaignNameOutput, error)
	UpdateCampaignNameRequest(*connectcampaigns.UpdateCampaignNameInput) (*request.Request, *connectcampaigns.UpdateCampaignNameOutput)

	UpdateCampaignOutboundCallConfig(*connectcampaigns.UpdateCampaignOutboundCallConfigInput) (*connectcampaigns.UpdateCampaignOutboundCallConfigOutput, error)
	UpdateCampaignOutboundCallConfigWithContext(aws.Context, *connectcampaigns.UpdateCampaignOutboundCallConfigInput, ...request.Option) (*connectcampaigns.UpdateCampaignOutboundCallConfigOutput, error)
	UpdateCampaignOutboundCallConfigRequest(*connectcampaigns.UpdateCampaignOutboundCallConfigInput) (*request.Request, *connectcampaigns.UpdateCampaignOutboundCallConfigOutput)
}

var _ ConnectCampaignsAPI = (*connectcampaigns.ConnectCampaigns)(nil)
