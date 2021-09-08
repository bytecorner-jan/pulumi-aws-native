// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html
type Nodegroup struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-amitype
	AmiType pulumi.StringPtrOutput `pulumi:"amiType"`
	Arn     pulumi.StringOutput    `pulumi:"arn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-capacitytype
	CapacityType pulumi.StringPtrOutput `pulumi:"capacityType"`
	ClusterName  pulumi.StringOutput    `pulumi:"clusterName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-disksize
	DiskSize pulumi.Float64PtrOutput `pulumi:"diskSize"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-forceupdateenabled
	ForceUpdateEnabled pulumi.BoolPtrOutput `pulumi:"forceUpdateEnabled"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-instancetypes
	InstanceTypes pulumi.StringArrayOutput `pulumi:"instanceTypes"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-labels
	Labels pulumi.AnyOutput `pulumi:"labels"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-launchtemplate
	LaunchTemplate NodegroupLaunchTemplateSpecificationPtrOutput `pulumi:"launchTemplate"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-noderole
	NodeRole      pulumi.StringOutput `pulumi:"nodeRole"`
	NodegroupName pulumi.StringOutput `pulumi:"nodegroupName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-releaseversion
	ReleaseVersion pulumi.StringPtrOutput `pulumi:"releaseVersion"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-remoteaccess
	RemoteAccess NodegroupRemoteAccessPtrOutput `pulumi:"remoteAccess"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-scalingconfig
	ScalingConfig NodegroupScalingConfigPtrOutput `pulumi:"scalingConfig"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-subnets
	Subnets pulumi.StringArrayOutput `pulumi:"subnets"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-tags
	Tags pulumi.AnyOutput `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-taints
	Taints NodegroupTaintArrayOutput `pulumi:"taints"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-updateconfig
	UpdateConfig NodegroupUpdateConfigPtrOutput `pulumi:"updateConfig"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-version
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewNodegroup registers a new resource with the given unique name, arguments, and options.
func NewNodegroup(ctx *pulumi.Context,
	name string, args *NodegroupArgs, opts ...pulumi.ResourceOption) (*Nodegroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.NodeRole == nil {
		return nil, errors.New("invalid value for required argument 'NodeRole'")
	}
	if args.Subnets == nil {
		return nil, errors.New("invalid value for required argument 'Subnets'")
	}
	var resource Nodegroup
	err := ctx.RegisterResource("aws-native:eks:Nodegroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodegroup gets an existing Nodegroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodegroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodegroupState, opts ...pulumi.ResourceOption) (*Nodegroup, error) {
	var resource Nodegroup
	err := ctx.ReadResource("aws-native:eks:Nodegroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Nodegroup resources.
type nodegroupState struct {
}

type NodegroupState struct {
}

func (NodegroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodegroupState)(nil)).Elem()
}

type nodegroupArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-amitype
	AmiType *string `pulumi:"amiType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-capacitytype
	CapacityType *string `pulumi:"capacityType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-clustername
	ClusterName string `pulumi:"clusterName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-disksize
	DiskSize *float64 `pulumi:"diskSize"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-forceupdateenabled
	ForceUpdateEnabled *bool `pulumi:"forceUpdateEnabled"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-instancetypes
	InstanceTypes []string `pulumi:"instanceTypes"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-labels
	Labels interface{} `pulumi:"labels"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-launchtemplate
	LaunchTemplate *NodegroupLaunchTemplateSpecification `pulumi:"launchTemplate"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-noderole
	NodeRole string `pulumi:"nodeRole"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-nodegroupname
	NodegroupName *string `pulumi:"nodegroupName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-releaseversion
	ReleaseVersion *string `pulumi:"releaseVersion"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-remoteaccess
	RemoteAccess *NodegroupRemoteAccess `pulumi:"remoteAccess"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-scalingconfig
	ScalingConfig *NodegroupScalingConfig `pulumi:"scalingConfig"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-subnets
	Subnets []string `pulumi:"subnets"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-tags
	Tags interface{} `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-taints
	Taints []NodegroupTaint `pulumi:"taints"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-updateconfig
	UpdateConfig *NodegroupUpdateConfig `pulumi:"updateConfig"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-version
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Nodegroup resource.
type NodegroupArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-amitype
	AmiType pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-capacitytype
	CapacityType pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-clustername
	ClusterName pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-disksize
	DiskSize pulumi.Float64PtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-forceupdateenabled
	ForceUpdateEnabled pulumi.BoolPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-instancetypes
	InstanceTypes pulumi.StringArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-labels
	Labels pulumi.Input
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-launchtemplate
	LaunchTemplate NodegroupLaunchTemplateSpecificationPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-noderole
	NodeRole pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-nodegroupname
	NodegroupName pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-releaseversion
	ReleaseVersion pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-remoteaccess
	RemoteAccess NodegroupRemoteAccessPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-scalingconfig
	ScalingConfig NodegroupScalingConfigPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-subnets
	Subnets pulumi.StringArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-tags
	Tags pulumi.Input
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-taints
	Taints NodegroupTaintArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-updateconfig
	UpdateConfig NodegroupUpdateConfigPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-version
	Version pulumi.StringPtrInput
}

func (NodegroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodegroupArgs)(nil)).Elem()
}

type NodegroupInput interface {
	pulumi.Input

	ToNodegroupOutput() NodegroupOutput
	ToNodegroupOutputWithContext(ctx context.Context) NodegroupOutput
}

func (*Nodegroup) ElementType() reflect.Type {
	return reflect.TypeOf((*Nodegroup)(nil))
}

func (i *Nodegroup) ToNodegroupOutput() NodegroupOutput {
	return i.ToNodegroupOutputWithContext(context.Background())
}

func (i *Nodegroup) ToNodegroupOutputWithContext(ctx context.Context) NodegroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodegroupOutput)
}

type NodegroupOutput struct{ *pulumi.OutputState }

func (NodegroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Nodegroup)(nil))
}

func (o NodegroupOutput) ToNodegroupOutput() NodegroupOutput {
	return o
}

func (o NodegroupOutput) ToNodegroupOutputWithContext(ctx context.Context) NodegroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NodegroupOutput{})
}
