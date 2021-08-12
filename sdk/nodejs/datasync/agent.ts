// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html
 */
export class Agent extends pulumi.CustomResource {
    /**
     * Get an existing Agent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Agent {
        return new Agent(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:DataSync:Agent';

    /**
     * Returns true if the given object is an instance of Agent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Agent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Agent.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-activationkey
     */
    public readonly activationKey!: pulumi.Output<string>;
    public /*out*/ readonly agentArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-agentname
     */
    public readonly agentName!: pulumi.Output<string | undefined>;
    public /*out*/ readonly endpointType!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-securitygrouparns
     */
    public readonly securityGroupArns!: pulumi.Output<string[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-subnetarns
     */
    public readonly subnetArns!: pulumi.Output<string[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-vpcendpointid
     */
    public readonly vpcEndpointId!: pulumi.Output<string | undefined>;

    /**
     * Create a Agent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AgentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.activationKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'activationKey'");
            }
            inputs["activationKey"] = args ? args.activationKey : undefined;
            inputs["agentName"] = args ? args.agentName : undefined;
            inputs["securityGroupArns"] = args ? args.securityGroupArns : undefined;
            inputs["subnetArns"] = args ? args.subnetArns : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcEndpointId"] = args ? args.vpcEndpointId : undefined;
            inputs["agentArn"] = undefined /*out*/;
            inputs["endpointType"] = undefined /*out*/;
        } else {
            inputs["activationKey"] = undefined /*out*/;
            inputs["agentArn"] = undefined /*out*/;
            inputs["agentName"] = undefined /*out*/;
            inputs["endpointType"] = undefined /*out*/;
            inputs["securityGroupArns"] = undefined /*out*/;
            inputs["subnetArns"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["vpcEndpointId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Agent.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Agent resource.
 */
export interface AgentArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-activationkey
     */
    activationKey: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-agentname
     */
    agentName?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-securitygrouparns
     */
    securityGroupArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-subnetarns
     */
    subnetArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-tags
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-vpcendpointid
     */
    vpcEndpointId?: pulumi.Input<string>;
}
