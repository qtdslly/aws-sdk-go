// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdbelastic

import (
	"github.com/qtdslly/aws-sdk-go/aws"
	"github.com/qtdslly/aws-sdk-go/aws/client"
	"github.com/qtdslly/aws-sdk-go/aws/client/metadata"
	"github.com/qtdslly/aws-sdk-go/aws/request"
	"github.com/qtdslly/aws-sdk-go/aws/signer/v4"
	"github.com/qtdslly/aws-sdk-go/private/protocol"
	"github.com/qtdslly/aws-sdk-go/private/protocol/restjson"
)

// DocDBElastic provides the API operation methods for making requests to
// Amazon DocumentDB Elastic Clusters. See this package's package overview docs
// for details on the service.
//
// DocDBElastic methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type DocDBElastic struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "DocDB Elastic" // Name of service.
	EndpointsID = "docdb-elastic" // ID to lookup a service endpoint with.
	ServiceID   = "DocDB Elastic" // ServiceID is a unique identifier of a specific service.
)

// New creates a new instance of the DocDBElastic client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//
//	mySession := session.Must(session.NewSession())
//
//	// Create a DocDBElastic client from just a session.
//	svc := docdbelastic.New(mySession)
//
//	// Create a DocDBElastic client with additional configuration
//	svc := docdbelastic.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *DocDBElastic {
	c := p.ClientConfig(EndpointsID, cfgs...)
	if c.SigningNameDerived || len(c.SigningName) == 0 {
		c.SigningName = "docdb-elastic"
	}
	return newClient(*c.Config, c.Handlers, c.PartitionID, c.Endpoint, c.SigningRegion, c.SigningName, c.ResolvedRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, partitionID, endpoint, signingRegion, signingName, resolvedRegion string) *DocDBElastic {
	svc := &DocDBElastic{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:    ServiceName,
				ServiceID:      ServiceID,
				SigningName:    signingName,
				SigningRegion:  signingRegion,
				PartitionID:    partitionID,
				Endpoint:       endpoint,
				APIVersion:     "2022-11-28",
				ResolvedRegion: resolvedRegion,
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(
		protocol.NewUnmarshalErrorHandler(restjson.NewUnmarshalTypedError(exceptionFromCode)).NamedHandler(),
	)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a DocDBElastic operation and runs any
// custom request initialization.
func (c *DocDBElastic) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
