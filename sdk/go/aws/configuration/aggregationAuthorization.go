// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Config::AggregationAuthorization
type AggregationAuthorization struct {
	pulumi.CustomResourceState

	AuthorizedAccountId pulumi.StringOutput                    `pulumi:"authorizedAccountId"`
	AuthorizedAwsRegion pulumi.StringOutput                    `pulumi:"authorizedAwsRegion"`
	Tags                AggregationAuthorizationTagArrayOutput `pulumi:"tags"`
}

// NewAggregationAuthorization registers a new resource with the given unique name, arguments, and options.
func NewAggregationAuthorization(ctx *pulumi.Context,
	name string, args *AggregationAuthorizationArgs, opts ...pulumi.ResourceOption) (*AggregationAuthorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthorizedAccountId == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizedAccountId'")
	}
	if args.AuthorizedAwsRegion == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizedAwsRegion'")
	}
	var resource AggregationAuthorization
	err := ctx.RegisterResource("aws-native:configuration:AggregationAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAggregationAuthorization gets an existing AggregationAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAggregationAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AggregationAuthorizationState, opts ...pulumi.ResourceOption) (*AggregationAuthorization, error) {
	var resource AggregationAuthorization
	err := ctx.ReadResource("aws-native:configuration:AggregationAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AggregationAuthorization resources.
type aggregationAuthorizationState struct {
}

type AggregationAuthorizationState struct {
}

func (AggregationAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*aggregationAuthorizationState)(nil)).Elem()
}

type aggregationAuthorizationArgs struct {
	AuthorizedAccountId string                        `pulumi:"authorizedAccountId"`
	AuthorizedAwsRegion string                        `pulumi:"authorizedAwsRegion"`
	Tags                []AggregationAuthorizationTag `pulumi:"tags"`
}

// The set of arguments for constructing a AggregationAuthorization resource.
type AggregationAuthorizationArgs struct {
	AuthorizedAccountId pulumi.StringInput
	AuthorizedAwsRegion pulumi.StringInput
	Tags                AggregationAuthorizationTagArrayInput
}

func (AggregationAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aggregationAuthorizationArgs)(nil)).Elem()
}

type AggregationAuthorizationInput interface {
	pulumi.Input

	ToAggregationAuthorizationOutput() AggregationAuthorizationOutput
	ToAggregationAuthorizationOutputWithContext(ctx context.Context) AggregationAuthorizationOutput
}

func (*AggregationAuthorization) ElementType() reflect.Type {
	return reflect.TypeOf((*AggregationAuthorization)(nil))
}

func (i *AggregationAuthorization) ToAggregationAuthorizationOutput() AggregationAuthorizationOutput {
	return i.ToAggregationAuthorizationOutputWithContext(context.Background())
}

func (i *AggregationAuthorization) ToAggregationAuthorizationOutputWithContext(ctx context.Context) AggregationAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregationAuthorizationOutput)
}

type AggregationAuthorizationOutput struct{ *pulumi.OutputState }

func (AggregationAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AggregationAuthorization)(nil))
}

func (o AggregationAuthorizationOutput) ToAggregationAuthorizationOutput() AggregationAuthorizationOutput {
	return o
}

func (o AggregationAuthorizationOutput) ToAggregationAuthorizationOutputWithContext(ctx context.Context) AggregationAuthorizationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AggregationAuthorizationOutput{})
}
