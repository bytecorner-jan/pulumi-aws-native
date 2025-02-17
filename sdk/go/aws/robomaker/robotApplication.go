// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package robomaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::RoboMaker::RobotApplication
//
// Deprecated: RobotApplication is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type RobotApplication struct {
	pulumi.CustomResourceState

	Arn                pulumi.StringOutput                      `pulumi:"arn"`
	CurrentRevisionId  pulumi.StringPtrOutput                   `pulumi:"currentRevisionId"`
	Name               pulumi.StringPtrOutput                   `pulumi:"name"`
	RobotSoftwareSuite RobotApplicationRobotSoftwareSuiteOutput `pulumi:"robotSoftwareSuite"`
	Sources            RobotApplicationSourceConfigArrayOutput  `pulumi:"sources"`
	Tags               pulumi.AnyOutput                         `pulumi:"tags"`
}

// NewRobotApplication registers a new resource with the given unique name, arguments, and options.
func NewRobotApplication(ctx *pulumi.Context,
	name string, args *RobotApplicationArgs, opts ...pulumi.ResourceOption) (*RobotApplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RobotSoftwareSuite == nil {
		return nil, errors.New("invalid value for required argument 'RobotSoftwareSuite'")
	}
	if args.Sources == nil {
		return nil, errors.New("invalid value for required argument 'Sources'")
	}
	var resource RobotApplication
	err := ctx.RegisterResource("aws-native:robomaker:RobotApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRobotApplication gets an existing RobotApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRobotApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RobotApplicationState, opts ...pulumi.ResourceOption) (*RobotApplication, error) {
	var resource RobotApplication
	err := ctx.ReadResource("aws-native:robomaker:RobotApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RobotApplication resources.
type robotApplicationState struct {
}

type RobotApplicationState struct {
}

func (RobotApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*robotApplicationState)(nil)).Elem()
}

type robotApplicationArgs struct {
	CurrentRevisionId  *string                            `pulumi:"currentRevisionId"`
	Name               *string                            `pulumi:"name"`
	RobotSoftwareSuite RobotApplicationRobotSoftwareSuite `pulumi:"robotSoftwareSuite"`
	Sources            []RobotApplicationSourceConfig     `pulumi:"sources"`
	Tags               interface{}                        `pulumi:"tags"`
}

// The set of arguments for constructing a RobotApplication resource.
type RobotApplicationArgs struct {
	CurrentRevisionId  pulumi.StringPtrInput
	Name               pulumi.StringPtrInput
	RobotSoftwareSuite RobotApplicationRobotSoftwareSuiteInput
	Sources            RobotApplicationSourceConfigArrayInput
	Tags               pulumi.Input
}

func (RobotApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*robotApplicationArgs)(nil)).Elem()
}

type RobotApplicationInput interface {
	pulumi.Input

	ToRobotApplicationOutput() RobotApplicationOutput
	ToRobotApplicationOutputWithContext(ctx context.Context) RobotApplicationOutput
}

func (*RobotApplication) ElementType() reflect.Type {
	return reflect.TypeOf((*RobotApplication)(nil))
}

func (i *RobotApplication) ToRobotApplicationOutput() RobotApplicationOutput {
	return i.ToRobotApplicationOutputWithContext(context.Background())
}

func (i *RobotApplication) ToRobotApplicationOutputWithContext(ctx context.Context) RobotApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RobotApplicationOutput)
}

type RobotApplicationOutput struct{ *pulumi.OutputState }

func (RobotApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RobotApplication)(nil))
}

func (o RobotApplicationOutput) ToRobotApplicationOutput() RobotApplicationOutput {
	return o
}

func (o RobotApplicationOutput) ToRobotApplicationOutputWithContext(ctx context.Context) RobotApplicationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RobotApplicationInput)(nil)).Elem(), &RobotApplication{})
	pulumi.RegisterOutputType(RobotApplicationOutput{})
}
