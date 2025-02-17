// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An example resource schema demonstrating some basic constructs and validation rules.
type ProfilePermission struct {
	pulumi.CustomResourceState

	Action         pulumi.StringOutput    `pulumi:"action"`
	Principal      pulumi.StringOutput    `pulumi:"principal"`
	ProfileName    pulumi.StringOutput    `pulumi:"profileName"`
	ProfileVersion pulumi.StringPtrOutput `pulumi:"profileVersion"`
	StatementId    pulumi.StringOutput    `pulumi:"statementId"`
}

// NewProfilePermission registers a new resource with the given unique name, arguments, and options.
func NewProfilePermission(ctx *pulumi.Context,
	name string, args *ProfilePermissionArgs, opts ...pulumi.ResourceOption) (*ProfilePermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.StatementId == nil {
		return nil, errors.New("invalid value for required argument 'StatementId'")
	}
	var resource ProfilePermission
	err := ctx.RegisterResource("aws-native:signer:ProfilePermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfilePermission gets an existing ProfilePermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfilePermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfilePermissionState, opts ...pulumi.ResourceOption) (*ProfilePermission, error) {
	var resource ProfilePermission
	err := ctx.ReadResource("aws-native:signer:ProfilePermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProfilePermission resources.
type profilePermissionState struct {
}

type ProfilePermissionState struct {
}

func (ProfilePermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*profilePermissionState)(nil)).Elem()
}

type profilePermissionArgs struct {
	Action         string  `pulumi:"action"`
	Principal      string  `pulumi:"principal"`
	ProfileName    string  `pulumi:"profileName"`
	ProfileVersion *string `pulumi:"profileVersion"`
	StatementId    string  `pulumi:"statementId"`
}

// The set of arguments for constructing a ProfilePermission resource.
type ProfilePermissionArgs struct {
	Action         pulumi.StringInput
	Principal      pulumi.StringInput
	ProfileName    pulumi.StringInput
	ProfileVersion pulumi.StringPtrInput
	StatementId    pulumi.StringInput
}

func (ProfilePermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profilePermissionArgs)(nil)).Elem()
}

type ProfilePermissionInput interface {
	pulumi.Input

	ToProfilePermissionOutput() ProfilePermissionOutput
	ToProfilePermissionOutputWithContext(ctx context.Context) ProfilePermissionOutput
}

func (*ProfilePermission) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfilePermission)(nil))
}

func (i *ProfilePermission) ToProfilePermissionOutput() ProfilePermissionOutput {
	return i.ToProfilePermissionOutputWithContext(context.Background())
}

func (i *ProfilePermission) ToProfilePermissionOutputWithContext(ctx context.Context) ProfilePermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfilePermissionOutput)
}

type ProfilePermissionOutput struct{ *pulumi.OutputState }

func (ProfilePermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfilePermission)(nil))
}

func (o ProfilePermissionOutput) ToProfilePermissionOutput() ProfilePermissionOutput {
	return o
}

func (o ProfilePermissionOutput) ToProfilePermissionOutputWithContext(ctx context.Context) ProfilePermissionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfilePermissionInput)(nil)).Elem(), &ProfilePermission{})
	pulumi.RegisterOutputType(ProfilePermissionOutput{})
}
