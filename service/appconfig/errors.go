// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appconfig

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// The input fails to satisfy the constraints specified by an Amazon Web Services
	// service.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// The request could not be processed because of conflict in the current state
	// of the resource.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// There was an internal failure in the AppConfig service.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodePayloadTooLargeException for service response error code
	// "PayloadTooLargeException".
	//
	// The configuration size is too large.
	ErrCodePayloadTooLargeException = "PayloadTooLargeException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The requested resource could not be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// The number of hosted configuration versions exceeds the limit for the AppConfig
	// hosted configuration store. Delete one or more versions and try again.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"BadRequestException":           newErrorBadRequestException,
	"ConflictException":             newErrorConflictException,
	"InternalServerException":       newErrorInternalServerException,
	"PayloadTooLargeException":      newErrorPayloadTooLargeException,
	"ResourceNotFoundException":     newErrorResourceNotFoundException,
	"ServiceQuotaExceededException": newErrorServiceQuotaExceededException,
}