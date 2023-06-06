// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationautoscaling_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/qtdslly/aws-sdk-go/aws"
	"github.com/qtdslly/aws-sdk-go/aws/awserr"
	"github.com/qtdslly/aws-sdk-go/aws/session"
	"github.com/qtdslly/aws-sdk-go/service/applicationautoscaling"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To delete a scaling policy
// This example deletes a scaling policy for the Amazon ECS service called web-app,
// which is running in the default cluster.
func ExampleApplicationAutoScaling_DeleteScalingPolicy_shared00() {
	svc := applicationautoscaling.New(session.New())
	input := &applicationautoscaling.DeleteScalingPolicyInput{
		PolicyName:        aws.String("web-app-cpu-lt-25"),
		ResourceId:        aws.String("service/default/web-app"),
		ScalableDimension: aws.String("ecs:service:DesiredCount"),
		ServiceNamespace:  aws.String("ecs"),
	}

	result, err := svc.DeleteScalingPolicy(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeObjectNotFoundException:
				fmt.Println(applicationautoscaling.ErrCodeObjectNotFoundException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete a scheduled action
// This example deletes a scheduled action for the AppStream 2.0 fleet called sample-fleet.
func ExampleApplicationAutoScaling_DeleteScheduledAction_shared00() {
	svc := applicationautoscaling.New(session.New())
	input := &applicationautoscaling.DeleteScheduledActionInput{
		ResourceId:          aws.String("fleet/sample-fleet"),
		ScalableDimension:   aws.String("appstream:fleet:DesiredCapacity"),
		ScheduledActionName: aws.String("my-recurring-action"),
		ServiceNamespace:    aws.String("appstream"),
	}

	result, err := svc.DeleteScheduledAction(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeObjectNotFoundException:
				fmt.Println(applicationautoscaling.ErrCodeObjectNotFoundException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To deregister a scalable target
// This example deregisters a scalable target for an Amazon ECS service called web-app
// that is running in the default cluster.
func ExampleApplicationAutoScaling_DeregisterScalableTarget_shared00() {
	svc := applicationautoscaling.New(session.New())
	input := &applicationautoscaling.DeregisterScalableTargetInput{
		ResourceId:        aws.String("service/default/web-app"),
		ScalableDimension: aws.String("ecs:service:DesiredCount"),
		ServiceNamespace:  aws.String("ecs"),
	}

	result, err := svc.DeregisterScalableTarget(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeObjectNotFoundException:
				fmt.Println(applicationautoscaling.ErrCodeObjectNotFoundException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe scalable targets
// This example describes the scalable targets for the ECS service namespace.
func ExampleApplicationAutoScaling_DescribeScalableTargets_shared00() {
	svc := applicationautoscaling.New(session.New())
	input := &applicationautoscaling.DescribeScalableTargetsInput{
		ServiceNamespace: aws.String("ecs"),
	}

	result, err := svc.DescribeScalableTargets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeInvalidNextTokenException:
				fmt.Println(applicationautoscaling.ErrCodeInvalidNextTokenException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe scaling activities for a scalable target
// This example describes the scaling activities for an Amazon ECS service called web-app
// that is running in the default cluster.
func ExampleApplicationAutoScaling_DescribeScalingActivities_shared00() {
	svc := applicationautoscaling.New(session.New())
	input := &applicationautoscaling.DescribeScalingActivitiesInput{
		ResourceId:        aws.String("service/default/web-app"),
		ScalableDimension: aws.String("ecs:service:DesiredCount"),
		ServiceNamespace:  aws.String("ecs"),
	}

	result, err := svc.DescribeScalingActivities(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeInvalidNextTokenException:
				fmt.Println(applicationautoscaling.ErrCodeInvalidNextTokenException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe scaling policies
// This example describes the scaling policies for the ECS service namespace.
func ExampleApplicationAutoScaling_DescribeScalingPolicies_shared00() {
	svc := applicationautoscaling.New(session.New())
	input := &applicationautoscaling.DescribeScalingPoliciesInput{
		ServiceNamespace: aws.String("ecs"),
	}

	result, err := svc.DescribeScalingPolicies(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeFailedResourceAccessException:
				fmt.Println(applicationautoscaling.ErrCodeFailedResourceAccessException, aerr.Error())
			case applicationautoscaling.ErrCodeInvalidNextTokenException:
				fmt.Println(applicationautoscaling.ErrCodeInvalidNextTokenException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe scheduled actions
// This example describes the scheduled actions for the dynamodb service namespace.
func ExampleApplicationAutoScaling_DescribeScheduledActions_shared00() {
	svc := applicationautoscaling.New(session.New())
	input := &applicationautoscaling.DescribeScheduledActionsInput{
		ServiceNamespace: aws.String("dynamodb"),
	}

	result, err := svc.DescribeScheduledActions(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeInvalidNextTokenException:
				fmt.Println(applicationautoscaling.ErrCodeInvalidNextTokenException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To list tags for a scalable target
// This example lists the tag key names and values that are attached to the scalable
// target specified by its ARN.
func ExampleApplicationAutoScaling_ListTagsForResource_shared00() {
	svc := applicationautoscaling.New(session.New())
	input := &applicationautoscaling.ListTagsForResourceInput{
		ResourceARN: aws.String("arn:aws:application-autoscaling:us-west-2:123456789012:scalable-target/1234abcd56ab78cd901ef1234567890ab123"),
	}

	result, err := svc.ListTagsForResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeResourceNotFoundException:
				fmt.Println(applicationautoscaling.ErrCodeResourceNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To apply a target tracking scaling policy with a predefined metric specification
// The following example applies a target tracking scaling policy with a predefined
// metric specification to an Amazon ECS service called web-app in the default cluster.
// The policy keeps the average CPU utilization of the service at 75 percent, with scale-out
// and scale-in cooldown periods of 60 seconds.
func ExampleApplicationAutoScaling_PutScalingPolicy_shared00() {
	svc := applicationautoscaling.New(session.New())
	input := &applicationautoscaling.PutScalingPolicyInput{
		PolicyName:        aws.String("cpu75-target-tracking-scaling-policy"),
		PolicyType:        aws.String("TargetTrackingScaling"),
		ResourceId:        aws.String("service/default/web-app"),
		ScalableDimension: aws.String("ecs:service:DesiredCount"),
		ServiceNamespace:  aws.String("ecs"),
		TargetTrackingScalingPolicyConfiguration: &applicationautoscaling.TargetTrackingScalingPolicyConfiguration{
			PredefinedMetricSpecification: &applicationautoscaling.PredefinedMetricSpecification{
				PredefinedMetricType: aws.String("ECSServiceAverageCPUUtilization"),
			},
			ScaleInCooldown:  aws.Int64(60),
			ScaleOutCooldown: aws.Int64(60),
			TargetValue:      aws.Float64(75.000000),
		},
	}

	result, err := svc.PutScalingPolicy(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeLimitExceededException:
				fmt.Println(applicationautoscaling.ErrCodeLimitExceededException, aerr.Error())
			case applicationautoscaling.ErrCodeObjectNotFoundException:
				fmt.Println(applicationautoscaling.ErrCodeObjectNotFoundException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeFailedResourceAccessException:
				fmt.Println(applicationautoscaling.ErrCodeFailedResourceAccessException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a recurring scheduled action
// This example adds a scheduled action to a DynamoDB table called TestTable to scale
// out on a recurring schedule. On the specified schedule (every day at 12:15pm UTC),
// if the current capacity is below the value specified for MinCapacity, Application
// Auto Scaling scales out to the value specified by MinCapacity.
func ExampleApplicationAutoScaling_PutScheduledAction_shared00() {
	svc := applicationautoscaling.New(session.New())
	input := &applicationautoscaling.PutScheduledActionInput{
		ResourceId:        aws.String("table/TestTable"),
		ScalableDimension: aws.String("dynamodb:table:WriteCapacityUnits"),
		ScalableTargetAction: &applicationautoscaling.ScalableTargetAction{
			MinCapacity: aws.Int64(6),
		},
		Schedule:            aws.String("cron(15 12 * * ? *)"),
		ScheduledActionName: aws.String("my-recurring-action"),
		ServiceNamespace:    aws.String("dynamodb"),
	}

	result, err := svc.PutScheduledAction(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeLimitExceededException:
				fmt.Println(applicationautoscaling.ErrCodeLimitExceededException, aerr.Error())
			case applicationautoscaling.ErrCodeObjectNotFoundException:
				fmt.Println(applicationautoscaling.ErrCodeObjectNotFoundException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To register an ECS service as a scalable target
// This example registers a scalable target from an Amazon ECS service called web-app
// that is running on the default cluster, with a minimum desired count of 1 task and
// a maximum desired count of 10 tasks.
func ExampleApplicationAutoScaling_RegisterScalableTarget_shared00() {
	svc := applicationautoscaling.New(session.New())
	input := &applicationautoscaling.RegisterScalableTargetInput{
		MaxCapacity:       aws.Int64(10),
		MinCapacity:       aws.Int64(1),
		ResourceId:        aws.String("service/default/web-app"),
		ScalableDimension: aws.String("ecs:service:DesiredCount"),
		ServiceNamespace:  aws.String("ecs"),
	}

	result, err := svc.RegisterScalableTarget(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeLimitExceededException:
				fmt.Println(applicationautoscaling.ErrCodeLimitExceededException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To add a tag to a scalable target
// This example adds a tag with the key name "environment" and the value "production"
// to the scalable target specified by its ARN.
func ExampleApplicationAutoScaling_TagResource_shared00() {
	svc := applicationautoscaling.New(session.New())
	input := &applicationautoscaling.TagResourceInput{
		ResourceARN: aws.String("arn:aws:application-autoscaling:us-west-2:123456789012:scalable-target/1234abcd56ab78cd901ef1234567890ab123"),
		Tags: map[string]*string{
			"environment": aws.String("production"),
		},
	}

	result, err := svc.TagResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeResourceNotFoundException:
				fmt.Println(applicationautoscaling.ErrCodeResourceNotFoundException, aerr.Error())
			case applicationautoscaling.ErrCodeTooManyTagsException:
				fmt.Println(applicationautoscaling.ErrCodeTooManyTagsException, aerr.Error())
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To remove a tag from a scalable target
// This example removes the tag pair with the key name "environment" from the scalable
// target specified by its ARN.
func ExampleApplicationAutoScaling_UntagResource_shared00() {
	svc := applicationautoscaling.New(session.New())
	input := &applicationautoscaling.UntagResourceInput{
		ResourceARN: aws.String("arn:aws:application-autoscaling:us-west-2:123456789012:scalable-target/1234abcd56ab78cd901ef1234567890ab123"),
		TagKeys: []*string{
			aws.String("environment"),
		},
	}

	result, err := svc.UntagResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeResourceNotFoundException:
				fmt.Println(applicationautoscaling.ErrCodeResourceNotFoundException, aerr.Error())
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}