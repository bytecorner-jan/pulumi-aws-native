// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::EC2::NetworkInsightsAnalysis
 */
export class NetworkInsightsAnalysis extends pulumi.CustomResource {
    /**
     * Get an existing NetworkInsightsAnalysis resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NetworkInsightsAnalysis {
        return new NetworkInsightsAnalysis(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:NetworkInsightsAnalysis';

    /**
     * Returns true if the given object is an instance of NetworkInsightsAnalysis.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkInsightsAnalysis {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkInsightsAnalysis.__pulumiType;
    }

    public /*out*/ readonly alternatePathHints!: pulumi.Output<outputs.ec2.NetworkInsightsAnalysisAlternatePathHint[]>;
    public /*out*/ readonly explanations!: pulumi.Output<outputs.ec2.NetworkInsightsAnalysisExplanation[]>;
    public readonly filterInArns!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly forwardPathComponents!: pulumi.Output<outputs.ec2.NetworkInsightsAnalysisPathComponent[]>;
    public /*out*/ readonly networkInsightsAnalysisArn!: pulumi.Output<string>;
    public /*out*/ readonly networkInsightsAnalysisId!: pulumi.Output<string>;
    public readonly networkInsightsPathId!: pulumi.Output<string>;
    public /*out*/ readonly networkPathFound!: pulumi.Output<boolean>;
    public /*out*/ readonly returnPathComponents!: pulumi.Output<outputs.ec2.NetworkInsightsAnalysisPathComponent[]>;
    public /*out*/ readonly startDate!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<enums.ec2.NetworkInsightsAnalysisStatus>;
    public /*out*/ readonly statusMessage!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.ec2.NetworkInsightsAnalysisTag[] | undefined>;

    /**
     * Create a NetworkInsightsAnalysis resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkInsightsAnalysisArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.networkInsightsPathId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkInsightsPathId'");
            }
            inputs["filterInArns"] = args ? args.filterInArns : undefined;
            inputs["networkInsightsPathId"] = args ? args.networkInsightsPathId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["alternatePathHints"] = undefined /*out*/;
            inputs["explanations"] = undefined /*out*/;
            inputs["forwardPathComponents"] = undefined /*out*/;
            inputs["networkInsightsAnalysisArn"] = undefined /*out*/;
            inputs["networkInsightsAnalysisId"] = undefined /*out*/;
            inputs["networkPathFound"] = undefined /*out*/;
            inputs["returnPathComponents"] = undefined /*out*/;
            inputs["startDate"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusMessage"] = undefined /*out*/;
        } else {
            inputs["alternatePathHints"] = undefined /*out*/;
            inputs["explanations"] = undefined /*out*/;
            inputs["filterInArns"] = undefined /*out*/;
            inputs["forwardPathComponents"] = undefined /*out*/;
            inputs["networkInsightsAnalysisArn"] = undefined /*out*/;
            inputs["networkInsightsAnalysisId"] = undefined /*out*/;
            inputs["networkInsightsPathId"] = undefined /*out*/;
            inputs["networkPathFound"] = undefined /*out*/;
            inputs["returnPathComponents"] = undefined /*out*/;
            inputs["startDate"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusMessage"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NetworkInsightsAnalysis.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NetworkInsightsAnalysis resource.
 */
export interface NetworkInsightsAnalysisArgs {
    filterInArns?: pulumi.Input<pulumi.Input<string>[]>;
    networkInsightsPathId: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.ec2.NetworkInsightsAnalysisTagArgs>[]>;
}
