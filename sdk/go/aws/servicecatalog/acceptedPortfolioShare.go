// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-acceptedportfolioshare.html
type AcceptedPortfolioShare struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-acceptedportfolioshare.html#cfn-servicecatalog-acceptedportfolioshare-acceptlanguage
	AcceptLanguage pulumi.StringPtrOutput `pulumi:"acceptLanguage"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-acceptedportfolioshare.html#cfn-servicecatalog-acceptedportfolioshare-portfolioid
	PortfolioId pulumi.StringOutput `pulumi:"portfolioId"`
}

// NewAcceptedPortfolioShare registers a new resource with the given unique name, arguments, and options.
func NewAcceptedPortfolioShare(ctx *pulumi.Context,
	name string, args *AcceptedPortfolioShareArgs, opts ...pulumi.ResourceOption) (*AcceptedPortfolioShare, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PortfolioId == nil {
		return nil, errors.New("invalid value for required argument 'PortfolioId'")
	}
	var resource AcceptedPortfolioShare
	err := ctx.RegisterResource("aws-native:servicecatalog:AcceptedPortfolioShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAcceptedPortfolioShare gets an existing AcceptedPortfolioShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAcceptedPortfolioShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AcceptedPortfolioShareState, opts ...pulumi.ResourceOption) (*AcceptedPortfolioShare, error) {
	var resource AcceptedPortfolioShare
	err := ctx.ReadResource("aws-native:servicecatalog:AcceptedPortfolioShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AcceptedPortfolioShare resources.
type acceptedPortfolioShareState struct {
}

type AcceptedPortfolioShareState struct {
}

func (AcceptedPortfolioShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*acceptedPortfolioShareState)(nil)).Elem()
}

type acceptedPortfolioShareArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-acceptedportfolioshare.html#cfn-servicecatalog-acceptedportfolioshare-acceptlanguage
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-acceptedportfolioshare.html#cfn-servicecatalog-acceptedportfolioshare-portfolioid
	PortfolioId string `pulumi:"portfolioId"`
}

// The set of arguments for constructing a AcceptedPortfolioShare resource.
type AcceptedPortfolioShareArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-acceptedportfolioshare.html#cfn-servicecatalog-acceptedportfolioshare-acceptlanguage
	AcceptLanguage pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-acceptedportfolioshare.html#cfn-servicecatalog-acceptedportfolioshare-portfolioid
	PortfolioId pulumi.StringInput
}

func (AcceptedPortfolioShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*acceptedPortfolioShareArgs)(nil)).Elem()
}

type AcceptedPortfolioShareInput interface {
	pulumi.Input

	ToAcceptedPortfolioShareOutput() AcceptedPortfolioShareOutput
	ToAcceptedPortfolioShareOutputWithContext(ctx context.Context) AcceptedPortfolioShareOutput
}

func (*AcceptedPortfolioShare) ElementType() reflect.Type {
	return reflect.TypeOf((*AcceptedPortfolioShare)(nil))
}

func (i *AcceptedPortfolioShare) ToAcceptedPortfolioShareOutput() AcceptedPortfolioShareOutput {
	return i.ToAcceptedPortfolioShareOutputWithContext(context.Background())
}

func (i *AcceptedPortfolioShare) ToAcceptedPortfolioShareOutputWithContext(ctx context.Context) AcceptedPortfolioShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcceptedPortfolioShareOutput)
}

type AcceptedPortfolioShareOutput struct{ *pulumi.OutputState }

func (AcceptedPortfolioShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AcceptedPortfolioShare)(nil))
}

func (o AcceptedPortfolioShareOutput) ToAcceptedPortfolioShareOutput() AcceptedPortfolioShareOutput {
	return o
}

func (o AcceptedPortfolioShareOutput) ToAcceptedPortfolioShareOutputWithContext(ctx context.Context) AcceptedPortfolioShareOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AcceptedPortfolioShareOutput{})
}
