// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticinference

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// Raised when a malformed input has been provided to the API.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// Raised when an unexpected error occurred during request processing.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Raised when the requested resource cannot be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"BadRequestException":       newErrorBadRequestException,
	"InternalServerException":   newErrorInternalServerException,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
}
