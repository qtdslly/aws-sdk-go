// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

//go:build go1.16 && integration
// +build go1.16,integration

package waf_test

import (
	"context"
	"testing"
	"time"

	"github.com/qtdslly/aws-sdk-go/aws"
	"github.com/qtdslly/aws-sdk-go/aws/awserr"
	"github.com/qtdslly/aws-sdk-go/aws/request"
	"github.com/qtdslly/aws-sdk-go/awstesting/integration"
	"github.com/qtdslly/aws-sdk-go/service/waf"
)

var _ aws.Config
var _ awserr.Error
var _ request.Request

func TestInteg_00_ListRules(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-east-1")
	svc := waf.New(sess)
	params := &waf.ListRulesInput{
		Limit: aws.Int64(20),
	}
	_, err := svc.ListRulesWithContext(ctx, params, func(r *request.Request) {
		r.Handlers.Validate.RemoveByName("core.ValidateParametersHandler")
	})
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
