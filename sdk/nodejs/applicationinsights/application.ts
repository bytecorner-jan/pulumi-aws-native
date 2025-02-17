// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::ApplicationInsights::Application
 */
export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:applicationinsights:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    /**
     * The ARN of the ApplicationInsights application.
     */
    public /*out*/ readonly applicationARN!: pulumi.Output<string>;
    /**
     * If set to true, application will be configured with recommended monitoring configuration.
     */
    public readonly autoConfigurationEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether Application Insights can listen to CloudWatch events for the application resources.
     */
    public readonly cWEMonitorEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The monitoring settings of the components.
     */
    public readonly componentMonitoringSettings!: pulumi.Output<outputs.applicationinsights.ApplicationComponentMonitoringSetting[] | undefined>;
    /**
     * The custom grouped components.
     */
    public readonly customComponents!: pulumi.Output<outputs.applicationinsights.ApplicationCustomComponent[] | undefined>;
    /**
     * The log pattern sets.
     */
    public readonly logPatternSets!: pulumi.Output<outputs.applicationinsights.ApplicationLogPatternSet[] | undefined>;
    /**
     * When set to true, creates opsItems for any problems detected on an application.
     */
    public readonly opsCenterEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The SNS topic provided to Application Insights that is associated to the created opsItem.
     */
    public readonly opsItemSNSTopicArn!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The tags of Application Insights application.
     */
    public readonly tags!: pulumi.Output<outputs.applicationinsights.ApplicationTag[] | undefined>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["autoConfigurationEnabled"] = args ? args.autoConfigurationEnabled : undefined;
            inputs["cWEMonitorEnabled"] = args ? args.cWEMonitorEnabled : undefined;
            inputs["componentMonitoringSettings"] = args ? args.componentMonitoringSettings : undefined;
            inputs["customComponents"] = args ? args.customComponents : undefined;
            inputs["logPatternSets"] = args ? args.logPatternSets : undefined;
            inputs["opsCenterEnabled"] = args ? args.opsCenterEnabled : undefined;
            inputs["opsItemSNSTopicArn"] = args ? args.opsItemSNSTopicArn : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["applicationARN"] = undefined /*out*/;
        } else {
            inputs["applicationARN"] = undefined /*out*/;
            inputs["autoConfigurationEnabled"] = undefined /*out*/;
            inputs["cWEMonitorEnabled"] = undefined /*out*/;
            inputs["componentMonitoringSettings"] = undefined /*out*/;
            inputs["customComponents"] = undefined /*out*/;
            inputs["logPatternSets"] = undefined /*out*/;
            inputs["opsCenterEnabled"] = undefined /*out*/;
            inputs["opsItemSNSTopicArn"] = undefined /*out*/;
            inputs["resourceGroupName"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Application.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * If set to true, application will be configured with recommended monitoring configuration.
     */
    autoConfigurationEnabled?: pulumi.Input<boolean>;
    /**
     * Indicates whether Application Insights can listen to CloudWatch events for the application resources.
     */
    cWEMonitorEnabled?: pulumi.Input<boolean>;
    /**
     * The monitoring settings of the components.
     */
    componentMonitoringSettings?: pulumi.Input<pulumi.Input<inputs.applicationinsights.ApplicationComponentMonitoringSettingArgs>[]>;
    /**
     * The custom grouped components.
     */
    customComponents?: pulumi.Input<pulumi.Input<inputs.applicationinsights.ApplicationCustomComponentArgs>[]>;
    /**
     * The log pattern sets.
     */
    logPatternSets?: pulumi.Input<pulumi.Input<inputs.applicationinsights.ApplicationLogPatternSetArgs>[]>;
    /**
     * When set to true, creates opsItems for any problems detected on an application.
     */
    opsCenterEnabled?: pulumi.Input<boolean>;
    /**
     * The SNS topic provided to Application Insights that is associated to the created opsItem.
     */
    opsItemSNSTopicArn?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The tags of Application Insights application.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.applicationinsights.ApplicationTagArgs>[]>;
}
