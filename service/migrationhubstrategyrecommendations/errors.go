// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhubstrategyrecommendations

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// The user does not have permission to perform the action. Check the AWS Identity
	// and Access Management (IAM) policy associated with this user.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// Exception to indicate that there is an ongoing task when a new task is created.
	// Return when once the existing tasks are complete.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeDependencyException for service response error code
	// "DependencyException".
	//
	// Dependency encountered an error.
	ErrCodeDependencyException = "DependencyException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// The server experienced an internal error. Try again.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified ID in the request is not found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceLinkedRoleLockClientException for service response error code
	// "ServiceLinkedRoleLockClientException".
	//
	// Exception to indicate that the service-linked role (SLR) is locked.
	ErrCodeServiceLinkedRoleLockClientException = "ServiceLinkedRoleLockClientException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// The AWS account has reached its quota of imports. Contact AWS Support to
	// increase the quota for this account.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied due to request throttling.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The request body isn't valid.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":                newErrorAccessDeniedException,
	"ConflictException":                    newErrorConflictException,
	"DependencyException":                  newErrorDependencyException,
	"InternalServerException":              newErrorInternalServerException,
	"ResourceNotFoundException":            newErrorResourceNotFoundException,
	"ServiceLinkedRoleLockClientException": newErrorServiceLinkedRoleLockClientException,
	"ServiceQuotaExceededException":        newErrorServiceQuotaExceededException,
	"ThrottlingException":                  newErrorThrottlingException,
	"ValidationException":                  newErrorValidationException,
}
