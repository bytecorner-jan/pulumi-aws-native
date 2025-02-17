// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package auditmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An entity that defines the scope of audit evidence collected by AWS Audit Manager.
type Assessment struct {
	pulumi.CustomResourceState

	Arn                          pulumi.StringOutput                   `pulumi:"arn"`
	AssessmentId                 pulumi.StringOutput                   `pulumi:"assessmentId"`
	AssessmentReportsDestination AssessmentReportsDestinationPtrOutput `pulumi:"assessmentReportsDestination"`
	AwsAccount                   AssessmentAWSAccountPtrOutput         `pulumi:"awsAccount"`
	CreationTime                 pulumi.Float64Output                  `pulumi:"creationTime"`
	// The list of delegations.
	Delegations AssessmentDelegationArrayOutput `pulumi:"delegations"`
	Description pulumi.StringPtrOutput          `pulumi:"description"`
	FrameworkId pulumi.StringPtrOutput          `pulumi:"frameworkId"`
	Name        pulumi.StringPtrOutput          `pulumi:"name"`
	// The list of roles for the specified assessment.
	Roles  AssessmentRoleArrayOutput `pulumi:"roles"`
	Scope  AssessmentScopePtrOutput  `pulumi:"scope"`
	Status AssessmentStatusPtrOutput `pulumi:"status"`
	// The tags associated with the assessment.
	Tags AssessmentTagArrayOutput `pulumi:"tags"`
}

// NewAssessment registers a new resource with the given unique name, arguments, and options.
func NewAssessment(ctx *pulumi.Context,
	name string, args *AssessmentArgs, opts ...pulumi.ResourceOption) (*Assessment, error) {
	if args == nil {
		args = &AssessmentArgs{}
	}

	var resource Assessment
	err := ctx.RegisterResource("aws-native:auditmanager:Assessment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssessment gets an existing Assessment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssessment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssessmentState, opts ...pulumi.ResourceOption) (*Assessment, error) {
	var resource Assessment
	err := ctx.ReadResource("aws-native:auditmanager:Assessment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Assessment resources.
type assessmentState struct {
}

type AssessmentState struct {
}

func (AssessmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentState)(nil)).Elem()
}

type assessmentArgs struct {
	AssessmentReportsDestination *AssessmentReportsDestination `pulumi:"assessmentReportsDestination"`
	AwsAccount                   *AssessmentAWSAccount         `pulumi:"awsAccount"`
	Description                  *string                       `pulumi:"description"`
	FrameworkId                  *string                       `pulumi:"frameworkId"`
	Name                         *string                       `pulumi:"name"`
	// The list of roles for the specified assessment.
	Roles  []AssessmentRole  `pulumi:"roles"`
	Scope  *AssessmentScope  `pulumi:"scope"`
	Status *AssessmentStatus `pulumi:"status"`
	// The tags associated with the assessment.
	Tags []AssessmentTag `pulumi:"tags"`
}

// The set of arguments for constructing a Assessment resource.
type AssessmentArgs struct {
	AssessmentReportsDestination AssessmentReportsDestinationPtrInput
	AwsAccount                   AssessmentAWSAccountPtrInput
	Description                  pulumi.StringPtrInput
	FrameworkId                  pulumi.StringPtrInput
	Name                         pulumi.StringPtrInput
	// The list of roles for the specified assessment.
	Roles  AssessmentRoleArrayInput
	Scope  AssessmentScopePtrInput
	Status AssessmentStatusPtrInput
	// The tags associated with the assessment.
	Tags AssessmentTagArrayInput
}

func (AssessmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentArgs)(nil)).Elem()
}

type AssessmentInput interface {
	pulumi.Input

	ToAssessmentOutput() AssessmentOutput
	ToAssessmentOutputWithContext(ctx context.Context) AssessmentOutput
}

func (*Assessment) ElementType() reflect.Type {
	return reflect.TypeOf((*Assessment)(nil))
}

func (i *Assessment) ToAssessmentOutput() AssessmentOutput {
	return i.ToAssessmentOutputWithContext(context.Background())
}

func (i *Assessment) ToAssessmentOutputWithContext(ctx context.Context) AssessmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentOutput)
}

type AssessmentOutput struct{ *pulumi.OutputState }

func (AssessmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Assessment)(nil))
}

func (o AssessmentOutput) ToAssessmentOutput() AssessmentOutput {
	return o
}

func (o AssessmentOutput) ToAssessmentOutputWithContext(ctx context.Context) AssessmentOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentInput)(nil)).Elem(), &Assessment{})
	pulumi.RegisterOutputType(AssessmentOutput{})
}
