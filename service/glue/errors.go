// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// Access to a resource was denied.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeAlreadyExistsException for service response error code
	// "AlreadyExistsException".
	//
	// A resource to be created or added already exists.
	ErrCodeAlreadyExistsException = "AlreadyExistsException"

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// Two processes are trying to modify a resource simultaneously.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeConcurrentRunsExceededException for service response error code
	// "ConcurrentRunsExceededException".
	//
	// Too many jobs are being run concurrently.
	ErrCodeConcurrentRunsExceededException = "ConcurrentRunsExceededException"

	// ErrCodeConditionCheckFailureException for service response error code
	// "ConditionCheckFailureException".
	//
	// A specified condition was not satisfied.
	ErrCodeConditionCheckFailureException = "ConditionCheckFailureException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// The CreatePartitions API was called on a table that has indexes enabled.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeCrawlerNotRunningException for service response error code
	// "CrawlerNotRunningException".
	//
	// The specified crawler is not running.
	ErrCodeCrawlerNotRunningException = "CrawlerNotRunningException"

	// ErrCodeCrawlerRunningException for service response error code
	// "CrawlerRunningException".
	//
	// The operation cannot be performed because the crawler is already running.
	ErrCodeCrawlerRunningException = "CrawlerRunningException"

	// ErrCodeCrawlerStoppingException for service response error code
	// "CrawlerStoppingException".
	//
	// The specified crawler is stopping.
	ErrCodeCrawlerStoppingException = "CrawlerStoppingException"

	// ErrCodeEncryptionException for service response error code
	// "GlueEncryptionException".
	//
	// An encryption operation failed.
	ErrCodeEncryptionException = "GlueEncryptionException"

	// ErrCodeEntityNotFoundException for service response error code
	// "EntityNotFoundException".
	//
	// A specified entity does not exist
	ErrCodeEntityNotFoundException = "EntityNotFoundException"

	// ErrCodeFederatedResourceAlreadyExistsException for service response error code
	// "FederatedResourceAlreadyExistsException".
	//
	// A federated resource already exists.
	ErrCodeFederatedResourceAlreadyExistsException = "FederatedResourceAlreadyExistsException"

	// ErrCodeFederationSourceException for service response error code
	// "FederationSourceException".
	//
	// A federation source failed.
	ErrCodeFederationSourceException = "FederationSourceException"

	// ErrCodeFederationSourceRetryableException for service response error code
	// "FederationSourceRetryableException".
	ErrCodeFederationSourceRetryableException = "FederationSourceRetryableException"

	// ErrCodeIdempotentParameterMismatchException for service response error code
	// "IdempotentParameterMismatchException".
	//
	// The same unique identifier was associated with two different records.
	ErrCodeIdempotentParameterMismatchException = "IdempotentParameterMismatchException"

	// ErrCodeIllegalBlueprintStateException for service response error code
	// "IllegalBlueprintStateException".
	//
	// The blueprint is in an invalid state to perform a requested operation.
	ErrCodeIllegalBlueprintStateException = "IllegalBlueprintStateException"

	// ErrCodeIllegalSessionStateException for service response error code
	// "IllegalSessionStateException".
	//
	// The session is in an invalid state to perform a requested operation.
	ErrCodeIllegalSessionStateException = "IllegalSessionStateException"

	// ErrCodeIllegalWorkflowStateException for service response error code
	// "IllegalWorkflowStateException".
	//
	// The workflow is in an invalid state to perform a requested operation.
	ErrCodeIllegalWorkflowStateException = "IllegalWorkflowStateException"

	// ErrCodeInternalServiceException for service response error code
	// "InternalServiceException".
	//
	// An internal service error occurred.
	ErrCodeInternalServiceException = "InternalServiceException"

	// ErrCodeInvalidInputException for service response error code
	// "InvalidInputException".
	//
	// The input provided was not valid.
	ErrCodeInvalidInputException = "InvalidInputException"

	// ErrCodeInvalidStateException for service response error code
	// "InvalidStateException".
	//
	// An error that indicates your data is in an invalid state.
	ErrCodeInvalidStateException = "InvalidStateException"

	// ErrCodeMLTransformNotReadyException for service response error code
	// "MLTransformNotReadyException".
	//
	// The machine learning transform is not ready to run.
	ErrCodeMLTransformNotReadyException = "MLTransformNotReadyException"

	// ErrCodeNoScheduleException for service response error code
	// "NoScheduleException".
	//
	// There is no applicable schedule.
	ErrCodeNoScheduleException = "NoScheduleException"

	// ErrCodeOperationTimeoutException for service response error code
	// "OperationTimeoutException".
	//
	// The operation timed out.
	ErrCodeOperationTimeoutException = "OperationTimeoutException"

	// ErrCodePermissionTypeMismatchException for service response error code
	// "PermissionTypeMismatchException".
	//
	// The operation timed out.
	ErrCodePermissionTypeMismatchException = "PermissionTypeMismatchException"

	// ErrCodeResourceNotReadyException for service response error code
	// "ResourceNotReadyException".
	//
	// A resource was not ready for a transaction.
	ErrCodeResourceNotReadyException = "ResourceNotReadyException"

	// ErrCodeResourceNumberLimitExceededException for service response error code
	// "ResourceNumberLimitExceededException".
	//
	// A resource numerical limit was exceeded.
	ErrCodeResourceNumberLimitExceededException = "ResourceNumberLimitExceededException"

	// ErrCodeSchedulerNotRunningException for service response error code
	// "SchedulerNotRunningException".
	//
	// The specified scheduler is not running.
	ErrCodeSchedulerNotRunningException = "SchedulerNotRunningException"

	// ErrCodeSchedulerRunningException for service response error code
	// "SchedulerRunningException".
	//
	// The specified scheduler is already running.
	ErrCodeSchedulerRunningException = "SchedulerRunningException"

	// ErrCodeSchedulerTransitioningException for service response error code
	// "SchedulerTransitioningException".
	//
	// The specified scheduler is transitioning.
	ErrCodeSchedulerTransitioningException = "SchedulerTransitioningException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// A value could not be validated.
	ErrCodeValidationException = "ValidationException"

	// ErrCodeVersionMismatchException for service response error code
	// "VersionMismatchException".
	//
	// There was a version conflict.
	ErrCodeVersionMismatchException = "VersionMismatchException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":                   newErrorAccessDeniedException,
	"AlreadyExistsException":                  newErrorAlreadyExistsException,
	"ConcurrentModificationException":         newErrorConcurrentModificationException,
	"ConcurrentRunsExceededException":         newErrorConcurrentRunsExceededException,
	"ConditionCheckFailureException":          newErrorConditionCheckFailureException,
	"ConflictException":                       newErrorConflictException,
	"CrawlerNotRunningException":              newErrorCrawlerNotRunningException,
	"CrawlerRunningException":                 newErrorCrawlerRunningException,
	"CrawlerStoppingException":                newErrorCrawlerStoppingException,
	"GlueEncryptionException":                 newErrorEncryptionException,
	"EntityNotFoundException":                 newErrorEntityNotFoundException,
	"FederatedResourceAlreadyExistsException": newErrorFederatedResourceAlreadyExistsException,
	"FederationSourceException":               newErrorFederationSourceException,
	"FederationSourceRetryableException":      newErrorFederationSourceRetryableException,
	"IdempotentParameterMismatchException":    newErrorIdempotentParameterMismatchException,
	"IllegalBlueprintStateException":          newErrorIllegalBlueprintStateException,
	"IllegalSessionStateException":            newErrorIllegalSessionStateException,
	"IllegalWorkflowStateException":           newErrorIllegalWorkflowStateException,
	"InternalServiceException":                newErrorInternalServiceException,
	"InvalidInputException":                   newErrorInvalidInputException,
	"InvalidStateException":                   newErrorInvalidStateException,
	"MLTransformNotReadyException":            newErrorMLTransformNotReadyException,
	"NoScheduleException":                     newErrorNoScheduleException,
	"OperationTimeoutException":               newErrorOperationTimeoutException,
	"PermissionTypeMismatchException":         newErrorPermissionTypeMismatchException,
	"ResourceNotReadyException":               newErrorResourceNotReadyException,
	"ResourceNumberLimitExceededException":    newErrorResourceNumberLimitExceededException,
	"SchedulerNotRunningException":            newErrorSchedulerNotRunningException,
	"SchedulerRunningException":               newErrorSchedulerRunningException,
	"SchedulerTransitioningException":         newErrorSchedulerTransitioningException,
	"ValidationException":                     newErrorValidationException,
	"VersionMismatchException":                newErrorVersionMismatchException,
}
