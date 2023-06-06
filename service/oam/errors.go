// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package oam

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// A resource was in an inconsistent state during an update or a deletion.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServiceFault for service response error code
	// "InternalServiceFault".
	//
	// Unexpected error while processing the request. Retry the request.
	ErrCodeInternalServiceFault = "InternalServiceFault"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// A parameter is specified incorrectly.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeMissingRequiredParameterException for service response error code
	// "MissingRequiredParameterException".
	//
	// A required parameter is missing from the request.
	ErrCodeMissingRequiredParameterException = "MissingRequiredParameterException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The request references a resource that does not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// The request would cause a service quota to be exceeded.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTagsException".
	//
	// A resource can have no more than 50 tags.
	ErrCodeTooManyTagsException = "TooManyTagsException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The value of a parameter in the request caused an error.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ConflictException":                 newErrorConflictException,
	"InternalServiceFault":              newErrorInternalServiceFault,
	"InvalidParameterException":         newErrorInvalidParameterException,
	"MissingRequiredParameterException": newErrorMissingRequiredParameterException,
	"ResourceNotFoundException":         newErrorResourceNotFoundException,
	"ServiceQuotaExceededException":     newErrorServiceQuotaExceededException,
	"TooManyTagsException":              newErrorTooManyTagsException,
	"ValidationException":               newErrorValidationException,
}
