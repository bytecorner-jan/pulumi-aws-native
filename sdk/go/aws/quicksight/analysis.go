// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the AWS::QuickSight::Analysis Resource Type.
type Analysis struct {
	pulumi.CustomResourceState

	AnalysisId pulumi.StringOutput `pulumi:"analysisId"`
	// <p>The Amazon Resource Name (ARN) of the analysis.</p>
	Arn          pulumi.StringOutput `pulumi:"arn"`
	AwsAccountId pulumi.StringOutput `pulumi:"awsAccountId"`
	// <p>The time that the analysis was created.</p>
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// <p>The ARNs of the datasets of the analysis.</p>
	DataSetArns pulumi.StringArrayOutput `pulumi:"dataSetArns"`
	// <p>Errors associated with the analysis.</p>
	Errors AnalysisErrorArrayOutput `pulumi:"errors"`
	// <p>The time that the analysis was last updated.</p>
	LastUpdatedTime pulumi.StringOutput `pulumi:"lastUpdatedTime"`
	// <p>The descriptive name of the analysis.</p>
	Name       pulumi.StringPtrOutput      `pulumi:"name"`
	Parameters AnalysisParametersPtrOutput `pulumi:"parameters"`
	// <p>A structure that describes the principals and the resource-level permissions on an
	//             analysis. You can use the <code>Permissions</code> structure to grant permissions by
	//             providing a list of AWS Identity and Access Management (IAM) action information for each
	//             principal listed by Amazon Resource Name (ARN). </p>
	//
	//         <p>To specify no permissions, omit <code>Permissions</code>.</p>
	Permissions AnalysisResourcePermissionArrayOutput `pulumi:"permissions"`
	// <p>A list of the associated sheets with the unique identifier and name of each sheet.</p>
	Sheets       AnalysisSheetArrayOutput     `pulumi:"sheets"`
	SourceEntity AnalysisSourceEntityOutput   `pulumi:"sourceEntity"`
	Status       AnalysisResourceStatusOutput `pulumi:"status"`
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the
	//             analysis.</p>
	Tags AnalysisTagArrayOutput `pulumi:"tags"`
	// <p>The ARN of the theme of the analysis.</p>
	ThemeArn pulumi.StringPtrOutput `pulumi:"themeArn"`
}

// NewAnalysis registers a new resource with the given unique name, arguments, and options.
func NewAnalysis(ctx *pulumi.Context,
	name string, args *AnalysisArgs, opts ...pulumi.ResourceOption) (*Analysis, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AnalysisId == nil {
		return nil, errors.New("invalid value for required argument 'AnalysisId'")
	}
	if args.AwsAccountId == nil {
		return nil, errors.New("invalid value for required argument 'AwsAccountId'")
	}
	if args.SourceEntity == nil {
		return nil, errors.New("invalid value for required argument 'SourceEntity'")
	}
	var resource Analysis
	err := ctx.RegisterResource("aws-native:quicksight:Analysis", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalysis gets an existing Analysis resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalysis(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalysisState, opts ...pulumi.ResourceOption) (*Analysis, error) {
	var resource Analysis
	err := ctx.ReadResource("aws-native:quicksight:Analysis", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Analysis resources.
type analysisState struct {
}

type AnalysisState struct {
}

func (AnalysisState) ElementType() reflect.Type {
	return reflect.TypeOf((*analysisState)(nil)).Elem()
}

type analysisArgs struct {
	AnalysisId   string `pulumi:"analysisId"`
	AwsAccountId string `pulumi:"awsAccountId"`
	// <p>Errors associated with the analysis.</p>
	Errors []AnalysisError `pulumi:"errors"`
	// <p>The descriptive name of the analysis.</p>
	Name       *string             `pulumi:"name"`
	Parameters *AnalysisParameters `pulumi:"parameters"`
	// <p>A structure that describes the principals and the resource-level permissions on an
	//             analysis. You can use the <code>Permissions</code> structure to grant permissions by
	//             providing a list of AWS Identity and Access Management (IAM) action information for each
	//             principal listed by Amazon Resource Name (ARN). </p>
	//
	//         <p>To specify no permissions, omit <code>Permissions</code>.</p>
	Permissions  []AnalysisResourcePermission `pulumi:"permissions"`
	SourceEntity AnalysisSourceEntity         `pulumi:"sourceEntity"`
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the
	//             analysis.</p>
	Tags []AnalysisTag `pulumi:"tags"`
	// <p>The ARN of the theme of the analysis.</p>
	ThemeArn *string `pulumi:"themeArn"`
}

// The set of arguments for constructing a Analysis resource.
type AnalysisArgs struct {
	AnalysisId   pulumi.StringInput
	AwsAccountId pulumi.StringInput
	// <p>Errors associated with the analysis.</p>
	Errors AnalysisErrorArrayInput
	// <p>The descriptive name of the analysis.</p>
	Name       pulumi.StringPtrInput
	Parameters AnalysisParametersPtrInput
	// <p>A structure that describes the principals and the resource-level permissions on an
	//             analysis. You can use the <code>Permissions</code> structure to grant permissions by
	//             providing a list of AWS Identity and Access Management (IAM) action information for each
	//             principal listed by Amazon Resource Name (ARN). </p>
	//
	//         <p>To specify no permissions, omit <code>Permissions</code>.</p>
	Permissions  AnalysisResourcePermissionArrayInput
	SourceEntity AnalysisSourceEntityInput
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the
	//             analysis.</p>
	Tags AnalysisTagArrayInput
	// <p>The ARN of the theme of the analysis.</p>
	ThemeArn pulumi.StringPtrInput
}

func (AnalysisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analysisArgs)(nil)).Elem()
}

type AnalysisInput interface {
	pulumi.Input

	ToAnalysisOutput() AnalysisOutput
	ToAnalysisOutputWithContext(ctx context.Context) AnalysisOutput
}

func (*Analysis) ElementType() reflect.Type {
	return reflect.TypeOf((*Analysis)(nil))
}

func (i *Analysis) ToAnalysisOutput() AnalysisOutput {
	return i.ToAnalysisOutputWithContext(context.Background())
}

func (i *Analysis) ToAnalysisOutputWithContext(ctx context.Context) AnalysisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalysisOutput)
}

type AnalysisOutput struct{ *pulumi.OutputState }

func (AnalysisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Analysis)(nil))
}

func (o AnalysisOutput) ToAnalysisOutput() AnalysisOutput {
	return o
}

func (o AnalysisOutput) ToAnalysisOutputWithContext(ctx context.Context) AnalysisOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnalysisInput)(nil)).Elem(), &Analysis{})
	pulumi.RegisterOutputType(AnalysisOutput{})
}
