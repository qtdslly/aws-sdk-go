// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeDuplicateResourceException for service response error code
	// "DuplicateResourceException".
	//
	// The specified resource is a duplicate.
	ErrCodeDuplicateResourceException = "DuplicateResourceException"

	// ErrCodeInvalidParametersException for service response error code
	// "InvalidParametersException".
	//
	// One or more parameters provided to the operation are not valid.
	ErrCodeInvalidParametersException = "InvalidParametersException"

	// ErrCodeInvalidStateException for service response error code
	// "InvalidStateException".
	//
	// An attempt was made to modify a resource that is in a state that is not valid.
	// Check your resources to ensure that they are in valid states before retrying
	// the operation.
	ErrCodeInvalidStateException = "InvalidStateException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The current limits of the service would have been exceeded by this operation.
	// Decrease your resource use or increase your service limits and retry the
	// operation.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeOperationNotSupportedException for service response error code
	// "OperationNotSupportedException".
	//
	// The operation is not supported.
	ErrCodeOperationNotSupportedException = "OperationNotSupportedException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// A resource that is currently in use. Ensure that the resource is not in use
	// and retry the operation.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource was not found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTagOptionNotMigratedException for service response error code
	// "TagOptionNotMigratedException".
	//
	// An operation requiring TagOptions failed because the TagOptions migration
	// process has not been performed for this account. Use the Amazon Web Services
	// Management Console to perform the migration process before retrying the operation.
	ErrCodeTagOptionNotMigratedException = "TagOptionNotMigratedException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"DuplicateResourceException":     newErrorDuplicateResourceException,
	"InvalidParametersException":     newErrorInvalidParametersException,
	"InvalidStateException":          newErrorInvalidStateException,
	"LimitExceededException":         newErrorLimitExceededException,
	"OperationNotSupportedException": newErrorOperationNotSupportedException,
	"ResourceInUseException":         newErrorResourceInUseException,
	"ResourceNotFoundException":      newErrorResourceNotFoundException,
	"TagOptionNotMigratedException":  newErrorTagOptionNotMigratedException,
}
