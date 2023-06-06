// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mq

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// Returns information about an error.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// Returns information about an error.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeForbiddenException for service response error code
	// "ForbiddenException".
	//
	// Returns information about an error.
	ErrCodeForbiddenException = "ForbiddenException"

	// ErrCodeInternalServerErrorException for service response error code
	// "InternalServerErrorException".
	//
	// Returns information about an error.
	ErrCodeInternalServerErrorException = "InternalServerErrorException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// Returns information about an error.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeUnauthorizedException for service response error code
	// "UnauthorizedException".
	//
	// Returns information about an error.
	ErrCodeUnauthorizedException = "UnauthorizedException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"BadRequestException":          newErrorBadRequestException,
	"ConflictException":            newErrorConflictException,
	"ForbiddenException":           newErrorForbiddenException,
	"InternalServerErrorException": newErrorInternalServerErrorException,
	"NotFoundException":            newErrorNotFoundException,
	"UnauthorizedException":        newErrorUnauthorizedException,
}
