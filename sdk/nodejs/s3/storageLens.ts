// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-storagelens.html
 */
export class StorageLens extends pulumi.CustomResource {
    /**
     * Get an existing StorageLens resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StorageLens {
        return new StorageLens(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:S3:StorageLens';

    /**
     * Returns true if the given object is an instance of StorageLens.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StorageLens {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StorageLens.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-storagelens.html#cfn-s3-storagelens-storagelensconfiguration
     */
    public readonly storageLensConfiguration!: pulumi.Output<outputs.S3.StorageLensStorageLensConfiguration>;
    public /*out*/ readonly storageLensConfigurationStorageLensArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-storagelens.html#cfn-s3-storagelens-tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a StorageLens resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StorageLensArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.storageLensConfiguration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageLensConfiguration'");
            }
            inputs["storageLensConfiguration"] = args ? args.storageLensConfiguration : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["storageLensConfigurationStorageLensArn"] = undefined /*out*/;
        } else {
            inputs["storageLensConfiguration"] = undefined /*out*/;
            inputs["storageLensConfigurationStorageLensArn"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(StorageLens.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a StorageLens resource.
 */
export interface StorageLensArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-storagelens.html#cfn-s3-storagelens-storagelensconfiguration
     */
    storageLensConfiguration: pulumi.Input<inputs.S3.StorageLensStorageLensConfigurationArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-storagelens.html#cfn-s3-storagelens-tags
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
