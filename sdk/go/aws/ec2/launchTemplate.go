// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html
type LaunchTemplate struct {
	pulumi.CustomResourceState

	DefaultVersionNumber pulumi.StringOutput `pulumi:"defaultVersionNumber"`
	LatestVersionNumber  pulumi.StringOutput `pulumi:"latestVersionNumber"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-launchtemplatedata
	LaunchTemplateData LaunchTemplateLaunchTemplateDataPtrOutput `pulumi:"launchTemplateData"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-launchtemplatename
	LaunchTemplateName pulumi.StringPtrOutput `pulumi:"launchTemplateName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-tagspecifications
	TagSpecifications LaunchTemplateLaunchTemplateTagSpecificationArrayOutput `pulumi:"tagSpecifications"`
}

// NewLaunchTemplate registers a new resource with the given unique name, arguments, and options.
func NewLaunchTemplate(ctx *pulumi.Context,
	name string, args *LaunchTemplateArgs, opts ...pulumi.ResourceOption) (*LaunchTemplate, error) {
	if args == nil {
		args = &LaunchTemplateArgs{}
	}

	var resource LaunchTemplate
	err := ctx.RegisterResource("aws-native:ec2:LaunchTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLaunchTemplate gets an existing LaunchTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLaunchTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LaunchTemplateState, opts ...pulumi.ResourceOption) (*LaunchTemplate, error) {
	var resource LaunchTemplate
	err := ctx.ReadResource("aws-native:ec2:LaunchTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LaunchTemplate resources.
type launchTemplateState struct {
}

type LaunchTemplateState struct {
}

func (LaunchTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*launchTemplateState)(nil)).Elem()
}

type launchTemplateArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-launchtemplatedata
	LaunchTemplateData *LaunchTemplateLaunchTemplateData `pulumi:"launchTemplateData"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-launchtemplatename
	LaunchTemplateName *string `pulumi:"launchTemplateName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-tagspecifications
	TagSpecifications []LaunchTemplateLaunchTemplateTagSpecification `pulumi:"tagSpecifications"`
}

// The set of arguments for constructing a LaunchTemplate resource.
type LaunchTemplateArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-launchtemplatedata
	LaunchTemplateData LaunchTemplateLaunchTemplateDataPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-launchtemplatename
	LaunchTemplateName pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-tagspecifications
	TagSpecifications LaunchTemplateLaunchTemplateTagSpecificationArrayInput
}

func (LaunchTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*launchTemplateArgs)(nil)).Elem()
}

type LaunchTemplateInput interface {
	pulumi.Input

	ToLaunchTemplateOutput() LaunchTemplateOutput
	ToLaunchTemplateOutputWithContext(ctx context.Context) LaunchTemplateOutput
}

func (*LaunchTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((*LaunchTemplate)(nil))
}

func (i *LaunchTemplate) ToLaunchTemplateOutput() LaunchTemplateOutput {
	return i.ToLaunchTemplateOutputWithContext(context.Background())
}

func (i *LaunchTemplate) ToLaunchTemplateOutputWithContext(ctx context.Context) LaunchTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateOutput)
}

type LaunchTemplateOutput struct{ *pulumi.OutputState }

func (LaunchTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LaunchTemplate)(nil))
}

func (o LaunchTemplateOutput) ToLaunchTemplateOutput() LaunchTemplateOutput {
	return o
}

func (o LaunchTemplateOutput) ToLaunchTemplateOutputWithContext(ctx context.Context) LaunchTemplateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LaunchTemplateOutput{})
}
