// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudcontrolapi

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAlreadyExistsException for service response error code
	// "AlreadyExistsException".
	//
	// The resource with the name requested already exists.
	ErrCodeAlreadyExistsException = "AlreadyExistsException"

	// ErrCodeClientTokenConflictException for service response error code
	// "ClientTokenConflictException".
	//
	// The specified client token has already been used in another resource request.
	//
	// It's best practice for client tokens to be unique for each resource operation
	// request. However, client token expire after 36 hours.
	ErrCodeClientTokenConflictException = "ClientTokenConflictException"

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// The resource is currently being modified by another operation.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeConcurrentOperationException for service response error code
	// "ConcurrentOperationException".
	//
	// Another resource operation is currently being performed on this resource.
	ErrCodeConcurrentOperationException = "ConcurrentOperationException"

	// ErrCodeGeneralServiceException for service response error code
	// "GeneralServiceException".
	//
	// The resource handler has returned that the downstream service generated an
	// error that doesn't map to any other handler error code.
	ErrCodeGeneralServiceException = "GeneralServiceException"

	// ErrCodeHandlerFailureException for service response error code
	// "HandlerFailureException".
	//
	// The resource handler has failed without a returning a more specific error
	// code. This can include timeouts.
	ErrCodeHandlerFailureException = "HandlerFailureException"

	// ErrCodeHandlerInternalFailureException for service response error code
	// "HandlerInternalFailureException".
	//
	// The resource handler has returned that an unexpected error occurred within
	// the resource handler.
	ErrCodeHandlerInternalFailureException = "HandlerInternalFailureException"

	// ErrCodeInvalidCredentialsException for service response error code
	// "InvalidCredentialsException".
	//
	// The resource handler has returned that the credentials provided by the user
	// are invalid.
	ErrCodeInvalidCredentialsException = "InvalidCredentialsException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// The resource handler has returned that invalid input from the user has generated
	// a generic exception.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeNetworkFailureException for service response error code
	// "NetworkFailureException".
	//
	// The resource handler has returned that the request couldn't be completed
	// due to networking issues, such as a failure to receive a response from the
	// server.
	ErrCodeNetworkFailureException = "NetworkFailureException"

	// ErrCodeNotStabilizedException for service response error code
	// "NotStabilizedException".
	//
	// The resource handler has returned that the downstream resource failed to
	// complete all of its ready-state checks.
	ErrCodeNotStabilizedException = "NotStabilizedException"

	// ErrCodeNotUpdatableException for service response error code
	// "NotUpdatableException".
	//
	// One or more properties included in this resource operation are defined as
	// create-only, and therefore can't be updated.
	ErrCodeNotUpdatableException = "NotUpdatableException"

	// ErrCodePrivateTypeException for service response error code
	// "PrivateTypeException".
	//
	// Cloud Control API hasn't received a valid response from the resource handler,
	// due to a configuration error. This includes issues such as the resource handler
	// returning an invalid response, or timing out.
	ErrCodePrivateTypeException = "PrivateTypeException"

	// ErrCodeRequestTokenNotFoundException for service response error code
	// "RequestTokenNotFoundException".
	//
	// A resource operation with the specified request token can't be found.
	ErrCodeRequestTokenNotFoundException = "RequestTokenNotFoundException"

	// ErrCodeResourceConflictException for service response error code
	// "ResourceConflictException".
	//
	// The resource is temporarily unavailable to be acted upon. For example, if
	// the resource is currently undergoing an operation and can't be acted upon
	// until that operation is finished.
	ErrCodeResourceConflictException = "ResourceConflictException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// A resource with the specified identifier can't be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceInternalErrorException for service response error code
	// "ServiceInternalErrorException".
	//
	// The resource handler has returned that the downstream service returned an
	// internal error, typically with a 5XX HTTP status code.
	ErrCodeServiceInternalErrorException = "ServiceInternalErrorException"

	// ErrCodeServiceLimitExceededException for service response error code
	// "ServiceLimitExceededException".
	//
	// The resource handler has returned that a non-transient resource limit was
	// reached on the service side.
	ErrCodeServiceLimitExceededException = "ServiceLimitExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied due to request throttling.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeTypeNotFoundException for service response error code
	// "TypeNotFoundException".
	//
	// The specified extension doesn't exist in the CloudFormation registry.
	ErrCodeTypeNotFoundException = "TypeNotFoundException"

	// ErrCodeUnsupportedActionException for service response error code
	// "UnsupportedActionException".
	//
	// The specified resource doesn't support this resource operation.
	ErrCodeUnsupportedActionException = "UnsupportedActionException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AlreadyExistsException":          newErrorAlreadyExistsException,
	"ClientTokenConflictException":    newErrorClientTokenConflictException,
	"ConcurrentModificationException": newErrorConcurrentModificationException,
	"ConcurrentOperationException":    newErrorConcurrentOperationException,
	"GeneralServiceException":         newErrorGeneralServiceException,
	"HandlerFailureException":         newErrorHandlerFailureException,
	"HandlerInternalFailureException": newErrorHandlerInternalFailureException,
	"InvalidCredentialsException":     newErrorInvalidCredentialsException,
	"InvalidRequestException":         newErrorInvalidRequestException,
	"NetworkFailureException":         newErrorNetworkFailureException,
	"NotStabilizedException":          newErrorNotStabilizedException,
	"NotUpdatableException":           newErrorNotUpdatableException,
	"PrivateTypeException":            newErrorPrivateTypeException,
	"RequestTokenNotFoundException":   newErrorRequestTokenNotFoundException,
	"ResourceConflictException":       newErrorResourceConflictException,
	"ResourceNotFoundException":       newErrorResourceNotFoundException,
	"ServiceInternalErrorException":   newErrorServiceInternalErrorException,
	"ServiceLimitExceededException":   newErrorServiceLimitExceededException,
	"ThrottlingException":             newErrorThrottlingException,
	"TypeNotFoundException":           newErrorTypeNotFoundException,
	"UnsupportedActionException":      newErrorUnsupportedActionException,
}
