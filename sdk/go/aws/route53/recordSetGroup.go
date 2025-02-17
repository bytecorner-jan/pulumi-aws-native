// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Route53::RecordSetGroup
//
// Deprecated: RecordSetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type RecordSetGroup struct {
	pulumi.CustomResourceState

	Comment        pulumi.StringPtrOutput             `pulumi:"comment"`
	HostedZoneId   pulumi.StringPtrOutput             `pulumi:"hostedZoneId"`
	HostedZoneName pulumi.StringPtrOutput             `pulumi:"hostedZoneName"`
	RecordSets     RecordSetGroupRecordSetArrayOutput `pulumi:"recordSets"`
}

// NewRecordSetGroup registers a new resource with the given unique name, arguments, and options.
func NewRecordSetGroup(ctx *pulumi.Context,
	name string, args *RecordSetGroupArgs, opts ...pulumi.ResourceOption) (*RecordSetGroup, error) {
	if args == nil {
		args = &RecordSetGroupArgs{}
	}

	var resource RecordSetGroup
	err := ctx.RegisterResource("aws-native:route53:RecordSetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRecordSetGroup gets an existing RecordSetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRecordSetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecordSetGroupState, opts ...pulumi.ResourceOption) (*RecordSetGroup, error) {
	var resource RecordSetGroup
	err := ctx.ReadResource("aws-native:route53:RecordSetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RecordSetGroup resources.
type recordSetGroupState struct {
}

type RecordSetGroupState struct {
}

func (RecordSetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetGroupState)(nil)).Elem()
}

type recordSetGroupArgs struct {
	Comment        *string                   `pulumi:"comment"`
	HostedZoneId   *string                   `pulumi:"hostedZoneId"`
	HostedZoneName *string                   `pulumi:"hostedZoneName"`
	RecordSets     []RecordSetGroupRecordSet `pulumi:"recordSets"`
}

// The set of arguments for constructing a RecordSetGroup resource.
type RecordSetGroupArgs struct {
	Comment        pulumi.StringPtrInput
	HostedZoneId   pulumi.StringPtrInput
	HostedZoneName pulumi.StringPtrInput
	RecordSets     RecordSetGroupRecordSetArrayInput
}

func (RecordSetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetGroupArgs)(nil)).Elem()
}

type RecordSetGroupInput interface {
	pulumi.Input

	ToRecordSetGroupOutput() RecordSetGroupOutput
	ToRecordSetGroupOutputWithContext(ctx context.Context) RecordSetGroupOutput
}

func (*RecordSetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordSetGroup)(nil))
}

func (i *RecordSetGroup) ToRecordSetGroupOutput() RecordSetGroupOutput {
	return i.ToRecordSetGroupOutputWithContext(context.Background())
}

func (i *RecordSetGroup) ToRecordSetGroupOutputWithContext(ctx context.Context) RecordSetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordSetGroupOutput)
}

type RecordSetGroupOutput struct{ *pulumi.OutputState }

func (RecordSetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordSetGroup)(nil))
}

func (o RecordSetGroupOutput) ToRecordSetGroupOutput() RecordSetGroupOutput {
	return o
}

func (o RecordSetGroupOutput) ToRecordSetGroupOutputWithContext(ctx context.Context) RecordSetGroupOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RecordSetGroupInput)(nil)).Elem(), &RecordSetGroup{})
	pulumi.RegisterOutputType(RecordSetGroupOutput{})
}
