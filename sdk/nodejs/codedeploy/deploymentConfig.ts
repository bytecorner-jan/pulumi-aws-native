// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html
 */
export class DeploymentConfig extends pulumi.CustomResource {
    /**
     * Get an existing DeploymentConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DeploymentConfig {
        return new DeploymentConfig(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:codedeploy:DeploymentConfig';

    /**
     * Returns true if the given object is an instance of DeploymentConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeploymentConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeploymentConfig.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-computeplatform
     */
    public readonly computePlatform!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-deploymentconfigname
     */
    public readonly deploymentConfigName!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-minimumhealthyhosts
     */
    public readonly minimumHealthyHosts!: pulumi.Output<outputs.codedeploy.DeploymentConfigMinimumHealthyHosts | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-trafficroutingconfig
     */
    public readonly trafficRoutingConfig!: pulumi.Output<outputs.codedeploy.DeploymentConfigTrafficRoutingConfig | undefined>;

    /**
     * Create a DeploymentConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DeploymentConfigArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["computePlatform"] = args ? args.computePlatform : undefined;
            inputs["deploymentConfigName"] = args ? args.deploymentConfigName : undefined;
            inputs["minimumHealthyHosts"] = args ? args.minimumHealthyHosts : undefined;
            inputs["trafficRoutingConfig"] = args ? args.trafficRoutingConfig : undefined;
        } else {
            inputs["computePlatform"] = undefined /*out*/;
            inputs["deploymentConfigName"] = undefined /*out*/;
            inputs["minimumHealthyHosts"] = undefined /*out*/;
            inputs["trafficRoutingConfig"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DeploymentConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DeploymentConfig resource.
 */
export interface DeploymentConfigArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-computeplatform
     */
    computePlatform?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-deploymentconfigname
     */
    deploymentConfigName?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-minimumhealthyhosts
     */
    minimumHealthyHosts?: pulumi.Input<inputs.codedeploy.DeploymentConfigMinimumHealthyHostsArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-trafficroutingconfig
     */
    trafficRoutingConfig?: pulumi.Input<inputs.codedeploy.DeploymentConfigTrafficRoutingConfigArgs>;
}
