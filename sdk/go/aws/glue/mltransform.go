// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html
type MLTransform struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-glueversion
	GlueVersion pulumi.StringPtrOutput `pulumi:"glueVersion"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-inputrecordtables
	InputRecordTables MLTransformInputRecordTablesOutput `pulumi:"inputRecordTables"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-maxcapacity
	MaxCapacity pulumi.Float64PtrOutput `pulumi:"maxCapacity"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-maxretries
	MaxRetries pulumi.IntPtrOutput `pulumi:"maxRetries"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-name
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-numberofworkers
	NumberOfWorkers pulumi.IntPtrOutput `pulumi:"numberOfWorkers"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-role
	Role pulumi.StringOutput `pulumi:"role"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-tags
	Tags pulumi.AnyOutput `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-timeout
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-transformencryption
	TransformEncryption MLTransformTransformEncryptionPtrOutput `pulumi:"transformEncryption"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-transformparameters
	TransformParameters MLTransformTransformParametersOutput `pulumi:"transformParameters"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-workertype
	WorkerType pulumi.StringPtrOutput `pulumi:"workerType"`
}

// NewMLTransform registers a new resource with the given unique name, arguments, and options.
func NewMLTransform(ctx *pulumi.Context,
	name string, args *MLTransformArgs, opts ...pulumi.ResourceOption) (*MLTransform, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InputRecordTables == nil {
		return nil, errors.New("invalid value for required argument 'InputRecordTables'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.TransformParameters == nil {
		return nil, errors.New("invalid value for required argument 'TransformParameters'")
	}
	var resource MLTransform
	err := ctx.RegisterResource("aws-native:glue:MLTransform", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMLTransform gets an existing MLTransform resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMLTransform(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MLTransformState, opts ...pulumi.ResourceOption) (*MLTransform, error) {
	var resource MLTransform
	err := ctx.ReadResource("aws-native:glue:MLTransform", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MLTransform resources.
type mltransformState struct {
}

type MLTransformState struct {
}

func (MLTransformState) ElementType() reflect.Type {
	return reflect.TypeOf((*mltransformState)(nil)).Elem()
}

type mltransformArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-description
	Description *string `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-glueversion
	GlueVersion *string `pulumi:"glueVersion"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-inputrecordtables
	InputRecordTables MLTransformInputRecordTables `pulumi:"inputRecordTables"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-maxcapacity
	MaxCapacity *float64 `pulumi:"maxCapacity"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-maxretries
	MaxRetries *int `pulumi:"maxRetries"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-name
	Name *string `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-numberofworkers
	NumberOfWorkers *int `pulumi:"numberOfWorkers"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-role
	Role string `pulumi:"role"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-tags
	Tags interface{} `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-timeout
	Timeout *int `pulumi:"timeout"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-transformencryption
	TransformEncryption *MLTransformTransformEncryption `pulumi:"transformEncryption"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-transformparameters
	TransformParameters MLTransformTransformParameters `pulumi:"transformParameters"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-workertype
	WorkerType *string `pulumi:"workerType"`
}

// The set of arguments for constructing a MLTransform resource.
type MLTransformArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-description
	Description pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-glueversion
	GlueVersion pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-inputrecordtables
	InputRecordTables MLTransformInputRecordTablesInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-maxcapacity
	MaxCapacity pulumi.Float64PtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-maxretries
	MaxRetries pulumi.IntPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-name
	Name pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-numberofworkers
	NumberOfWorkers pulumi.IntPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-role
	Role pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-tags
	Tags pulumi.Input
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-timeout
	Timeout pulumi.IntPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-transformencryption
	TransformEncryption MLTransformTransformEncryptionPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-transformparameters
	TransformParameters MLTransformTransformParametersInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-workertype
	WorkerType pulumi.StringPtrInput
}

func (MLTransformArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mltransformArgs)(nil)).Elem()
}

type MLTransformInput interface {
	pulumi.Input

	ToMLTransformOutput() MLTransformOutput
	ToMLTransformOutputWithContext(ctx context.Context) MLTransformOutput
}

func (*MLTransform) ElementType() reflect.Type {
	return reflect.TypeOf((*MLTransform)(nil))
}

func (i *MLTransform) ToMLTransformOutput() MLTransformOutput {
	return i.ToMLTransformOutputWithContext(context.Background())
}

func (i *MLTransform) ToMLTransformOutputWithContext(ctx context.Context) MLTransformOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MLTransformOutput)
}

type MLTransformOutput struct{ *pulumi.OutputState }

func (MLTransformOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MLTransform)(nil))
}

func (o MLTransformOutput) ToMLTransformOutput() MLTransformOutput {
	return o
}

func (o MLTransformOutput) ToMLTransformOutputWithContext(ctx context.Context) MLTransformOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MLTransformOutput{})
}
