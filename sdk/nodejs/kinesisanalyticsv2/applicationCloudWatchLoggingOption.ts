// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption
 *
 * @deprecated ApplicationCloudWatchLoggingOption is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class ApplicationCloudWatchLoggingOption extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationCloudWatchLoggingOption resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApplicationCloudWatchLoggingOption {
        pulumi.log.warn("ApplicationCloudWatchLoggingOption is deprecated: ApplicationCloudWatchLoggingOption is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new ApplicationCloudWatchLoggingOption(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:kinesisanalyticsv2:ApplicationCloudWatchLoggingOption';

    /**
     * Returns true if the given object is an instance of ApplicationCloudWatchLoggingOption.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationCloudWatchLoggingOption {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationCloudWatchLoggingOption.__pulumiType;
    }

    public readonly applicationName!: pulumi.Output<string>;
    public readonly cloudWatchLoggingOption!: pulumi.Output<outputs.kinesisanalyticsv2.ApplicationCloudWatchLoggingOptionCloudWatchLoggingOption>;

    /**
     * Create a ApplicationCloudWatchLoggingOption resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated ApplicationCloudWatchLoggingOption is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: ApplicationCloudWatchLoggingOptionArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ApplicationCloudWatchLoggingOption is deprecated: ApplicationCloudWatchLoggingOption is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.applicationName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationName'");
            }
            if ((!args || args.cloudWatchLoggingOption === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudWatchLoggingOption'");
            }
            inputs["applicationName"] = args ? args.applicationName : undefined;
            inputs["cloudWatchLoggingOption"] = args ? args.cloudWatchLoggingOption : undefined;
        } else {
            inputs["applicationName"] = undefined /*out*/;
            inputs["cloudWatchLoggingOption"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ApplicationCloudWatchLoggingOption.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApplicationCloudWatchLoggingOption resource.
 */
export interface ApplicationCloudWatchLoggingOptionArgs {
    applicationName: pulumi.Input<string>;
    cloudWatchLoggingOption: pulumi.Input<inputs.kinesisanalyticsv2.ApplicationCloudWatchLoggingOptionCloudWatchLoggingOptionArgs>;
}
