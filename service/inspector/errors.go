// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have required permissions to access the requested resource.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeAgentsAlreadyRunningAssessmentException for service response error code
	// "AgentsAlreadyRunningAssessmentException".
	//
	// You started an assessment run, but one of the instances is already participating
	// in another assessment run.
	ErrCodeAgentsAlreadyRunningAssessmentException = "AgentsAlreadyRunningAssessmentException"

	// ErrCodeAssessmentRunInProgressException for service response error code
	// "AssessmentRunInProgressException".
	//
	// You cannot perform a specified action if an assessment run is currently in
	// progress.
	ErrCodeAssessmentRunInProgressException = "AssessmentRunInProgressException"

	// ErrCodeInternalException for service response error code
	// "InternalException".
	//
	// Internal server error.
	ErrCodeInternalException = "InternalException"

	// ErrCodeInvalidCrossAccountRoleException for service response error code
	// "InvalidCrossAccountRoleException".
	//
	// Amazon Inspector cannot assume the cross-account role that it needs to list
	// your EC2 instances during the assessment run.
	ErrCodeInvalidCrossAccountRoleException = "InvalidCrossAccountRoleException"

	// ErrCodeInvalidInputException for service response error code
	// "InvalidInputException".
	//
	// The request was rejected because an invalid or out-of-range value was supplied
	// for an input parameter.
	ErrCodeInvalidInputException = "InvalidInputException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The request was rejected because it attempted to create resources beyond
	// the current AWS account limits. The error code describes the limit exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeNoSuchEntityException for service response error code
	// "NoSuchEntityException".
	//
	// The request was rejected because it referenced an entity that does not exist.
	// The error code describes the entity.
	ErrCodeNoSuchEntityException = "NoSuchEntityException"

	// ErrCodePreviewGenerationInProgressException for service response error code
	// "PreviewGenerationInProgressException".
	//
	// The request is rejected. The specified assessment template is currently generating
	// an exclusions preview.
	ErrCodePreviewGenerationInProgressException = "PreviewGenerationInProgressException"

	// ErrCodeServiceTemporarilyUnavailableException for service response error code
	// "ServiceTemporarilyUnavailableException".
	//
	// The serice is temporary unavailable.
	ErrCodeServiceTemporarilyUnavailableException = "ServiceTemporarilyUnavailableException"

	// ErrCodeUnsupportedFeatureException for service response error code
	// "UnsupportedFeatureException".
	//
	// Used by the GetAssessmentReport API. The request was rejected because you
	// tried to generate a report for an assessment run that existed before reporting
	// was supported in Amazon Inspector. You can only generate reports for assessment
	// runs that took place or will take place after generating reports in Amazon
	// Inspector became available.
	ErrCodeUnsupportedFeatureException = "UnsupportedFeatureException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":                   newErrorAccessDeniedException,
	"AgentsAlreadyRunningAssessmentException": newErrorAgentsAlreadyRunningAssessmentException,
	"AssessmentRunInProgressException":        newErrorAssessmentRunInProgressException,
	"InternalException":                       newErrorInternalException,
	"InvalidCrossAccountRoleException":        newErrorInvalidCrossAccountRoleException,
	"InvalidInputException":                   newErrorInvalidInputException,
	"LimitExceededException":                  newErrorLimitExceededException,
	"NoSuchEntityException":                   newErrorNoSuchEntityException,
	"PreviewGenerationInProgressException":    newErrorPreviewGenerationInProgressException,
	"ServiceTemporarilyUnavailableException":  newErrorServiceTemporarilyUnavailableException,
	"UnsupportedFeatureException":             newErrorUnsupportedFeatureException,
}
