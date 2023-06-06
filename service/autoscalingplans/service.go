// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscalingplans

import (
	"github.com/qtdslly/aws-sdk-go/aws"
	"github.com/qtdslly/aws-sdk-go/aws/client"
	"github.com/qtdslly/aws-sdk-go/aws/client/metadata"
	"github.com/qtdslly/aws-sdk-go/aws/request"
	"github.com/qtdslly/aws-sdk-go/aws/signer/v4"
	"github.com/qtdslly/aws-sdk-go/private/protocol"
	"github.com/qtdslly/aws-sdk-go/private/protocol/jsonrpc"
)

// AutoScalingPlans provides the API operation methods for making requests to
// AWS Auto Scaling Plans. See this package's package overview docs
// for details on the service.
//
// AutoScalingPlans methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type AutoScalingPlans struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "autoscaling"        // Name of service.
	EndpointsID = "autoscaling-plans"  // ID to lookup a service endpoint with.
	ServiceID   = "Auto Scaling Plans" // ServiceID is a unique identifier of a specific service.
)

// New creates a new instance of the AutoScalingPlans client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//
//	mySession := session.Must(session.NewSession())
//
//	// Create a AutoScalingPlans client from just a session.
//	svc := autoscalingplans.New(mySession)
//
//	// Create a AutoScalingPlans client with additional configuration
//	svc := autoscalingplans.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *AutoScalingPlans {
	c := p.ClientConfig(EndpointsID, cfgs...)
	if c.SigningNameDerived || len(c.SigningName) == 0 {
		c.SigningName = "autoscaling-plans"
	}
	return newClient(*c.Config, c.Handlers, c.PartitionID, c.Endpoint, c.SigningRegion, c.SigningName, c.ResolvedRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, partitionID, endpoint, signingRegion, signingName, resolvedRegion string) *AutoScalingPlans {
	svc := &AutoScalingPlans{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:    ServiceName,
				ServiceID:      ServiceID,
				SigningName:    signingName,
				SigningRegion:  signingRegion,
				PartitionID:    partitionID,
				Endpoint:       endpoint,
				APIVersion:     "2018-01-06",
				ResolvedRegion: resolvedRegion,
				JSONVersion:    "1.1",
				TargetPrefix:   "AnyScaleScalingPlannerFrontendService",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(
		protocol.NewUnmarshalErrorHandler(jsonrpc.NewUnmarshalTypedError(exceptionFromCode)).NamedHandler(),
	)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a AutoScalingPlans operation and runs any
// custom request initialization.
func (c *AutoScalingPlans) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
