// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acmpca

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeCertificateMismatchException for service response error code
	// "CertificateMismatchException".
	//
	// The certificate authority certificate you are importing does not comply with
	// conditions specified in the certificate that signed it.
	ErrCodeCertificateMismatchException = "CertificateMismatchException"

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// A previous update to your private CA is still ongoing.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeInvalidArgsException for service response error code
	// "InvalidArgsException".
	//
	// One or more of the specified arguments was not valid.
	ErrCodeInvalidArgsException = "InvalidArgsException"

	// ErrCodeInvalidArnException for service response error code
	// "InvalidArnException".
	//
	// The requested Amazon Resource Name (ARN) does not refer to an existing resource.
	ErrCodeInvalidArnException = "InvalidArnException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The token specified in the NextToken argument is not valid. Use the token
	// returned from your previous call to ListCertificateAuthorities (https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListCertificateAuthorities.html).
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidPolicyException for service response error code
	// "InvalidPolicyException".
	//
	// The resource policy is invalid or is missing a required statement. For general
	// information about IAM policy and statement structure, see Overview of JSON
	// Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policies-json).
	ErrCodeInvalidPolicyException = "InvalidPolicyException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// The request action cannot be performed or is prohibited.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeInvalidStateException for service response error code
	// "InvalidStateException".
	//
	// The state of the private CA does not allow this action to occur.
	ErrCodeInvalidStateException = "InvalidStateException"

	// ErrCodeInvalidTagException for service response error code
	// "InvalidTagException".
	//
	// The tag associated with the CA is not valid. The invalid argument is contained
	// in the message field.
	ErrCodeInvalidTagException = "InvalidTagException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// An Amazon Web Services Private CA quota has been exceeded. See the exception
	// message returned to determine the quota that was exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeLockoutPreventedException for service response error code
	// "LockoutPreventedException".
	//
	// The current action was prevented because it would lock the caller out from
	// performing subsequent actions. Verify that the specified parameters would
	// not result in the caller being denied access to the resource.
	ErrCodeLockoutPreventedException = "LockoutPreventedException"

	// ErrCodeMalformedCSRException for service response error code
	// "MalformedCSRException".
	//
	// The certificate signing request is invalid.
	ErrCodeMalformedCSRException = "MalformedCSRException"

	// ErrCodeMalformedCertificateException for service response error code
	// "MalformedCertificateException".
	//
	// One or more fields in the certificate are invalid.
	ErrCodeMalformedCertificateException = "MalformedCertificateException"

	// ErrCodePermissionAlreadyExistsException for service response error code
	// "PermissionAlreadyExistsException".
	//
	// The designated permission has already been given to the user.
	ErrCodePermissionAlreadyExistsException = "PermissionAlreadyExistsException"

	// ErrCodeRequestAlreadyProcessedException for service response error code
	// "RequestAlreadyProcessedException".
	//
	// Your request has already been completed.
	ErrCodeRequestAlreadyProcessedException = "RequestAlreadyProcessedException"

	// ErrCodeRequestFailedException for service response error code
	// "RequestFailedException".
	//
	// The request has failed for an unspecified reason.
	ErrCodeRequestFailedException = "RequestFailedException"

	// ErrCodeRequestInProgressException for service response error code
	// "RequestInProgressException".
	//
	// Your request is already in progress.
	ErrCodeRequestInProgressException = "RequestInProgressException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// A resource such as a private CA, S3 bucket, certificate, audit report, or
	// policy cannot be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTagsException".
	//
	// You can associate up to 50 tags with a private CA. Exception information
	// is contained in the exception message field.
	ErrCodeTooManyTagsException = "TooManyTagsException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"CertificateMismatchException":     newErrorCertificateMismatchException,
	"ConcurrentModificationException":  newErrorConcurrentModificationException,
	"InvalidArgsException":             newErrorInvalidArgsException,
	"InvalidArnException":              newErrorInvalidArnException,
	"InvalidNextTokenException":        newErrorInvalidNextTokenException,
	"InvalidPolicyException":           newErrorInvalidPolicyException,
	"InvalidRequestException":          newErrorInvalidRequestException,
	"InvalidStateException":            newErrorInvalidStateException,
	"InvalidTagException":              newErrorInvalidTagException,
	"LimitExceededException":           newErrorLimitExceededException,
	"LockoutPreventedException":        newErrorLockoutPreventedException,
	"MalformedCSRException":            newErrorMalformedCSRException,
	"MalformedCertificateException":    newErrorMalformedCertificateException,
	"PermissionAlreadyExistsException": newErrorPermissionAlreadyExistsException,
	"RequestAlreadyProcessedException": newErrorRequestAlreadyProcessedException,
	"RequestFailedException":           newErrorRequestFailedException,
	"RequestInProgressException":       newErrorRequestInProgressException,
	"ResourceNotFoundException":        newErrorResourceNotFoundException,
	"TooManyTagsException":             newErrorTooManyTagsException,
}
