// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplacecommerceanalytics

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeException for service response error code
	// "MarketplaceCommerceAnalyticsException".
	//
	// This exception is thrown when an internal service error occurs.
	ErrCodeException = "MarketplaceCommerceAnalyticsException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"MarketplaceCommerceAnalyticsException": newErrorException,
}
