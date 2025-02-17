// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssmincidents

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource type definition for AWS::SSMIncidents::ResponsePlan
type ResponsePlan struct {
	pulumi.CustomResourceState

	// The list of actions.
	Actions ResponsePlanActionArrayOutput `pulumi:"actions"`
	// The ARN of the response plan.
	Arn         pulumi.StringOutput              `pulumi:"arn"`
	ChatChannel ResponsePlanChatChannelPtrOutput `pulumi:"chatChannel"`
	// The display name of the response plan.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The list of engagements to use.
	Engagements      pulumi.StringArrayOutput           `pulumi:"engagements"`
	IncidentTemplate ResponsePlanIncidentTemplateOutput `pulumi:"incidentTemplate"`
	// The name of the response plan.
	Name pulumi.StringOutput `pulumi:"name"`
	// The tags to apply to the response plan.
	Tags ResponsePlanTagArrayOutput `pulumi:"tags"`
}

// NewResponsePlan registers a new resource with the given unique name, arguments, and options.
func NewResponsePlan(ctx *pulumi.Context,
	name string, args *ResponsePlanArgs, opts ...pulumi.ResourceOption) (*ResponsePlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IncidentTemplate == nil {
		return nil, errors.New("invalid value for required argument 'IncidentTemplate'")
	}
	var resource ResponsePlan
	err := ctx.RegisterResource("aws-native:ssmincidents:ResponsePlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResponsePlan gets an existing ResponsePlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResponsePlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResponsePlanState, opts ...pulumi.ResourceOption) (*ResponsePlan, error) {
	var resource ResponsePlan
	err := ctx.ReadResource("aws-native:ssmincidents:ResponsePlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResponsePlan resources.
type responsePlanState struct {
}

type ResponsePlanState struct {
}

func (ResponsePlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*responsePlanState)(nil)).Elem()
}

type responsePlanArgs struct {
	// The list of actions.
	Actions     []ResponsePlanAction     `pulumi:"actions"`
	ChatChannel *ResponsePlanChatChannel `pulumi:"chatChannel"`
	// The display name of the response plan.
	DisplayName *string `pulumi:"displayName"`
	// The list of engagements to use.
	Engagements      []string                     `pulumi:"engagements"`
	IncidentTemplate ResponsePlanIncidentTemplate `pulumi:"incidentTemplate"`
	// The name of the response plan.
	Name *string `pulumi:"name"`
	// The tags to apply to the response plan.
	Tags []ResponsePlanTag `pulumi:"tags"`
}

// The set of arguments for constructing a ResponsePlan resource.
type ResponsePlanArgs struct {
	// The list of actions.
	Actions     ResponsePlanActionArrayInput
	ChatChannel ResponsePlanChatChannelPtrInput
	// The display name of the response plan.
	DisplayName pulumi.StringPtrInput
	// The list of engagements to use.
	Engagements      pulumi.StringArrayInput
	IncidentTemplate ResponsePlanIncidentTemplateInput
	// The name of the response plan.
	Name pulumi.StringPtrInput
	// The tags to apply to the response plan.
	Tags ResponsePlanTagArrayInput
}

func (ResponsePlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*responsePlanArgs)(nil)).Elem()
}

type ResponsePlanInput interface {
	pulumi.Input

	ToResponsePlanOutput() ResponsePlanOutput
	ToResponsePlanOutputWithContext(ctx context.Context) ResponsePlanOutput
}

func (*ResponsePlan) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponsePlan)(nil))
}

func (i *ResponsePlan) ToResponsePlanOutput() ResponsePlanOutput {
	return i.ToResponsePlanOutputWithContext(context.Background())
}

func (i *ResponsePlan) ToResponsePlanOutputWithContext(ctx context.Context) ResponsePlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponsePlanOutput)
}

type ResponsePlanOutput struct{ *pulumi.OutputState }

func (ResponsePlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponsePlan)(nil))
}

func (o ResponsePlanOutput) ToResponsePlanOutput() ResponsePlanOutput {
	return o
}

func (o ResponsePlanOutput) ToResponsePlanOutputWithContext(ctx context.Context) ResponsePlanOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResponsePlanInput)(nil)).Elem(), &ResponsePlan{})
	pulumi.RegisterOutputType(ResponsePlanOutput{})
}
