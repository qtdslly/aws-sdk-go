// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package vpclatticeiface provides an interface to enable mocking the Amazon VPC Lattice service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package vpclatticeiface

import (
	"github.com/qtdslly/aws-sdk-go/aws"
	"github.com/qtdslly/aws-sdk-go/aws/request"
	"github.com/qtdslly/aws-sdk-go/service/vpclattice"
)

// VPCLatticeAPI provides an interface to enable mocking the
// vpclattice.VPCLattice service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon VPC Lattice.
//	func myFunc(svc vpclatticeiface.VPCLatticeAPI) bool {
//	    // Make svc.BatchUpdateRule request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := vpclattice.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockVPCLatticeClient struct {
//	    vpclatticeiface.VPCLatticeAPI
//	}
//	func (m *mockVPCLatticeClient) BatchUpdateRule(input *vpclattice.BatchUpdateRuleInput) (*vpclattice.BatchUpdateRuleOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockVPCLatticeClient{}
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
type VPCLatticeAPI interface {
	BatchUpdateRule(*vpclattice.BatchUpdateRuleInput) (*vpclattice.BatchUpdateRuleOutput, error)
	BatchUpdateRuleWithContext(aws.Context, *vpclattice.BatchUpdateRuleInput, ...request.Option) (*vpclattice.BatchUpdateRuleOutput, error)
	BatchUpdateRuleRequest(*vpclattice.BatchUpdateRuleInput) (*request.Request, *vpclattice.BatchUpdateRuleOutput)

	CreateAccessLogSubscription(*vpclattice.CreateAccessLogSubscriptionInput) (*vpclattice.CreateAccessLogSubscriptionOutput, error)
	CreateAccessLogSubscriptionWithContext(aws.Context, *vpclattice.CreateAccessLogSubscriptionInput, ...request.Option) (*vpclattice.CreateAccessLogSubscriptionOutput, error)
	CreateAccessLogSubscriptionRequest(*vpclattice.CreateAccessLogSubscriptionInput) (*request.Request, *vpclattice.CreateAccessLogSubscriptionOutput)

	CreateListener(*vpclattice.CreateListenerInput) (*vpclattice.CreateListenerOutput, error)
	CreateListenerWithContext(aws.Context, *vpclattice.CreateListenerInput, ...request.Option) (*vpclattice.CreateListenerOutput, error)
	CreateListenerRequest(*vpclattice.CreateListenerInput) (*request.Request, *vpclattice.CreateListenerOutput)

	CreateRule(*vpclattice.CreateRuleInput) (*vpclattice.CreateRuleOutput, error)
	CreateRuleWithContext(aws.Context, *vpclattice.CreateRuleInput, ...request.Option) (*vpclattice.CreateRuleOutput, error)
	CreateRuleRequest(*vpclattice.CreateRuleInput) (*request.Request, *vpclattice.CreateRuleOutput)

	CreateService(*vpclattice.CreateServiceInput) (*vpclattice.CreateServiceOutput, error)
	CreateServiceWithContext(aws.Context, *vpclattice.CreateServiceInput, ...request.Option) (*vpclattice.CreateServiceOutput, error)
	CreateServiceRequest(*vpclattice.CreateServiceInput) (*request.Request, *vpclattice.CreateServiceOutput)

	CreateServiceNetwork(*vpclattice.CreateServiceNetworkInput) (*vpclattice.CreateServiceNetworkOutput, error)
	CreateServiceNetworkWithContext(aws.Context, *vpclattice.CreateServiceNetworkInput, ...request.Option) (*vpclattice.CreateServiceNetworkOutput, error)
	CreateServiceNetworkRequest(*vpclattice.CreateServiceNetworkInput) (*request.Request, *vpclattice.CreateServiceNetworkOutput)

	CreateServiceNetworkServiceAssociation(*vpclattice.CreateServiceNetworkServiceAssociationInput) (*vpclattice.CreateServiceNetworkServiceAssociationOutput, error)
	CreateServiceNetworkServiceAssociationWithContext(aws.Context, *vpclattice.CreateServiceNetworkServiceAssociationInput, ...request.Option) (*vpclattice.CreateServiceNetworkServiceAssociationOutput, error)
	CreateServiceNetworkServiceAssociationRequest(*vpclattice.CreateServiceNetworkServiceAssociationInput) (*request.Request, *vpclattice.CreateServiceNetworkServiceAssociationOutput)

	CreateServiceNetworkVpcAssociation(*vpclattice.CreateServiceNetworkVpcAssociationInput) (*vpclattice.CreateServiceNetworkVpcAssociationOutput, error)
	CreateServiceNetworkVpcAssociationWithContext(aws.Context, *vpclattice.CreateServiceNetworkVpcAssociationInput, ...request.Option) (*vpclattice.CreateServiceNetworkVpcAssociationOutput, error)
	CreateServiceNetworkVpcAssociationRequest(*vpclattice.CreateServiceNetworkVpcAssociationInput) (*request.Request, *vpclattice.CreateServiceNetworkVpcAssociationOutput)

	CreateTargetGroup(*vpclattice.CreateTargetGroupInput) (*vpclattice.CreateTargetGroupOutput, error)
	CreateTargetGroupWithContext(aws.Context, *vpclattice.CreateTargetGroupInput, ...request.Option) (*vpclattice.CreateTargetGroupOutput, error)
	CreateTargetGroupRequest(*vpclattice.CreateTargetGroupInput) (*request.Request, *vpclattice.CreateTargetGroupOutput)

	DeleteAccessLogSubscription(*vpclattice.DeleteAccessLogSubscriptionInput) (*vpclattice.DeleteAccessLogSubscriptionOutput, error)
	DeleteAccessLogSubscriptionWithContext(aws.Context, *vpclattice.DeleteAccessLogSubscriptionInput, ...request.Option) (*vpclattice.DeleteAccessLogSubscriptionOutput, error)
	DeleteAccessLogSubscriptionRequest(*vpclattice.DeleteAccessLogSubscriptionInput) (*request.Request, *vpclattice.DeleteAccessLogSubscriptionOutput)

	DeleteAuthPolicy(*vpclattice.DeleteAuthPolicyInput) (*vpclattice.DeleteAuthPolicyOutput, error)
	DeleteAuthPolicyWithContext(aws.Context, *vpclattice.DeleteAuthPolicyInput, ...request.Option) (*vpclattice.DeleteAuthPolicyOutput, error)
	DeleteAuthPolicyRequest(*vpclattice.DeleteAuthPolicyInput) (*request.Request, *vpclattice.DeleteAuthPolicyOutput)

	DeleteListener(*vpclattice.DeleteListenerInput) (*vpclattice.DeleteListenerOutput, error)
	DeleteListenerWithContext(aws.Context, *vpclattice.DeleteListenerInput, ...request.Option) (*vpclattice.DeleteListenerOutput, error)
	DeleteListenerRequest(*vpclattice.DeleteListenerInput) (*request.Request, *vpclattice.DeleteListenerOutput)

	DeleteResourcePolicy(*vpclattice.DeleteResourcePolicyInput) (*vpclattice.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyWithContext(aws.Context, *vpclattice.DeleteResourcePolicyInput, ...request.Option) (*vpclattice.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyRequest(*vpclattice.DeleteResourcePolicyInput) (*request.Request, *vpclattice.DeleteResourcePolicyOutput)

	DeleteRule(*vpclattice.DeleteRuleInput) (*vpclattice.DeleteRuleOutput, error)
	DeleteRuleWithContext(aws.Context, *vpclattice.DeleteRuleInput, ...request.Option) (*vpclattice.DeleteRuleOutput, error)
	DeleteRuleRequest(*vpclattice.DeleteRuleInput) (*request.Request, *vpclattice.DeleteRuleOutput)

	DeleteService(*vpclattice.DeleteServiceInput) (*vpclattice.DeleteServiceOutput, error)
	DeleteServiceWithContext(aws.Context, *vpclattice.DeleteServiceInput, ...request.Option) (*vpclattice.DeleteServiceOutput, error)
	DeleteServiceRequest(*vpclattice.DeleteServiceInput) (*request.Request, *vpclattice.DeleteServiceOutput)

	DeleteServiceNetwork(*vpclattice.DeleteServiceNetworkInput) (*vpclattice.DeleteServiceNetworkOutput, error)
	DeleteServiceNetworkWithContext(aws.Context, *vpclattice.DeleteServiceNetworkInput, ...request.Option) (*vpclattice.DeleteServiceNetworkOutput, error)
	DeleteServiceNetworkRequest(*vpclattice.DeleteServiceNetworkInput) (*request.Request, *vpclattice.DeleteServiceNetworkOutput)

	DeleteServiceNetworkServiceAssociation(*vpclattice.DeleteServiceNetworkServiceAssociationInput) (*vpclattice.DeleteServiceNetworkServiceAssociationOutput, error)
	DeleteServiceNetworkServiceAssociationWithContext(aws.Context, *vpclattice.DeleteServiceNetworkServiceAssociationInput, ...request.Option) (*vpclattice.DeleteServiceNetworkServiceAssociationOutput, error)
	DeleteServiceNetworkServiceAssociationRequest(*vpclattice.DeleteServiceNetworkServiceAssociationInput) (*request.Request, *vpclattice.DeleteServiceNetworkServiceAssociationOutput)

	DeleteServiceNetworkVpcAssociation(*vpclattice.DeleteServiceNetworkVpcAssociationInput) (*vpclattice.DeleteServiceNetworkVpcAssociationOutput, error)
	DeleteServiceNetworkVpcAssociationWithContext(aws.Context, *vpclattice.DeleteServiceNetworkVpcAssociationInput, ...request.Option) (*vpclattice.DeleteServiceNetworkVpcAssociationOutput, error)
	DeleteServiceNetworkVpcAssociationRequest(*vpclattice.DeleteServiceNetworkVpcAssociationInput) (*request.Request, *vpclattice.DeleteServiceNetworkVpcAssociationOutput)

	DeleteTargetGroup(*vpclattice.DeleteTargetGroupInput) (*vpclattice.DeleteTargetGroupOutput, error)
	DeleteTargetGroupWithContext(aws.Context, *vpclattice.DeleteTargetGroupInput, ...request.Option) (*vpclattice.DeleteTargetGroupOutput, error)
	DeleteTargetGroupRequest(*vpclattice.DeleteTargetGroupInput) (*request.Request, *vpclattice.DeleteTargetGroupOutput)

	DeregisterTargets(*vpclattice.DeregisterTargetsInput) (*vpclattice.DeregisterTargetsOutput, error)
	DeregisterTargetsWithContext(aws.Context, *vpclattice.DeregisterTargetsInput, ...request.Option) (*vpclattice.DeregisterTargetsOutput, error)
	DeregisterTargetsRequest(*vpclattice.DeregisterTargetsInput) (*request.Request, *vpclattice.DeregisterTargetsOutput)

	GetAccessLogSubscription(*vpclattice.GetAccessLogSubscriptionInput) (*vpclattice.GetAccessLogSubscriptionOutput, error)
	GetAccessLogSubscriptionWithContext(aws.Context, *vpclattice.GetAccessLogSubscriptionInput, ...request.Option) (*vpclattice.GetAccessLogSubscriptionOutput, error)
	GetAccessLogSubscriptionRequest(*vpclattice.GetAccessLogSubscriptionInput) (*request.Request, *vpclattice.GetAccessLogSubscriptionOutput)

	GetAuthPolicy(*vpclattice.GetAuthPolicyInput) (*vpclattice.GetAuthPolicyOutput, error)
	GetAuthPolicyWithContext(aws.Context, *vpclattice.GetAuthPolicyInput, ...request.Option) (*vpclattice.GetAuthPolicyOutput, error)
	GetAuthPolicyRequest(*vpclattice.GetAuthPolicyInput) (*request.Request, *vpclattice.GetAuthPolicyOutput)

	GetListener(*vpclattice.GetListenerInput) (*vpclattice.GetListenerOutput, error)
	GetListenerWithContext(aws.Context, *vpclattice.GetListenerInput, ...request.Option) (*vpclattice.GetListenerOutput, error)
	GetListenerRequest(*vpclattice.GetListenerInput) (*request.Request, *vpclattice.GetListenerOutput)

	GetResourcePolicy(*vpclattice.GetResourcePolicyInput) (*vpclattice.GetResourcePolicyOutput, error)
	GetResourcePolicyWithContext(aws.Context, *vpclattice.GetResourcePolicyInput, ...request.Option) (*vpclattice.GetResourcePolicyOutput, error)
	GetResourcePolicyRequest(*vpclattice.GetResourcePolicyInput) (*request.Request, *vpclattice.GetResourcePolicyOutput)

	GetRule(*vpclattice.GetRuleInput) (*vpclattice.GetRuleOutput, error)
	GetRuleWithContext(aws.Context, *vpclattice.GetRuleInput, ...request.Option) (*vpclattice.GetRuleOutput, error)
	GetRuleRequest(*vpclattice.GetRuleInput) (*request.Request, *vpclattice.GetRuleOutput)

	GetService(*vpclattice.GetServiceInput) (*vpclattice.GetServiceOutput, error)
	GetServiceWithContext(aws.Context, *vpclattice.GetServiceInput, ...request.Option) (*vpclattice.GetServiceOutput, error)
	GetServiceRequest(*vpclattice.GetServiceInput) (*request.Request, *vpclattice.GetServiceOutput)

	GetServiceNetwork(*vpclattice.GetServiceNetworkInput) (*vpclattice.GetServiceNetworkOutput, error)
	GetServiceNetworkWithContext(aws.Context, *vpclattice.GetServiceNetworkInput, ...request.Option) (*vpclattice.GetServiceNetworkOutput, error)
	GetServiceNetworkRequest(*vpclattice.GetServiceNetworkInput) (*request.Request, *vpclattice.GetServiceNetworkOutput)

	GetServiceNetworkServiceAssociation(*vpclattice.GetServiceNetworkServiceAssociationInput) (*vpclattice.GetServiceNetworkServiceAssociationOutput, error)
	GetServiceNetworkServiceAssociationWithContext(aws.Context, *vpclattice.GetServiceNetworkServiceAssociationInput, ...request.Option) (*vpclattice.GetServiceNetworkServiceAssociationOutput, error)
	GetServiceNetworkServiceAssociationRequest(*vpclattice.GetServiceNetworkServiceAssociationInput) (*request.Request, *vpclattice.GetServiceNetworkServiceAssociationOutput)

	GetServiceNetworkVpcAssociation(*vpclattice.GetServiceNetworkVpcAssociationInput) (*vpclattice.GetServiceNetworkVpcAssociationOutput, error)
	GetServiceNetworkVpcAssociationWithContext(aws.Context, *vpclattice.GetServiceNetworkVpcAssociationInput, ...request.Option) (*vpclattice.GetServiceNetworkVpcAssociationOutput, error)
	GetServiceNetworkVpcAssociationRequest(*vpclattice.GetServiceNetworkVpcAssociationInput) (*request.Request, *vpclattice.GetServiceNetworkVpcAssociationOutput)

	GetTargetGroup(*vpclattice.GetTargetGroupInput) (*vpclattice.GetTargetGroupOutput, error)
	GetTargetGroupWithContext(aws.Context, *vpclattice.GetTargetGroupInput, ...request.Option) (*vpclattice.GetTargetGroupOutput, error)
	GetTargetGroupRequest(*vpclattice.GetTargetGroupInput) (*request.Request, *vpclattice.GetTargetGroupOutput)

	ListAccessLogSubscriptions(*vpclattice.ListAccessLogSubscriptionsInput) (*vpclattice.ListAccessLogSubscriptionsOutput, error)
	ListAccessLogSubscriptionsWithContext(aws.Context, *vpclattice.ListAccessLogSubscriptionsInput, ...request.Option) (*vpclattice.ListAccessLogSubscriptionsOutput, error)
	ListAccessLogSubscriptionsRequest(*vpclattice.ListAccessLogSubscriptionsInput) (*request.Request, *vpclattice.ListAccessLogSubscriptionsOutput)

	ListAccessLogSubscriptionsPages(*vpclattice.ListAccessLogSubscriptionsInput, func(*vpclattice.ListAccessLogSubscriptionsOutput, bool) bool) error
	ListAccessLogSubscriptionsPagesWithContext(aws.Context, *vpclattice.ListAccessLogSubscriptionsInput, func(*vpclattice.ListAccessLogSubscriptionsOutput, bool) bool, ...request.Option) error

	ListListeners(*vpclattice.ListListenersInput) (*vpclattice.ListListenersOutput, error)
	ListListenersWithContext(aws.Context, *vpclattice.ListListenersInput, ...request.Option) (*vpclattice.ListListenersOutput, error)
	ListListenersRequest(*vpclattice.ListListenersInput) (*request.Request, *vpclattice.ListListenersOutput)

	ListListenersPages(*vpclattice.ListListenersInput, func(*vpclattice.ListListenersOutput, bool) bool) error
	ListListenersPagesWithContext(aws.Context, *vpclattice.ListListenersInput, func(*vpclattice.ListListenersOutput, bool) bool, ...request.Option) error

	ListRules(*vpclattice.ListRulesInput) (*vpclattice.ListRulesOutput, error)
	ListRulesWithContext(aws.Context, *vpclattice.ListRulesInput, ...request.Option) (*vpclattice.ListRulesOutput, error)
	ListRulesRequest(*vpclattice.ListRulesInput) (*request.Request, *vpclattice.ListRulesOutput)

	ListRulesPages(*vpclattice.ListRulesInput, func(*vpclattice.ListRulesOutput, bool) bool) error
	ListRulesPagesWithContext(aws.Context, *vpclattice.ListRulesInput, func(*vpclattice.ListRulesOutput, bool) bool, ...request.Option) error

	ListServiceNetworkServiceAssociations(*vpclattice.ListServiceNetworkServiceAssociationsInput) (*vpclattice.ListServiceNetworkServiceAssociationsOutput, error)
	ListServiceNetworkServiceAssociationsWithContext(aws.Context, *vpclattice.ListServiceNetworkServiceAssociationsInput, ...request.Option) (*vpclattice.ListServiceNetworkServiceAssociationsOutput, error)
	ListServiceNetworkServiceAssociationsRequest(*vpclattice.ListServiceNetworkServiceAssociationsInput) (*request.Request, *vpclattice.ListServiceNetworkServiceAssociationsOutput)

	ListServiceNetworkServiceAssociationsPages(*vpclattice.ListServiceNetworkServiceAssociationsInput, func(*vpclattice.ListServiceNetworkServiceAssociationsOutput, bool) bool) error
	ListServiceNetworkServiceAssociationsPagesWithContext(aws.Context, *vpclattice.ListServiceNetworkServiceAssociationsInput, func(*vpclattice.ListServiceNetworkServiceAssociationsOutput, bool) bool, ...request.Option) error

	ListServiceNetworkVpcAssociations(*vpclattice.ListServiceNetworkVpcAssociationsInput) (*vpclattice.ListServiceNetworkVpcAssociationsOutput, error)
	ListServiceNetworkVpcAssociationsWithContext(aws.Context, *vpclattice.ListServiceNetworkVpcAssociationsInput, ...request.Option) (*vpclattice.ListServiceNetworkVpcAssociationsOutput, error)
	ListServiceNetworkVpcAssociationsRequest(*vpclattice.ListServiceNetworkVpcAssociationsInput) (*request.Request, *vpclattice.ListServiceNetworkVpcAssociationsOutput)

	ListServiceNetworkVpcAssociationsPages(*vpclattice.ListServiceNetworkVpcAssociationsInput, func(*vpclattice.ListServiceNetworkVpcAssociationsOutput, bool) bool) error
	ListServiceNetworkVpcAssociationsPagesWithContext(aws.Context, *vpclattice.ListServiceNetworkVpcAssociationsInput, func(*vpclattice.ListServiceNetworkVpcAssociationsOutput, bool) bool, ...request.Option) error

	ListServiceNetworks(*vpclattice.ListServiceNetworksInput) (*vpclattice.ListServiceNetworksOutput, error)
	ListServiceNetworksWithContext(aws.Context, *vpclattice.ListServiceNetworksInput, ...request.Option) (*vpclattice.ListServiceNetworksOutput, error)
	ListServiceNetworksRequest(*vpclattice.ListServiceNetworksInput) (*request.Request, *vpclattice.ListServiceNetworksOutput)

	ListServiceNetworksPages(*vpclattice.ListServiceNetworksInput, func(*vpclattice.ListServiceNetworksOutput, bool) bool) error
	ListServiceNetworksPagesWithContext(aws.Context, *vpclattice.ListServiceNetworksInput, func(*vpclattice.ListServiceNetworksOutput, bool) bool, ...request.Option) error

	ListServices(*vpclattice.ListServicesInput) (*vpclattice.ListServicesOutput, error)
	ListServicesWithContext(aws.Context, *vpclattice.ListServicesInput, ...request.Option) (*vpclattice.ListServicesOutput, error)
	ListServicesRequest(*vpclattice.ListServicesInput) (*request.Request, *vpclattice.ListServicesOutput)

	ListServicesPages(*vpclattice.ListServicesInput, func(*vpclattice.ListServicesOutput, bool) bool) error
	ListServicesPagesWithContext(aws.Context, *vpclattice.ListServicesInput, func(*vpclattice.ListServicesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*vpclattice.ListTagsForResourceInput) (*vpclattice.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *vpclattice.ListTagsForResourceInput, ...request.Option) (*vpclattice.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*vpclattice.ListTagsForResourceInput) (*request.Request, *vpclattice.ListTagsForResourceOutput)

	ListTargetGroups(*vpclattice.ListTargetGroupsInput) (*vpclattice.ListTargetGroupsOutput, error)
	ListTargetGroupsWithContext(aws.Context, *vpclattice.ListTargetGroupsInput, ...request.Option) (*vpclattice.ListTargetGroupsOutput, error)
	ListTargetGroupsRequest(*vpclattice.ListTargetGroupsInput) (*request.Request, *vpclattice.ListTargetGroupsOutput)

	ListTargetGroupsPages(*vpclattice.ListTargetGroupsInput, func(*vpclattice.ListTargetGroupsOutput, bool) bool) error
	ListTargetGroupsPagesWithContext(aws.Context, *vpclattice.ListTargetGroupsInput, func(*vpclattice.ListTargetGroupsOutput, bool) bool, ...request.Option) error

	ListTargets(*vpclattice.ListTargetsInput) (*vpclattice.ListTargetsOutput, error)
	ListTargetsWithContext(aws.Context, *vpclattice.ListTargetsInput, ...request.Option) (*vpclattice.ListTargetsOutput, error)
	ListTargetsRequest(*vpclattice.ListTargetsInput) (*request.Request, *vpclattice.ListTargetsOutput)

	ListTargetsPages(*vpclattice.ListTargetsInput, func(*vpclattice.ListTargetsOutput, bool) bool) error
	ListTargetsPagesWithContext(aws.Context, *vpclattice.ListTargetsInput, func(*vpclattice.ListTargetsOutput, bool) bool, ...request.Option) error

	PutAuthPolicy(*vpclattice.PutAuthPolicyInput) (*vpclattice.PutAuthPolicyOutput, error)
	PutAuthPolicyWithContext(aws.Context, *vpclattice.PutAuthPolicyInput, ...request.Option) (*vpclattice.PutAuthPolicyOutput, error)
	PutAuthPolicyRequest(*vpclattice.PutAuthPolicyInput) (*request.Request, *vpclattice.PutAuthPolicyOutput)

	PutResourcePolicy(*vpclattice.PutResourcePolicyInput) (*vpclattice.PutResourcePolicyOutput, error)
	PutResourcePolicyWithContext(aws.Context, *vpclattice.PutResourcePolicyInput, ...request.Option) (*vpclattice.PutResourcePolicyOutput, error)
	PutResourcePolicyRequest(*vpclattice.PutResourcePolicyInput) (*request.Request, *vpclattice.PutResourcePolicyOutput)

	RegisterTargets(*vpclattice.RegisterTargetsInput) (*vpclattice.RegisterTargetsOutput, error)
	RegisterTargetsWithContext(aws.Context, *vpclattice.RegisterTargetsInput, ...request.Option) (*vpclattice.RegisterTargetsOutput, error)
	RegisterTargetsRequest(*vpclattice.RegisterTargetsInput) (*request.Request, *vpclattice.RegisterTargetsOutput)

	TagResource(*vpclattice.TagResourceInput) (*vpclattice.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *vpclattice.TagResourceInput, ...request.Option) (*vpclattice.TagResourceOutput, error)
	TagResourceRequest(*vpclattice.TagResourceInput) (*request.Request, *vpclattice.TagResourceOutput)

	UntagResource(*vpclattice.UntagResourceInput) (*vpclattice.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *vpclattice.UntagResourceInput, ...request.Option) (*vpclattice.UntagResourceOutput, error)
	UntagResourceRequest(*vpclattice.UntagResourceInput) (*request.Request, *vpclattice.UntagResourceOutput)

	UpdateAccessLogSubscription(*vpclattice.UpdateAccessLogSubscriptionInput) (*vpclattice.UpdateAccessLogSubscriptionOutput, error)
	UpdateAccessLogSubscriptionWithContext(aws.Context, *vpclattice.UpdateAccessLogSubscriptionInput, ...request.Option) (*vpclattice.UpdateAccessLogSubscriptionOutput, error)
	UpdateAccessLogSubscriptionRequest(*vpclattice.UpdateAccessLogSubscriptionInput) (*request.Request, *vpclattice.UpdateAccessLogSubscriptionOutput)

	UpdateListener(*vpclattice.UpdateListenerInput) (*vpclattice.UpdateListenerOutput, error)
	UpdateListenerWithContext(aws.Context, *vpclattice.UpdateListenerInput, ...request.Option) (*vpclattice.UpdateListenerOutput, error)
	UpdateListenerRequest(*vpclattice.UpdateListenerInput) (*request.Request, *vpclattice.UpdateListenerOutput)

	UpdateRule(*vpclattice.UpdateRuleInput) (*vpclattice.UpdateRuleOutput, error)
	UpdateRuleWithContext(aws.Context, *vpclattice.UpdateRuleInput, ...request.Option) (*vpclattice.UpdateRuleOutput, error)
	UpdateRuleRequest(*vpclattice.UpdateRuleInput) (*request.Request, *vpclattice.UpdateRuleOutput)

	UpdateService(*vpclattice.UpdateServiceInput) (*vpclattice.UpdateServiceOutput, error)
	UpdateServiceWithContext(aws.Context, *vpclattice.UpdateServiceInput, ...request.Option) (*vpclattice.UpdateServiceOutput, error)
	UpdateServiceRequest(*vpclattice.UpdateServiceInput) (*request.Request, *vpclattice.UpdateServiceOutput)

	UpdateServiceNetwork(*vpclattice.UpdateServiceNetworkInput) (*vpclattice.UpdateServiceNetworkOutput, error)
	UpdateServiceNetworkWithContext(aws.Context, *vpclattice.UpdateServiceNetworkInput, ...request.Option) (*vpclattice.UpdateServiceNetworkOutput, error)
	UpdateServiceNetworkRequest(*vpclattice.UpdateServiceNetworkInput) (*request.Request, *vpclattice.UpdateServiceNetworkOutput)

	UpdateServiceNetworkVpcAssociation(*vpclattice.UpdateServiceNetworkVpcAssociationInput) (*vpclattice.UpdateServiceNetworkVpcAssociationOutput, error)
	UpdateServiceNetworkVpcAssociationWithContext(aws.Context, *vpclattice.UpdateServiceNetworkVpcAssociationInput, ...request.Option) (*vpclattice.UpdateServiceNetworkVpcAssociationOutput, error)
	UpdateServiceNetworkVpcAssociationRequest(*vpclattice.UpdateServiceNetworkVpcAssociationInput) (*request.Request, *vpclattice.UpdateServiceNetworkVpcAssociationOutput)

	UpdateTargetGroup(*vpclattice.UpdateTargetGroupInput) (*vpclattice.UpdateTargetGroupOutput, error)
	UpdateTargetGroupWithContext(aws.Context, *vpclattice.UpdateTargetGroupInput, ...request.Option) (*vpclattice.UpdateTargetGroupOutput, error)
	UpdateTargetGroupRequest(*vpclattice.UpdateTargetGroupInput) (*request.Request, *vpclattice.UpdateTargetGroupOutput)
}

var _ VPCLatticeAPI = (*vpclattice.VPCLattice)(nil)
