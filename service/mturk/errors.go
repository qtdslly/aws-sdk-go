// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeRequestError for service response error code
	// "RequestError".
	//
	// Your request is invalid.
	ErrCodeRequestError = "RequestError"

	// ErrCodeServiceFault for service response error code
	// "ServiceFault".
	//
	// Amazon Mechanical Turk is temporarily unable to process your request. Try
	// your call again.
	ErrCodeServiceFault = "ServiceFault"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"RequestError": newErrorRequestError,
	"ServiceFault": newErrorServiceFault,
}