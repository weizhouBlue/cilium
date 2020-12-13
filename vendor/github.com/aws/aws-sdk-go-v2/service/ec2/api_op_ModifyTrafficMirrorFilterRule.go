// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyTrafficMirrorFilterRuleInput struct {
	_ struct{} `type:"structure"`

	// The description to assign to the Traffic Mirror rule.
	Description *string `type:"string"`

	// The destination CIDR block to assign to the Traffic Mirror rule.
	DestinationCidrBlock *string `type:"string"`

	// The destination ports that are associated with the Traffic Mirror rule.
	DestinationPortRange *TrafficMirrorPortRangeRequest `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The protocol, for example TCP, to assign to the Traffic Mirror rule.
	Protocol *int64 `type:"integer"`

	// The properties that you want to remove from the Traffic Mirror filter rule.
	//
	// When you remove a property from a Traffic Mirror filter rule, the property
	// is set to the default.
	RemoveFields []TrafficMirrorFilterRuleField `locationName:"RemoveField" type:"list"`

	// The action to assign to the rule.
	RuleAction TrafficMirrorRuleAction `type:"string" enum:"true"`

	// The number of the Traffic Mirror rule. This number must be unique for each
	// Traffic Mirror rule in a given direction. The rules are processed in ascending
	// order by rule number.
	RuleNumber *int64 `type:"integer"`

	// The source CIDR block to assign to the Traffic Mirror rule.
	SourceCidrBlock *string `type:"string"`

	// The port range to assign to the Traffic Mirror rule.
	SourcePortRange *TrafficMirrorPortRangeRequest `type:"structure"`

	// The type of traffic (ingress | egress) to assign to the rule.
	TrafficDirection TrafficDirection `type:"string" enum:"true"`

	// The ID of the Traffic Mirror rule.
	//
	// TrafficMirrorFilterRuleId is a required field
	TrafficMirrorFilterRuleId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyTrafficMirrorFilterRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyTrafficMirrorFilterRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyTrafficMirrorFilterRuleInput"}

	if s.TrafficMirrorFilterRuleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrafficMirrorFilterRuleId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyTrafficMirrorFilterRuleOutput struct {
	_ struct{} `type:"structure"`

	// Modifies a Traffic Mirror rule.
	TrafficMirrorFilterRule *TrafficMirrorFilterRule `locationName:"trafficMirrorFilterRule" type:"structure"`
}

// String returns the string representation
func (s ModifyTrafficMirrorFilterRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyTrafficMirrorFilterRule = "ModifyTrafficMirrorFilterRule"

// ModifyTrafficMirrorFilterRuleRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the specified Traffic Mirror rule.
//
// DestinationCidrBlock and SourceCidrBlock must both be an IPv4 range or an
// IPv6 range.
//
//    // Example sending a request using ModifyTrafficMirrorFilterRuleRequest.
//    req := client.ModifyTrafficMirrorFilterRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyTrafficMirrorFilterRule
func (c *Client) ModifyTrafficMirrorFilterRuleRequest(input *ModifyTrafficMirrorFilterRuleInput) ModifyTrafficMirrorFilterRuleRequest {
	op := &aws.Operation{
		Name:       opModifyTrafficMirrorFilterRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyTrafficMirrorFilterRuleInput{}
	}

	req := c.newRequest(op, input, &ModifyTrafficMirrorFilterRuleOutput{})

	return ModifyTrafficMirrorFilterRuleRequest{Request: req, Input: input, Copy: c.ModifyTrafficMirrorFilterRuleRequest}
}

// ModifyTrafficMirrorFilterRuleRequest is the request type for the
// ModifyTrafficMirrorFilterRule API operation.
type ModifyTrafficMirrorFilterRuleRequest struct {
	*aws.Request
	Input *ModifyTrafficMirrorFilterRuleInput
	Copy  func(*ModifyTrafficMirrorFilterRuleInput) ModifyTrafficMirrorFilterRuleRequest
}

// Send marshals and sends the ModifyTrafficMirrorFilterRule API request.
func (r ModifyTrafficMirrorFilterRuleRequest) Send(ctx context.Context) (*ModifyTrafficMirrorFilterRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyTrafficMirrorFilterRuleResponse{
		ModifyTrafficMirrorFilterRuleOutput: r.Request.Data.(*ModifyTrafficMirrorFilterRuleOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyTrafficMirrorFilterRuleResponse is the response type for the
// ModifyTrafficMirrorFilterRule API operation.
type ModifyTrafficMirrorFilterRuleResponse struct {
	*ModifyTrafficMirrorFilterRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyTrafficMirrorFilterRule request.
func (r *ModifyTrafficMirrorFilterRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
