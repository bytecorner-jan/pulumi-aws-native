// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html
type BackupVault struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-accesspolicy
	AccessPolicy    pulumi.AnyOutput    `pulumi:"accessPolicy"`
	BackupVaultArn  pulumi.StringOutput `pulumi:"backupVaultArn"`
	BackupVaultName pulumi.StringOutput `pulumi:"backupVaultName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-backupvaulttags
	BackupVaultTags pulumi.StringMapOutput `pulumi:"backupVaultTags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-encryptionkeyarn
	EncryptionKeyArn pulumi.StringPtrOutput `pulumi:"encryptionKeyArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-notifications
	Notifications BackupVaultNotificationObjectTypePtrOutput `pulumi:"notifications"`
}

// NewBackupVault registers a new resource with the given unique name, arguments, and options.
func NewBackupVault(ctx *pulumi.Context,
	name string, args *BackupVaultArgs, opts ...pulumi.ResourceOption) (*BackupVault, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupVaultName == nil {
		return nil, errors.New("invalid value for required argument 'BackupVaultName'")
	}
	var resource BackupVault
	err := ctx.RegisterResource("aws-native:Backup:BackupVault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupVault gets an existing BackupVault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupVaultState, opts ...pulumi.ResourceOption) (*BackupVault, error) {
	var resource BackupVault
	err := ctx.ReadResource("aws-native:Backup:BackupVault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupVault resources.
type backupVaultState struct {
}

type BackupVaultState struct {
}

func (BackupVaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupVaultState)(nil)).Elem()
}

type backupVaultArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-accesspolicy
	AccessPolicy interface{} `pulumi:"accessPolicy"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-backupvaultname
	BackupVaultName string `pulumi:"backupVaultName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-backupvaulttags
	BackupVaultTags map[string]string `pulumi:"backupVaultTags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-encryptionkeyarn
	EncryptionKeyArn *string `pulumi:"encryptionKeyArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-notifications
	Notifications *BackupVaultNotificationObjectType `pulumi:"notifications"`
}

// The set of arguments for constructing a BackupVault resource.
type BackupVaultArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-accesspolicy
	AccessPolicy pulumi.Input
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-backupvaultname
	BackupVaultName pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-backupvaulttags
	BackupVaultTags pulumi.StringMapInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-encryptionkeyarn
	EncryptionKeyArn pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-notifications
	Notifications BackupVaultNotificationObjectTypePtrInput
}

func (BackupVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupVaultArgs)(nil)).Elem()
}

type BackupVaultInput interface {
	pulumi.Input

	ToBackupVaultOutput() BackupVaultOutput
	ToBackupVaultOutputWithContext(ctx context.Context) BackupVaultOutput
}

func (*BackupVault) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupVault)(nil))
}

func (i *BackupVault) ToBackupVaultOutput() BackupVaultOutput {
	return i.ToBackupVaultOutputWithContext(context.Background())
}

func (i *BackupVault) ToBackupVaultOutputWithContext(ctx context.Context) BackupVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupVaultOutput)
}

type BackupVaultOutput struct{ *pulumi.OutputState }

func (BackupVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupVault)(nil))
}

func (o BackupVaultOutput) ToBackupVaultOutput() BackupVaultOutput {
	return o
}

func (o BackupVaultOutput) ToBackupVaultOutputWithContext(ctx context.Context) BackupVaultOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BackupVaultOutput{})
}
