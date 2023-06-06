// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package internetmonitor

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You don't have sufficient permission to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// A bad request was received.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// The requested resource is in use.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerErrorException for service response error code
	// "InternalServerErrorException".
	//
	// There was an internal server error.
	ErrCodeInternalServerErrorException = "InternalServerErrorException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// An internal error occurred.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The request exceeded a service quota.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// The request specifies something that doesn't exist.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The request specifies a resource that doesn't exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied due to request throttling.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// There were too many requests.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// Invalid request.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":        newErrorAccessDeniedException,
	"BadRequestException":          newErrorBadRequestException,
	"ConflictException":            newErrorConflictException,
	"InternalServerErrorException": newErrorInternalServerErrorException,
	"InternalServerException":      newErrorInternalServerException,
	"LimitExceededException":       newErrorLimitExceededException,
	"NotFoundException":            newErrorNotFoundException,
	"ResourceNotFoundException":    newErrorResourceNotFoundException,
	"ThrottlingException":          newErrorThrottlingException,
	"TooManyRequestsException":     newErrorTooManyRequestsException,
	"ValidationException":          newErrorValidationException,
}
