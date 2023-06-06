// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appconfigdata

import (
	"fmt"

	"github.com/qtdslly/aws-sdk-go/aws"
	"github.com/qtdslly/aws-sdk-go/aws/awsutil"
	"github.com/qtdslly/aws-sdk-go/aws/request"
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const opGetLatestConfiguration = "GetLatestConfiguration"

// GetLatestConfigurationRequest generates a "aws/request.Request" representing the
// client's request for the GetLatestConfiguration operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetLatestConfiguration for more information on using the GetLatestConfiguration
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the GetLatestConfigurationRequest method.
//	req, resp := client.GetLatestConfigurationRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/appconfigdata-2021-11-11/GetLatestConfiguration
func (c *AppConfigData) GetLatestConfigurationRequest(input *GetLatestConfigurationInput) (req *request.Request, output *GetLatestConfigurationOutput) {
	op := &request.Operation{
		Name:       opGetLatestConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/configuration",
	}

	if input == nil {
		input = &GetLatestConfigurationInput{}
	}

	output = &GetLatestConfigurationOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetLatestConfiguration API operation for AWS AppConfig Data.
//
// Retrieves the latest deployed configuration. This API may return empty configuration
// data if the client already has the latest version. For more information about
// this API action and to view example CLI commands that show how to use it
// with the StartConfigurationSession API action, see Retrieving the configuration
// (http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-retrieving-the-configuration)
// in the AppConfig User Guide.
//
// Note the following important information.
//
//   - Each configuration token is only valid for one call to GetLatestConfiguration.
//     The GetLatestConfiguration response includes a NextPollConfigurationToken
//     that should always replace the token used for the just-completed call
//     in preparation for the next one.
//
//   - GetLatestConfiguration is a priced call. For more information, see Pricing
//     (https://aws.amazon.com/systems-manager/pricing/).
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS AppConfig Data's
// API operation GetLatestConfiguration for usage and error information.
//
// Returned Error Types:
//
//   - ThrottlingException
//     The request was denied due to request throttling.
//
//   - ResourceNotFoundException
//     The requested resource could not be found.
//
//   - BadRequestException
//     The input fails to satisfy the constraints specified by the service.
//
//   - InternalServerException
//     There was an internal failure in the service.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/appconfigdata-2021-11-11/GetLatestConfiguration
func (c *AppConfigData) GetLatestConfiguration(input *GetLatestConfigurationInput) (*GetLatestConfigurationOutput, error) {
	req, out := c.GetLatestConfigurationRequest(input)
	return out, req.Send()
}

// GetLatestConfigurationWithContext is the same as GetLatestConfiguration with the addition of
// the ability to pass a context and additional request options.
//
// See GetLatestConfiguration for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AppConfigData) GetLatestConfigurationWithContext(ctx aws.Context, input *GetLatestConfigurationInput, opts ...request.Option) (*GetLatestConfigurationOutput, error) {
	req, out := c.GetLatestConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStartConfigurationSession = "StartConfigurationSession"

// StartConfigurationSessionRequest generates a "aws/request.Request" representing the
// client's request for the StartConfigurationSession operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See StartConfigurationSession for more information on using the StartConfigurationSession
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the StartConfigurationSessionRequest method.
//	req, resp := client.StartConfigurationSessionRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/appconfigdata-2021-11-11/StartConfigurationSession
func (c *AppConfigData) StartConfigurationSessionRequest(input *StartConfigurationSessionInput) (req *request.Request, output *StartConfigurationSessionOutput) {
	op := &request.Operation{
		Name:       opStartConfigurationSession,
		HTTPMethod: "POST",
		HTTPPath:   "/configurationsessions",
	}

	if input == nil {
		input = &StartConfigurationSessionInput{}
	}

	output = &StartConfigurationSessionOutput{}
	req = c.newRequest(op, input, output)
	return
}

// StartConfigurationSession API operation for AWS AppConfig Data.
//
// Starts a configuration session used to retrieve a deployed configuration.
// For more information about this API action and to view example CLI commands
// that show how to use it with the GetLatestConfiguration API action, see Retrieving
// the configuration (http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-retrieving-the-configuration)
// in the AppConfig User Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS AppConfig Data's
// API operation StartConfigurationSession for usage and error information.
//
// Returned Error Types:
//
//   - ThrottlingException
//     The request was denied due to request throttling.
//
//   - ResourceNotFoundException
//     The requested resource could not be found.
//
//   - BadRequestException
//     The input fails to satisfy the constraints specified by the service.
//
//   - InternalServerException
//     There was an internal failure in the service.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/appconfigdata-2021-11-11/StartConfigurationSession
func (c *AppConfigData) StartConfigurationSession(input *StartConfigurationSessionInput) (*StartConfigurationSessionOutput, error) {
	req, out := c.StartConfigurationSessionRequest(input)
	return out, req.Send()
}

// StartConfigurationSessionWithContext is the same as StartConfigurationSession with the addition of
// the ability to pass a context and additional request options.
//
// See StartConfigurationSession for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AppConfigData) StartConfigurationSessionWithContext(ctx aws.Context, input *StartConfigurationSessionInput, opts ...request.Option) (*StartConfigurationSessionOutput, error) {
	req, out := c.StartConfigurationSessionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// Detailed information about the input that failed to satisfy the constraints
// specified by a call.
type BadRequestDetails struct {
	_ struct{} `type:"structure"`

	// One or more specified parameters are not valid for the call.
	InvalidParameters map[string]*InvalidParameterDetail `type:"map"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s BadRequestDetails) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s BadRequestDetails) GoString() string {
	return s.String()
}

// SetInvalidParameters sets the InvalidParameters field's value.
func (s *BadRequestDetails) SetInvalidParameters(v map[string]*InvalidParameterDetail) *BadRequestDetails {
	s.InvalidParameters = v
	return s
}

// The input fails to satisfy the constraints specified by the service.
type BadRequestException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	// Details describing why the request was invalid.
	Details *BadRequestDetails `type:"structure"`

	Message_ *string `locationName:"Message" type:"string"`

	// Code indicating the reason the request was invalid.
	Reason *string `type:"string" enum:"BadRequestReason"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s BadRequestException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s BadRequestException) GoString() string {
	return s.String()
}

func newErrorBadRequestException(v protocol.ResponseMetadata) error {
	return &BadRequestException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *BadRequestException) Code() string {
	return "BadRequestException"
}

// Message returns the exception's message.
func (s *BadRequestException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *BadRequestException) OrigErr() error {
	return nil
}

func (s *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s\n%s", s.Code(), s.Message(), s.String())
}

// Status code returns the HTTP status code for the request's response error.
func (s *BadRequestException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *BadRequestException) RequestID() string {
	return s.RespMetadata.RequestID
}

type GetLatestConfigurationInput struct {
	_ struct{} `type:"structure" nopayload:"true"`

	// Token describing the current state of the configuration session. To obtain
	// a token, first call the StartConfigurationSession API. Note that every call
	// to GetLatestConfiguration will return a new ConfigurationToken (NextPollConfigurationToken
	// in the response) and must be provided to subsequent GetLatestConfiguration
	// API calls.
	//
	// This token should only be used once. To support long poll use cases, the
	// token is valid for up to 24 hours. If a GetLatestConfiguration call uses
	// an expired token, the system returns BadRequestException.
	//
	// ConfigurationToken is a required field
	ConfigurationToken *string `location:"querystring" locationName:"configuration_token" type:"string" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GetLatestConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GetLatestConfigurationInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLatestConfigurationInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetLatestConfigurationInput"}
	if s.ConfigurationToken == nil {
		invalidParams.Add(request.NewErrParamRequired("ConfigurationToken"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetConfigurationToken sets the ConfigurationToken field's value.
func (s *GetLatestConfigurationInput) SetConfigurationToken(v string) *GetLatestConfigurationInput {
	s.ConfigurationToken = &v
	return s
}

type GetLatestConfigurationOutput struct {
	_ struct{} `type:"structure" payload:"Configuration"`

	// The data of the configuration. This may be empty if the client already has
	// the latest version of configuration.
	//
	// Configuration is a sensitive parameter and its value will be
	// replaced with "sensitive" in string returned by GetLatestConfigurationOutput's
	// String and GoString methods.
	Configuration []byte `type:"blob" sensitive:"true"`

	// A standard MIME type describing the format of the configuration content.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// The latest token describing the current state of the configuration session.
	// This must be provided to the next call to GetLatestConfiguration.
	//
	// This token should only be used once. To support long poll use cases, the
	// token is valid for up to 24 hours. If a GetLatestConfiguration call uses
	// an expired token, the system returns BadRequestException.
	NextPollConfigurationToken *string `location:"header" locationName:"Next-Poll-Configuration-Token" type:"string"`

	// The amount of time the client should wait before polling for configuration
	// updates again. Use RequiredMinimumPollIntervalInSeconds to set the desired
	// poll interval.
	NextPollIntervalInSeconds *int64 `location:"header" locationName:"Next-Poll-Interval-In-Seconds" type:"integer"`

	// The user-defined label for the AppConfig hosted configuration version. This
	// attribute doesn't apply if the configuration is not from an AppConfig hosted
	// configuration version. If the client already has the latest version of the
	// configuration data, this value is empty.
	VersionLabel *string `location:"header" locationName:"Version-Label" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GetLatestConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GetLatestConfigurationOutput) GoString() string {
	return s.String()
}

// SetConfiguration sets the Configuration field's value.
func (s *GetLatestConfigurationOutput) SetConfiguration(v []byte) *GetLatestConfigurationOutput {
	s.Configuration = v
	return s
}

// SetContentType sets the ContentType field's value.
func (s *GetLatestConfigurationOutput) SetContentType(v string) *GetLatestConfigurationOutput {
	s.ContentType = &v
	return s
}

// SetNextPollConfigurationToken sets the NextPollConfigurationToken field's value.
func (s *GetLatestConfigurationOutput) SetNextPollConfigurationToken(v string) *GetLatestConfigurationOutput {
	s.NextPollConfigurationToken = &v
	return s
}

// SetNextPollIntervalInSeconds sets the NextPollIntervalInSeconds field's value.
func (s *GetLatestConfigurationOutput) SetNextPollIntervalInSeconds(v int64) *GetLatestConfigurationOutput {
	s.NextPollIntervalInSeconds = &v
	return s
}

// SetVersionLabel sets the VersionLabel field's value.
func (s *GetLatestConfigurationOutput) SetVersionLabel(v string) *GetLatestConfigurationOutput {
	s.VersionLabel = &v
	return s
}

// There was an internal failure in the service.
type InternalServerException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InternalServerException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InternalServerException) GoString() string {
	return s.String()
}

func newErrorInternalServerException(v protocol.ResponseMetadata) error {
	return &InternalServerException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *InternalServerException) Code() string {
	return "InternalServerException"
}

// Message returns the exception's message.
func (s *InternalServerException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *InternalServerException) OrigErr() error {
	return nil
}

func (s *InternalServerException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *InternalServerException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *InternalServerException) RequestID() string {
	return s.RespMetadata.RequestID
}

// Information about an invalid parameter.
type InvalidParameterDetail struct {
	_ struct{} `type:"structure"`

	// The reason the parameter is invalid.
	Problem *string `type:"string" enum:"InvalidParameterProblem"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InvalidParameterDetail) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InvalidParameterDetail) GoString() string {
	return s.String()
}

// SetProblem sets the Problem field's value.
func (s *InvalidParameterDetail) SetProblem(v string) *InvalidParameterDetail {
	s.Problem = &v
	return s
}

// The requested resource could not be found.
type ResourceNotFoundException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`

	// A map indicating which parameters in the request reference the resource that
	// was not found.
	ReferencedBy map[string]*string `type:"map"`

	// The type of resource that was not found.
	ResourceType *string `type:"string" enum:"ResourceType"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ResourceNotFoundException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ResourceNotFoundException) GoString() string {
	return s.String()
}

func newErrorResourceNotFoundException(v protocol.ResponseMetadata) error {
	return &ResourceNotFoundException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ResourceNotFoundException) Code() string {
	return "ResourceNotFoundException"
}

// Message returns the exception's message.
func (s *ResourceNotFoundException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ResourceNotFoundException) OrigErr() error {
	return nil
}

func (s *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s\n%s", s.Code(), s.Message(), s.String())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ResourceNotFoundException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ResourceNotFoundException) RequestID() string {
	return s.RespMetadata.RequestID
}

type StartConfigurationSessionInput struct {
	_ struct{} `type:"structure"`

	// The application ID or the application name.
	//
	// ApplicationIdentifier is a required field
	ApplicationIdentifier *string `min:"1" type:"string" required:"true"`

	// The configuration profile ID or the configuration profile name.
	//
	// ConfigurationProfileIdentifier is a required field
	ConfigurationProfileIdentifier *string `min:"1" type:"string" required:"true"`

	// The environment ID or the environment name.
	//
	// EnvironmentIdentifier is a required field
	EnvironmentIdentifier *string `min:"1" type:"string" required:"true"`

	// Sets a constraint on a session. If you specify a value of, for example, 60
	// seconds, then the client that established the session can't call GetLatestConfiguration
	// more frequently than every 60 seconds.
	RequiredMinimumPollIntervalInSeconds *int64 `min:"15" type:"integer"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s StartConfigurationSessionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s StartConfigurationSessionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartConfigurationSessionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "StartConfigurationSessionInput"}
	if s.ApplicationIdentifier == nil {
		invalidParams.Add(request.NewErrParamRequired("ApplicationIdentifier"))
	}
	if s.ApplicationIdentifier != nil && len(*s.ApplicationIdentifier) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ApplicationIdentifier", 1))
	}
	if s.ConfigurationProfileIdentifier == nil {
		invalidParams.Add(request.NewErrParamRequired("ConfigurationProfileIdentifier"))
	}
	if s.ConfigurationProfileIdentifier != nil && len(*s.ConfigurationProfileIdentifier) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ConfigurationProfileIdentifier", 1))
	}
	if s.EnvironmentIdentifier == nil {
		invalidParams.Add(request.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if s.EnvironmentIdentifier != nil && len(*s.EnvironmentIdentifier) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("EnvironmentIdentifier", 1))
	}
	if s.RequiredMinimumPollIntervalInSeconds != nil && *s.RequiredMinimumPollIntervalInSeconds < 15 {
		invalidParams.Add(request.NewErrParamMinValue("RequiredMinimumPollIntervalInSeconds", 15))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetApplicationIdentifier sets the ApplicationIdentifier field's value.
func (s *StartConfigurationSessionInput) SetApplicationIdentifier(v string) *StartConfigurationSessionInput {
	s.ApplicationIdentifier = &v
	return s
}

// SetConfigurationProfileIdentifier sets the ConfigurationProfileIdentifier field's value.
func (s *StartConfigurationSessionInput) SetConfigurationProfileIdentifier(v string) *StartConfigurationSessionInput {
	s.ConfigurationProfileIdentifier = &v
	return s
}

// SetEnvironmentIdentifier sets the EnvironmentIdentifier field's value.
func (s *StartConfigurationSessionInput) SetEnvironmentIdentifier(v string) *StartConfigurationSessionInput {
	s.EnvironmentIdentifier = &v
	return s
}

// SetRequiredMinimumPollIntervalInSeconds sets the RequiredMinimumPollIntervalInSeconds field's value.
func (s *StartConfigurationSessionInput) SetRequiredMinimumPollIntervalInSeconds(v int64) *StartConfigurationSessionInput {
	s.RequiredMinimumPollIntervalInSeconds = &v
	return s
}

type StartConfigurationSessionOutput struct {
	_ struct{} `type:"structure"`

	// Token encapsulating state about the configuration session. Provide this token
	// to the GetLatestConfiguration API to retrieve configuration data.
	//
	// This token should only be used once in your first call to GetLatestConfiguration.
	// You must use the new token in the GetLatestConfiguration response (NextPollConfigurationToken)
	// in each subsequent call to GetLatestConfiguration.
	//
	// The InitialConfigurationToken and NextPollConfigurationToken should only
	// be used once. To support long poll use cases, the tokens are valid for up
	// to 24 hours. If a GetLatestConfiguration call uses an expired token, the
	// system returns BadRequestException.
	InitialConfigurationToken *string `type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s StartConfigurationSessionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s StartConfigurationSessionOutput) GoString() string {
	return s.String()
}

// SetInitialConfigurationToken sets the InitialConfigurationToken field's value.
func (s *StartConfigurationSessionOutput) SetInitialConfigurationToken(v string) *StartConfigurationSessionOutput {
	s.InitialConfigurationToken = &v
	return s
}

// The request was denied due to request throttling.
type ThrottlingException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ThrottlingException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ThrottlingException) GoString() string {
	return s.String()
}

func newErrorThrottlingException(v protocol.ResponseMetadata) error {
	return &ThrottlingException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ThrottlingException) Code() string {
	return "ThrottlingException"
}

// Message returns the exception's message.
func (s *ThrottlingException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ThrottlingException) OrigErr() error {
	return nil
}

func (s *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ThrottlingException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ThrottlingException) RequestID() string {
	return s.RespMetadata.RequestID
}

const (
	// BadRequestReasonInvalidParameters is a BadRequestReason enum value
	BadRequestReasonInvalidParameters = "InvalidParameters"
)

// BadRequestReason_Values returns all elements of the BadRequestReason enum
func BadRequestReason_Values() []string {
	return []string{
		BadRequestReasonInvalidParameters,
	}
}

const (
	// InvalidParameterProblemCorrupted is a InvalidParameterProblem enum value
	InvalidParameterProblemCorrupted = "Corrupted"

	// InvalidParameterProblemExpired is a InvalidParameterProblem enum value
	InvalidParameterProblemExpired = "Expired"

	// InvalidParameterProblemPollIntervalNotSatisfied is a InvalidParameterProblem enum value
	InvalidParameterProblemPollIntervalNotSatisfied = "PollIntervalNotSatisfied"
)

// InvalidParameterProblem_Values returns all elements of the InvalidParameterProblem enum
func InvalidParameterProblem_Values() []string {
	return []string{
		InvalidParameterProblemCorrupted,
		InvalidParameterProblemExpired,
		InvalidParameterProblemPollIntervalNotSatisfied,
	}
}

const (
	// ResourceTypeApplication is a ResourceType enum value
	ResourceTypeApplication = "Application"

	// ResourceTypeConfigurationProfile is a ResourceType enum value
	ResourceTypeConfigurationProfile = "ConfigurationProfile"

	// ResourceTypeDeployment is a ResourceType enum value
	ResourceTypeDeployment = "Deployment"

	// ResourceTypeEnvironment is a ResourceType enum value
	ResourceTypeEnvironment = "Environment"

	// ResourceTypeConfiguration is a ResourceType enum value
	ResourceTypeConfiguration = "Configuration"
)

// ResourceType_Values returns all elements of the ResourceType enum
func ResourceType_Values() []string {
	return []string{
		ResourceTypeApplication,
		ResourceTypeConfigurationProfile,
		ResourceTypeDeployment,
		ResourceTypeEnvironment,
		ResourceTypeConfiguration,
	}
}
