// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ElastiCache::ReplicationGroup
//
// Deprecated: ReplicationGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ReplicationGroup struct {
	pulumi.CustomResourceState

	AtRestEncryptionEnabled      pulumi.BoolPtrOutput                                       `pulumi:"atRestEncryptionEnabled"`
	AuthToken                    pulumi.StringPtrOutput                                     `pulumi:"authToken"`
	AutoMinorVersionUpgrade      pulumi.BoolPtrOutput                                       `pulumi:"autoMinorVersionUpgrade"`
	AutomaticFailoverEnabled     pulumi.BoolPtrOutput                                       `pulumi:"automaticFailoverEnabled"`
	CacheNodeType                pulumi.StringPtrOutput                                     `pulumi:"cacheNodeType"`
	CacheParameterGroupName      pulumi.StringPtrOutput                                     `pulumi:"cacheParameterGroupName"`
	CacheSecurityGroupNames      pulumi.StringArrayOutput                                   `pulumi:"cacheSecurityGroupNames"`
	CacheSubnetGroupName         pulumi.StringPtrOutput                                     `pulumi:"cacheSubnetGroupName"`
	ConfigurationEndPointAddress pulumi.StringPtrOutput                                     `pulumi:"configurationEndPointAddress"`
	ConfigurationEndPointPort    pulumi.StringPtrOutput                                     `pulumi:"configurationEndPointPort"`
	Engine                       pulumi.StringPtrOutput                                     `pulumi:"engine"`
	EngineVersion                pulumi.StringPtrOutput                                     `pulumi:"engineVersion"`
	GlobalReplicationGroupId     pulumi.StringPtrOutput                                     `pulumi:"globalReplicationGroupId"`
	KmsKeyId                     pulumi.StringPtrOutput                                     `pulumi:"kmsKeyId"`
	LogDeliveryConfigurations    ReplicationGroupLogDeliveryConfigurationRequestArrayOutput `pulumi:"logDeliveryConfigurations"`
	MultiAZEnabled               pulumi.BoolPtrOutput                                       `pulumi:"multiAZEnabled"`
	NodeGroupConfiguration       ReplicationGroupNodeGroupConfigurationArrayOutput          `pulumi:"nodeGroupConfiguration"`
	NotificationTopicArn         pulumi.StringPtrOutput                                     `pulumi:"notificationTopicArn"`
	NumCacheClusters             pulumi.IntPtrOutput                                        `pulumi:"numCacheClusters"`
	NumNodeGroups                pulumi.IntPtrOutput                                        `pulumi:"numNodeGroups"`
	Port                         pulumi.IntPtrOutput                                        `pulumi:"port"`
	PreferredCacheClusterAZs     pulumi.StringArrayOutput                                   `pulumi:"preferredCacheClusterAZs"`
	PreferredMaintenanceWindow   pulumi.StringPtrOutput                                     `pulumi:"preferredMaintenanceWindow"`
	PrimaryClusterId             pulumi.StringPtrOutput                                     `pulumi:"primaryClusterId"`
	PrimaryEndPointAddress       pulumi.StringPtrOutput                                     `pulumi:"primaryEndPointAddress"`
	PrimaryEndPointPort          pulumi.StringPtrOutput                                     `pulumi:"primaryEndPointPort"`
	ReadEndPointAddresses        pulumi.StringPtrOutput                                     `pulumi:"readEndPointAddresses"`
	ReadEndPointAddressesList    pulumi.StringArrayOutput                                   `pulumi:"readEndPointAddressesList"`
	ReadEndPointPorts            pulumi.StringPtrOutput                                     `pulumi:"readEndPointPorts"`
	ReadEndPointPortsList        pulumi.StringArrayOutput                                   `pulumi:"readEndPointPortsList"`
	ReaderEndPointAddress        pulumi.StringPtrOutput                                     `pulumi:"readerEndPointAddress"`
	ReaderEndPointPort           pulumi.StringPtrOutput                                     `pulumi:"readerEndPointPort"`
	ReplicasPerNodeGroup         pulumi.IntPtrOutput                                        `pulumi:"replicasPerNodeGroup"`
	ReplicationGroupDescription  pulumi.StringOutput                                        `pulumi:"replicationGroupDescription"`
	ReplicationGroupId           pulumi.StringOutput                                        `pulumi:"replicationGroupId"`
	SecurityGroupIds             pulumi.StringArrayOutput                                   `pulumi:"securityGroupIds"`
	SnapshotArns                 pulumi.StringArrayOutput                                   `pulumi:"snapshotArns"`
	SnapshotName                 pulumi.StringPtrOutput                                     `pulumi:"snapshotName"`
	SnapshotRetentionLimit       pulumi.IntPtrOutput                                        `pulumi:"snapshotRetentionLimit"`
	SnapshotWindow               pulumi.StringPtrOutput                                     `pulumi:"snapshotWindow"`
	SnapshottingClusterId        pulumi.StringPtrOutput                                     `pulumi:"snapshottingClusterId"`
	Tags                         ReplicationGroupTagArrayOutput                             `pulumi:"tags"`
	TransitEncryptionEnabled     pulumi.BoolPtrOutput                                       `pulumi:"transitEncryptionEnabled"`
	UserGroupIds                 pulumi.StringArrayOutput                                   `pulumi:"userGroupIds"`
}

