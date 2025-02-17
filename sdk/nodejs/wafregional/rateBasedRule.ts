// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::WAFRegional::RateBasedRule
 *
 * @deprecated RateBasedRule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class RateBasedRule extends pulumi.CustomResource {
    /**
     * Get an existing RateBasedRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RateBasedRule {
        pulumi.log.warn("RateBasedRule is deprecated: RateBasedRule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new RateBasedRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:wafregional:RateBasedRule';

    /**
     * Returns true if the given object is an instance of RateBasedRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RateBasedRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RateBasedRule.__pulumiType;
    }

    public readonly matchPredicates!: pulumi.Output<outputs.wafregional.RateBasedRulePredicate[] | undefined>;
    public readonly metricName!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly rateKey!: pulumi.Output<string>;
    public readonly rateLimit!: pulumi.Output<number>;

    /**
     * Create a RateBasedRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated RateBasedRule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: RateBasedRuleArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("RateBasedRule is deprecated: RateBasedRule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.metricName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metricName'");
            }
            if ((!args || args.rateKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rateKey'");
            }
            if ((!args || args.rateLimit === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rateLimit'");
            }
            inputs["matchPredicates"] = args ? args.matchPredicates : undefined;
            inputs["metricName"] = args ? args.metricName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["rateKey"] = args ? args.rateKey : undefined;
            inputs["rateLimit"] = args ? args.rateLimit : undefined;
        } else {
            inputs["matchPredicates"] = undefined /*out*/;
            inputs["metricName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["rateKey"] = undefined /*out*/;
            inputs["rateLimit"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RateBasedRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RateBasedRule resource.
 */
export interface RateBasedRuleArgs {
    matchPredicates?: pulumi.Input<pulumi.Input<inputs.wafregional.RateBasedRulePredicateArgs>[]>;
    metricName: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    rateKey: pulumi.Input<string>;
    rateLimit: pulumi.Input<number>;
}
