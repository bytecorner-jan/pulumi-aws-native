// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CloudWatch::AnomalyDetector
 *
 * @deprecated AnomalyDetector is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class AnomalyDetector extends pulumi.CustomResource {
    /**
     * Get an existing AnomalyDetector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AnomalyDetector {
        pulumi.log.warn("AnomalyDetector is deprecated: AnomalyDetector is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new AnomalyDetector(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cloudwatch:AnomalyDetector';

    /**
     * Returns true if the given object is an instance of AnomalyDetector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AnomalyDetector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AnomalyDetector.__pulumiType;
    }

    public readonly configuration!: pulumi.Output<outputs.cloudwatch.AnomalyDetectorConfiguration | undefined>;
    public readonly dimensions!: pulumi.Output<outputs.cloudwatch.AnomalyDetectorDimension[] | undefined>;
    public readonly metricName!: pulumi.Output<string>;
    public readonly namespace!: pulumi.Output<string>;
    public readonly stat!: pulumi.Output<string>;

    /**
     * Create a AnomalyDetector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated AnomalyDetector is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: AnomalyDetectorArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("AnomalyDetector is deprecated: AnomalyDetector is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.metricName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metricName'");
            }
            if ((!args || args.namespace === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespace'");
            }
            if ((!args || args.stat === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stat'");
            }
            inputs["configuration"] = args ? args.configuration : undefined;
            inputs["dimensions"] = args ? args.dimensions : undefined;
            inputs["metricName"] = args ? args.metricName : undefined;
            inputs["namespace"] = args ? args.namespace : undefined;
            inputs["stat"] = args ? args.stat : undefined;
        } else {
            inputs["configuration"] = undefined /*out*/;
            inputs["dimensions"] = undefined /*out*/;
            inputs["metricName"] = undefined /*out*/;
            inputs["namespace"] = undefined /*out*/;
            inputs["stat"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AnomalyDetector.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AnomalyDetector resource.
 */
export interface AnomalyDetectorArgs {
    configuration?: pulumi.Input<inputs.cloudwatch.AnomalyDetectorConfigurationArgs>;
    dimensions?: pulumi.Input<pulumi.Input<inputs.cloudwatch.AnomalyDetectorDimensionArgs>[]>;
    metricName: pulumi.Input<string>;
    namespace: pulumi.Input<string>;
    stat: pulumi.Input<string>;
}
