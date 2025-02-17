// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codecommit

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CodeCommit::Repository
//
// Deprecated: Repository is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Repository struct {
	pulumi.CustomResourceState

	Arn                   pulumi.StringOutput          `pulumi:"arn"`
	CloneUrlHttp          pulumi.StringOutput          `pulumi:"cloneUrlHttp"`
	CloneUrlSsh           pulumi.StringOutput          `pulumi:"cloneUrlSsh"`
	Code                  RepositoryCodePtrOutput      `pulumi:"code"`
	Name                  pulumi.StringOutput          `pulumi:"name"`
	RepositoryDescription pulumi.StringPtrOutput       `pulumi:"repositoryDescription"`
	RepositoryName        pulumi.StringOutput          `pulumi:"repositoryName"`
	Tags                  RepositoryTagArrayOutput     `pulumi:"tags"`
	Triggers              RepositoryTriggerArrayOutput `pulumi:"triggers"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil {
		args = &RepositoryArgs{}
	}

	var resource Repository
	err := ctx.RegisterResource("aws-native:codecommit:Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryState, opts ...pulumi.ResourceOption) (*Repository, error) {
	var resource Repository
	err := ctx.ReadResource("aws-native:codecommit:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
}

type RepositoryState struct {
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	Code                  *RepositoryCode     `pulumi:"code"`
	RepositoryDescription *string             `pulumi:"repositoryDescription"`
	RepositoryName        *string             `pulumi:"repositoryName"`
	Tags                  []RepositoryTag     `pulumi:"tags"`
	Triggers              []RepositoryTrigger `pulumi:"triggers"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	Code                  RepositoryCodePtrInput
	RepositoryDescription pulumi.StringPtrInput
	RepositoryName        pulumi.StringPtrInput
	Tags                  RepositoryTagArrayInput
	Triggers              RepositoryTriggerArrayInput
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryArgs)(nil)).Elem()
}

type RepositoryInput interface {
	pulumi.Input

	ToRepositoryOutput() RepositoryOutput
	ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput
}

func (*Repository) ElementType() reflect.Type {
	return reflect.TypeOf((*Repository)(nil))
}

func (i *Repository) ToRepositoryOutput() RepositoryOutput {
	return i.ToRepositoryOutputWithContext(context.Background())
}

func (i *Repository) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryOutput)
}

type RepositoryOutput struct{ *pulumi.OutputState }

func (RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Repository)(nil))
}

func (o RepositoryOutput) ToRepositoryOutput() RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryInput)(nil)).Elem(), &Repository{})
	pulumi.RegisterOutputType(RepositoryOutput{})
}
