// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudformationiface provides an interface to enable mocking the AWS CloudFormation service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudformationiface

import (
	"github.com/qtdslly/aws-sdk-go/aws"
	"github.com/qtdslly/aws-sdk-go/aws/request"
	"github.com/qtdslly/aws-sdk-go/service/cloudformation"
)

// CloudFormationAPI provides an interface to enable mocking the
// cloudformation.CloudFormation service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS CloudFormation.
//	func myFunc(svc cloudformationiface.CloudFormationAPI) bool {
//	    // Make svc.ActivateOrganizationsAccess request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := cloudformation.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockCloudFormationClient struct {
//	    cloudformationiface.CloudFormationAPI
//	}
//	func (m *mockCloudFormationClient) ActivateOrganizationsAccess(input *cloudformation.ActivateOrganizationsAccessInput) (*cloudformation.ActivateOrganizationsAccessOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockCloudFormationClient{}
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
type CloudFormationAPI interface {
	ActivateOrganizationsAccess(*cloudformation.ActivateOrganizationsAccessInput) (*cloudformation.ActivateOrganizationsAccessOutput, error)
	ActivateOrganizationsAccessWithContext(aws.Context, *cloudformation.ActivateOrganizationsAccessInput, ...request.Option) (*cloudformation.ActivateOrganizationsAccessOutput, error)
	ActivateOrganizationsAccessRequest(*cloudformation.ActivateOrganizationsAccessInput) (*request.Request, *cloudformation.ActivateOrganizationsAccessOutput)

	ActivateType(*cloudformation.ActivateTypeInput) (*cloudformation.ActivateTypeOutput, error)
	ActivateTypeWithContext(aws.Context, *cloudformation.ActivateTypeInput, ...request.Option) (*cloudformation.ActivateTypeOutput, error)
	ActivateTypeRequest(*cloudformation.ActivateTypeInput) (*request.Request, *cloudformation.ActivateTypeOutput)

	BatchDescribeTypeConfigurations(*cloudformation.BatchDescribeTypeConfigurationsInput) (*cloudformation.BatchDescribeTypeConfigurationsOutput, error)
	BatchDescribeTypeConfigurationsWithContext(aws.Context, *cloudformation.BatchDescribeTypeConfigurationsInput, ...request.Option) (*cloudformation.BatchDescribeTypeConfigurationsOutput, error)
	BatchDescribeTypeConfigurationsRequest(*cloudformation.BatchDescribeTypeConfigurationsInput) (*request.Request, *cloudformation.BatchDescribeTypeConfigurationsOutput)

	CancelUpdateStack(*cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error)
	CancelUpdateStackWithContext(aws.Context, *cloudformation.CancelUpdateStackInput, ...request.Option) (*cloudformation.CancelUpdateStackOutput, error)
	CancelUpdateStackRequest(*cloudformation.CancelUpdateStackInput) (*request.Request, *cloudformation.CancelUpdateStackOutput)

	ContinueUpdateRollback(*cloudformation.ContinueUpdateRollbackInput) (*cloudformation.ContinueUpdateRollbackOutput, error)
	ContinueUpdateRollbackWithContext(aws.Context, *cloudformation.ContinueUpdateRollbackInput, ...request.Option) (*cloudformation.ContinueUpdateRollbackOutput, error)
	ContinueUpdateRollbackRequest(*cloudformation.ContinueUpdateRollbackInput) (*request.Request, *cloudformation.ContinueUpdateRollbackOutput)

	CreateChangeSet(*cloudformation.CreateChangeSetInput) (*cloudformation.CreateChangeSetOutput, error)
	CreateChangeSetWithContext(aws.Context, *cloudformation.CreateChangeSetInput, ...request.Option) (*cloudformation.CreateChangeSetOutput, error)
	CreateChangeSetRequest(*cloudformation.CreateChangeSetInput) (*request.Request, *cloudformation.CreateChangeSetOutput)

	CreateStack(*cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error)
	CreateStackWithContext(aws.Context, *cloudformation.CreateStackInput, ...request.Option) (*cloudformation.CreateStackOutput, error)
	CreateStackRequest(*cloudformation.CreateStackInput) (*request.Request, *cloudformation.CreateStackOutput)

	CreateStackInstances(*cloudformation.CreateStackInstancesInput) (*cloudformation.CreateStackInstancesOutput, error)
	CreateStackInstancesWithContext(aws.Context, *cloudformation.CreateStackInstancesInput, ...request.Option) (*cloudformation.CreateStackInstancesOutput, error)
	CreateStackInstancesRequest(*cloudformation.CreateStackInstancesInput) (*request.Request, *cloudformation.CreateStackInstancesOutput)

	CreateStackSet(*cloudformation.CreateStackSetInput) (*cloudformation.CreateStackSetOutput, error)
	CreateStackSetWithContext(aws.Context, *cloudformation.CreateStackSetInput, ...request.Option) (*cloudformation.CreateStackSetOutput, error)
	CreateStackSetRequest(*cloudformation.CreateStackSetInput) (*request.Request, *cloudformation.CreateStackSetOutput)

	DeactivateOrganizationsAccess(*cloudformation.DeactivateOrganizationsAccessInput) (*cloudformation.DeactivateOrganizationsAccessOutput, error)
	DeactivateOrganizationsAccessWithContext(aws.Context, *cloudformation.DeactivateOrganizationsAccessInput, ...request.Option) (*cloudformation.DeactivateOrganizationsAccessOutput, error)
	DeactivateOrganizationsAccessRequest(*cloudformation.DeactivateOrganizationsAccessInput) (*request.Request, *cloudformation.DeactivateOrganizationsAccessOutput)

	DeactivateType(*cloudformation.DeactivateTypeInput) (*cloudformation.DeactivateTypeOutput, error)
	DeactivateTypeWithContext(aws.Context, *cloudformation.DeactivateTypeInput, ...request.Option) (*cloudformation.DeactivateTypeOutput, error)
	DeactivateTypeRequest(*cloudformation.DeactivateTypeInput) (*request.Request, *cloudformation.DeactivateTypeOutput)

	DeleteChangeSet(*cloudformation.DeleteChangeSetInput) (*cloudformation.DeleteChangeSetOutput, error)
	DeleteChangeSetWithContext(aws.Context, *cloudformation.DeleteChangeSetInput, ...request.Option) (*cloudformation.DeleteChangeSetOutput, error)
	DeleteChangeSetRequest(*cloudformation.DeleteChangeSetInput) (*request.Request, *cloudformation.DeleteChangeSetOutput)

	DeleteStack(*cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error)
	DeleteStackWithContext(aws.Context, *cloudformation.DeleteStackInput, ...request.Option) (*cloudformation.DeleteStackOutput, error)
	DeleteStackRequest(*cloudformation.DeleteStackInput) (*request.Request, *cloudformation.DeleteStackOutput)

	DeleteStackInstances(*cloudformation.DeleteStackInstancesInput) (*cloudformation.DeleteStackInstancesOutput, error)
	DeleteStackInstancesWithContext(aws.Context, *cloudformation.DeleteStackInstancesInput, ...request.Option) (*cloudformation.DeleteStackInstancesOutput, error)
	DeleteStackInstancesRequest(*cloudformation.DeleteStackInstancesInput) (*request.Request, *cloudformation.DeleteStackInstancesOutput)

	DeleteStackSet(*cloudformation.DeleteStackSetInput) (*cloudformation.DeleteStackSetOutput, error)
	DeleteStackSetWithContext(aws.Context, *cloudformation.DeleteStackSetInput, ...request.Option) (*cloudformation.DeleteStackSetOutput, error)
	DeleteStackSetRequest(*cloudformation.DeleteStackSetInput) (*request.Request, *cloudformation.DeleteStackSetOutput)

	DeregisterType(*cloudformation.DeregisterTypeInput) (*cloudformation.DeregisterTypeOutput, error)
	DeregisterTypeWithContext(aws.Context, *cloudformation.DeregisterTypeInput, ...request.Option) (*cloudformation.DeregisterTypeOutput, error)
	DeregisterTypeRequest(*cloudformation.DeregisterTypeInput) (*request.Request, *cloudformation.DeregisterTypeOutput)

	DescribeAccountLimits(*cloudformation.DescribeAccountLimitsInput) (*cloudformation.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsWithContext(aws.Context, *cloudformation.DescribeAccountLimitsInput, ...request.Option) (*cloudformation.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsRequest(*cloudformation.DescribeAccountLimitsInput) (*request.Request, *cloudformation.DescribeAccountLimitsOutput)

	DescribeAccountLimitsPages(*cloudformation.DescribeAccountLimitsInput, func(*cloudformation.DescribeAccountLimitsOutput, bool) bool) error
	DescribeAccountLimitsPagesWithContext(aws.Context, *cloudformation.DescribeAccountLimitsInput, func(*cloudformation.DescribeAccountLimitsOutput, bool) bool, ...request.Option) error

	DescribeChangeSet(*cloudformation.DescribeChangeSetInput) (*cloudformation.DescribeChangeSetOutput, error)
	DescribeChangeSetWithContext(aws.Context, *cloudformation.DescribeChangeSetInput, ...request.Option) (*cloudformation.DescribeChangeSetOutput, error)
	DescribeChangeSetRequest(*cloudformation.DescribeChangeSetInput) (*request.Request, *cloudformation.DescribeChangeSetOutput)

	DescribeChangeSetHooks(*cloudformation.DescribeChangeSetHooksInput) (*cloudformation.DescribeChangeSetHooksOutput, error)
	DescribeChangeSetHooksWithContext(aws.Context, *cloudformation.DescribeChangeSetHooksInput, ...request.Option) (*cloudformation.DescribeChangeSetHooksOutput, error)
	DescribeChangeSetHooksRequest(*cloudformation.DescribeChangeSetHooksInput) (*request.Request, *cloudformation.DescribeChangeSetHooksOutput)

	DescribeOrganizationsAccess(*cloudformation.DescribeOrganizationsAccessInput) (*cloudformation.DescribeOrganizationsAccessOutput, error)
	DescribeOrganizationsAccessWithContext(aws.Context, *cloudformation.DescribeOrganizationsAccessInput, ...request.Option) (*cloudformation.DescribeOrganizationsAccessOutput, error)
	DescribeOrganizationsAccessRequest(*cloudformation.DescribeOrganizationsAccessInput) (*request.Request, *cloudformation.DescribeOrganizationsAccessOutput)

	DescribePublisher(*cloudformation.DescribePublisherInput) (*cloudformation.DescribePublisherOutput, error)
	DescribePublisherWithContext(aws.Context, *cloudformation.DescribePublisherInput, ...request.Option) (*cloudformation.DescribePublisherOutput, error)
	DescribePublisherRequest(*cloudformation.DescribePublisherInput) (*request.Request, *cloudformation.DescribePublisherOutput)

	DescribeStackDriftDetectionStatus(*cloudformation.DescribeStackDriftDetectionStatusInput) (*cloudformation.DescribeStackDriftDetectionStatusOutput, error)
	DescribeStackDriftDetectionStatusWithContext(aws.Context, *cloudformation.DescribeStackDriftDetectionStatusInput, ...request.Option) (*cloudformation.DescribeStackDriftDetectionStatusOutput, error)
	DescribeStackDriftDetectionStatusRequest(*cloudformation.DescribeStackDriftDetectionStatusInput) (*request.Request, *cloudformation.DescribeStackDriftDetectionStatusOutput)

	DescribeStackEvents(*cloudformation.DescribeStackEventsInput) (*cloudformation.DescribeStackEventsOutput, error)
	DescribeStackEventsWithContext(aws.Context, *cloudformation.DescribeStackEventsInput, ...request.Option) (*cloudformation.DescribeStackEventsOutput, error)
	DescribeStackEventsRequest(*cloudformation.DescribeStackEventsInput) (*request.Request, *cloudformation.DescribeStackEventsOutput)

	DescribeStackEventsPages(*cloudformation.DescribeStackEventsInput, func(*cloudformation.DescribeStackEventsOutput, bool) bool) error
	DescribeStackEventsPagesWithContext(aws.Context, *cloudformation.DescribeStackEventsInput, func(*cloudformation.DescribeStackEventsOutput, bool) bool, ...request.Option) error

	DescribeStackInstance(*cloudformation.DescribeStackInstanceInput) (*cloudformation.DescribeStackInstanceOutput, error)
	DescribeStackInstanceWithContext(aws.Context, *cloudformation.DescribeStackInstanceInput, ...request.Option) (*cloudformation.DescribeStackInstanceOutput, error)
	DescribeStackInstanceRequest(*cloudformation.DescribeStackInstanceInput) (*request.Request, *cloudformation.DescribeStackInstanceOutput)

	DescribeStackResource(*cloudformation.DescribeStackResourceInput) (*cloudformation.DescribeStackResourceOutput, error)
	DescribeStackResourceWithContext(aws.Context, *cloudformation.DescribeStackResourceInput, ...request.Option) (*cloudformation.DescribeStackResourceOutput, error)
	DescribeStackResourceRequest(*cloudformation.DescribeStackResourceInput) (*request.Request, *cloudformation.DescribeStackResourceOutput)

	DescribeStackResourceDrifts(*cloudformation.DescribeStackResourceDriftsInput) (*cloudformation.DescribeStackResourceDriftsOutput, error)
	DescribeStackResourceDriftsWithContext(aws.Context, *cloudformation.DescribeStackResourceDriftsInput, ...request.Option) (*cloudformation.DescribeStackResourceDriftsOutput, error)
	DescribeStackResourceDriftsRequest(*cloudformation.DescribeStackResourceDriftsInput) (*request.Request, *cloudformation.DescribeStackResourceDriftsOutput)

	DescribeStackResourceDriftsPages(*cloudformation.DescribeStackResourceDriftsInput, func(*cloudformation.DescribeStackResourceDriftsOutput, bool) bool) error
	DescribeStackResourceDriftsPagesWithContext(aws.Context, *cloudformation.DescribeStackResourceDriftsInput, func(*cloudformation.DescribeStackResourceDriftsOutput, bool) bool, ...request.Option) error

	DescribeStackResources(*cloudformation.DescribeStackResourcesInput) (*cloudformation.DescribeStackResourcesOutput, error)
	DescribeStackResourcesWithContext(aws.Context, *cloudformation.DescribeStackResourcesInput, ...request.Option) (*cloudformation.DescribeStackResourcesOutput, error)
	DescribeStackResourcesRequest(*cloudformation.DescribeStackResourcesInput) (*request.Request, *cloudformation.DescribeStackResourcesOutput)

	DescribeStackSet(*cloudformation.DescribeStackSetInput) (*cloudformation.DescribeStackSetOutput, error)
	DescribeStackSetWithContext(aws.Context, *cloudformation.DescribeStackSetInput, ...request.Option) (*cloudformation.DescribeStackSetOutput, error)
	DescribeStackSetRequest(*cloudformation.DescribeStackSetInput) (*request.Request, *cloudformation.DescribeStackSetOutput)

	DescribeStackSetOperation(*cloudformation.DescribeStackSetOperationInput) (*cloudformation.DescribeStackSetOperationOutput, error)
	DescribeStackSetOperationWithContext(aws.Context, *cloudformation.DescribeStackSetOperationInput, ...request.Option) (*cloudformation.DescribeStackSetOperationOutput, error)
	DescribeStackSetOperationRequest(*cloudformation.DescribeStackSetOperationInput) (*request.Request, *cloudformation.DescribeStackSetOperationOutput)

	DescribeStacks(*cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error)
	DescribeStacksWithContext(aws.Context, *cloudformation.DescribeStacksInput, ...request.Option) (*cloudformation.DescribeStacksOutput, error)
	DescribeStacksRequest(*cloudformation.DescribeStacksInput) (*request.Request, *cloudformation.DescribeStacksOutput)

	DescribeStacksPages(*cloudformation.DescribeStacksInput, func(*cloudformation.DescribeStacksOutput, bool) bool) error
	DescribeStacksPagesWithContext(aws.Context, *cloudformation.DescribeStacksInput, func(*cloudformation.DescribeStacksOutput, bool) bool, ...request.Option) error

	DescribeType(*cloudformation.DescribeTypeInput) (*cloudformation.DescribeTypeOutput, error)
	DescribeTypeWithContext(aws.Context, *cloudformation.DescribeTypeInput, ...request.Option) (*cloudformation.DescribeTypeOutput, error)
	DescribeTypeRequest(*cloudformation.DescribeTypeInput) (*request.Request, *cloudformation.DescribeTypeOutput)

	DescribeTypeRegistration(*cloudformation.DescribeTypeRegistrationInput) (*cloudformation.DescribeTypeRegistrationOutput, error)
	DescribeTypeRegistrationWithContext(aws.Context, *cloudformation.DescribeTypeRegistrationInput, ...request.Option) (*cloudformation.DescribeTypeRegistrationOutput, error)
	DescribeTypeRegistrationRequest(*cloudformation.DescribeTypeRegistrationInput) (*request.Request, *cloudformation.DescribeTypeRegistrationOutput)

	DetectStackDrift(*cloudformation.DetectStackDriftInput) (*cloudformation.DetectStackDriftOutput, error)
	DetectStackDriftWithContext(aws.Context, *cloudformation.DetectStackDriftInput, ...request.Option) (*cloudformation.DetectStackDriftOutput, error)
	DetectStackDriftRequest(*cloudformation.DetectStackDriftInput) (*request.Request, *cloudformation.DetectStackDriftOutput)

	DetectStackResourceDrift(*cloudformation.DetectStackResourceDriftInput) (*cloudformation.DetectStackResourceDriftOutput, error)
	DetectStackResourceDriftWithContext(aws.Context, *cloudformation.DetectStackResourceDriftInput, ...request.Option) (*cloudformation.DetectStackResourceDriftOutput, error)
	DetectStackResourceDriftRequest(*cloudformation.DetectStackResourceDriftInput) (*request.Request, *cloudformation.DetectStackResourceDriftOutput)

	DetectStackSetDrift(*cloudformation.DetectStackSetDriftInput) (*cloudformation.DetectStackSetDriftOutput, error)
	DetectStackSetDriftWithContext(aws.Context, *cloudformation.DetectStackSetDriftInput, ...request.Option) (*cloudformation.DetectStackSetDriftOutput, error)
	DetectStackSetDriftRequest(*cloudformation.DetectStackSetDriftInput) (*request.Request, *cloudformation.DetectStackSetDriftOutput)

	EstimateTemplateCost(*cloudformation.EstimateTemplateCostInput) (*cloudformation.EstimateTemplateCostOutput, error)
	EstimateTemplateCostWithContext(aws.Context, *cloudformation.EstimateTemplateCostInput, ...request.Option) (*cloudformation.EstimateTemplateCostOutput, error)
	EstimateTemplateCostRequest(*cloudformation.EstimateTemplateCostInput) (*request.Request, *cloudformation.EstimateTemplateCostOutput)

	ExecuteChangeSet(*cloudformation.ExecuteChangeSetInput) (*cloudformation.ExecuteChangeSetOutput, error)
	ExecuteChangeSetWithContext(aws.Context, *cloudformation.ExecuteChangeSetInput, ...request.Option) (*cloudformation.ExecuteChangeSetOutput, error)
	ExecuteChangeSetRequest(*cloudformation.ExecuteChangeSetInput) (*request.Request, *cloudformation.ExecuteChangeSetOutput)

	GetStackPolicy(*cloudformation.GetStackPolicyInput) (*cloudformation.GetStackPolicyOutput, error)
	GetStackPolicyWithContext(aws.Context, *cloudformation.GetStackPolicyInput, ...request.Option) (*cloudformation.GetStackPolicyOutput, error)
	GetStackPolicyRequest(*cloudformation.GetStackPolicyInput) (*request.Request, *cloudformation.GetStackPolicyOutput)

	GetTemplate(*cloudformation.GetTemplateInput) (*cloudformation.GetTemplateOutput, error)
	GetTemplateWithContext(aws.Context, *cloudformation.GetTemplateInput, ...request.Option) (*cloudformation.GetTemplateOutput, error)
	GetTemplateRequest(*cloudformation.GetTemplateInput) (*request.Request, *cloudformation.GetTemplateOutput)

	GetTemplateSummary(*cloudformation.GetTemplateSummaryInput) (*cloudformation.GetTemplateSummaryOutput, error)
	GetTemplateSummaryWithContext(aws.Context, *cloudformation.GetTemplateSummaryInput, ...request.Option) (*cloudformation.GetTemplateSummaryOutput, error)
	GetTemplateSummaryRequest(*cloudformation.GetTemplateSummaryInput) (*request.Request, *cloudformation.GetTemplateSummaryOutput)

	ImportStacksToStackSet(*cloudformation.ImportStacksToStackSetInput) (*cloudformation.ImportStacksToStackSetOutput, error)
	ImportStacksToStackSetWithContext(aws.Context, *cloudformation.ImportStacksToStackSetInput, ...request.Option) (*cloudformation.ImportStacksToStackSetOutput, error)
	ImportStacksToStackSetRequest(*cloudformation.ImportStacksToStackSetInput) (*request.Request, *cloudformation.ImportStacksToStackSetOutput)

	ListChangeSets(*cloudformation.ListChangeSetsInput) (*cloudformation.ListChangeSetsOutput, error)
	ListChangeSetsWithContext(aws.Context, *cloudformation.ListChangeSetsInput, ...request.Option) (*cloudformation.ListChangeSetsOutput, error)
	ListChangeSetsRequest(*cloudformation.ListChangeSetsInput) (*request.Request, *cloudformation.ListChangeSetsOutput)

	ListChangeSetsPages(*cloudformation.ListChangeSetsInput, func(*cloudformation.ListChangeSetsOutput, bool) bool) error
	ListChangeSetsPagesWithContext(aws.Context, *cloudformation.ListChangeSetsInput, func(*cloudformation.ListChangeSetsOutput, bool) bool, ...request.Option) error

	ListExports(*cloudformation.ListExportsInput) (*cloudformation.ListExportsOutput, error)
	ListExportsWithContext(aws.Context, *cloudformation.ListExportsInput, ...request.Option) (*cloudformation.ListExportsOutput, error)
	ListExportsRequest(*cloudformation.ListExportsInput) (*request.Request, *cloudformation.ListExportsOutput)

	ListExportsPages(*cloudformation.ListExportsInput, func(*cloudformation.ListExportsOutput, bool) bool) error
	ListExportsPagesWithContext(aws.Context, *cloudformation.ListExportsInput, func(*cloudformation.ListExportsOutput, bool) bool, ...request.Option) error

	ListImports(*cloudformation.ListImportsInput) (*cloudformation.ListImportsOutput, error)
	ListImportsWithContext(aws.Context, *cloudformation.ListImportsInput, ...request.Option) (*cloudformation.ListImportsOutput, error)
	ListImportsRequest(*cloudformation.ListImportsInput) (*request.Request, *cloudformation.ListImportsOutput)

	ListImportsPages(*cloudformation.ListImportsInput, func(*cloudformation.ListImportsOutput, bool) bool) error
	ListImportsPagesWithContext(aws.Context, *cloudformation.ListImportsInput, func(*cloudformation.ListImportsOutput, bool) bool, ...request.Option) error

	ListStackInstances(*cloudformation.ListStackInstancesInput) (*cloudformation.ListStackInstancesOutput, error)
	ListStackInstancesWithContext(aws.Context, *cloudformation.ListStackInstancesInput, ...request.Option) (*cloudformation.ListStackInstancesOutput, error)
	ListStackInstancesRequest(*cloudformation.ListStackInstancesInput) (*request.Request, *cloudformation.ListStackInstancesOutput)

	ListStackInstancesPages(*cloudformation.ListStackInstancesInput, func(*cloudformation.ListStackInstancesOutput, bool) bool) error
	ListStackInstancesPagesWithContext(aws.Context, *cloudformation.ListStackInstancesInput, func(*cloudformation.ListStackInstancesOutput, bool) bool, ...request.Option) error

	ListStackResources(*cloudformation.ListStackResourcesInput) (*cloudformation.ListStackResourcesOutput, error)
	ListStackResourcesWithContext(aws.Context, *cloudformation.ListStackResourcesInput, ...request.Option) (*cloudformation.ListStackResourcesOutput, error)
	ListStackResourcesRequest(*cloudformation.ListStackResourcesInput) (*request.Request, *cloudformation.ListStackResourcesOutput)

	ListStackResourcesPages(*cloudformation.ListStackResourcesInput, func(*cloudformation.ListStackResourcesOutput, bool) bool) error
	ListStackResourcesPagesWithContext(aws.Context, *cloudformation.ListStackResourcesInput, func(*cloudformation.ListStackResourcesOutput, bool) bool, ...request.Option) error

	ListStackSetOperationResults(*cloudformation.ListStackSetOperationResultsInput) (*cloudformation.ListStackSetOperationResultsOutput, error)
	ListStackSetOperationResultsWithContext(aws.Context, *cloudformation.ListStackSetOperationResultsInput, ...request.Option) (*cloudformation.ListStackSetOperationResultsOutput, error)
	ListStackSetOperationResultsRequest(*cloudformation.ListStackSetOperationResultsInput) (*request.Request, *cloudformation.ListStackSetOperationResultsOutput)

	ListStackSetOperationResultsPages(*cloudformation.ListStackSetOperationResultsInput, func(*cloudformation.ListStackSetOperationResultsOutput, bool) bool) error
	ListStackSetOperationResultsPagesWithContext(aws.Context, *cloudformation.ListStackSetOperationResultsInput, func(*cloudformation.ListStackSetOperationResultsOutput, bool) bool, ...request.Option) error

	ListStackSetOperations(*cloudformation.ListStackSetOperationsInput) (*cloudformation.ListStackSetOperationsOutput, error)
	ListStackSetOperationsWithContext(aws.Context, *cloudformation.ListStackSetOperationsInput, ...request.Option) (*cloudformation.ListStackSetOperationsOutput, error)
	ListStackSetOperationsRequest(*cloudformation.ListStackSetOperationsInput) (*request.Request, *cloudformation.ListStackSetOperationsOutput)

	ListStackSetOperationsPages(*cloudformation.ListStackSetOperationsInput, func(*cloudformation.ListStackSetOperationsOutput, bool) bool) error
	ListStackSetOperationsPagesWithContext(aws.Context, *cloudformation.ListStackSetOperationsInput, func(*cloudformation.ListStackSetOperationsOutput, bool) bool, ...request.Option) error

	ListStackSets(*cloudformation.ListStackSetsInput) (*cloudformation.ListStackSetsOutput, error)
	ListStackSetsWithContext(aws.Context, *cloudformation.ListStackSetsInput, ...request.Option) (*cloudformation.ListStackSetsOutput, error)
	ListStackSetsRequest(*cloudformation.ListStackSetsInput) (*request.Request, *cloudformation.ListStackSetsOutput)

	ListStackSetsPages(*cloudformation.ListStackSetsInput, func(*cloudformation.ListStackSetsOutput, bool) bool) error
	ListStackSetsPagesWithContext(aws.Context, *cloudformation.ListStackSetsInput, func(*cloudformation.ListStackSetsOutput, bool) bool, ...request.Option) error

	ListStacks(*cloudformation.ListStacksInput) (*cloudformation.ListStacksOutput, error)
	ListStacksWithContext(aws.Context, *cloudformation.ListStacksInput, ...request.Option) (*cloudformation.ListStacksOutput, error)
	ListStacksRequest(*cloudformation.ListStacksInput) (*request.Request, *cloudformation.ListStacksOutput)

	ListStacksPages(*cloudformation.ListStacksInput, func(*cloudformation.ListStacksOutput, bool) bool) error
	ListStacksPagesWithContext(aws.Context, *cloudformation.ListStacksInput, func(*cloudformation.ListStacksOutput, bool) bool, ...request.Option) error

	ListTypeRegistrations(*cloudformation.ListTypeRegistrationsInput) (*cloudformation.ListTypeRegistrationsOutput, error)
	ListTypeRegistrationsWithContext(aws.Context, *cloudformation.ListTypeRegistrationsInput, ...request.Option) (*cloudformation.ListTypeRegistrationsOutput, error)
	ListTypeRegistrationsRequest(*cloudformation.ListTypeRegistrationsInput) (*request.Request, *cloudformation.ListTypeRegistrationsOutput)

	ListTypeRegistrationsPages(*cloudformation.ListTypeRegistrationsInput, func(*cloudformation.ListTypeRegistrationsOutput, bool) bool) error
	ListTypeRegistrationsPagesWithContext(aws.Context, *cloudformation.ListTypeRegistrationsInput, func(*cloudformation.ListTypeRegistrationsOutput, bool) bool, ...request.Option) error

	ListTypeVersions(*cloudformation.ListTypeVersionsInput) (*cloudformation.ListTypeVersionsOutput, error)
	ListTypeVersionsWithContext(aws.Context, *cloudformation.ListTypeVersionsInput, ...request.Option) (*cloudformation.ListTypeVersionsOutput, error)
	ListTypeVersionsRequest(*cloudformation.ListTypeVersionsInput) (*request.Request, *cloudformation.ListTypeVersionsOutput)

	ListTypeVersionsPages(*cloudformation.ListTypeVersionsInput, func(*cloudformation.ListTypeVersionsOutput, bool) bool) error
	ListTypeVersionsPagesWithContext(aws.Context, *cloudformation.ListTypeVersionsInput, func(*cloudformation.ListTypeVersionsOutput, bool) bool, ...request.Option) error

	ListTypes(*cloudformation.ListTypesInput) (*cloudformation.ListTypesOutput, error)
	ListTypesWithContext(aws.Context, *cloudformation.ListTypesInput, ...request.Option) (*cloudformation.ListTypesOutput, error)
	ListTypesRequest(*cloudformation.ListTypesInput) (*request.Request, *cloudformation.ListTypesOutput)

	ListTypesPages(*cloudformation.ListTypesInput, func(*cloudformation.ListTypesOutput, bool) bool) error
	ListTypesPagesWithContext(aws.Context, *cloudformation.ListTypesInput, func(*cloudformation.ListTypesOutput, bool) bool, ...request.Option) error

	PublishType(*cloudformation.PublishTypeInput) (*cloudformation.PublishTypeOutput, error)
	PublishTypeWithContext(aws.Context, *cloudformation.PublishTypeInput, ...request.Option) (*cloudformation.PublishTypeOutput, error)
	PublishTypeRequest(*cloudformation.PublishTypeInput) (*request.Request, *cloudformation.PublishTypeOutput)

	RecordHandlerProgress(*cloudformation.RecordHandlerProgressInput) (*cloudformation.RecordHandlerProgressOutput, error)
	RecordHandlerProgressWithContext(aws.Context, *cloudformation.RecordHandlerProgressInput, ...request.Option) (*cloudformation.RecordHandlerProgressOutput, error)
	RecordHandlerProgressRequest(*cloudformation.RecordHandlerProgressInput) (*request.Request, *cloudformation.RecordHandlerProgressOutput)

	RegisterPublisher(*cloudformation.RegisterPublisherInput) (*cloudformation.RegisterPublisherOutput, error)
	RegisterPublisherWithContext(aws.Context, *cloudformation.RegisterPublisherInput, ...request.Option) (*cloudformation.RegisterPublisherOutput, error)
	RegisterPublisherRequest(*cloudformation.RegisterPublisherInput) (*request.Request, *cloudformation.RegisterPublisherOutput)

	RegisterType(*cloudformation.RegisterTypeInput) (*cloudformation.RegisterTypeOutput, error)
	RegisterTypeWithContext(aws.Context, *cloudformation.RegisterTypeInput, ...request.Option) (*cloudformation.RegisterTypeOutput, error)
	RegisterTypeRequest(*cloudformation.RegisterTypeInput) (*request.Request, *cloudformation.RegisterTypeOutput)

	RollbackStack(*cloudformation.RollbackStackInput) (*cloudformation.RollbackStackOutput, error)
	RollbackStackWithContext(aws.Context, *cloudformation.RollbackStackInput, ...request.Option) (*cloudformation.RollbackStackOutput, error)
	RollbackStackRequest(*cloudformation.RollbackStackInput) (*request.Request, *cloudformation.RollbackStackOutput)

	SetStackPolicy(*cloudformation.SetStackPolicyInput) (*cloudformation.SetStackPolicyOutput, error)
	SetStackPolicyWithContext(aws.Context, *cloudformation.SetStackPolicyInput, ...request.Option) (*cloudformation.SetStackPolicyOutput, error)
	SetStackPolicyRequest(*cloudformation.SetStackPolicyInput) (*request.Request, *cloudformation.SetStackPolicyOutput)

	SetTypeConfiguration(*cloudformation.SetTypeConfigurationInput) (*cloudformation.SetTypeConfigurationOutput, error)
	SetTypeConfigurationWithContext(aws.Context, *cloudformation.SetTypeConfigurationInput, ...request.Option) (*cloudformation.SetTypeConfigurationOutput, error)
	SetTypeConfigurationRequest(*cloudformation.SetTypeConfigurationInput) (*request.Request, *cloudformation.SetTypeConfigurationOutput)

	SetTypeDefaultVersion(*cloudformation.SetTypeDefaultVersionInput) (*cloudformation.SetTypeDefaultVersionOutput, error)
	SetTypeDefaultVersionWithContext(aws.Context, *cloudformation.SetTypeDefaultVersionInput, ...request.Option) (*cloudformation.SetTypeDefaultVersionOutput, error)
	SetTypeDefaultVersionRequest(*cloudformation.SetTypeDefaultVersionInput) (*request.Request, *cloudformation.SetTypeDefaultVersionOutput)

	SignalResource(*cloudformation.SignalResourceInput) (*cloudformation.SignalResourceOutput, error)
	SignalResourceWithContext(aws.Context, *cloudformation.SignalResourceInput, ...request.Option) (*cloudformation.SignalResourceOutput, error)
	SignalResourceRequest(*cloudformation.SignalResourceInput) (*request.Request, *cloudformation.SignalResourceOutput)

	StopStackSetOperation(*cloudformation.StopStackSetOperationInput) (*cloudformation.StopStackSetOperationOutput, error)
	StopStackSetOperationWithContext(aws.Context, *cloudformation.StopStackSetOperationInput, ...request.Option) (*cloudformation.StopStackSetOperationOutput, error)
	StopStackSetOperationRequest(*cloudformation.StopStackSetOperationInput) (*request.Request, *cloudformation.StopStackSetOperationOutput)

	TestType(*cloudformation.TestTypeInput) (*cloudformation.TestTypeOutput, error)
	TestTypeWithContext(aws.Context, *cloudformation.TestTypeInput, ...request.Option) (*cloudformation.TestTypeOutput, error)
	TestTypeRequest(*cloudformation.TestTypeInput) (*request.Request, *cloudformation.TestTypeOutput)

	UpdateStack(*cloudformation.UpdateStackInput) (*cloudformation.UpdateStackOutput, error)
	UpdateStackWithContext(aws.Context, *cloudformation.UpdateStackInput, ...request.Option) (*cloudformation.UpdateStackOutput, error)
	UpdateStackRequest(*cloudformation.UpdateStackInput) (*request.Request, *cloudformation.UpdateStackOutput)

	UpdateStackInstances(*cloudformation.UpdateStackInstancesInput) (*cloudformation.UpdateStackInstancesOutput, error)
	UpdateStackInstancesWithContext(aws.Context, *cloudformation.UpdateStackInstancesInput, ...request.Option) (*cloudformation.UpdateStackInstancesOutput, error)
	UpdateStackInstancesRequest(*cloudformation.UpdateStackInstancesInput) (*request.Request, *cloudformation.UpdateStackInstancesOutput)

	UpdateStackSet(*cloudformation.UpdateStackSetInput) (*cloudformation.UpdateStackSetOutput, error)
	UpdateStackSetWithContext(aws.Context, *cloudformation.UpdateStackSetInput, ...request.Option) (*cloudformation.UpdateStackSetOutput, error)
	UpdateStackSetRequest(*cloudformation.UpdateStackSetInput) (*request.Request, *cloudformation.UpdateStackSetOutput)

	UpdateTerminationProtection(*cloudformation.UpdateTerminationProtectionInput) (*cloudformation.UpdateTerminationProtectionOutput, error)
	UpdateTerminationProtectionWithContext(aws.Context, *cloudformation.UpdateTerminationProtectionInput, ...request.Option) (*cloudformation.UpdateTerminationProtectionOutput, error)
	UpdateTerminationProtectionRequest(*cloudformation.UpdateTerminationProtectionInput) (*request.Request, *cloudformation.UpdateTerminationProtectionOutput)

	ValidateTemplate(*cloudformation.ValidateTemplateInput) (*cloudformation.ValidateTemplateOutput, error)
	ValidateTemplateWithContext(aws.Context, *cloudformation.ValidateTemplateInput, ...request.Option) (*cloudformation.ValidateTemplateOutput, error)
	ValidateTemplateRequest(*cloudformation.ValidateTemplateInput) (*request.Request, *cloudformation.ValidateTemplateOutput)

	WaitUntilChangeSetCreateComplete(*cloudformation.DescribeChangeSetInput) error
	WaitUntilChangeSetCreateCompleteWithContext(aws.Context, *cloudformation.DescribeChangeSetInput, ...request.WaiterOption) error

	WaitUntilStackCreateComplete(*cloudformation.DescribeStacksInput) error
	WaitUntilStackCreateCompleteWithContext(aws.Context, *cloudformation.DescribeStacksInput, ...request.WaiterOption) error

	WaitUntilStackDeleteComplete(*cloudformation.DescribeStacksInput) error
	WaitUntilStackDeleteCompleteWithContext(aws.Context, *cloudformation.DescribeStacksInput, ...request.WaiterOption) error

	WaitUntilStackExists(*cloudformation.DescribeStacksInput) error
	WaitUntilStackExistsWithContext(aws.Context, *cloudformation.DescribeStacksInput, ...request.WaiterOption) error

	WaitUntilStackImportComplete(*cloudformation.DescribeStacksInput) error
	WaitUntilStackImportCompleteWithContext(aws.Context, *cloudformation.DescribeStacksInput, ...request.WaiterOption) error

	WaitUntilStackRollbackComplete(*cloudformation.DescribeStacksInput) error
	WaitUntilStackRollbackCompleteWithContext(aws.Context, *cloudformation.DescribeStacksInput, ...request.WaiterOption) error

	WaitUntilStackUpdateComplete(*cloudformation.DescribeStacksInput) error
	WaitUntilStackUpdateCompleteWithContext(aws.Context, *cloudformation.DescribeStacksInput, ...request.WaiterOption) error

	WaitUntilTypeRegistrationComplete(*cloudformation.DescribeTypeRegistrationInput) error
	WaitUntilTypeRegistrationCompleteWithContext(aws.Context, *cloudformation.DescribeTypeRegistrationInput, ...request.WaiterOption) error
}

var _ CloudFormationAPI = (*cloudformation.CloudFormation)(nil)
