// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html
type LaunchTemplateConstraint struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-acceptlanguage
	AcceptLanguage pulumi.StringPtrOutput `pulumi:"acceptLanguage"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-portfolioid
	PortfolioId pulumi.StringOutput `pulumi:"portfolioId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-productid
	ProductId pulumi.StringOutput `pulumi:"productId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-rules
	Rules pulumi.StringOutput `pulumi:"rules"`
}

// NewLaunchTemplateConstraint registers a new resource with the given unique name, arguments, and options.
func NewLaunchTemplateConstraint(ctx *pulumi.Context,
	name string, args *LaunchTemplateConstraintArgs, opts ...pulumi.ResourceOption) (*LaunchTemplateConstraint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PortfolioId == nil {
		return nil, errors.New("invalid value for required argument 'PortfolioId'")
	}
	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	var resource LaunchTemplateConstraint
	err := ctx.RegisterResource("aws-native:servicecatalog:LaunchTemplateConstraint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLaunchTemplateConstraint gets an existing LaunchTemplateConstraint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLaunchTemplateConstraint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LaunchTemplateConstraintState, opts ...pulumi.ResourceOption) (*LaunchTemplateConstraint, error) {
	var resource LaunchTemplateConstraint
	err := ctx.ReadResource("aws-native:servicecatalog:LaunchTemplateConstraint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LaunchTemplateConstraint resources.
type launchTemplateConstraintState struct {
}

type LaunchTemplateConstraintState struct {
}

func (LaunchTemplateConstraintState) ElementType() reflect.Type {
	return reflect.TypeOf((*launchTemplateConstraintState)(nil)).Elem()
}

type launchTemplateConstraintArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-acceptlanguage
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-description
	Description *string `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-portfolioid
	PortfolioId string `pulumi:"portfolioId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-productid
	ProductId string `pulumi:"productId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-rules
	Rules string `pulumi:"rules"`
}

// The set of arguments for constructing a LaunchTemplateConstraint resource.
type LaunchTemplateConstraintArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-acceptlanguage
	AcceptLanguage pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-description
	Description pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-portfolioid
	PortfolioId pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-productid
	ProductId pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-rules
	Rules pulumi.StringInput
}

func (LaunchTemplateConstraintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*launchTemplateConstraintArgs)(nil)).Elem()
}

type LaunchTemplateConstraintInput interface {
	pulumi.Input

	ToLaunchTemplateConstraintOutput() LaunchTemplateConstraintOutput
	ToLaunchTemplateConstraintOutputWithContext(ctx context.Context) LaunchTemplateConstraintOutput
}

func (*LaunchTemplateConstraint) ElementType() reflect.Type {
	return reflect.TypeOf((*LaunchTemplateConstraint)(nil))
}

func (i *LaunchTemplateConstraint) ToLaunchTemplateConstraintOutput() LaunchTemplateConstraintOutput {
	return i.ToLaunchTemplateConstraintOutputWithContext(context.Background())
}

func (i *LaunchTemplateConstraint) ToLaunchTemplateConstraintOutputWithContext(ctx context.Context) LaunchTemplateConstraintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateConstraintOutput)
}

type LaunchTemplateConstraintOutput struct{ *pulumi.OutputState }

func (LaunchTemplateConstraintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LaunchTemplateConstraint)(nil))
}

func (o LaunchTemplateConstraintOutput) ToLaunchTemplateConstraintOutput() LaunchTemplateConstraintOutput {
	return o
}

func (o LaunchTemplateConstraintOutput) ToLaunchTemplateConstraintOutputWithContext(ctx context.Context) LaunchTemplateConstraintOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LaunchTemplateConstraintOutput{})
}
