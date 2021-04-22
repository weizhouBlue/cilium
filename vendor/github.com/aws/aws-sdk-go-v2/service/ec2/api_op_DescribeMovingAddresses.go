// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes your Elastic IP addresses that are being moved to the EC2-VPC
// platform, or that are being restored to the EC2-Classic platform. This request
// does not return information about any other Elastic IP addresses in your
// account.
func (c *Client) DescribeMovingAddresses(ctx context.Context, params *DescribeMovingAddressesInput, optFns ...func(*Options)) (*DescribeMovingAddressesOutput, error) {
	if params == nil {
		params = &DescribeMovingAddressesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMovingAddresses", params, optFns, addOperationDescribeMovingAddressesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMovingAddressesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMovingAddressesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// One or more filters.
	//
	// * moving-status - The status of the Elastic IP address
	// (MovingToVpc | RestoringToClassic).
	Filters []types.Filter

	// The maximum number of results to return for the request in a single page. The
	// remaining results of the initial request can be seen by sending another request
	// with the returned NextToken value. This value can be between 5 and 1000; if
	// MaxResults is given a value outside of this range, an error is returned.
	// Default: If no value is provided, the default is 1000.
	MaxResults int32

	// The token for the next page of results.
	NextToken *string

	// One or more Elastic IP addresses.
	PublicIps []string
}

type DescribeMovingAddressesOutput struct {

	// The status for each Elastic IP address.
	MovingAddressStatuses []types.MovingAddressStatus

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeMovingAddressesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeMovingAddresses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeMovingAddresses{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMovingAddresses(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeMovingAddressesAPIClient is a client that implements the
// DescribeMovingAddresses operation.
type DescribeMovingAddressesAPIClient interface {
	DescribeMovingAddresses(context.Context, *DescribeMovingAddressesInput, ...func(*Options)) (*DescribeMovingAddressesOutput, error)
}

var _ DescribeMovingAddressesAPIClient = (*Client)(nil)

// DescribeMovingAddressesPaginatorOptions is the paginator options for
// DescribeMovingAddresses
type DescribeMovingAddressesPaginatorOptions struct {
	// The maximum number of results to return for the request in a single page. The
	// remaining results of the initial request can be seen by sending another request
	// with the returned NextToken value. This value can be between 5 and 1000; if
	// MaxResults is given a value outside of this range, an error is returned.
	// Default: If no value is provided, the default is 1000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMovingAddressesPaginator is a paginator for DescribeMovingAddresses
type DescribeMovingAddressesPaginator struct {
	options   DescribeMovingAddressesPaginatorOptions
	client    DescribeMovingAddressesAPIClient
	params    *DescribeMovingAddressesInput
	nextToken *string
	firstPage bool
}

// NewDescribeMovingAddressesPaginator returns a new
// DescribeMovingAddressesPaginator
func NewDescribeMovingAddressesPaginator(client DescribeMovingAddressesAPIClient, params *DescribeMovingAddressesInput, optFns ...func(*DescribeMovingAddressesPaginatorOptions)) *DescribeMovingAddressesPaginator {
	if params == nil {
		params = &DescribeMovingAddressesInput{}
	}

	options := DescribeMovingAddressesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeMovingAddressesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMovingAddressesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeMovingAddresses page.
func (p *DescribeMovingAddressesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMovingAddressesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeMovingAddresses(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeMovingAddresses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeMovingAddresses",
	}
}
