// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupplan.html
 */
export class BackupPlan extends pulumi.CustomResource {
    /**
     * Get an existing BackupPlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BackupPlan {
        return new BackupPlan(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:Backup:BackupPlan';

    /**
     * Returns true if the given object is an instance of BackupPlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackupPlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackupPlan.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupplan.html#cfn-backup-backupplan-backupplan
     */
    public readonly backupPlan!: pulumi.Output<outputs.Backup.BackupPlanBackupPlanResourceType>;
    public /*out*/ readonly backupPlanArn!: pulumi.Output<string>;
    public /*out*/ readonly backupPlanId!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupplan.html#cfn-backup-backupplan-backupplantags
     */
    public readonly backupPlanTags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly versionId!: pulumi.Output<string>;

    /**
     * Create a BackupPlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupPlanArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.backupPlan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backupPlan'");
            }
            inputs["backupPlan"] = args ? args.backupPlan : undefined;
            inputs["backupPlanTags"] = args ? args.backupPlanTags : undefined;
            inputs["backupPlanArn"] = undefined /*out*/;
            inputs["backupPlanId"] = undefined /*out*/;
            inputs["versionId"] = undefined /*out*/;
        } else {
            inputs["backupPlan"] = undefined /*out*/;
            inputs["backupPlanArn"] = undefined /*out*/;
            inputs["backupPlanId"] = undefined /*out*/;
            inputs["backupPlanTags"] = undefined /*out*/;
            inputs["versionId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BackupPlan.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a BackupPlan resource.
 */
export interface BackupPlanArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupplan.html#cfn-backup-backupplan-backupplan
     */
    backupPlan: pulumi.Input<inputs.Backup.BackupPlanBackupPlanResourceTypeArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupplan.html#cfn-backup-backupplan-backupplantags
     */
    backupPlanTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
