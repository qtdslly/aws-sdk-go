// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeInternalServerErrorException for service response error code
	// "InternalServerErrorException".
	//
	// The service is temporarily unavailable.
	ErrCodeInternalServerErrorException = "InternalServerErrorException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// The request is not valid.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeResourceAlreadyExistsException for service response error code
	// "ResourceAlreadyExistsException".
	//
	// The resource already exists.
	ErrCodeResourceAlreadyExistsException = "ResourceAlreadyExistsException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The requested resource was not found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// The number of requests exceeds the limit.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"

	// ErrCodeUnauthorizedException for service response error code
	// "UnauthorizedException".
	//
	// You are not authorized to perform this action.
	ErrCodeUnauthorizedException = "UnauthorizedException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InternalServerErrorException":   newErrorInternalServerErrorException,
	"InvalidRequestException":        newErrorInvalidRequestException,
	"ResourceAlreadyExistsException": newErrorResourceAlreadyExistsException,
	"ResourceNotFoundException":      newErrorResourceNotFoundException,
	"TooManyRequestsException":       newErrorTooManyRequestsException,
	"UnauthorizedException":          newErrorUnauthorizedException,
}
