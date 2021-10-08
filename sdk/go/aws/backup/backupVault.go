// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Backup::BackupVault
type BackupVault struct {
	pulumi.CustomResourceState

	AccessPolicy      pulumi.AnyOutput                           `pulumi:"accessPolicy"`
	BackupVaultArn    pulumi.StringOutput                        `pulumi:"backupVaultArn"`
	BackupVaultName   pulumi.StringOutput                        `pulumi:"backupVaultName"`
	BackupVaultTags   pulumi.AnyOutput                           `pulumi:"backupVaultTags"`
	EncryptionKeyArn  pulumi.StringPtrOutput                     `pulumi:"encryptionKeyArn"`
	LockConfiguration BackupVaultLockConfigurationTypePtrOutput  `pulumi:"lockConfiguration"`
	Notifications     BackupVaultNotificationObjectTypePtrOutput `pulumi:"notifications"`
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
	err := ctx.RegisterResource("aws-native:backup:BackupVault", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:backup:BackupVault", name, id, state, &resource, opts...)
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
	AccessPolicy      interface{}                        `pulumi:"accessPolicy"`
	BackupVaultName   string                             `pulumi:"backupVaultName"`
	BackupVaultTags   interface{}                        `pulumi:"backupVaultTags"`
	EncryptionKeyArn  *string                            `pulumi:"encryptionKeyArn"`
	LockConfiguration *BackupVaultLockConfigurationType  `pulumi:"lockConfiguration"`
	Notifications     *BackupVaultNotificationObjectType `pulumi:"notifications"`
}

// The set of arguments for constructing a BackupVault resource.
type BackupVaultArgs struct {
	AccessPolicy      pulumi.Input
	BackupVaultName   pulumi.StringInput
	BackupVaultTags   pulumi.Input
	EncryptionKeyArn  pulumi.StringPtrInput
	LockConfiguration BackupVaultLockConfigurationTypePtrInput
	Notifications     BackupVaultNotificationObjectTypePtrInput
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
