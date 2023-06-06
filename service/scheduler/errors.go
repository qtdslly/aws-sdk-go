// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package scheduler

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// Updating or deleting the resource can cause an inconsistent state.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// Unexpected error encountered while processing the request.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The request references a resource which does not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// The request exceeds a service quota.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied due to request throttling.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The input fails to satisfy the constraints specified by an AWS service.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ConflictException":             newErrorConflictException,
	"InternalServerException":       newErrorInternalServerException,
	"ResourceNotFoundException":     newErrorResourceNotFoundException,
	"ServiceQuotaExceededException": newErrorServiceQuotaExceededException,
	"ThrottlingException":           newErrorThrottlingException,
	"ValidationException":           newErrorValidationException,
}
