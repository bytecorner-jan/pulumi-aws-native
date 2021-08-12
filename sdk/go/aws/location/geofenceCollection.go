// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package location

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html
type GeofenceCollection struct {
	pulumi.CustomResourceState

	CollectionArn pulumi.StringOutput `pulumi:"collectionArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-collectionname
	CollectionName pulumi.StringOutput `pulumi:"collectionName"`
	CreateTime     pulumi.StringOutput `pulumi:"createTime"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-kmskeyid
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-pricingplan
	PricingPlan pulumi.StringOutput `pulumi:"pricingPlan"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-pricingplandatasource
	PricingPlanDataSource pulumi.StringPtrOutput `pulumi:"pricingPlanDataSource"`
	UpdateTime            pulumi.StringOutput    `pulumi:"updateTime"`
}

// NewGeofenceCollection registers a new resource with the given unique name, arguments, and options.
func NewGeofenceCollection(ctx *pulumi.Context,
	name string, args *GeofenceCollectionArgs, opts ...pulumi.ResourceOption) (*GeofenceCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CollectionName == nil {
		return nil, errors.New("invalid value for required argument 'CollectionName'")
	}
	if args.PricingPlan == nil {
		return nil, errors.New("invalid value for required argument 'PricingPlan'")
	}
	var resource GeofenceCollection
	err := ctx.RegisterResource("aws-native:Location:GeofenceCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGeofenceCollection gets an existing GeofenceCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGeofenceCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GeofenceCollectionState, opts ...pulumi.ResourceOption) (*GeofenceCollection, error) {
	var resource GeofenceCollection
	err := ctx.ReadResource("aws-native:Location:GeofenceCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GeofenceCollection resources.
type geofenceCollectionState struct {
}

type GeofenceCollectionState struct {
}

func (GeofenceCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*geofenceCollectionState)(nil)).Elem()
}

type geofenceCollectionArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-collectionname
	CollectionName string `pulumi:"collectionName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-description
	Description *string `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-kmskeyid
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-pricingplan
	PricingPlan string `pulumi:"pricingPlan"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-pricingplandatasource
	PricingPlanDataSource *string `pulumi:"pricingPlanDataSource"`
}

// The set of arguments for constructing a GeofenceCollection resource.
type GeofenceCollectionArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-collectionname
	CollectionName pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-description
	Description pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-kmskeyid
	KmsKeyId pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-pricingplan
	PricingPlan pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-pricingplandatasource
	PricingPlanDataSource pulumi.StringPtrInput
}

func (GeofenceCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*geofenceCollectionArgs)(nil)).Elem()
}

type GeofenceCollectionInput interface {
	pulumi.Input

	ToGeofenceCollectionOutput() GeofenceCollectionOutput
	ToGeofenceCollectionOutputWithContext(ctx context.Context) GeofenceCollectionOutput
}

func (*GeofenceCollection) ElementType() reflect.Type {
	return reflect.TypeOf((*GeofenceCollection)(nil))
}

func (i *GeofenceCollection) ToGeofenceCollectionOutput() GeofenceCollectionOutput {
	return i.ToGeofenceCollectionOutputWithContext(context.Background())
}

func (i *GeofenceCollection) ToGeofenceCollectionOutputWithContext(ctx context.Context) GeofenceCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeofenceCollectionOutput)
}

type GeofenceCollectionOutput struct{ *pulumi.OutputState }

func (GeofenceCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GeofenceCollection)(nil))
}

func (o GeofenceCollectionOutput) ToGeofenceCollectionOutput() GeofenceCollectionOutput {
	return o
}

func (o GeofenceCollectionOutput) ToGeofenceCollectionOutputWithContext(ctx context.Context) GeofenceCollectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GeofenceCollectionOutput{})
}
