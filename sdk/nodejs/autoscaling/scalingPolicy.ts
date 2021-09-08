// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html
 */
export class ScalingPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ScalingPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ScalingPolicy {
        return new ScalingPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:autoscaling:ScalingPolicy';

    /**
     * Returns true if the given object is an instance of ScalingPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScalingPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScalingPolicy.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-adjustmenttype
     */
    public readonly adjustmentType!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-autoscalinggroupname
     */
    public readonly autoScalingGroupName!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-cooldown
     */
    public readonly cooldown!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-estimatedinstancewarmup
     */
    public readonly estimatedInstanceWarmup!: pulumi.Output<number | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-metricaggregationtype
     */
    public readonly metricAggregationType!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-minadjustmentmagnitude
     */
    public readonly minAdjustmentMagnitude!: pulumi.Output<number | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-policytype
     */
    public readonly policyType!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-scalingadjustment
     */
    public readonly scalingAdjustment!: pulumi.Output<number | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-stepadjustments
     */
    public readonly stepAdjustments!: pulumi.Output<outputs.autoscaling.ScalingPolicyStepAdjustment[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-autoscaling-scalingpolicy-targettrackingconfiguration
     */
    public readonly targetTrackingConfiguration!: pulumi.Output<outputs.autoscaling.ScalingPolicyTargetTrackingConfiguration | undefined>;

    /**
     * Create a ScalingPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScalingPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.autoScalingGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoScalingGroupName'");
            }
            inputs["adjustmentType"] = args ? args.adjustmentType : undefined;
            inputs["autoScalingGroupName"] = args ? args.autoScalingGroupName : undefined;
            inputs["cooldown"] = args ? args.cooldown : undefined;
            inputs["estimatedInstanceWarmup"] = args ? args.estimatedInstanceWarmup : undefined;
            inputs["metricAggregationType"] = args ? args.metricAggregationType : undefined;
            inputs["minAdjustmentMagnitude"] = args ? args.minAdjustmentMagnitude : undefined;
            inputs["policyType"] = args ? args.policyType : undefined;
            inputs["scalingAdjustment"] = args ? args.scalingAdjustment : undefined;
            inputs["stepAdjustments"] = args ? args.stepAdjustments : undefined;
            inputs["targetTrackingConfiguration"] = args ? args.targetTrackingConfiguration : undefined;
        } else {
            inputs["adjustmentType"] = undefined /*out*/;
            inputs["autoScalingGroupName"] = undefined /*out*/;
            inputs["cooldown"] = undefined /*out*/;
            inputs["estimatedInstanceWarmup"] = undefined /*out*/;
            inputs["metricAggregationType"] = undefined /*out*/;
            inputs["minAdjustmentMagnitude"] = undefined /*out*/;
            inputs["policyType"] = undefined /*out*/;
            inputs["scalingAdjustment"] = undefined /*out*/;
            inputs["stepAdjustments"] = undefined /*out*/;
            inputs["targetTrackingConfiguration"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ScalingPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ScalingPolicy resource.
 */
export interface ScalingPolicyArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-adjustmenttype
     */
    adjustmentType?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-autoscalinggroupname
     */
    autoScalingGroupName: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-cooldown
     */
    cooldown?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-estimatedinstancewarmup
     */
    estimatedInstanceWarmup?: pulumi.Input<number>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-metricaggregationtype
     */
    metricAggregationType?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-minadjustmentmagnitude
     */
    minAdjustmentMagnitude?: pulumi.Input<number>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-policytype
     */
    policyType?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-scalingadjustment
     */
    scalingAdjustment?: pulumi.Input<number>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-stepadjustments
     */
    stepAdjustments?: pulumi.Input<pulumi.Input<inputs.autoscaling.ScalingPolicyStepAdjustmentArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-autoscaling-scalingpolicy-targettrackingconfiguration
     */
    targetTrackingConfiguration?: pulumi.Input<inputs.autoscaling.ScalingPolicyTargetTrackingConfigurationArgs>;
}
