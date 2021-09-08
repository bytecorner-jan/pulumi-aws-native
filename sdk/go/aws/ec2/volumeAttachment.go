// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html
type VolumeAttachment struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html#cfn-ec2-ebs-volumeattachment-device
	Device pulumi.StringOutput `pulumi:"device"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html#cfn-ec2-ebs-volumeattachment-instanceid
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html#cfn-ec2-ebs-volumeattachment-volumeid
	VolumeId pulumi.StringOutput `pulumi:"volumeId"`
}

// NewVolumeAttachment registers a new resource with the given unique name, arguments, and options.
func NewVolumeAttachment(ctx *pulumi.Context,
	name string, args *VolumeAttachmentArgs, opts ...pulumi.ResourceOption) (*VolumeAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Device == nil {
		return nil, errors.New("invalid value for required argument 'Device'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.VolumeId == nil {
		return nil, errors.New("invalid value for required argument 'VolumeId'")
	}
	var resource VolumeAttachment
	err := ctx.RegisterResource("aws-native:ec2:VolumeAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeAttachment gets an existing VolumeAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeAttachmentState, opts ...pulumi.ResourceOption) (*VolumeAttachment, error) {
	var resource VolumeAttachment
	err := ctx.ReadResource("aws-native:ec2:VolumeAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeAttachment resources.
type volumeAttachmentState struct {
}

type VolumeAttachmentState struct {
}

func (VolumeAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeAttachmentState)(nil)).Elem()
}

type volumeAttachmentArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html#cfn-ec2-ebs-volumeattachment-device
	Device string `pulumi:"device"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html#cfn-ec2-ebs-volumeattachment-instanceid
	InstanceId string `pulumi:"instanceId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html#cfn-ec2-ebs-volumeattachment-volumeid
	VolumeId string `pulumi:"volumeId"`
}

// The set of arguments for constructing a VolumeAttachment resource.
type VolumeAttachmentArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html#cfn-ec2-ebs-volumeattachment-device
	Device pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html#cfn-ec2-ebs-volumeattachment-instanceid
	InstanceId pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html#cfn-ec2-ebs-volumeattachment-volumeid
	VolumeId pulumi.StringInput
}

func (VolumeAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeAttachmentArgs)(nil)).Elem()
}

type VolumeAttachmentInput interface {
	pulumi.Input

	ToVolumeAttachmentOutput() VolumeAttachmentOutput
	ToVolumeAttachmentOutputWithContext(ctx context.Context) VolumeAttachmentOutput
}

func (*VolumeAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachment)(nil))
}

func (i *VolumeAttachment) ToVolumeAttachmentOutput() VolumeAttachmentOutput {
	return i.ToVolumeAttachmentOutputWithContext(context.Background())
}

func (i *VolumeAttachment) ToVolumeAttachmentOutputWithContext(ctx context.Context) VolumeAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentOutput)
}

type VolumeAttachmentOutput struct{ *pulumi.OutputState }

func (VolumeAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachment)(nil))
}

func (o VolumeAttachmentOutput) ToVolumeAttachmentOutput() VolumeAttachmentOutput {
	return o
}

func (o VolumeAttachmentOutput) ToVolumeAttachmentOutputWithContext(ctx context.Context) VolumeAttachmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VolumeAttachmentOutput{})
}
