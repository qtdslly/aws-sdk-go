// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// An exception indicating Amazon Fraud Detector does not have the needed permissions.
	// This can occur if you submit a request, such as PutExternalModel, that specifies
	// a role that is not in your account.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// An exception indicating there was a conflict during a delete operation.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// An exception indicating an internal server error.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// An exception indicating the specified resource was not found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeResourceUnavailableException for service response error code
	// "ResourceUnavailableException".
	//
	// An exception indicating that the attached customer-owned (external) model
	// threw an exception when Amazon Fraud Detector invoked the model.
	ErrCodeResourceUnavailableException = "ResourceUnavailableException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// An exception indicating a throttling error.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// An exception indicating a specified value is not allowed.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":        newErrorAccessDeniedException,
	"ConflictException":            newErrorConflictException,
	"InternalServerException":      newErrorInternalServerException,
	"ResourceNotFoundException":    newErrorResourceNotFoundException,
	"ResourceUnavailableException": newErrorResourceUnavailableException,
	"ThrottlingException":          newErrorThrottlingException,
	"ValidationException":          newErrorValidationException,
}
