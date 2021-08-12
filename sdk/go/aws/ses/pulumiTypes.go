// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html
type ContactListTopic struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html#cfn-ses-contactlist-topic-defaultsubscriptionstatus
	DefaultSubscriptionStatus string `pulumi:"defaultSubscriptionStatus"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html#cfn-ses-contactlist-topic-description
	Description *string `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html#cfn-ses-contactlist-topic-displayname
	DisplayName string `pulumi:"displayName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html#cfn-ses-contactlist-topic-topicname
	TopicName string `pulumi:"topicName"`
}

// ContactListTopicInput is an input type that accepts ContactListTopicArgs and ContactListTopicOutput values.
// You can construct a concrete instance of `ContactListTopicInput` via:
//
//          ContactListTopicArgs{...}
type ContactListTopicInput interface {
	pulumi.Input

	ToContactListTopicOutput() ContactListTopicOutput
	ToContactListTopicOutputWithContext(context.Context) ContactListTopicOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html
type ContactListTopicArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html#cfn-ses-contactlist-topic-defaultsubscriptionstatus
	DefaultSubscriptionStatus pulumi.StringInput `pulumi:"defaultSubscriptionStatus"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html#cfn-ses-contactlist-topic-description
	Description pulumi.StringPtrInput `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html#cfn-ses-contactlist-topic-displayname
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html#cfn-ses-contactlist-topic-topicname
	TopicName pulumi.StringInput `pulumi:"topicName"`
}

func (ContactListTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactListTopic)(nil)).Elem()
}

func (i ContactListTopicArgs) ToContactListTopicOutput() ContactListTopicOutput {
	return i.ToContactListTopicOutputWithContext(context.Background())
}

func (i ContactListTopicArgs) ToContactListTopicOutputWithContext(ctx context.Context) ContactListTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactListTopicOutput)
}

// ContactListTopicArrayInput is an input type that accepts ContactListTopicArray and ContactListTopicArrayOutput values.
// You can construct a concrete instance of `ContactListTopicArrayInput` via:
//
//          ContactListTopicArray{ ContactListTopicArgs{...} }
type ContactListTopicArrayInput interface {
	pulumi.Input

	ToContactListTopicArrayOutput() ContactListTopicArrayOutput
	ToContactListTopicArrayOutputWithContext(context.Context) ContactListTopicArrayOutput
}

type ContactListTopicArray []ContactListTopicInput

func (ContactListTopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContactListTopic)(nil)).Elem()
}

func (i ContactListTopicArray) ToContactListTopicArrayOutput() ContactListTopicArrayOutput {
	return i.ToContactListTopicArrayOutputWithContext(context.Background())
}

func (i ContactListTopicArray) ToContactListTopicArrayOutputWithContext(ctx context.Context) ContactListTopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactListTopicArrayOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html
type ContactListTopicOutput struct{ *pulumi.OutputState }

func (ContactListTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactListTopic)(nil)).Elem()
}

func (o ContactListTopicOutput) ToContactListTopicOutput() ContactListTopicOutput {
	return o
}

func (o ContactListTopicOutput) ToContactListTopicOutputWithContext(ctx context.Context) ContactListTopicOutput {
	return o
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html#cfn-ses-contactlist-topic-defaultsubscriptionstatus
func (o ContactListTopicOutput) DefaultSubscriptionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v ContactListTopic) string { return v.DefaultSubscriptionStatus }).(pulumi.StringOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html#cfn-ses-contactlist-topic-description
func (o ContactListTopicOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactListTopic) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html#cfn-ses-contactlist-topic-displayname
func (o ContactListTopicOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ContactListTopic) string { return v.DisplayName }).(pulumi.StringOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-contactlist-topic.html#cfn-ses-contactlist-topic-topicname
func (o ContactListTopicOutput) TopicName() pulumi.StringOutput {
	return o.ApplyT(func(v ContactListTopic) string { return v.TopicName }).(pulumi.StringOutput)
}

type ContactListTopicArrayOutput struct{ *pulumi.OutputState }

func (ContactListTopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContactListTopic)(nil)).Elem()
}

func (o ContactListTopicArrayOutput) ToContactListTopicArrayOutput() ContactListTopicArrayOutput {
	return o
}

func (o ContactListTopicArrayOutput) ToContactListTopicArrayOutputWithContext(ctx context.Context) ContactListTopicArrayOutput {
	return o
}

func (o ContactListTopicArrayOutput) Index(i pulumi.IntInput) ContactListTopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContactListTopic {
		return vs[0].([]ContactListTopic)[vs[1].(int)]
	}).(ContactListTopicOutput)
}

func init() {
	pulumi.RegisterOutputType(ContactListTopicOutput{})
	pulumi.RegisterOutputType(ContactListTopicArrayOutput{})
}
