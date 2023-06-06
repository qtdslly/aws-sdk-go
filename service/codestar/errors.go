// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestar

import (
	"github.com/qtdslly/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// Another modification is being made. That modification must complete before
	// you can make your change.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The next token is not valid.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidServiceRoleException for service response error code
	// "InvalidServiceRoleException".
	//
	// The service role is not valid.
	ErrCodeInvalidServiceRoleException = "InvalidServiceRoleException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// A resource limit has been exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeProjectAlreadyExistsException for service response error code
	// "ProjectAlreadyExistsException".
	//
	// An AWS CodeStar project with the same ID already exists in this region for
	// the AWS account. AWS CodeStar project IDs must be unique within a region
	// for the AWS account.
	ErrCodeProjectAlreadyExistsException = "ProjectAlreadyExistsException"

	// ErrCodeProjectConfigurationException for service response error code
	// "ProjectConfigurationException".
	//
	// Project configuration information is required but not specified.
	ErrCodeProjectConfigurationException = "ProjectConfigurationException"

	// ErrCodeProjectCreationFailedException for service response error code
	// "ProjectCreationFailedException".
	//
	// The project creation request was valid, but a nonspecific exception or error
	// occurred during project creation. The project could not be created in AWS
	// CodeStar.
	ErrCodeProjectCreationFailedException = "ProjectCreationFailedException"

	// ErrCodeProjectNotFoundException for service response error code
	// "ProjectNotFoundException".
	//
	// The specified AWS CodeStar project was not found.
	ErrCodeProjectNotFoundException = "ProjectNotFoundException"

	// ErrCodeTeamMemberAlreadyAssociatedException for service response error code
	// "TeamMemberAlreadyAssociatedException".
	//
	// The team member is already associated with a role in this project.
	ErrCodeTeamMemberAlreadyAssociatedException = "TeamMemberAlreadyAssociatedException"

	// ErrCodeTeamMemberNotFoundException for service response error code
	// "TeamMemberNotFoundException".
	//
	// The specified team member was not found.
	ErrCodeTeamMemberNotFoundException = "TeamMemberNotFoundException"

	// ErrCodeUserProfileAlreadyExistsException for service response error code
	// "UserProfileAlreadyExistsException".
	//
	// A user profile with that name already exists in this region for the AWS account.
	// AWS CodeStar user profile names must be unique within a region for the AWS
	// account.
	ErrCodeUserProfileAlreadyExistsException = "UserProfileAlreadyExistsException"

	// ErrCodeUserProfileNotFoundException for service response error code
	// "UserProfileNotFoundException".
	//
	// The user profile was not found.
	ErrCodeUserProfileNotFoundException = "UserProfileNotFoundException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The specified input is either not valid, or it could not be validated.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ConcurrentModificationException":      newErrorConcurrentModificationException,
	"InvalidNextTokenException":            newErrorInvalidNextTokenException,
	"InvalidServiceRoleException":          newErrorInvalidServiceRoleException,
	"LimitExceededException":               newErrorLimitExceededException,
	"ProjectAlreadyExistsException":        newErrorProjectAlreadyExistsException,
	"ProjectConfigurationException":        newErrorProjectConfigurationException,
	"ProjectCreationFailedException":       newErrorProjectCreationFailedException,
	"ProjectNotFoundException":             newErrorProjectNotFoundException,
	"TeamMemberAlreadyAssociatedException": newErrorTeamMemberAlreadyAssociatedException,
	"TeamMemberNotFoundException":          newErrorTeamMemberNotFoundException,
	"UserProfileAlreadyExistsException":    newErrorUserProfileAlreadyExistsException,
	"UserProfileNotFoundException":         newErrorUserProfileNotFoundException,
	"ValidationException":                  newErrorValidationException,
}
