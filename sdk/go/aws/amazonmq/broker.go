// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package amazonmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html
type Broker struct {
	pulumi.CustomResourceState

	AmqpEndpoints pulumi.StringArrayOutput `pulumi:"amqpEndpoints"`
	Arn           pulumi.StringOutput      `pulumi:"arn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-authenticationstrategy
	AuthenticationStrategy pulumi.StringPtrOutput `pulumi:"authenticationStrategy"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-autominorversionupgrade
	AutoMinorVersionUpgrade pulumi.BoolOutput `pulumi:"autoMinorVersionUpgrade"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-brokername
	BrokerName pulumi.StringOutput `pulumi:"brokerName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-configuration
	Configuration         BrokerConfigurationIdPtrOutput `pulumi:"configuration"`
	ConfigurationId       pulumi.StringOutput            `pulumi:"configurationId"`
	ConfigurationRevision pulumi.IntOutput               `pulumi:"configurationRevision"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-deploymentmode
	DeploymentMode pulumi.StringOutput `pulumi:"deploymentMode"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-encryptionoptions
	EncryptionOptions BrokerEncryptionOptionsPtrOutput `pulumi:"encryptionOptions"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-enginetype
	EngineType pulumi.StringOutput `pulumi:"engineType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-engineversion
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-hostinstancetype
	HostInstanceType pulumi.StringOutput      `pulumi:"hostInstanceType"`
	IpAddresses      pulumi.StringArrayOutput `pulumi:"ipAddresses"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-ldapservermetadata
	LdapServerMetadata BrokerLdapServerMetadataPtrOutput `pulumi:"ldapServerMetadata"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-logs
	Logs BrokerLogListPtrOutput `pulumi:"logs"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-maintenancewindowstarttime
	MaintenanceWindowStartTime BrokerMaintenanceWindowPtrOutput `pulumi:"maintenanceWindowStartTime"`
	MqttEndpoints              pulumi.StringArrayOutput         `pulumi:"mqttEndpoints"`
	OpenWireEndpoints          pulumi.StringArrayOutput         `pulumi:"openWireEndpoints"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-publiclyaccessible
	PubliclyAccessible pulumi.BoolOutput `pulumi:"publiclyAccessible"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-securitygroups
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	StompEndpoints pulumi.StringArrayOutput `pulumi:"stompEndpoints"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-storagetype
	StorageType pulumi.StringPtrOutput `pulumi:"storageType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-subnetids
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-tags
	Tags BrokerTagsEntryArrayOutput `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-users
	Users        BrokerUserArrayOutput    `pulumi:"users"`
	WssEndpoints pulumi.StringArrayOutput `pulumi:"wssEndpoints"`
}

