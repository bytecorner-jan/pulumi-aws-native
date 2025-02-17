// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Route53::HealthCheck.
 */
export class HealthCheck extends pulumi.CustomResource {
    /**
     * Get an existing HealthCheck resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): HealthCheck {
        return new HealthCheck(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:route53:HealthCheck';

    /**
     * Returns true if the given object is an instance of HealthCheck.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HealthCheck {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HealthCheck.__pulumiType;
    }

    /**
     * A complex type that contains information about the health check.
     */
    public readonly healthCheckConfig!: pulumi.Output<outputs.route53.HealthCheckConfigProperties>;
    public /*out*/ readonly healthCheckId!: pulumi.Output<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly healthCheckTags!: pulumi.Output<outputs.route53.HealthCheckTag[] | undefined>;

    /**
     * Create a HealthCheck resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HealthCheckArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.healthCheckConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'healthCheckConfig'");
            }
            inputs["healthCheckConfig"] = args ? args.healthCheckConfig : undefined;
            inputs["healthCheckTags"] = args ? args.healthCheckTags : undefined;
            inputs["healthCheckId"] = undefined /*out*/;
        } else {
            inputs["healthCheckConfig"] = undefined /*out*/;
            inputs["healthCheckId"] = undefined /*out*/;
            inputs["healthCheckTags"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(HealthCheck.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a HealthCheck resource.
 */
export interface HealthCheckArgs {
    /**
     * A complex type that contains information about the health check.
     */
    healthCheckConfig: pulumi.Input<inputs.route53.HealthCheckConfigPropertiesArgs>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    healthCheckTags?: pulumi.Input<pulumi.Input<inputs.route53.HealthCheckTagArgs>[]>;
}
