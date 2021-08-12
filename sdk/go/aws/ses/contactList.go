// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html
type ContactList struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html#cfn-ses-contactlist-contactlistname
	ContactListName pulumi.StringPtrOutput `pulumi:"contactListName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html#cfn-ses-contactlist-description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html#cfn-ses-contactlist-tags
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html#cfn-ses-contactlist-topics
	Topics ContactListTopicArrayOutput `pulumi:"topics"`
}

// NewContactList registers a new resource with the given unique name, arguments, and options.
func NewContactList(ctx *pulumi.Context,
	name string, args *ContactListArgs, opts ...pulumi.ResourceOption) (*ContactList, error) {
	if args == nil {
		args = &ContactListArgs{}
	}

	var resource ContactList
	err := ctx.RegisterResource("aws-native:SES:ContactList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContactList gets an existing ContactList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContactList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContactListState, opts ...pulumi.ResourceOption) (*ContactList, error) {
	var resource ContactList
	err := ctx.ReadResource("aws-native:SES:ContactList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContactList resources.
type contactListState struct {
}

type ContactListState struct {
}

func (ContactListState) ElementType() reflect.Type {
	return reflect.TypeOf((*contactListState)(nil)).Elem()
}

type contactListArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html#cfn-ses-contactlist-contactlistname
	ContactListName *string `pulumi:"contactListName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html#cfn-ses-contactlist-description
	Description *string `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html#cfn-ses-contactlist-tags
	Tags []aws.Tag `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html#cfn-ses-contactlist-topics
	Topics []ContactListTopic `pulumi:"topics"`
}

// The set of arguments for constructing a ContactList resource.
type ContactListArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html#cfn-ses-contactlist-contactlistname
	ContactListName pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html#cfn-ses-contactlist-description
	Description pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html#cfn-ses-contactlist-tags
	Tags aws.TagArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-contactlist.html#cfn-ses-contactlist-topics
	Topics ContactListTopicArrayInput
}

func (ContactListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contactListArgs)(nil)).Elem()
}

type ContactListInput interface {
	pulumi.Input

	ToContactListOutput() ContactListOutput
	ToContactListOutputWithContext(ctx context.Context) ContactListOutput
}

func (*ContactList) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactList)(nil))
}

func (i *ContactList) ToContactListOutput() ContactListOutput {
	return i.ToContactListOutputWithContext(context.Background())
}

func (i *ContactList) ToContactListOutputWithContext(ctx context.Context) ContactListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactListOutput)
}

type ContactListOutput struct{ *pulumi.OutputState }

func (ContactListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactList)(nil))
}

func (o ContactListOutput) ToContactListOutput() ContactListOutput {
	return o
}

func (o ContactListOutput) ToContactListOutputWithContext(ctx context.Context) ContactListOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ContactListOutput{})
}