// NewBroker registers a new resource with the given unique name, arguments, and options.
func NewBroker(ctx *pulumi.Context,
	name string, args *BrokerArgs, opts ...pulumi.ResourceOption) (*Broker, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoMinorVersionUpgrade == nil {
		return nil, errors.New("invalid value for required argument 'AutoMinorVersionUpgrade'")
	}
	if args.BrokerName == nil {
		return nil, errors.New("invalid value for required argument 'BrokerName'")
	}
	if args.DeploymentMode == nil {
		return nil, errors.New("invalid value for required argument 'DeploymentMode'")
	}
	if args.EngineType == nil {
		return nil, errors.New("invalid value for required argument 'EngineType'")
	}
	if args.EngineVersion == nil {
		return nil, errors.New("invalid value for required argument 'EngineVersion'")
	}
	if args.HostInstanceType == nil {
		return nil, errors.New("invalid value for required argument 'HostInstanceType'")
	}
	if args.PubliclyAccessible == nil {
		return nil, errors.New("invalid value for required argument 'PubliclyAccessible'")
	}
	if args.Users == nil {
		return nil, errors.New("invalid value for required argument 'Users'")
	}
	var resource Broker
	err := ctx.RegisterResource("aws-native:amazonmq:Broker", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBroker gets an existing Broker resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBroker(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BrokerState, opts ...pulumi.ResourceOption) (*Broker, error) {
	var resource Broker
	err := ctx.ReadResource("aws-native:amazonmq:Broker", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Broker resources.
type brokerState struct {
}

type BrokerState struct {
}

func (BrokerState) ElementType() reflect.Type {
	return reflect.TypeOf((*brokerState)(nil)).Elem()
}

type brokerArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-authenticationstrategy
	AuthenticationStrategy *string `pulumi:"authenticationStrategy"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-autominorversionupgrade
	AutoMinorVersionUpgrade bool `pulumi:"autoMinorVersionUpgrade"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-brokername
	BrokerName string `pulumi:"brokerName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-configuration
	Configuration *BrokerConfigurationId `pulumi:"configuration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-deploymentmode
	DeploymentMode string `pulumi:"deploymentMode"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-encryptionoptions
	EncryptionOptions *BrokerEncryptionOptions `pulumi:"encryptionOptions"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-enginetype
	EngineType string `pulumi:"engineType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-engineversion
	EngineVersion string `pulumi:"engineVersion"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-hostinstancetype
	HostInstanceType string `pulumi:"hostInstanceType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-ldapservermetadata
	LdapServerMetadata *BrokerLdapServerMetadata `pulumi:"ldapServerMetadata"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-logs
	Logs *BrokerLogList `pulumi:"logs"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-maintenancewindowstarttime
	MaintenanceWindowStartTime *BrokerMaintenanceWindow `pulumi:"maintenanceWindowStartTime"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-publiclyaccessible
	PubliclyAccessible bool `pulumi:"publiclyAccessible"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-securitygroups
	SecurityGroups []string `pulumi:"securityGroups"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-storagetype
	StorageType *string `pulumi:"storageType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-subnetids
	SubnetIds []string `pulumi:"subnetIds"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-tags
	Tags []BrokerTagsEntry `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-users
	Users []BrokerUser `pulumi:"users"`
}

// The set of arguments for constructing a Broker resource.
type BrokerArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-authenticationstrategy
	AuthenticationStrategy pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-autominorversionupgrade
	AutoMinorVersionUpgrade pulumi.BoolInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-brokername
	BrokerName pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-configuration
	Configuration BrokerConfigurationIdPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-deploymentmode
	DeploymentMode pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-encryptionoptions
	EncryptionOptions BrokerEncryptionOptionsPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-enginetype
	EngineType pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-engineversion
	EngineVersion pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-hostinstancetype
	HostInstanceType pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-ldapservermetadata
	LdapServerMetadata BrokerLdapServerMetadataPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-logs
	Logs BrokerLogListPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-maintenancewindowstarttime
	MaintenanceWindowStartTime BrokerMaintenanceWindowPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-publiclyaccessible
	PubliclyAccessible pulumi.BoolInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-securitygroups
	SecurityGroups pulumi.StringArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-storagetype
	StorageType pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-subnetids
	SubnetIds pulumi.StringArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-tags
	Tags BrokerTagsEntryArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-users
	Users BrokerUserArrayInput
}

func (BrokerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*brokerArgs)(nil)).Elem()
}

type BrokerInput interface {
	pulumi.Input

	ToBrokerOutput() BrokerOutput
	ToBrokerOutputWithContext(ctx context.Context) BrokerOutput
}

func (*Broker) ElementType() reflect.Type {
	return reflect.TypeOf((*Broker)(nil))
}

func (i *Broker) ToBrokerOutput() BrokerOutput {
	return i.ToBrokerOutputWithContext(context.Background())
}

func (i *Broker) ToBrokerOutputWithContext(ctx context.Context) BrokerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BrokerOutput)
}

type BrokerOutput struct{ *pulumi.OutputState }

func (BrokerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Broker)(nil))
}

func (o BrokerOutput) ToBrokerOutput() BrokerOutput {
	return o
}

func (o BrokerOutput) ToBrokerOutputWithContext(ctx context.Context) BrokerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BrokerOutput{})
}
