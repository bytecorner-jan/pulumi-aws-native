// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html
type IPSet struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-activate
	Activate pulumi.BoolOutput `pulumi:"activate"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-detectorid
	DetectorId pulumi.StringOutput `pulumi:"detectorId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-format
	Format pulumi.StringOutput `pulumi:"format"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-location
	Location pulumi.StringOutput `pulumi:"location"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-name
	Name pulumi.StringPtrOutput `pulumi:"name"`
}

// NewIPSet registers a new resource with the given unique name, arguments, and options.
func NewIPSet(ctx *pulumi.Context,
	name string, args *IPSetArgs, opts ...pulumi.ResourceOption) (*IPSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Activate == nil {
		return nil, errors.New("invalid value for required argument 'Activate'")
	}
	if args.DetectorId == nil {
		return nil, errors.New("invalid value for required argument 'DetectorId'")
	}
	if args.Format == nil {
		return nil, errors.New("invalid value for required argument 'Format'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	var resource IPSet
	err := ctx.RegisterResource("aws-native:guardduty:IPSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIPSet gets an existing IPSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIPSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IPSetState, opts ...pulumi.ResourceOption) (*IPSet, error) {
	var resource IPSet
	err := ctx.ReadResource("aws-native:guardduty:IPSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IPSet resources.
type ipsetState struct {
}

type IPSetState struct {
}

func (IPSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsetState)(nil)).Elem()
}

type ipsetArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-activate
	Activate bool `pulumi:"activate"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-detectorid
	DetectorId string `pulumi:"detectorId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-format
	Format string `pulumi:"format"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-location
	Location string `pulumi:"location"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-name
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a IPSet resource.
type IPSetArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-activate
	Activate pulumi.BoolInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-detectorid
	DetectorId pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-format
	Format pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-location
	Location pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-name
	Name pulumi.StringPtrInput
}

func (IPSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsetArgs)(nil)).Elem()
}

type IPSetInput interface {
	pulumi.Input

	ToIPSetOutput() IPSetOutput
	ToIPSetOutputWithContext(ctx context.Context) IPSetOutput
}

func (*IPSet) ElementType() reflect.Type {
	return reflect.TypeOf((*IPSet)(nil))
}

func (i *IPSet) ToIPSetOutput() IPSetOutput {
	return i.ToIPSetOutputWithContext(context.Background())
}

func (i *IPSet) ToIPSetOutputWithContext(ctx context.Context) IPSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPSetOutput)
}

type IPSetOutput struct{ *pulumi.OutputState }

func (IPSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPSet)(nil))
}

func (o IPSetOutput) ToIPSetOutput() IPSetOutput {
	return o
}

func (o IPSetOutput) ToIPSetOutputWithContext(ctx context.Context) IPSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IPSetOutput{})
}
