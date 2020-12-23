// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package robomaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplicationversion.html
type RobotApplicationVersion struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes RobotApplicationVersionAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties RobotApplicationVersionPropertiesOutput `pulumi:"properties"`
}

// NewRobotApplicationVersion registers a new resource with the given unique name, arguments, and options.
func NewRobotApplicationVersion(ctx *pulumi.Context,
	name string, args *RobotApplicationVersionArgs, opts ...pulumi.ResourceOption) (*RobotApplicationVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource RobotApplicationVersion
	err := ctx.RegisterResource("cloudformation:RoboMaker:RobotApplicationVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRobotApplicationVersion gets an existing RobotApplicationVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRobotApplicationVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RobotApplicationVersionState, opts ...pulumi.ResourceOption) (*RobotApplicationVersion, error) {
	var resource RobotApplicationVersion
	err := ctx.ReadResource("cloudformation:RoboMaker:RobotApplicationVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RobotApplicationVersion resources.
type robotApplicationVersionState struct {
	// The attributes associated with the resource
	Attributes *RobotApplicationVersionAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *RobotApplicationVersionProperties `pulumi:"properties"`
}

type RobotApplicationVersionState struct {
	// The attributes associated with the resource
	Attributes RobotApplicationVersionAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties RobotApplicationVersionPropertiesPtrInput
}

func (RobotApplicationVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*robotApplicationVersionState)(nil)).Elem()
}

type robotApplicationVersionArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties RobotApplicationVersionProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a RobotApplicationVersion resource.
type RobotApplicationVersionArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties RobotApplicationVersionPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (RobotApplicationVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*robotApplicationVersionArgs)(nil)).Elem()
}

type RobotApplicationVersionInput interface {
	pulumi.Input

	ToRobotApplicationVersionOutput() RobotApplicationVersionOutput
	ToRobotApplicationVersionOutputWithContext(ctx context.Context) RobotApplicationVersionOutput
}

func (*RobotApplicationVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*RobotApplicationVersion)(nil))
}

func (i *RobotApplicationVersion) ToRobotApplicationVersionOutput() RobotApplicationVersionOutput {
	return i.ToRobotApplicationVersionOutputWithContext(context.Background())
}

func (i *RobotApplicationVersion) ToRobotApplicationVersionOutputWithContext(ctx context.Context) RobotApplicationVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RobotApplicationVersionOutput)
}

type RobotApplicationVersionOutput struct {
	*pulumi.OutputState
}

func (RobotApplicationVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RobotApplicationVersion)(nil))
}

func (o RobotApplicationVersionOutput) ToRobotApplicationVersionOutput() RobotApplicationVersionOutput {
	return o
}

func (o RobotApplicationVersionOutput) ToRobotApplicationVersionOutputWithContext(ctx context.Context) RobotApplicationVersionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RobotApplicationVersionOutput{})
}
