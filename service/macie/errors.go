// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// (Discontinued) You do not have required permissions to access the requested
	// resource.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeInternalException for service response error code
	// "InternalException".
	//
	// (Discontinued) Internal server error.
	ErrCodeInternalException = "InternalException"

	// ErrCodeInvalidInputException for service response error code
	// "InvalidInputException".
	//
	// (Discontinued) The request was rejected because an invalid or out-of-range
	// value was supplied for an input parameter.
	ErrCodeInvalidInputException = "InvalidInputException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// (Discontinued) The request was rejected because it attempted to create resources
	// beyond the current Amazon Web Services account quotas. The error code describes
	// the quota exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":  newErrorAccessDeniedException,
	"InternalException":      newErrorInternalException,
	"InvalidInputException":  newErrorInvalidInputException,
	"LimitExceededException": newErrorLimitExceededException,
}
