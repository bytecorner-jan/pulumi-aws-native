// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Backup::BackupVault
 */
export class BackupVault extends pulumi.CustomResource {
    /**
     * Get an existing BackupVault resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BackupVault {
        return new BackupVault(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:backup:BackupVault';

    /**
     * Returns true if the given object is an instance of BackupVault.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackupVault {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackupVault.__pulumiType;
    }

    public readonly accessPolicy!: pulumi.Output<any | undefined>;
    public /*out*/ readonly backupVaultArn!: pulumi.Output<string>;
    public readonly backupVaultName!: pulumi.Output<string>;
    public readonly backupVaultTags!: pulumi.Output<any | undefined>;
    public readonly encryptionKeyArn!: pulumi.Output<string | undefined>;
    public readonly lockConfiguration!: pulumi.Output<outputs.backup.BackupVaultLockConfigurationType | undefined>;
    public readonly notifications!: pulumi.Output<outputs.backup.BackupVaultNotificationObjectType | undefined>;

    /**
     * Create a BackupVault resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupVaultArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.backupVaultName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backupVaultName'");
            }
            inputs["accessPolicy"] = args ? args.accessPolicy : undefined;
            inputs["backupVaultName"] = args ? args.backupVaultName : undefined;
            inputs["backupVaultTags"] = args ? args.backupVaultTags : undefined;
            inputs["encryptionKeyArn"] = args ? args.encryptionKeyArn : undefined;
            inputs["lockConfiguration"] = args ? args.lockConfiguration : undefined;
            inputs["notifications"] = args ? args.notifications : undefined;
            inputs["backupVaultArn"] = undefined /*out*/;
        } else {
            inputs["accessPolicy"] = undefined /*out*/;
            inputs["backupVaultArn"] = undefined /*out*/;
            inputs["backupVaultName"] = undefined /*out*/;
            inputs["backupVaultTags"] = undefined /*out*/;
            inputs["encryptionKeyArn"] = undefined /*out*/;
            inputs["lockConfiguration"] = undefined /*out*/;
            inputs["notifications"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BackupVault.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a BackupVault resource.
 */
export interface BackupVaultArgs {
    accessPolicy?: any;
    backupVaultName: pulumi.Input<string>;
    backupVaultTags?: any;
    encryptionKeyArn?: pulumi.Input<string>;
    lockConfiguration?: pulumi.Input<inputs.backup.BackupVaultLockConfigurationTypeArgs>;
    notifications?: pulumi.Input<inputs.backup.BackupVaultNotificationObjectTypeArgs>;
}
