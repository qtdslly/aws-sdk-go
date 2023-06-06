// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package chimesdkidentityiface provides an interface to enable mocking the Amazon Chime SDK Identity service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package chimesdkidentityiface

import (
	"github.com/qtdslly/aws-sdk-go/aws"
	"github.com/qtdslly/aws-sdk-go/aws/request"
	"github.com/qtdslly/aws-sdk-go/service/chimesdkidentity"
)

// ChimeSDKIdentityAPI provides an interface to enable mocking the
// chimesdkidentity.ChimeSDKIdentity service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Chime SDK Identity.
//	func myFunc(svc chimesdkidentityiface.ChimeSDKIdentityAPI) bool {
//	    // Make svc.CreateAppInstance request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := chimesdkidentity.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockChimeSDKIdentityClient struct {
//	    chimesdkidentityiface.ChimeSDKIdentityAPI
//	}
//	func (m *mockChimeSDKIdentityClient) CreateAppInstance(input *chimesdkidentity.CreateAppInstanceInput) (*chimesdkidentity.CreateAppInstanceOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockChimeSDKIdentityClient{}
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
type ChimeSDKIdentityAPI interface {
	CreateAppInstance(*chimesdkidentity.CreateAppInstanceInput) (*chimesdkidentity.CreateAppInstanceOutput, error)
	CreateAppInstanceWithContext(aws.Context, *chimesdkidentity.CreateAppInstanceInput, ...request.Option) (*chimesdkidentity.CreateAppInstanceOutput, error)
	CreateAppInstanceRequest(*chimesdkidentity.CreateAppInstanceInput) (*request.Request, *chimesdkidentity.CreateAppInstanceOutput)

	CreateAppInstanceAdmin(*chimesdkidentity.CreateAppInstanceAdminInput) (*chimesdkidentity.CreateAppInstanceAdminOutput, error)
	CreateAppInstanceAdminWithContext(aws.Context, *chimesdkidentity.CreateAppInstanceAdminInput, ...request.Option) (*chimesdkidentity.CreateAppInstanceAdminOutput, error)
	CreateAppInstanceAdminRequest(*chimesdkidentity.CreateAppInstanceAdminInput) (*request.Request, *chimesdkidentity.CreateAppInstanceAdminOutput)

	CreateAppInstanceBot(*chimesdkidentity.CreateAppInstanceBotInput) (*chimesdkidentity.CreateAppInstanceBotOutput, error)
	CreateAppInstanceBotWithContext(aws.Context, *chimesdkidentity.CreateAppInstanceBotInput, ...request.Option) (*chimesdkidentity.CreateAppInstanceBotOutput, error)
	CreateAppInstanceBotRequest(*chimesdkidentity.CreateAppInstanceBotInput) (*request.Request, *chimesdkidentity.CreateAppInstanceBotOutput)

	CreateAppInstanceUser(*chimesdkidentity.CreateAppInstanceUserInput) (*chimesdkidentity.CreateAppInstanceUserOutput, error)
	CreateAppInstanceUserWithContext(aws.Context, *chimesdkidentity.CreateAppInstanceUserInput, ...request.Option) (*chimesdkidentity.CreateAppInstanceUserOutput, error)
	CreateAppInstanceUserRequest(*chimesdkidentity.CreateAppInstanceUserInput) (*request.Request, *chimesdkidentity.CreateAppInstanceUserOutput)

	DeleteAppInstance(*chimesdkidentity.DeleteAppInstanceInput) (*chimesdkidentity.DeleteAppInstanceOutput, error)
	DeleteAppInstanceWithContext(aws.Context, *chimesdkidentity.DeleteAppInstanceInput, ...request.Option) (*chimesdkidentity.DeleteAppInstanceOutput, error)
	DeleteAppInstanceRequest(*chimesdkidentity.DeleteAppInstanceInput) (*request.Request, *chimesdkidentity.DeleteAppInstanceOutput)

	DeleteAppInstanceAdmin(*chimesdkidentity.DeleteAppInstanceAdminInput) (*chimesdkidentity.DeleteAppInstanceAdminOutput, error)
	DeleteAppInstanceAdminWithContext(aws.Context, *chimesdkidentity.DeleteAppInstanceAdminInput, ...request.Option) (*chimesdkidentity.DeleteAppInstanceAdminOutput, error)
	DeleteAppInstanceAdminRequest(*chimesdkidentity.DeleteAppInstanceAdminInput) (*request.Request, *chimesdkidentity.DeleteAppInstanceAdminOutput)

	DeleteAppInstanceBot(*chimesdkidentity.DeleteAppInstanceBotInput) (*chimesdkidentity.DeleteAppInstanceBotOutput, error)
	DeleteAppInstanceBotWithContext(aws.Context, *chimesdkidentity.DeleteAppInstanceBotInput, ...request.Option) (*chimesdkidentity.DeleteAppInstanceBotOutput, error)
	DeleteAppInstanceBotRequest(*chimesdkidentity.DeleteAppInstanceBotInput) (*request.Request, *chimesdkidentity.DeleteAppInstanceBotOutput)

	DeleteAppInstanceUser(*chimesdkidentity.DeleteAppInstanceUserInput) (*chimesdkidentity.DeleteAppInstanceUserOutput, error)
	DeleteAppInstanceUserWithContext(aws.Context, *chimesdkidentity.DeleteAppInstanceUserInput, ...request.Option) (*chimesdkidentity.DeleteAppInstanceUserOutput, error)
	DeleteAppInstanceUserRequest(*chimesdkidentity.DeleteAppInstanceUserInput) (*request.Request, *chimesdkidentity.DeleteAppInstanceUserOutput)

	DeregisterAppInstanceUserEndpoint(*chimesdkidentity.DeregisterAppInstanceUserEndpointInput) (*chimesdkidentity.DeregisterAppInstanceUserEndpointOutput, error)
	DeregisterAppInstanceUserEndpointWithContext(aws.Context, *chimesdkidentity.DeregisterAppInstanceUserEndpointInput, ...request.Option) (*chimesdkidentity.DeregisterAppInstanceUserEndpointOutput, error)
	DeregisterAppInstanceUserEndpointRequest(*chimesdkidentity.DeregisterAppInstanceUserEndpointInput) (*request.Request, *chimesdkidentity.DeregisterAppInstanceUserEndpointOutput)

	DescribeAppInstance(*chimesdkidentity.DescribeAppInstanceInput) (*chimesdkidentity.DescribeAppInstanceOutput, error)
	DescribeAppInstanceWithContext(aws.Context, *chimesdkidentity.DescribeAppInstanceInput, ...request.Option) (*chimesdkidentity.DescribeAppInstanceOutput, error)
	DescribeAppInstanceRequest(*chimesdkidentity.DescribeAppInstanceInput) (*request.Request, *chimesdkidentity.DescribeAppInstanceOutput)

	DescribeAppInstanceAdmin(*chimesdkidentity.DescribeAppInstanceAdminInput) (*chimesdkidentity.DescribeAppInstanceAdminOutput, error)
	DescribeAppInstanceAdminWithContext(aws.Context, *chimesdkidentity.DescribeAppInstanceAdminInput, ...request.Option) (*chimesdkidentity.DescribeAppInstanceAdminOutput, error)
	DescribeAppInstanceAdminRequest(*chimesdkidentity.DescribeAppInstanceAdminInput) (*request.Request, *chimesdkidentity.DescribeAppInstanceAdminOutput)

	DescribeAppInstanceBot(*chimesdkidentity.DescribeAppInstanceBotInput) (*chimesdkidentity.DescribeAppInstanceBotOutput, error)
	DescribeAppInstanceBotWithContext(aws.Context, *chimesdkidentity.DescribeAppInstanceBotInput, ...request.Option) (*chimesdkidentity.DescribeAppInstanceBotOutput, error)
	DescribeAppInstanceBotRequest(*chimesdkidentity.DescribeAppInstanceBotInput) (*request.Request, *chimesdkidentity.DescribeAppInstanceBotOutput)

	DescribeAppInstanceUser(*chimesdkidentity.DescribeAppInstanceUserInput) (*chimesdkidentity.DescribeAppInstanceUserOutput, error)
	DescribeAppInstanceUserWithContext(aws.Context, *chimesdkidentity.DescribeAppInstanceUserInput, ...request.Option) (*chimesdkidentity.DescribeAppInstanceUserOutput, error)
	DescribeAppInstanceUserRequest(*chimesdkidentity.DescribeAppInstanceUserInput) (*request.Request, *chimesdkidentity.DescribeAppInstanceUserOutput)

	DescribeAppInstanceUserEndpoint(*chimesdkidentity.DescribeAppInstanceUserEndpointInput) (*chimesdkidentity.DescribeAppInstanceUserEndpointOutput, error)
	DescribeAppInstanceUserEndpointWithContext(aws.Context, *chimesdkidentity.DescribeAppInstanceUserEndpointInput, ...request.Option) (*chimesdkidentity.DescribeAppInstanceUserEndpointOutput, error)
	DescribeAppInstanceUserEndpointRequest(*chimesdkidentity.DescribeAppInstanceUserEndpointInput) (*request.Request, *chimesdkidentity.DescribeAppInstanceUserEndpointOutput)

	GetAppInstanceRetentionSettings(*chimesdkidentity.GetAppInstanceRetentionSettingsInput) (*chimesdkidentity.GetAppInstanceRetentionSettingsOutput, error)
	GetAppInstanceRetentionSettingsWithContext(aws.Context, *chimesdkidentity.GetAppInstanceRetentionSettingsInput, ...request.Option) (*chimesdkidentity.GetAppInstanceRetentionSettingsOutput, error)
	GetAppInstanceRetentionSettingsRequest(*chimesdkidentity.GetAppInstanceRetentionSettingsInput) (*request.Request, *chimesdkidentity.GetAppInstanceRetentionSettingsOutput)

	ListAppInstanceAdmins(*chimesdkidentity.ListAppInstanceAdminsInput) (*chimesdkidentity.ListAppInstanceAdminsOutput, error)
	ListAppInstanceAdminsWithContext(aws.Context, *chimesdkidentity.ListAppInstanceAdminsInput, ...request.Option) (*chimesdkidentity.ListAppInstanceAdminsOutput, error)
	ListAppInstanceAdminsRequest(*chimesdkidentity.ListAppInstanceAdminsInput) (*request.Request, *chimesdkidentity.ListAppInstanceAdminsOutput)

	ListAppInstanceAdminsPages(*chimesdkidentity.ListAppInstanceAdminsInput, func(*chimesdkidentity.ListAppInstanceAdminsOutput, bool) bool) error
	ListAppInstanceAdminsPagesWithContext(aws.Context, *chimesdkidentity.ListAppInstanceAdminsInput, func(*chimesdkidentity.ListAppInstanceAdminsOutput, bool) bool, ...request.Option) error

	ListAppInstanceBots(*chimesdkidentity.ListAppInstanceBotsInput) (*chimesdkidentity.ListAppInstanceBotsOutput, error)
	ListAppInstanceBotsWithContext(aws.Context, *chimesdkidentity.ListAppInstanceBotsInput, ...request.Option) (*chimesdkidentity.ListAppInstanceBotsOutput, error)
	ListAppInstanceBotsRequest(*chimesdkidentity.ListAppInstanceBotsInput) (*request.Request, *chimesdkidentity.ListAppInstanceBotsOutput)

	ListAppInstanceBotsPages(*chimesdkidentity.ListAppInstanceBotsInput, func(*chimesdkidentity.ListAppInstanceBotsOutput, bool) bool) error
	ListAppInstanceBotsPagesWithContext(aws.Context, *chimesdkidentity.ListAppInstanceBotsInput, func(*chimesdkidentity.ListAppInstanceBotsOutput, bool) bool, ...request.Option) error

	ListAppInstanceUserEndpoints(*chimesdkidentity.ListAppInstanceUserEndpointsInput) (*chimesdkidentity.ListAppInstanceUserEndpointsOutput, error)
	ListAppInstanceUserEndpointsWithContext(aws.Context, *chimesdkidentity.ListAppInstanceUserEndpointsInput, ...request.Option) (*chimesdkidentity.ListAppInstanceUserEndpointsOutput, error)
	ListAppInstanceUserEndpointsRequest(*chimesdkidentity.ListAppInstanceUserEndpointsInput) (*request.Request, *chimesdkidentity.ListAppInstanceUserEndpointsOutput)

	ListAppInstanceUserEndpointsPages(*chimesdkidentity.ListAppInstanceUserEndpointsInput, func(*chimesdkidentity.ListAppInstanceUserEndpointsOutput, bool) bool) error
	ListAppInstanceUserEndpointsPagesWithContext(aws.Context, *chimesdkidentity.ListAppInstanceUserEndpointsInput, func(*chimesdkidentity.ListAppInstanceUserEndpointsOutput, bool) bool, ...request.Option) error

	ListAppInstanceUsers(*chimesdkidentity.ListAppInstanceUsersInput) (*chimesdkidentity.ListAppInstanceUsersOutput, error)
	ListAppInstanceUsersWithContext(aws.Context, *chimesdkidentity.ListAppInstanceUsersInput, ...request.Option) (*chimesdkidentity.ListAppInstanceUsersOutput, error)
	ListAppInstanceUsersRequest(*chimesdkidentity.ListAppInstanceUsersInput) (*request.Request, *chimesdkidentity.ListAppInstanceUsersOutput)

	ListAppInstanceUsersPages(*chimesdkidentity.ListAppInstanceUsersInput, func(*chimesdkidentity.ListAppInstanceUsersOutput, bool) bool) error
	ListAppInstanceUsersPagesWithContext(aws.Context, *chimesdkidentity.ListAppInstanceUsersInput, func(*chimesdkidentity.ListAppInstanceUsersOutput, bool) bool, ...request.Option) error

	ListAppInstances(*chimesdkidentity.ListAppInstancesInput) (*chimesdkidentity.ListAppInstancesOutput, error)
	ListAppInstancesWithContext(aws.Context, *chimesdkidentity.ListAppInstancesInput, ...request.Option) (*chimesdkidentity.ListAppInstancesOutput, error)
	ListAppInstancesRequest(*chimesdkidentity.ListAppInstancesInput) (*request.Request, *chimesdkidentity.ListAppInstancesOutput)

	ListAppInstancesPages(*chimesdkidentity.ListAppInstancesInput, func(*chimesdkidentity.ListAppInstancesOutput, bool) bool) error
	ListAppInstancesPagesWithContext(aws.Context, *chimesdkidentity.ListAppInstancesInput, func(*chimesdkidentity.ListAppInstancesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*chimesdkidentity.ListTagsForResourceInput) (*chimesdkidentity.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *chimesdkidentity.ListTagsForResourceInput, ...request.Option) (*chimesdkidentity.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*chimesdkidentity.ListTagsForResourceInput) (*request.Request, *chimesdkidentity.ListTagsForResourceOutput)

	PutAppInstanceRetentionSettings(*chimesdkidentity.PutAppInstanceRetentionSettingsInput) (*chimesdkidentity.PutAppInstanceRetentionSettingsOutput, error)
	PutAppInstanceRetentionSettingsWithContext(aws.Context, *chimesdkidentity.PutAppInstanceRetentionSettingsInput, ...request.Option) (*chimesdkidentity.PutAppInstanceRetentionSettingsOutput, error)
	PutAppInstanceRetentionSettingsRequest(*chimesdkidentity.PutAppInstanceRetentionSettingsInput) (*request.Request, *chimesdkidentity.PutAppInstanceRetentionSettingsOutput)

	PutAppInstanceUserExpirationSettings(*chimesdkidentity.PutAppInstanceUserExpirationSettingsInput) (*chimesdkidentity.PutAppInstanceUserExpirationSettingsOutput, error)
	PutAppInstanceUserExpirationSettingsWithContext(aws.Context, *chimesdkidentity.PutAppInstanceUserExpirationSettingsInput, ...request.Option) (*chimesdkidentity.PutAppInstanceUserExpirationSettingsOutput, error)
	PutAppInstanceUserExpirationSettingsRequest(*chimesdkidentity.PutAppInstanceUserExpirationSettingsInput) (*request.Request, *chimesdkidentity.PutAppInstanceUserExpirationSettingsOutput)

	RegisterAppInstanceUserEndpoint(*chimesdkidentity.RegisterAppInstanceUserEndpointInput) (*chimesdkidentity.RegisterAppInstanceUserEndpointOutput, error)
	RegisterAppInstanceUserEndpointWithContext(aws.Context, *chimesdkidentity.RegisterAppInstanceUserEndpointInput, ...request.Option) (*chimesdkidentity.RegisterAppInstanceUserEndpointOutput, error)
	RegisterAppInstanceUserEndpointRequest(*chimesdkidentity.RegisterAppInstanceUserEndpointInput) (*request.Request, *chimesdkidentity.RegisterAppInstanceUserEndpointOutput)

	TagResource(*chimesdkidentity.TagResourceInput) (*chimesdkidentity.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *chimesdkidentity.TagResourceInput, ...request.Option) (*chimesdkidentity.TagResourceOutput, error)
	TagResourceRequest(*chimesdkidentity.TagResourceInput) (*request.Request, *chimesdkidentity.TagResourceOutput)

	UntagResource(*chimesdkidentity.UntagResourceInput) (*chimesdkidentity.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *chimesdkidentity.UntagResourceInput, ...request.Option) (*chimesdkidentity.UntagResourceOutput, error)
	UntagResourceRequest(*chimesdkidentity.UntagResourceInput) (*request.Request, *chimesdkidentity.UntagResourceOutput)

	UpdateAppInstance(*chimesdkidentity.UpdateAppInstanceInput) (*chimesdkidentity.UpdateAppInstanceOutput, error)
	UpdateAppInstanceWithContext(aws.Context, *chimesdkidentity.UpdateAppInstanceInput, ...request.Option) (*chimesdkidentity.UpdateAppInstanceOutput, error)
	UpdateAppInstanceRequest(*chimesdkidentity.UpdateAppInstanceInput) (*request.Request, *chimesdkidentity.UpdateAppInstanceOutput)

	UpdateAppInstanceBot(*chimesdkidentity.UpdateAppInstanceBotInput) (*chimesdkidentity.UpdateAppInstanceBotOutput, error)
	UpdateAppInstanceBotWithContext(aws.Context, *chimesdkidentity.UpdateAppInstanceBotInput, ...request.Option) (*chimesdkidentity.UpdateAppInstanceBotOutput, error)
	UpdateAppInstanceBotRequest(*chimesdkidentity.UpdateAppInstanceBotInput) (*request.Request, *chimesdkidentity.UpdateAppInstanceBotOutput)

	UpdateAppInstanceUser(*chimesdkidentity.UpdateAppInstanceUserInput) (*chimesdkidentity.UpdateAppInstanceUserOutput, error)
	UpdateAppInstanceUserWithContext(aws.Context, *chimesdkidentity.UpdateAppInstanceUserInput, ...request.Option) (*chimesdkidentity.UpdateAppInstanceUserOutput, error)
	UpdateAppInstanceUserRequest(*chimesdkidentity.UpdateAppInstanceUserInput) (*request.Request, *chimesdkidentity.UpdateAppInstanceUserOutput)

	UpdateAppInstanceUserEndpoint(*chimesdkidentity.UpdateAppInstanceUserEndpointInput) (*chimesdkidentity.UpdateAppInstanceUserEndpointOutput, error)
	UpdateAppInstanceUserEndpointWithContext(aws.Context, *chimesdkidentity.UpdateAppInstanceUserEndpointInput, ...request.Option) (*chimesdkidentity.UpdateAppInstanceUserEndpointOutput, error)
	UpdateAppInstanceUserEndpointRequest(*chimesdkidentity.UpdateAppInstanceUserEndpointInput) (*request.Request, *chimesdkidentity.UpdateAppInstanceUserEndpointOutput)
}

var _ ChimeSDKIdentityAPI = (*chimesdkidentity.ChimeSDKIdentity)(nil)