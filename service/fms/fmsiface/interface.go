// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package fmsiface provides an interface to enable mocking the Firewall Management Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package fmsiface

import (
	"github.com/qtdslly/aws-sdk-go/aws"
	"github.com/qtdslly/aws-sdk-go/aws/request"
	"github.com/qtdslly/aws-sdk-go/service/fms"
)

// FMSAPI provides an interface to enable mocking the
// fms.FMS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Firewall Management Service.
//	func myFunc(svc fmsiface.FMSAPI) bool {
//	    // Make svc.AssociateAdminAccount request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := fms.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockFMSClient struct {
//	    fmsiface.FMSAPI
//	}
//	func (m *mockFMSClient) AssociateAdminAccount(input *fms.AssociateAdminAccountInput) (*fms.AssociateAdminAccountOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockFMSClient{}
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
type FMSAPI interface {
	AssociateAdminAccount(*fms.AssociateAdminAccountInput) (*fms.AssociateAdminAccountOutput, error)
	AssociateAdminAccountWithContext(aws.Context, *fms.AssociateAdminAccountInput, ...request.Option) (*fms.AssociateAdminAccountOutput, error)
	AssociateAdminAccountRequest(*fms.AssociateAdminAccountInput) (*request.Request, *fms.AssociateAdminAccountOutput)

	AssociateThirdPartyFirewall(*fms.AssociateThirdPartyFirewallInput) (*fms.AssociateThirdPartyFirewallOutput, error)
	AssociateThirdPartyFirewallWithContext(aws.Context, *fms.AssociateThirdPartyFirewallInput, ...request.Option) (*fms.AssociateThirdPartyFirewallOutput, error)
	AssociateThirdPartyFirewallRequest(*fms.AssociateThirdPartyFirewallInput) (*request.Request, *fms.AssociateThirdPartyFirewallOutput)

	BatchAssociateResource(*fms.BatchAssociateResourceInput) (*fms.BatchAssociateResourceOutput, error)
	BatchAssociateResourceWithContext(aws.Context, *fms.BatchAssociateResourceInput, ...request.Option) (*fms.BatchAssociateResourceOutput, error)
	BatchAssociateResourceRequest(*fms.BatchAssociateResourceInput) (*request.Request, *fms.BatchAssociateResourceOutput)

	BatchDisassociateResource(*fms.BatchDisassociateResourceInput) (*fms.BatchDisassociateResourceOutput, error)
	BatchDisassociateResourceWithContext(aws.Context, *fms.BatchDisassociateResourceInput, ...request.Option) (*fms.BatchDisassociateResourceOutput, error)
	BatchDisassociateResourceRequest(*fms.BatchDisassociateResourceInput) (*request.Request, *fms.BatchDisassociateResourceOutput)

	DeleteAppsList(*fms.DeleteAppsListInput) (*fms.DeleteAppsListOutput, error)
	DeleteAppsListWithContext(aws.Context, *fms.DeleteAppsListInput, ...request.Option) (*fms.DeleteAppsListOutput, error)
	DeleteAppsListRequest(*fms.DeleteAppsListInput) (*request.Request, *fms.DeleteAppsListOutput)

	DeleteNotificationChannel(*fms.DeleteNotificationChannelInput) (*fms.DeleteNotificationChannelOutput, error)
	DeleteNotificationChannelWithContext(aws.Context, *fms.DeleteNotificationChannelInput, ...request.Option) (*fms.DeleteNotificationChannelOutput, error)
	DeleteNotificationChannelRequest(*fms.DeleteNotificationChannelInput) (*request.Request, *fms.DeleteNotificationChannelOutput)

	DeletePolicy(*fms.DeletePolicyInput) (*fms.DeletePolicyOutput, error)
	DeletePolicyWithContext(aws.Context, *fms.DeletePolicyInput, ...request.Option) (*fms.DeletePolicyOutput, error)
	DeletePolicyRequest(*fms.DeletePolicyInput) (*request.Request, *fms.DeletePolicyOutput)

	DeleteProtocolsList(*fms.DeleteProtocolsListInput) (*fms.DeleteProtocolsListOutput, error)
	DeleteProtocolsListWithContext(aws.Context, *fms.DeleteProtocolsListInput, ...request.Option) (*fms.DeleteProtocolsListOutput, error)
	DeleteProtocolsListRequest(*fms.DeleteProtocolsListInput) (*request.Request, *fms.DeleteProtocolsListOutput)

	DeleteResourceSet(*fms.DeleteResourceSetInput) (*fms.DeleteResourceSetOutput, error)
	DeleteResourceSetWithContext(aws.Context, *fms.DeleteResourceSetInput, ...request.Option) (*fms.DeleteResourceSetOutput, error)
	DeleteResourceSetRequest(*fms.DeleteResourceSetInput) (*request.Request, *fms.DeleteResourceSetOutput)

	DisassociateAdminAccount(*fms.DisassociateAdminAccountInput) (*fms.DisassociateAdminAccountOutput, error)
	DisassociateAdminAccountWithContext(aws.Context, *fms.DisassociateAdminAccountInput, ...request.Option) (*fms.DisassociateAdminAccountOutput, error)
	DisassociateAdminAccountRequest(*fms.DisassociateAdminAccountInput) (*request.Request, *fms.DisassociateAdminAccountOutput)

	DisassociateThirdPartyFirewall(*fms.DisassociateThirdPartyFirewallInput) (*fms.DisassociateThirdPartyFirewallOutput, error)
	DisassociateThirdPartyFirewallWithContext(aws.Context, *fms.DisassociateThirdPartyFirewallInput, ...request.Option) (*fms.DisassociateThirdPartyFirewallOutput, error)
	DisassociateThirdPartyFirewallRequest(*fms.DisassociateThirdPartyFirewallInput) (*request.Request, *fms.DisassociateThirdPartyFirewallOutput)

	GetAdminAccount(*fms.GetAdminAccountInput) (*fms.GetAdminAccountOutput, error)
	GetAdminAccountWithContext(aws.Context, *fms.GetAdminAccountInput, ...request.Option) (*fms.GetAdminAccountOutput, error)
	GetAdminAccountRequest(*fms.GetAdminAccountInput) (*request.Request, *fms.GetAdminAccountOutput)

	GetAdminScope(*fms.GetAdminScopeInput) (*fms.GetAdminScopeOutput, error)
	GetAdminScopeWithContext(aws.Context, *fms.GetAdminScopeInput, ...request.Option) (*fms.GetAdminScopeOutput, error)
	GetAdminScopeRequest(*fms.GetAdminScopeInput) (*request.Request, *fms.GetAdminScopeOutput)

	GetAppsList(*fms.GetAppsListInput) (*fms.GetAppsListOutput, error)
	GetAppsListWithContext(aws.Context, *fms.GetAppsListInput, ...request.Option) (*fms.GetAppsListOutput, error)
	GetAppsListRequest(*fms.GetAppsListInput) (*request.Request, *fms.GetAppsListOutput)

	GetComplianceDetail(*fms.GetComplianceDetailInput) (*fms.GetComplianceDetailOutput, error)
	GetComplianceDetailWithContext(aws.Context, *fms.GetComplianceDetailInput, ...request.Option) (*fms.GetComplianceDetailOutput, error)
	GetComplianceDetailRequest(*fms.GetComplianceDetailInput) (*request.Request, *fms.GetComplianceDetailOutput)

	GetNotificationChannel(*fms.GetNotificationChannelInput) (*fms.GetNotificationChannelOutput, error)
	GetNotificationChannelWithContext(aws.Context, *fms.GetNotificationChannelInput, ...request.Option) (*fms.GetNotificationChannelOutput, error)
	GetNotificationChannelRequest(*fms.GetNotificationChannelInput) (*request.Request, *fms.GetNotificationChannelOutput)

	GetPolicy(*fms.GetPolicyInput) (*fms.GetPolicyOutput, error)
	GetPolicyWithContext(aws.Context, *fms.GetPolicyInput, ...request.Option) (*fms.GetPolicyOutput, error)
	GetPolicyRequest(*fms.GetPolicyInput) (*request.Request, *fms.GetPolicyOutput)

	GetProtectionStatus(*fms.GetProtectionStatusInput) (*fms.GetProtectionStatusOutput, error)
	GetProtectionStatusWithContext(aws.Context, *fms.GetProtectionStatusInput, ...request.Option) (*fms.GetProtectionStatusOutput, error)
	GetProtectionStatusRequest(*fms.GetProtectionStatusInput) (*request.Request, *fms.GetProtectionStatusOutput)

	GetProtocolsList(*fms.GetProtocolsListInput) (*fms.GetProtocolsListOutput, error)
	GetProtocolsListWithContext(aws.Context, *fms.GetProtocolsListInput, ...request.Option) (*fms.GetProtocolsListOutput, error)
	GetProtocolsListRequest(*fms.GetProtocolsListInput) (*request.Request, *fms.GetProtocolsListOutput)

	GetResourceSet(*fms.GetResourceSetInput) (*fms.GetResourceSetOutput, error)
	GetResourceSetWithContext(aws.Context, *fms.GetResourceSetInput, ...request.Option) (*fms.GetResourceSetOutput, error)
	GetResourceSetRequest(*fms.GetResourceSetInput) (*request.Request, *fms.GetResourceSetOutput)

	GetThirdPartyFirewallAssociationStatus(*fms.GetThirdPartyFirewallAssociationStatusInput) (*fms.GetThirdPartyFirewallAssociationStatusOutput, error)
	GetThirdPartyFirewallAssociationStatusWithContext(aws.Context, *fms.GetThirdPartyFirewallAssociationStatusInput, ...request.Option) (*fms.GetThirdPartyFirewallAssociationStatusOutput, error)
	GetThirdPartyFirewallAssociationStatusRequest(*fms.GetThirdPartyFirewallAssociationStatusInput) (*request.Request, *fms.GetThirdPartyFirewallAssociationStatusOutput)

	GetViolationDetails(*fms.GetViolationDetailsInput) (*fms.GetViolationDetailsOutput, error)
	GetViolationDetailsWithContext(aws.Context, *fms.GetViolationDetailsInput, ...request.Option) (*fms.GetViolationDetailsOutput, error)
	GetViolationDetailsRequest(*fms.GetViolationDetailsInput) (*request.Request, *fms.GetViolationDetailsOutput)

	ListAdminAccountsForOrganization(*fms.ListAdminAccountsForOrganizationInput) (*fms.ListAdminAccountsForOrganizationOutput, error)
	ListAdminAccountsForOrganizationWithContext(aws.Context, *fms.ListAdminAccountsForOrganizationInput, ...request.Option) (*fms.ListAdminAccountsForOrganizationOutput, error)
	ListAdminAccountsForOrganizationRequest(*fms.ListAdminAccountsForOrganizationInput) (*request.Request, *fms.ListAdminAccountsForOrganizationOutput)

	ListAdminAccountsForOrganizationPages(*fms.ListAdminAccountsForOrganizationInput, func(*fms.ListAdminAccountsForOrganizationOutput, bool) bool) error
	ListAdminAccountsForOrganizationPagesWithContext(aws.Context, *fms.ListAdminAccountsForOrganizationInput, func(*fms.ListAdminAccountsForOrganizationOutput, bool) bool, ...request.Option) error

	ListAdminsManagingAccount(*fms.ListAdminsManagingAccountInput) (*fms.ListAdminsManagingAccountOutput, error)
	ListAdminsManagingAccountWithContext(aws.Context, *fms.ListAdminsManagingAccountInput, ...request.Option) (*fms.ListAdminsManagingAccountOutput, error)
	ListAdminsManagingAccountRequest(*fms.ListAdminsManagingAccountInput) (*request.Request, *fms.ListAdminsManagingAccountOutput)

	ListAdminsManagingAccountPages(*fms.ListAdminsManagingAccountInput, func(*fms.ListAdminsManagingAccountOutput, bool) bool) error
	ListAdminsManagingAccountPagesWithContext(aws.Context, *fms.ListAdminsManagingAccountInput, func(*fms.ListAdminsManagingAccountOutput, bool) bool, ...request.Option) error

	ListAppsLists(*fms.ListAppsListsInput) (*fms.ListAppsListsOutput, error)
	ListAppsListsWithContext(aws.Context, *fms.ListAppsListsInput, ...request.Option) (*fms.ListAppsListsOutput, error)
	ListAppsListsRequest(*fms.ListAppsListsInput) (*request.Request, *fms.ListAppsListsOutput)

	ListAppsListsPages(*fms.ListAppsListsInput, func(*fms.ListAppsListsOutput, bool) bool) error
	ListAppsListsPagesWithContext(aws.Context, *fms.ListAppsListsInput, func(*fms.ListAppsListsOutput, bool) bool, ...request.Option) error

	ListComplianceStatus(*fms.ListComplianceStatusInput) (*fms.ListComplianceStatusOutput, error)
	ListComplianceStatusWithContext(aws.Context, *fms.ListComplianceStatusInput, ...request.Option) (*fms.ListComplianceStatusOutput, error)
	ListComplianceStatusRequest(*fms.ListComplianceStatusInput) (*request.Request, *fms.ListComplianceStatusOutput)

	ListComplianceStatusPages(*fms.ListComplianceStatusInput, func(*fms.ListComplianceStatusOutput, bool) bool) error
	ListComplianceStatusPagesWithContext(aws.Context, *fms.ListComplianceStatusInput, func(*fms.ListComplianceStatusOutput, bool) bool, ...request.Option) error

	ListDiscoveredResources(*fms.ListDiscoveredResourcesInput) (*fms.ListDiscoveredResourcesOutput, error)
	ListDiscoveredResourcesWithContext(aws.Context, *fms.ListDiscoveredResourcesInput, ...request.Option) (*fms.ListDiscoveredResourcesOutput, error)
	ListDiscoveredResourcesRequest(*fms.ListDiscoveredResourcesInput) (*request.Request, *fms.ListDiscoveredResourcesOutput)

	ListMemberAccounts(*fms.ListMemberAccountsInput) (*fms.ListMemberAccountsOutput, error)
	ListMemberAccountsWithContext(aws.Context, *fms.ListMemberAccountsInput, ...request.Option) (*fms.ListMemberAccountsOutput, error)
	ListMemberAccountsRequest(*fms.ListMemberAccountsInput) (*request.Request, *fms.ListMemberAccountsOutput)

	ListMemberAccountsPages(*fms.ListMemberAccountsInput, func(*fms.ListMemberAccountsOutput, bool) bool) error
	ListMemberAccountsPagesWithContext(aws.Context, *fms.ListMemberAccountsInput, func(*fms.ListMemberAccountsOutput, bool) bool, ...request.Option) error

	ListPolicies(*fms.ListPoliciesInput) (*fms.ListPoliciesOutput, error)
	ListPoliciesWithContext(aws.Context, *fms.ListPoliciesInput, ...request.Option) (*fms.ListPoliciesOutput, error)
	ListPoliciesRequest(*fms.ListPoliciesInput) (*request.Request, *fms.ListPoliciesOutput)

	ListPoliciesPages(*fms.ListPoliciesInput, func(*fms.ListPoliciesOutput, bool) bool) error
	ListPoliciesPagesWithContext(aws.Context, *fms.ListPoliciesInput, func(*fms.ListPoliciesOutput, bool) bool, ...request.Option) error

	ListProtocolsLists(*fms.ListProtocolsListsInput) (*fms.ListProtocolsListsOutput, error)
	ListProtocolsListsWithContext(aws.Context, *fms.ListProtocolsListsInput, ...request.Option) (*fms.ListProtocolsListsOutput, error)
	ListProtocolsListsRequest(*fms.ListProtocolsListsInput) (*request.Request, *fms.ListProtocolsListsOutput)

	ListProtocolsListsPages(*fms.ListProtocolsListsInput, func(*fms.ListProtocolsListsOutput, bool) bool) error
	ListProtocolsListsPagesWithContext(aws.Context, *fms.ListProtocolsListsInput, func(*fms.ListProtocolsListsOutput, bool) bool, ...request.Option) error

	ListResourceSetResources(*fms.ListResourceSetResourcesInput) (*fms.ListResourceSetResourcesOutput, error)
	ListResourceSetResourcesWithContext(aws.Context, *fms.ListResourceSetResourcesInput, ...request.Option) (*fms.ListResourceSetResourcesOutput, error)
	ListResourceSetResourcesRequest(*fms.ListResourceSetResourcesInput) (*request.Request, *fms.ListResourceSetResourcesOutput)

	ListResourceSets(*fms.ListResourceSetsInput) (*fms.ListResourceSetsOutput, error)
	ListResourceSetsWithContext(aws.Context, *fms.ListResourceSetsInput, ...request.Option) (*fms.ListResourceSetsOutput, error)
	ListResourceSetsRequest(*fms.ListResourceSetsInput) (*request.Request, *fms.ListResourceSetsOutput)

	ListTagsForResource(*fms.ListTagsForResourceInput) (*fms.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *fms.ListTagsForResourceInput, ...request.Option) (*fms.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*fms.ListTagsForResourceInput) (*request.Request, *fms.ListTagsForResourceOutput)

	ListThirdPartyFirewallFirewallPolicies(*fms.ListThirdPartyFirewallFirewallPoliciesInput) (*fms.ListThirdPartyFirewallFirewallPoliciesOutput, error)
	ListThirdPartyFirewallFirewallPoliciesWithContext(aws.Context, *fms.ListThirdPartyFirewallFirewallPoliciesInput, ...request.Option) (*fms.ListThirdPartyFirewallFirewallPoliciesOutput, error)
	ListThirdPartyFirewallFirewallPoliciesRequest(*fms.ListThirdPartyFirewallFirewallPoliciesInput) (*request.Request, *fms.ListThirdPartyFirewallFirewallPoliciesOutput)

	ListThirdPartyFirewallFirewallPoliciesPages(*fms.ListThirdPartyFirewallFirewallPoliciesInput, func(*fms.ListThirdPartyFirewallFirewallPoliciesOutput, bool) bool) error
	ListThirdPartyFirewallFirewallPoliciesPagesWithContext(aws.Context, *fms.ListThirdPartyFirewallFirewallPoliciesInput, func(*fms.ListThirdPartyFirewallFirewallPoliciesOutput, bool) bool, ...request.Option) error

	PutAdminAccount(*fms.PutAdminAccountInput) (*fms.PutAdminAccountOutput, error)
	PutAdminAccountWithContext(aws.Context, *fms.PutAdminAccountInput, ...request.Option) (*fms.PutAdminAccountOutput, error)
	PutAdminAccountRequest(*fms.PutAdminAccountInput) (*request.Request, *fms.PutAdminAccountOutput)

	PutAppsList(*fms.PutAppsListInput) (*fms.PutAppsListOutput, error)
	PutAppsListWithContext(aws.Context, *fms.PutAppsListInput, ...request.Option) (*fms.PutAppsListOutput, error)
	PutAppsListRequest(*fms.PutAppsListInput) (*request.Request, *fms.PutAppsListOutput)

	PutNotificationChannel(*fms.PutNotificationChannelInput) (*fms.PutNotificationChannelOutput, error)
	PutNotificationChannelWithContext(aws.Context, *fms.PutNotificationChannelInput, ...request.Option) (*fms.PutNotificationChannelOutput, error)
	PutNotificationChannelRequest(*fms.PutNotificationChannelInput) (*request.Request, *fms.PutNotificationChannelOutput)

	PutPolicy(*fms.PutPolicyInput) (*fms.PutPolicyOutput, error)
	PutPolicyWithContext(aws.Context, *fms.PutPolicyInput, ...request.Option) (*fms.PutPolicyOutput, error)
	PutPolicyRequest(*fms.PutPolicyInput) (*request.Request, *fms.PutPolicyOutput)

	PutProtocolsList(*fms.PutProtocolsListInput) (*fms.PutProtocolsListOutput, error)
	PutProtocolsListWithContext(aws.Context, *fms.PutProtocolsListInput, ...request.Option) (*fms.PutProtocolsListOutput, error)
	PutProtocolsListRequest(*fms.PutProtocolsListInput) (*request.Request, *fms.PutProtocolsListOutput)

	PutResourceSet(*fms.PutResourceSetInput) (*fms.PutResourceSetOutput, error)
	PutResourceSetWithContext(aws.Context, *fms.PutResourceSetInput, ...request.Option) (*fms.PutResourceSetOutput, error)
	PutResourceSetRequest(*fms.PutResourceSetInput) (*request.Request, *fms.PutResourceSetOutput)

	TagResource(*fms.TagResourceInput) (*fms.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *fms.TagResourceInput, ...request.Option) (*fms.TagResourceOutput, error)
	TagResourceRequest(*fms.TagResourceInput) (*request.Request, *fms.TagResourceOutput)

	UntagResource(*fms.UntagResourceInput) (*fms.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *fms.UntagResourceInput, ...request.Option) (*fms.UntagResourceOutput, error)
	UntagResourceRequest(*fms.UntagResourceInput) (*request.Request, *fms.UntagResourceOutput)
}

var _ FMSAPI = (*fms.FMS)(nil)