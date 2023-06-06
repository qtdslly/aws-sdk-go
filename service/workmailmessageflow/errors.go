// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmailmessageflow

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeInvalidContentLocation for service response error code
	// "InvalidContentLocation".
	//
	// WorkMail could not access the updated email content. Possible reasons:
	//
	//    * You made the request in a region other than your S3 bucket region.
	//
	//    * The S3 bucket owner (https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-owner-condition.html)
	//    is not the same as the calling AWS account.
	//
	//    * You have an incomplete or missing S3 bucket policy. For more information
	//    about policies, see Updating message content with AWS Lambda (https://docs.aws.amazon.com/workmail/latest/adminguide/update-with-lambda.html)
	//    in the WorkMail Administrator Guide.
	ErrCodeInvalidContentLocation = "InvalidContentLocation"

	// ErrCodeMessageFrozen for service response error code
	// "MessageFrozen".
	//
	// The requested email is not eligible for update. This is usually the case
	// for a redirected email.
	ErrCodeMessageFrozen = "MessageFrozen"

	// ErrCodeMessageRejected for service response error code
	// "MessageRejected".
	//
	// The requested email could not be updated due to an error in the MIME content.
	// Check the error message for more information about what caused the error.
	ErrCodeMessageRejected = "MessageRejected"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The requested email message is not found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InvalidContentLocation":    newErrorInvalidContentLocation,
	"MessageFrozen":             newErrorMessageFrozen,
	"MessageRejected":           newErrorMessageRejected,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
}