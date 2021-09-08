// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ask

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ask-skill.html
type Skill struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ask-skill.html#cfn-ask-skill-authenticationconfiguration
	AuthenticationConfiguration SkillAuthenticationConfigurationOutput `pulumi:"authenticationConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ask-skill.html#cfn-ask-skill-skillpackage
	SkillPackage SkillSkillPackageOutput `pulumi:"skillPackage"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ask-skill.html#cfn-ask-skill-vendorid
	VendorId pulumi.StringOutput `pulumi:"vendorId"`
}

// NewSkill registers a new resource with the given unique name, arguments, and options.
func NewSkill(ctx *pulumi.Context,
	name string, args *SkillArgs, opts ...pulumi.ResourceOption) (*Skill, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthenticationConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'AuthenticationConfiguration'")
	}
	if args.SkillPackage == nil {
		return nil, errors.New("invalid value for required argument 'SkillPackage'")
	}
	if args.VendorId == nil {
		return nil, errors.New("invalid value for required argument 'VendorId'")
	}
	var resource Skill
	err := ctx.RegisterResource("aws-native:ask:Skill", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSkill gets an existing Skill resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSkill(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SkillState, opts ...pulumi.ResourceOption) (*Skill, error) {
	var resource Skill
	err := ctx.ReadResource("aws-native:ask:Skill", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Skill resources.
type skillState struct {
}

type SkillState struct {
}

func (SkillState) ElementType() reflect.Type {
	return reflect.TypeOf((*skillState)(nil)).Elem()
}

type skillArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ask-skill.html#cfn-ask-skill-authenticationconfiguration
	AuthenticationConfiguration SkillAuthenticationConfiguration `pulumi:"authenticationConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ask-skill.html#cfn-ask-skill-skillpackage
	SkillPackage SkillSkillPackage `pulumi:"skillPackage"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ask-skill.html#cfn-ask-skill-vendorid
	VendorId string `pulumi:"vendorId"`
}

// The set of arguments for constructing a Skill resource.
type SkillArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ask-skill.html#cfn-ask-skill-authenticationconfiguration
	AuthenticationConfiguration SkillAuthenticationConfigurationInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ask-skill.html#cfn-ask-skill-skillpackage
	SkillPackage SkillSkillPackageInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ask-skill.html#cfn-ask-skill-vendorid
	VendorId pulumi.StringInput
}

func (SkillArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*skillArgs)(nil)).Elem()
}

type SkillInput interface {
	pulumi.Input

	ToSkillOutput() SkillOutput
	ToSkillOutputWithContext(ctx context.Context) SkillOutput
}

func (*Skill) ElementType() reflect.Type {
	return reflect.TypeOf((*Skill)(nil))
}

func (i *Skill) ToSkillOutput() SkillOutput {
	return i.ToSkillOutputWithContext(context.Background())
}

func (i *Skill) ToSkillOutputWithContext(ctx context.Context) SkillOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkillOutput)
}

type SkillOutput struct{ *pulumi.OutputState }

func (SkillOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Skill)(nil))
}

func (o SkillOutput) ToSkillOutput() SkillOutput {
	return o
}

func (o SkillOutput) ToSkillOutputWithContext(ctx context.Context) SkillOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SkillOutput{})
}
