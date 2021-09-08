// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html
 */
export class InstanceFleetConfig extends pulumi.CustomResource {
    /**
     * Get an existing InstanceFleetConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): InstanceFleetConfig {
        return new InstanceFleetConfig(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:emr:InstanceFleetConfig';

    /**
     * Returns true if the given object is an instance of InstanceFleetConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceFleetConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceFleetConfig.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-clusterid
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-instancefleettype
     */
    public readonly instanceFleetType!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-instancetypeconfigs
     */
    public readonly instanceTypeConfigs!: pulumi.Output<outputs.emr.InstanceFleetConfigInstanceTypeConfig[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-launchspecifications
     */
    public readonly launchSpecifications!: pulumi.Output<outputs.emr.InstanceFleetConfigInstanceFleetProvisioningSpecifications | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-name
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-targetondemandcapacity
     */
    public readonly targetOnDemandCapacity!: pulumi.Output<number | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-targetspotcapacity
     */
    public readonly targetSpotCapacity!: pulumi.Output<number | undefined>;

    /**
     * Create a InstanceFleetConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceFleetConfigArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.instanceFleetType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceFleetType'");
            }
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["instanceFleetType"] = args ? args.instanceFleetType : undefined;
            inputs["instanceTypeConfigs"] = args ? args.instanceTypeConfigs : undefined;
            inputs["launchSpecifications"] = args ? args.launchSpecifications : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["targetOnDemandCapacity"] = args ? args.targetOnDemandCapacity : undefined;
            inputs["targetSpotCapacity"] = args ? args.targetSpotCapacity : undefined;
        } else {
            inputs["clusterId"] = undefined /*out*/;
            inputs["instanceFleetType"] = undefined /*out*/;
            inputs["instanceTypeConfigs"] = undefined /*out*/;
            inputs["launchSpecifications"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["targetOnDemandCapacity"] = undefined /*out*/;
            inputs["targetSpotCapacity"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(InstanceFleetConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a InstanceFleetConfig resource.
 */
export interface InstanceFleetConfigArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-clusterid
     */
    clusterId: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-instancefleettype
     */
    instanceFleetType: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-instancetypeconfigs
     */
    instanceTypeConfigs?: pulumi.Input<pulumi.Input<inputs.emr.InstanceFleetConfigInstanceTypeConfigArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-launchspecifications
     */
    launchSpecifications?: pulumi.Input<inputs.emr.InstanceFleetConfigInstanceFleetProvisioningSpecificationsArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-name
     */
    name?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-targetondemandcapacity
     */
    targetOnDemandCapacity?: pulumi.Input<number>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-targetspotcapacity
     */
    targetSpotCapacity?: pulumi.Input<number>;
}
