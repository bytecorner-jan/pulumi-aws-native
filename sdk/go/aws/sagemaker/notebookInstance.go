// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::NotebookInstance
//
// Deprecated: NotebookInstance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type NotebookInstance struct {
	pulumi.CustomResourceState

	AcceleratorTypes           pulumi.StringArrayOutput       `pulumi:"acceleratorTypes"`
	AdditionalCodeRepositories pulumi.StringArrayOutput       `pulumi:"additionalCodeRepositories"`
	DefaultCodeRepository      pulumi.StringPtrOutput         `pulumi:"defaultCodeRepository"`
	DirectInternetAccess       pulumi.StringPtrOutput         `pulumi:"directInternetAccess"`
	InstanceType               pulumi.StringOutput            `pulumi:"instanceType"`
	KmsKeyId                   pulumi.StringPtrOutput         `pulumi:"kmsKeyId"`
	LifecycleConfigName        pulumi.StringPtrOutput         `pulumi:"lifecycleConfigName"`
	NotebookInstanceName       pulumi.StringPtrOutput         `pulumi:"notebookInstanceName"`
	PlatformIdentifier         pulumi.StringPtrOutput         `pulumi:"platformIdentifier"`
	RoleArn                    pulumi.StringOutput            `pulumi:"roleArn"`
	RootAccess                 pulumi.StringPtrOutput         `pulumi:"rootAccess"`
	SecurityGroupIds           pulumi.StringArrayOutput       `pulumi:"securityGroupIds"`
	SubnetId                   pulumi.StringPtrOutput         `pulumi:"subnetId"`
	Tags                       NotebookInstanceTagArrayOutput `pulumi:"tags"`
	VolumeSizeInGB             pulumi.IntPtrOutput            `pulumi:"volumeSizeInGB"`
}

// NewNotebookInstance registers a new resource with the given unique name, arguments, and options.
func NewNotebookInstance(ctx *pulumi.Context,
	name string, args *NotebookInstanceArgs, opts ...pulumi.ResourceOption) (*NotebookInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource NotebookInstance
	err := ctx.RegisterResource("aws-native:sagemaker:NotebookInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotebookInstance gets an existing NotebookInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotebookInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebookInstanceState, opts ...pulumi.ResourceOption) (*NotebookInstance, error) {
	var resource NotebookInstance
	err := ctx.ReadResource("aws-native:sagemaker:NotebookInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotebookInstance resources.
type notebookInstanceState struct {
}

type NotebookInstanceState struct {
}

func (NotebookInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookInstanceState)(nil)).Elem()
}

type notebookInstanceArgs struct {
	AcceleratorTypes           []string              `pulumi:"acceleratorTypes"`
	AdditionalCodeRepositories []string              `pulumi:"additionalCodeRepositories"`
	DefaultCodeRepository      *string               `pulumi:"defaultCodeRepository"`
	DirectInternetAccess       *string               `pulumi:"directInternetAccess"`
	InstanceType               string                `pulumi:"instanceType"`
	KmsKeyId                   *string               `pulumi:"kmsKeyId"`
	LifecycleConfigName        *string               `pulumi:"lifecycleConfigName"`
	NotebookInstanceName       *string               `pulumi:"notebookInstanceName"`
	PlatformIdentifier         *string               `pulumi:"platformIdentifier"`
	RoleArn                    string                `pulumi:"roleArn"`
	RootAccess                 *string               `pulumi:"rootAccess"`
	SecurityGroupIds           []string              `pulumi:"securityGroupIds"`
	SubnetId                   *string               `pulumi:"subnetId"`
	Tags                       []NotebookInstanceTag `pulumi:"tags"`
	VolumeSizeInGB             *int                  `pulumi:"volumeSizeInGB"`
}

// The set of arguments for constructing a NotebookInstance resource.
type NotebookInstanceArgs struct {
	AcceleratorTypes           pulumi.StringArrayInput
	AdditionalCodeRepositories pulumi.StringArrayInput
	DefaultCodeRepository      pulumi.StringPtrInput
	DirectInternetAccess       pulumi.StringPtrInput
	InstanceType               pulumi.StringInput
	KmsKeyId                   pulumi.StringPtrInput
	LifecycleConfigName        pulumi.StringPtrInput
	NotebookInstanceName       pulumi.StringPtrInput
	PlatformIdentifier         pulumi.StringPtrInput
	RoleArn                    pulumi.StringInput
	RootAccess                 pulumi.StringPtrInput
	SecurityGroupIds           pulumi.StringArrayInput
	SubnetId                   pulumi.StringPtrInput
	Tags                       NotebookInstanceTagArrayInput
	VolumeSizeInGB             pulumi.IntPtrInput
}

func (NotebookInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookInstanceArgs)(nil)).Elem()
}

type NotebookInstanceInput interface {
	pulumi.Input

	ToNotebookInstanceOutput() NotebookInstanceOutput
	ToNotebookInstanceOutputWithContext(ctx context.Context) NotebookInstanceOutput
}

func (*NotebookInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookInstance)(nil))
}

func (i *NotebookInstance) ToNotebookInstanceOutput() NotebookInstanceOutput {
	return i.ToNotebookInstanceOutputWithContext(context.Background())
}

func (i *NotebookInstance) ToNotebookInstanceOutputWithContext(ctx context.Context) NotebookInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookInstanceOutput)
}

type NotebookInstanceOutput struct{ *pulumi.OutputState }

func (NotebookInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookInstance)(nil))
}

func (o NotebookInstanceOutput) ToNotebookInstanceOutput() NotebookInstanceOutput {
	return o
}

func (o NotebookInstanceOutput) ToNotebookInstanceOutputWithContext(ctx context.Context) NotebookInstanceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotebookInstanceInput)(nil)).Elem(), &NotebookInstance{})
	pulumi.RegisterOutputType(NotebookInstanceOutput{})
}
