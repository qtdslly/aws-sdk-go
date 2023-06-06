// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package snowball

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeClusterLimitExceededException for service response error code
	// "ClusterLimitExceededException".
	//
	// Job creation failed. Currently, clusters support five nodes. If you have
	// fewer than five nodes for your cluster and you have more nodes to create
	// for this cluster, try again and create jobs until your cluster has exactly
	// five nodes.
	ErrCodeClusterLimitExceededException = "ClusterLimitExceededException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// You get this exception when you call CreateReturnShippingLabel more than
	// once when other requests are not completed.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeEc2RequestFailedException for service response error code
	// "Ec2RequestFailedException".
	//
	// Your user lacks the necessary Amazon EC2 permissions to perform the attempted
	// action.
	ErrCodeEc2RequestFailedException = "Ec2RequestFailedException"

	// ErrCodeInvalidAddressException for service response error code
	// "InvalidAddressException".
	//
	// The address provided was invalid. Check the address with your region's carrier,
	// and try again.
	ErrCodeInvalidAddressException = "InvalidAddressException"

	// ErrCodeInvalidInputCombinationException for service response error code
	// "InvalidInputCombinationException".
	//
	// Job or cluster creation failed. One or more inputs were invalid. Confirm
	// that the CreateClusterRequest$SnowballType value supports your CreateJobRequest$JobType,
	// and try again.
	ErrCodeInvalidInputCombinationException = "InvalidInputCombinationException"

	// ErrCodeInvalidJobStateException for service response error code
	// "InvalidJobStateException".
	//
	// The action can't be performed because the job's current state doesn't allow
	// that action to be performed.
	ErrCodeInvalidJobStateException = "InvalidJobStateException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The NextToken string was altered unexpectedly, and the operation has stopped.
	// Run the operation without changing the NextToken string, and try again.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidResourceException for service response error code
	// "InvalidResourceException".
	//
	// The specified resource can't be found. Check the information you provided
	// in your last request, and try again.
	ErrCodeInvalidResourceException = "InvalidResourceException"

	// ErrCodeKMSRequestFailedException for service response error code
	// "KMSRequestFailedException".
	//
	// The provided Key Management Service key lacks the permissions to perform
	// the specified CreateJob or UpdateJob action.
	ErrCodeKMSRequestFailedException = "KMSRequestFailedException"

	// ErrCodeReturnShippingLabelAlreadyExistsException for service response error code
	// "ReturnShippingLabelAlreadyExistsException".
	//
	// You get this exception if you call CreateReturnShippingLabel and a valid
	// return shipping label already exists. In this case, use DescribeReturnShippingLabel
	// to get the URL.
	ErrCodeReturnShippingLabelAlreadyExistsException = "ReturnShippingLabelAlreadyExistsException"

	// ErrCodeUnsupportedAddressException for service response error code
	// "UnsupportedAddressException".
	//
	// The address is either outside the serviceable area for your region, or an
	// error occurred. Check the address with your region's carrier and try again.
	// If the issue persists, contact Amazon Web Services Support.
	ErrCodeUnsupportedAddressException = "UnsupportedAddressException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ClusterLimitExceededException":             newErrorClusterLimitExceededException,
	"ConflictException":                         newErrorConflictException,
	"Ec2RequestFailedException":                 newErrorEc2RequestFailedException,
	"InvalidAddressException":                   newErrorInvalidAddressException,
	"InvalidInputCombinationException":          newErrorInvalidInputCombinationException,
	"InvalidJobStateException":                  newErrorInvalidJobStateException,
	"InvalidNextTokenException":                 newErrorInvalidNextTokenException,
	"InvalidResourceException":                  newErrorInvalidResourceException,
	"KMSRequestFailedException":                 newErrorKMSRequestFailedException,
	"ReturnShippingLabelAlreadyExistsException": newErrorReturnShippingLabelAlreadyExistsException,
	"UnsupportedAddressException":               newErrorUnsupportedAddressException,
}