// NewReplicationGroup registers a new resource with the given unique name, arguments, and options.
func NewReplicationGroup(ctx *pulumi.Context,
	name string, args *ReplicationGroupArgs, opts ...pulumi.ResourceOption) (*ReplicationGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ReplicationGroupDescription == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationGroupDescription'")
	}
	var resource ReplicationGroup
	err := ctx.RegisterResource("aws-native:elasticache:ReplicationGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationGroup gets an existing ReplicationGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationGroupState, opts ...pulumi.ResourceOption) (*ReplicationGroup, error) {
	var resource ReplicationGroup
	err := ctx.ReadResource("aws-native:elasticache:ReplicationGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationGroup resources.
type replicationGroupState struct {
}

type ReplicationGroupState struct {
}

func (ReplicationGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationGroupState)(nil)).Elem()
}

type replicationGroupArgs struct {
	AtRestEncryptionEnabled      *bool                                             `pulumi:"atRestEncryptionEnabled"`
	AuthToken                    *string                                           `pulumi:"authToken"`
	AutoMinorVersionUpgrade      *bool                                             `pulumi:"autoMinorVersionUpgrade"`
	AutomaticFailoverEnabled     *bool                                             `pulumi:"automaticFailoverEnabled"`
	CacheNodeType                *string                                           `pulumi:"cacheNodeType"`
	CacheParameterGroupName      *string                                           `pulumi:"cacheParameterGroupName"`
	CacheSecurityGroupNames      []string                                          `pulumi:"cacheSecurityGroupNames"`
	CacheSubnetGroupName         *string                                           `pulumi:"cacheSubnetGroupName"`
	ConfigurationEndPointAddress *string                                           `pulumi:"configurationEndPointAddress"`
	ConfigurationEndPointPort    *string                                           `pulumi:"configurationEndPointPort"`
	Engine                       *string                                           `pulumi:"engine"`
	EngineVersion                *string                                           `pulumi:"engineVersion"`
	GlobalReplicationGroupId     *string                                           `pulumi:"globalReplicationGroupId"`
	KmsKeyId                     *string                                           `pulumi:"kmsKeyId"`
	LogDeliveryConfigurations    []ReplicationGroupLogDeliveryConfigurationRequest `pulumi:"logDeliveryConfigurations"`
	MultiAZEnabled               *bool                                             `pulumi:"multiAZEnabled"`
	NodeGroupConfiguration       []ReplicationGroupNodeGroupConfiguration          `pulumi:"nodeGroupConfiguration"`
	NotificationTopicArn         *string                                           `pulumi:"notificationTopicArn"`
	NumCacheClusters             *int                                              `pulumi:"numCacheClusters"`
	NumNodeGroups                *int                                              `pulumi:"numNodeGroups"`
	Port                         *int                                              `pulumi:"port"`
	PreferredCacheClusterAZs     []string                                          `pulumi:"preferredCacheClusterAZs"`
	PreferredMaintenanceWindow   *string                                           `pulumi:"preferredMaintenanceWindow"`
	PrimaryClusterId             *string                                           `pulumi:"primaryClusterId"`
	PrimaryEndPointAddress       *string                                           `pulumi:"primaryEndPointAddress"`
	PrimaryEndPointPort          *string                                           `pulumi:"primaryEndPointPort"`
	ReadEndPointAddresses        *string                                           `pulumi:"readEndPointAddresses"`
	ReadEndPointAddressesList    []string                                          `pulumi:"readEndPointAddressesList"`
	ReadEndPointPorts            *string                                           `pulumi:"readEndPointPorts"`
	ReadEndPointPortsList        []string                                          `pulumi:"readEndPointPortsList"`
	ReaderEndPointAddress        *string                                           `pulumi:"readerEndPointAddress"`
	ReaderEndPointPort           *string                                           `pulumi:"readerEndPointPort"`
	ReplicasPerNodeGroup         *int                                              `pulumi:"replicasPerNodeGroup"`
	ReplicationGroupDescription  string                                            `pulumi:"replicationGroupDescription"`
	SecurityGroupIds             []string                                          `pulumi:"securityGroupIds"`
	SnapshotArns                 []string                                          `pulumi:"snapshotArns"`
	SnapshotName                 *string                                           `pulumi:"snapshotName"`
	SnapshotRetentionLimit       *int                                              `pulumi:"snapshotRetentionLimit"`
	SnapshotWindow               *string                                           `pulumi:"snapshotWindow"`
	SnapshottingClusterId        *string                                           `pulumi:"snapshottingClusterId"`
	Tags                         []ReplicationGroupTag                             `pulumi:"tags"`
	TransitEncryptionEnabled     *bool                                             `pulumi:"transitEncryptionEnabled"`
	UserGroupIds                 []string                                          `pulumi:"userGroupIds"`
}

