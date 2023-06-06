// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package snsiface provides an interface to enable mocking the Amazon Simple Notification Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package snsiface

import (
	"github.com/qtdslly/aws-sdk-go/aws"
	"github.com/qtdslly/aws-sdk-go/aws/request"
	"github.com/qtdslly/aws-sdk-go/service/sns"
)

// SNSAPI provides an interface to enable mocking the
// sns.SNS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Simple Notification Service.
//	func myFunc(svc snsiface.SNSAPI) bool {
//	    // Make svc.AddPermission request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := sns.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockSNSClient struct {
//	    snsiface.SNSAPI
//	}
//	func (m *mockSNSClient) AddPermission(input *sns.AddPermissionInput) (*sns.AddPermissionOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockSNSClient{}
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
type SNSAPI interface {
	AddPermission(*sns.AddPermissionInput) (*sns.AddPermissionOutput, error)
	AddPermissionWithContext(aws.Context, *sns.AddPermissionInput, ...request.Option) (*sns.AddPermissionOutput, error)
	AddPermissionRequest(*sns.AddPermissionInput) (*request.Request, *sns.AddPermissionOutput)

	CheckIfPhoneNumberIsOptedOut(*sns.CheckIfPhoneNumberIsOptedOutInput) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error)
	CheckIfPhoneNumberIsOptedOutWithContext(aws.Context, *sns.CheckIfPhoneNumberIsOptedOutInput, ...request.Option) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error)
	CheckIfPhoneNumberIsOptedOutRequest(*sns.CheckIfPhoneNumberIsOptedOutInput) (*request.Request, *sns.CheckIfPhoneNumberIsOptedOutOutput)

	ConfirmSubscription(*sns.ConfirmSubscriptionInput) (*sns.ConfirmSubscriptionOutput, error)
	ConfirmSubscriptionWithContext(aws.Context, *sns.ConfirmSubscriptionInput, ...request.Option) (*sns.ConfirmSubscriptionOutput, error)
	ConfirmSubscriptionRequest(*sns.ConfirmSubscriptionInput) (*request.Request, *sns.ConfirmSubscriptionOutput)

	CreatePlatformApplication(*sns.CreatePlatformApplicationInput) (*sns.CreatePlatformApplicationOutput, error)
	CreatePlatformApplicationWithContext(aws.Context, *sns.CreatePlatformApplicationInput, ...request.Option) (*sns.CreatePlatformApplicationOutput, error)
	CreatePlatformApplicationRequest(*sns.CreatePlatformApplicationInput) (*request.Request, *sns.CreatePlatformApplicationOutput)

	CreatePlatformEndpoint(*sns.CreatePlatformEndpointInput) (*sns.CreatePlatformEndpointOutput, error)
	CreatePlatformEndpointWithContext(aws.Context, *sns.CreatePlatformEndpointInput, ...request.Option) (*sns.CreatePlatformEndpointOutput, error)
	CreatePlatformEndpointRequest(*sns.CreatePlatformEndpointInput) (*request.Request, *sns.CreatePlatformEndpointOutput)

	CreateSMSSandboxPhoneNumber(*sns.CreateSMSSandboxPhoneNumberInput) (*sns.CreateSMSSandboxPhoneNumberOutput, error)
	CreateSMSSandboxPhoneNumberWithContext(aws.Context, *sns.CreateSMSSandboxPhoneNumberInput, ...request.Option) (*sns.CreateSMSSandboxPhoneNumberOutput, error)
	CreateSMSSandboxPhoneNumberRequest(*sns.CreateSMSSandboxPhoneNumberInput) (*request.Request, *sns.CreateSMSSandboxPhoneNumberOutput)

	CreateTopic(*sns.CreateTopicInput) (*sns.CreateTopicOutput, error)
	CreateTopicWithContext(aws.Context, *sns.CreateTopicInput, ...request.Option) (*sns.CreateTopicOutput, error)
	CreateTopicRequest(*sns.CreateTopicInput) (*request.Request, *sns.CreateTopicOutput)

	DeleteEndpoint(*sns.DeleteEndpointInput) (*sns.DeleteEndpointOutput, error)
	DeleteEndpointWithContext(aws.Context, *sns.DeleteEndpointInput, ...request.Option) (*sns.DeleteEndpointOutput, error)
	DeleteEndpointRequest(*sns.DeleteEndpointInput) (*request.Request, *sns.DeleteEndpointOutput)

	DeletePlatformApplication(*sns.DeletePlatformApplicationInput) (*sns.DeletePlatformApplicationOutput, error)
	DeletePlatformApplicationWithContext(aws.Context, *sns.DeletePlatformApplicationInput, ...request.Option) (*sns.DeletePlatformApplicationOutput, error)
	DeletePlatformApplicationRequest(*sns.DeletePlatformApplicationInput) (*request.Request, *sns.DeletePlatformApplicationOutput)

	DeleteSMSSandboxPhoneNumber(*sns.DeleteSMSSandboxPhoneNumberInput) (*sns.DeleteSMSSandboxPhoneNumberOutput, error)
	DeleteSMSSandboxPhoneNumberWithContext(aws.Context, *sns.DeleteSMSSandboxPhoneNumberInput, ...request.Option) (*sns.DeleteSMSSandboxPhoneNumberOutput, error)
	DeleteSMSSandboxPhoneNumberRequest(*sns.DeleteSMSSandboxPhoneNumberInput) (*request.Request, *sns.DeleteSMSSandboxPhoneNumberOutput)

	DeleteTopic(*sns.DeleteTopicInput) (*sns.DeleteTopicOutput, error)
	DeleteTopicWithContext(aws.Context, *sns.DeleteTopicInput, ...request.Option) (*sns.DeleteTopicOutput, error)
	DeleteTopicRequest(*sns.DeleteTopicInput) (*request.Request, *sns.DeleteTopicOutput)

	GetDataProtectionPolicy(*sns.GetDataProtectionPolicyInput) (*sns.GetDataProtectionPolicyOutput, error)
	GetDataProtectionPolicyWithContext(aws.Context, *sns.GetDataProtectionPolicyInput, ...request.Option) (*sns.GetDataProtectionPolicyOutput, error)
	GetDataProtectionPolicyRequest(*sns.GetDataProtectionPolicyInput) (*request.Request, *sns.GetDataProtectionPolicyOutput)

	GetEndpointAttributes(*sns.GetEndpointAttributesInput) (*sns.GetEndpointAttributesOutput, error)
	GetEndpointAttributesWithContext(aws.Context, *sns.GetEndpointAttributesInput, ...request.Option) (*sns.GetEndpointAttributesOutput, error)
	GetEndpointAttributesRequest(*sns.GetEndpointAttributesInput) (*request.Request, *sns.GetEndpointAttributesOutput)

	GetPlatformApplicationAttributes(*sns.GetPlatformApplicationAttributesInput) (*sns.GetPlatformApplicationAttributesOutput, error)
	GetPlatformApplicationAttributesWithContext(aws.Context, *sns.GetPlatformApplicationAttributesInput, ...request.Option) (*sns.GetPlatformApplicationAttributesOutput, error)
	GetPlatformApplicationAttributesRequest(*sns.GetPlatformApplicationAttributesInput) (*request.Request, *sns.GetPlatformApplicationAttributesOutput)

	GetSMSAttributes(*sns.GetSMSAttributesInput) (*sns.GetSMSAttributesOutput, error)
	GetSMSAttributesWithContext(aws.Context, *sns.GetSMSAttributesInput, ...request.Option) (*sns.GetSMSAttributesOutput, error)
	GetSMSAttributesRequest(*sns.GetSMSAttributesInput) (*request.Request, *sns.GetSMSAttributesOutput)

	GetSMSSandboxAccountStatus(*sns.GetSMSSandboxAccountStatusInput) (*sns.GetSMSSandboxAccountStatusOutput, error)
	GetSMSSandboxAccountStatusWithContext(aws.Context, *sns.GetSMSSandboxAccountStatusInput, ...request.Option) (*sns.GetSMSSandboxAccountStatusOutput, error)
	GetSMSSandboxAccountStatusRequest(*sns.GetSMSSandboxAccountStatusInput) (*request.Request, *sns.GetSMSSandboxAccountStatusOutput)

	GetSubscriptionAttributes(*sns.GetSubscriptionAttributesInput) (*sns.GetSubscriptionAttributesOutput, error)
	GetSubscriptionAttributesWithContext(aws.Context, *sns.GetSubscriptionAttributesInput, ...request.Option) (*sns.GetSubscriptionAttributesOutput, error)
	GetSubscriptionAttributesRequest(*sns.GetSubscriptionAttributesInput) (*request.Request, *sns.GetSubscriptionAttributesOutput)

	GetTopicAttributes(*sns.GetTopicAttributesInput) (*sns.GetTopicAttributesOutput, error)
	GetTopicAttributesWithContext(aws.Context, *sns.GetTopicAttributesInput, ...request.Option) (*sns.GetTopicAttributesOutput, error)
	GetTopicAttributesRequest(*sns.GetTopicAttributesInput) (*request.Request, *sns.GetTopicAttributesOutput)

	ListEndpointsByPlatformApplication(*sns.ListEndpointsByPlatformApplicationInput) (*sns.ListEndpointsByPlatformApplicationOutput, error)
	ListEndpointsByPlatformApplicationWithContext(aws.Context, *sns.ListEndpointsByPlatformApplicationInput, ...request.Option) (*sns.ListEndpointsByPlatformApplicationOutput, error)
	ListEndpointsByPlatformApplicationRequest(*sns.ListEndpointsByPlatformApplicationInput) (*request.Request, *sns.ListEndpointsByPlatformApplicationOutput)

	ListEndpointsByPlatformApplicationPages(*sns.ListEndpointsByPlatformApplicationInput, func(*sns.ListEndpointsByPlatformApplicationOutput, bool) bool) error
	ListEndpointsByPlatformApplicationPagesWithContext(aws.Context, *sns.ListEndpointsByPlatformApplicationInput, func(*sns.ListEndpointsByPlatformApplicationOutput, bool) bool, ...request.Option) error

	ListOriginationNumbers(*sns.ListOriginationNumbersInput) (*sns.ListOriginationNumbersOutput, error)
	ListOriginationNumbersWithContext(aws.Context, *sns.ListOriginationNumbersInput, ...request.Option) (*sns.ListOriginationNumbersOutput, error)
	ListOriginationNumbersRequest(*sns.ListOriginationNumbersInput) (*request.Request, *sns.ListOriginationNumbersOutput)

	ListOriginationNumbersPages(*sns.ListOriginationNumbersInput, func(*sns.ListOriginationNumbersOutput, bool) bool) error
	ListOriginationNumbersPagesWithContext(aws.Context, *sns.ListOriginationNumbersInput, func(*sns.ListOriginationNumbersOutput, bool) bool, ...request.Option) error

	ListPhoneNumbersOptedOut(*sns.ListPhoneNumbersOptedOutInput) (*sns.ListPhoneNumbersOptedOutOutput, error)
	ListPhoneNumbersOptedOutWithContext(aws.Context, *sns.ListPhoneNumbersOptedOutInput, ...request.Option) (*sns.ListPhoneNumbersOptedOutOutput, error)
	ListPhoneNumbersOptedOutRequest(*sns.ListPhoneNumbersOptedOutInput) (*request.Request, *sns.ListPhoneNumbersOptedOutOutput)

	ListPhoneNumbersOptedOutPages(*sns.ListPhoneNumbersOptedOutInput, func(*sns.ListPhoneNumbersOptedOutOutput, bool) bool) error
	ListPhoneNumbersOptedOutPagesWithContext(aws.Context, *sns.ListPhoneNumbersOptedOutInput, func(*sns.ListPhoneNumbersOptedOutOutput, bool) bool, ...request.Option) error

	ListPlatformApplications(*sns.ListPlatformApplicationsInput) (*sns.ListPlatformApplicationsOutput, error)
	ListPlatformApplicationsWithContext(aws.Context, *sns.ListPlatformApplicationsInput, ...request.Option) (*sns.ListPlatformApplicationsOutput, error)
	ListPlatformApplicationsRequest(*sns.ListPlatformApplicationsInput) (*request.Request, *sns.ListPlatformApplicationsOutput)

	ListPlatformApplicationsPages(*sns.ListPlatformApplicationsInput, func(*sns.ListPlatformApplicationsOutput, bool) bool) error
	ListPlatformApplicationsPagesWithContext(aws.Context, *sns.ListPlatformApplicationsInput, func(*sns.ListPlatformApplicationsOutput, bool) bool, ...request.Option) error

	ListSMSSandboxPhoneNumbers(*sns.ListSMSSandboxPhoneNumbersInput) (*sns.ListSMSSandboxPhoneNumbersOutput, error)
	ListSMSSandboxPhoneNumbersWithContext(aws.Context, *sns.ListSMSSandboxPhoneNumbersInput, ...request.Option) (*sns.ListSMSSandboxPhoneNumbersOutput, error)
	ListSMSSandboxPhoneNumbersRequest(*sns.ListSMSSandboxPhoneNumbersInput) (*request.Request, *sns.ListSMSSandboxPhoneNumbersOutput)

	ListSMSSandboxPhoneNumbersPages(*sns.ListSMSSandboxPhoneNumbersInput, func(*sns.ListSMSSandboxPhoneNumbersOutput, bool) bool) error
	ListSMSSandboxPhoneNumbersPagesWithContext(aws.Context, *sns.ListSMSSandboxPhoneNumbersInput, func(*sns.ListSMSSandboxPhoneNumbersOutput, bool) bool, ...request.Option) error

	ListSubscriptions(*sns.ListSubscriptionsInput) (*sns.ListSubscriptionsOutput, error)
	ListSubscriptionsWithContext(aws.Context, *sns.ListSubscriptionsInput, ...request.Option) (*sns.ListSubscriptionsOutput, error)
	ListSubscriptionsRequest(*sns.ListSubscriptionsInput) (*request.Request, *sns.ListSubscriptionsOutput)

	ListSubscriptionsPages(*sns.ListSubscriptionsInput, func(*sns.ListSubscriptionsOutput, bool) bool) error
	ListSubscriptionsPagesWithContext(aws.Context, *sns.ListSubscriptionsInput, func(*sns.ListSubscriptionsOutput, bool) bool, ...request.Option) error

	ListSubscriptionsByTopic(*sns.ListSubscriptionsByTopicInput) (*sns.ListSubscriptionsByTopicOutput, error)
	ListSubscriptionsByTopicWithContext(aws.Context, *sns.ListSubscriptionsByTopicInput, ...request.Option) (*sns.ListSubscriptionsByTopicOutput, error)
	ListSubscriptionsByTopicRequest(*sns.ListSubscriptionsByTopicInput) (*request.Request, *sns.ListSubscriptionsByTopicOutput)

	ListSubscriptionsByTopicPages(*sns.ListSubscriptionsByTopicInput, func(*sns.ListSubscriptionsByTopicOutput, bool) bool) error
	ListSubscriptionsByTopicPagesWithContext(aws.Context, *sns.ListSubscriptionsByTopicInput, func(*sns.ListSubscriptionsByTopicOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*sns.ListTagsForResourceInput) (*sns.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *sns.ListTagsForResourceInput, ...request.Option) (*sns.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*sns.ListTagsForResourceInput) (*request.Request, *sns.ListTagsForResourceOutput)

	ListTopics(*sns.ListTopicsInput) (*sns.ListTopicsOutput, error)
	ListTopicsWithContext(aws.Context, *sns.ListTopicsInput, ...request.Option) (*sns.ListTopicsOutput, error)
	ListTopicsRequest(*sns.ListTopicsInput) (*request.Request, *sns.ListTopicsOutput)

	ListTopicsPages(*sns.ListTopicsInput, func(*sns.ListTopicsOutput, bool) bool) error
	ListTopicsPagesWithContext(aws.Context, *sns.ListTopicsInput, func(*sns.ListTopicsOutput, bool) bool, ...request.Option) error

	OptInPhoneNumber(*sns.OptInPhoneNumberInput) (*sns.OptInPhoneNumberOutput, error)
	OptInPhoneNumberWithContext(aws.Context, *sns.OptInPhoneNumberInput, ...request.Option) (*sns.OptInPhoneNumberOutput, error)
	OptInPhoneNumberRequest(*sns.OptInPhoneNumberInput) (*request.Request, *sns.OptInPhoneNumberOutput)

	Publish(*sns.PublishInput) (*sns.PublishOutput, error)
	PublishWithContext(aws.Context, *sns.PublishInput, ...request.Option) (*sns.PublishOutput, error)
	PublishRequest(*sns.PublishInput) (*request.Request, *sns.PublishOutput)

	PublishBatch(*sns.PublishBatchInput) (*sns.PublishBatchOutput, error)
	PublishBatchWithContext(aws.Context, *sns.PublishBatchInput, ...request.Option) (*sns.PublishBatchOutput, error)
	PublishBatchRequest(*sns.PublishBatchInput) (*request.Request, *sns.PublishBatchOutput)

	PutDataProtectionPolicy(*sns.PutDataProtectionPolicyInput) (*sns.PutDataProtectionPolicyOutput, error)
	PutDataProtectionPolicyWithContext(aws.Context, *sns.PutDataProtectionPolicyInput, ...request.Option) (*sns.PutDataProtectionPolicyOutput, error)
	PutDataProtectionPolicyRequest(*sns.PutDataProtectionPolicyInput) (*request.Request, *sns.PutDataProtectionPolicyOutput)

	RemovePermission(*sns.RemovePermissionInput) (*sns.RemovePermissionOutput, error)
	RemovePermissionWithContext(aws.Context, *sns.RemovePermissionInput, ...request.Option) (*sns.RemovePermissionOutput, error)
	RemovePermissionRequest(*sns.RemovePermissionInput) (*request.Request, *sns.RemovePermissionOutput)

	SetEndpointAttributes(*sns.SetEndpointAttributesInput) (*sns.SetEndpointAttributesOutput, error)
	SetEndpointAttributesWithContext(aws.Context, *sns.SetEndpointAttributesInput, ...request.Option) (*sns.SetEndpointAttributesOutput, error)
	SetEndpointAttributesRequest(*sns.SetEndpointAttributesInput) (*request.Request, *sns.SetEndpointAttributesOutput)

	SetPlatformApplicationAttributes(*sns.SetPlatformApplicationAttributesInput) (*sns.SetPlatformApplicationAttributesOutput, error)
	SetPlatformApplicationAttributesWithContext(aws.Context, *sns.SetPlatformApplicationAttributesInput, ...request.Option) (*sns.SetPlatformApplicationAttributesOutput, error)
	SetPlatformApplicationAttributesRequest(*sns.SetPlatformApplicationAttributesInput) (*request.Request, *sns.SetPlatformApplicationAttributesOutput)

	SetSMSAttributes(*sns.SetSMSAttributesInput) (*sns.SetSMSAttributesOutput, error)
	SetSMSAttributesWithContext(aws.Context, *sns.SetSMSAttributesInput, ...request.Option) (*sns.SetSMSAttributesOutput, error)
	SetSMSAttributesRequest(*sns.SetSMSAttributesInput) (*request.Request, *sns.SetSMSAttributesOutput)

	SetSubscriptionAttributes(*sns.SetSubscriptionAttributesInput) (*sns.SetSubscriptionAttributesOutput, error)
	SetSubscriptionAttributesWithContext(aws.Context, *sns.SetSubscriptionAttributesInput, ...request.Option) (*sns.SetSubscriptionAttributesOutput, error)
	SetSubscriptionAttributesRequest(*sns.SetSubscriptionAttributesInput) (*request.Request, *sns.SetSubscriptionAttributesOutput)

	SetTopicAttributes(*sns.SetTopicAttributesInput) (*sns.SetTopicAttributesOutput, error)
	SetTopicAttributesWithContext(aws.Context, *sns.SetTopicAttributesInput, ...request.Option) (*sns.SetTopicAttributesOutput, error)
	SetTopicAttributesRequest(*sns.SetTopicAttributesInput) (*request.Request, *sns.SetTopicAttributesOutput)

	Subscribe(*sns.SubscribeInput) (*sns.SubscribeOutput, error)
	SubscribeWithContext(aws.Context, *sns.SubscribeInput, ...request.Option) (*sns.SubscribeOutput, error)
	SubscribeRequest(*sns.SubscribeInput) (*request.Request, *sns.SubscribeOutput)

	TagResource(*sns.TagResourceInput) (*sns.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *sns.TagResourceInput, ...request.Option) (*sns.TagResourceOutput, error)
	TagResourceRequest(*sns.TagResourceInput) (*request.Request, *sns.TagResourceOutput)

	Unsubscribe(*sns.UnsubscribeInput) (*sns.UnsubscribeOutput, error)
	UnsubscribeWithContext(aws.Context, *sns.UnsubscribeInput, ...request.Option) (*sns.UnsubscribeOutput, error)
	UnsubscribeRequest(*sns.UnsubscribeInput) (*request.Request, *sns.UnsubscribeOutput)

	UntagResource(*sns.UntagResourceInput) (*sns.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *sns.UntagResourceInput, ...request.Option) (*sns.UntagResourceOutput, error)
	UntagResourceRequest(*sns.UntagResourceInput) (*request.Request, *sns.UntagResourceOutput)

	VerifySMSSandboxPhoneNumber(*sns.VerifySMSSandboxPhoneNumberInput) (*sns.VerifySMSSandboxPhoneNumberOutput, error)
	VerifySMSSandboxPhoneNumberWithContext(aws.Context, *sns.VerifySMSSandboxPhoneNumberInput, ...request.Option) (*sns.VerifySMSSandboxPhoneNumberOutput, error)
	VerifySMSSandboxPhoneNumberRequest(*sns.VerifySMSSandboxPhoneNumberInput) (*request.Request, *sns.VerifySMSSandboxPhoneNumberOutput)
}

var _ SNSAPI = (*sns.SNS)(nil)
