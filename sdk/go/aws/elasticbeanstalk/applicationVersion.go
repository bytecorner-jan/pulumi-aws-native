// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html
type ApplicationVersion struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-applicationname
	ApplicationName pulumi.StringOutput `pulumi:"applicationName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-sourcebundle
	SourceBundle ApplicationVersionSourceBundleOutput `pulumi:"sourceBundle"`
}

// NewApplicationVersion registers a new resource with the given unique name, arguments, and options.
func NewApplicationVersion(ctx *pulumi.Context,
	name string, args *ApplicationVersionArgs, opts ...pulumi.ResourceOption) (*ApplicationVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationName'")
	}
	if args.SourceBundle == nil {
		return nil, errors.New("invalid value for required argument 'SourceBundle'")
	}
	var resource ApplicationVersion
	err := ctx.RegisterResource("aws-native:elasticbeanstalk:ApplicationVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationVersion gets an existing ApplicationVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationVersionState, opts ...pulumi.ResourceOption) (*ApplicationVersion, error) {
	var resource ApplicationVersion
	err := ctx.ReadResource("aws-native:elasticbeanstalk:ApplicationVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationVersion resources.
type applicationVersionState struct {
}

type ApplicationVersionState struct {
}

func (ApplicationVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationVersionState)(nil)).Elem()
}

type applicationVersionArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-applicationname
	ApplicationName string `pulumi:"applicationName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-description
	Description *string `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-sourcebundle
	SourceBundle ApplicationVersionSourceBundle `pulumi:"sourceBundle"`
}

// The set of arguments for constructing a ApplicationVersion resource.
type ApplicationVersionArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-applicationname
	ApplicationName pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-description
	Description pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-sourcebundle
	SourceBundle ApplicationVersionSourceBundleInput
}

func (ApplicationVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationVersionArgs)(nil)).Elem()
}

type ApplicationVersionInput interface {
	pulumi.Input

	ToApplicationVersionOutput() ApplicationVersionOutput
	ToApplicationVersionOutputWithContext(ctx context.Context) ApplicationVersionOutput
}

func (*ApplicationVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationVersion)(nil))
}

func (i *ApplicationVersion) ToApplicationVersionOutput() ApplicationVersionOutput {
	return i.ToApplicationVersionOutputWithContext(context.Background())
}

func (i *ApplicationVersion) ToApplicationVersionOutputWithContext(ctx context.Context) ApplicationVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationVersionOutput)
}

type ApplicationVersionOutput struct{ *pulumi.OutputState }

func (ApplicationVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationVersion)(nil))
}

func (o ApplicationVersionOutput) ToApplicationVersionOutput() ApplicationVersionOutput {
	return o
}

func (o ApplicationVersionOutput) ToApplicationVersionOutputWithContext(ctx context.Context) ApplicationVersionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApplicationVersionOutput{})
}
