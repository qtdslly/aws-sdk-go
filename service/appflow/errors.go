// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appflow

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// AppFlow/Requester has invalid or missing permissions.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// There was a conflict when processing the request (for example, a flow with
	// the given name already exists within the account. Check for conflicting resource
	// names and try again.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeConnectorAuthenticationException for service response error code
	// "ConnectorAuthenticationException".
	//
	// An error occurred when authenticating with the connector endpoint.
	ErrCodeConnectorAuthenticationException = "ConnectorAuthenticationException"

	// ErrCodeConnectorServerException for service response error code
	// "ConnectorServerException".
	//
	// An error occurred when retrieving data from the connector endpoint.
	ErrCodeConnectorServerException = "ConnectorServerException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// An internal service error occurred during the processing of your request.
	// Try again later.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The resource specified in the request (such as the source or destination
	// connector profile) is not found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// The request would cause a service quota (such as the number of flows) to
	// be exceeded.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// API calls have exceeded the maximum allowed API request rate per account
	// and per Region.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeUnsupportedOperationException for service response error code
	// "UnsupportedOperationException".
	//
	// The requested operation is not supported for the current flow.
	ErrCodeUnsupportedOperationException = "UnsupportedOperationException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The request has invalid or missing parameters.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":            newErrorAccessDeniedException,
	"ConflictException":                newErrorConflictException,
	"ConnectorAuthenticationException": newErrorConnectorAuthenticationException,
	"ConnectorServerException":         newErrorConnectorServerException,
	"InternalServerException":          newErrorInternalServerException,
	"ResourceNotFoundException":        newErrorResourceNotFoundException,
	"ServiceQuotaExceededException":    newErrorServiceQuotaExceededException,
	"ThrottlingException":              newErrorThrottlingException,
	"UnsupportedOperationException":    newErrorUnsupportedOperationException,
	"ValidationException":              newErrorValidationException,
}