// The set of arguments for constructing a ReplicationGroup resource.
type ReplicationGroupArgs struct {
	AtRestEncryptionEnabled      pulumi.BoolPtrInput
	AuthToken                    pulumi.StringPtrInput
	AutoMinorVersionUpgrade      pulumi.BoolPtrInput
	AutomaticFailoverEnabled     pulumi.BoolPtrInput
	CacheNodeType                pulumi.StringPtrInput
	CacheParameterGroupName      pulumi.StringPtrInput
	CacheSecurityGroupNames      pulumi.StringArrayInput
	CacheSubnetGroupName         pulumi.StringPtrInput
	ConfigurationEndPointAddress pulumi.StringPtrInput
	ConfigurationEndPointPort    pulumi.StringPtrInput
	Engine                       pulumi.StringPtrInput
	EngineVersion                pulumi.StringPtrInput
	GlobalReplicationGroupId     pulumi.StringPtrInput
	KmsKeyId                     pulumi.StringPtrInput
	LogDeliveryConfigurations    ReplicationGroupLogDeliveryConfigurationRequestArrayInput
	MultiAZEnabled               pulumi.BoolPtrInput
	NodeGroupConfiguration       ReplicationGroupNodeGroupConfigurationArrayInput
	NotificationTopicArn         pulumi.StringPtrInput
	NumCacheClusters             pulumi.IntPtrInput
	NumNodeGroups                pulumi.IntPtrInput
	Port                         pulumi.IntPtrInput
	PreferredCacheClusterAZs     pulumi.StringArrayInput
	PreferredMaintenanceWindow   pulumi.StringPtrInput
	PrimaryClusterId             pulumi.StringPtrInput
	PrimaryEndPointAddress       pulumi.StringPtrInput
	PrimaryEndPointPort          pulumi.StringPtrInput
	ReadEndPointAddresses        pulumi.StringPtrInput
	ReadEndPointAddressesList    pulumi.StringArrayInput
	ReadEndPointPorts            pulumi.StringPtrInput
	ReadEndPointPortsList        pulumi.StringArrayInput
	ReaderEndPointAddress        pulumi.StringPtrInput
	ReaderEndPointPort           pulumi.StringPtrInput
	ReplicasPerNodeGroup         pulumi.IntPtrInput
	ReplicationGroupDescription  pulumi.StringInput
	SecurityGroupIds             pulumi.StringArrayInput
	SnapshotArns                 pulumi.StringArrayInput
	SnapshotName                 pulumi.StringPtrInput
	SnapshotRetentionLimit       pulumi.IntPtrInput
	SnapshotWindow               pulumi.StringPtrInput
	SnapshottingClusterId        pulumi.StringPtrInput
	Tags                         ReplicationGroupTagArrayInput
	TransitEncryptionEnabled     pulumi.BoolPtrInput
	UserGroupIds                 pulumi.StringArrayInput
}

func (ReplicationGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationGroupArgs)(nil)).Elem()
}

type ReplicationGroupInput interface {
	pulumi.Input

	ToReplicationGroupOutput() ReplicationGroupOutput
	ToReplicationGroupOutputWithContext(ctx context.Context) ReplicationGroupOutput
}

func (*ReplicationGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationGroup)(nil))
}

func (i *ReplicationGroup) ToReplicationGroupOutput() ReplicationGroupOutput {
	return i.ToReplicationGroupOutputWithContext(context.Background())
}

func (i *ReplicationGroup) ToReplicationGroupOutputWithContext(ctx context.Context) ReplicationGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationGroupOutput)
}

type ReplicationGroupOutput struct{ *pulumi.OutputState }

func (ReplicationGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationGroup)(nil))
}

func (o ReplicationGroupOutput) ToReplicationGroupOutput() ReplicationGroupOutput {
	return o
}

func (o ReplicationGroupOutput) ToReplicationGroupOutputWithContext(ctx context.Context) ReplicationGroupOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationGroupInput)(nil)).Elem(), &ReplicationGroup{})
	pulumi.RegisterOutputType(ReplicationGroupOutput{})
}
