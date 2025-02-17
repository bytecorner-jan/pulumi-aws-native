// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * An example resource schema demonstrating some basic constructs and validation rules.
 */
export class StudioSessionMapping extends pulumi.CustomResource {
    /**
     * Get an existing StudioSessionMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StudioSessionMapping {
        return new StudioSessionMapping(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:emr:StudioSessionMapping';

    /**
     * Returns true if the given object is an instance of StudioSessionMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StudioSessionMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StudioSessionMapping.__pulumiType;
    }

    /**
     * The name of the user or group. For more information, see UserName and DisplayName in the AWS SSO Identity Store API Reference. Either IdentityName or IdentityId must be specified.
     */
    public readonly identityName!: pulumi.Output<string>;
    /**
     * Specifies whether the identity to map to the Studio is a user or a group.
     */
    public readonly identityType!: pulumi.Output<enums.emr.StudioSessionMappingIdentityType>;
    /**
     * The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group. Session policies refine Studio user permissions without the need to use multiple IAM user roles.
     */
    public readonly sessionPolicyArn!: pulumi.Output<string>;
    /**
     * The ID of the Amazon EMR Studio to which the user or group will be mapped.
     */
    public readonly studioId!: pulumi.Output<string>;

    /**
     * Create a StudioSessionMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StudioSessionMappingArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.identityName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identityName'");
            }
            if ((!args || args.identityType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identityType'");
            }
            if ((!args || args.sessionPolicyArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sessionPolicyArn'");
            }
            if ((!args || args.studioId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'studioId'");
            }
            inputs["identityName"] = args ? args.identityName : undefined;
            inputs["identityType"] = args ? args.identityType : undefined;
            inputs["sessionPolicyArn"] = args ? args.sessionPolicyArn : undefined;
            inputs["studioId"] = args ? args.studioId : undefined;
        } else {
            inputs["identityName"] = undefined /*out*/;
            inputs["identityType"] = undefined /*out*/;
            inputs["sessionPolicyArn"] = undefined /*out*/;
            inputs["studioId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(StudioSessionMapping.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a StudioSessionMapping resource.
 */
export interface StudioSessionMappingArgs {
    /**
     * The name of the user or group. For more information, see UserName and DisplayName in the AWS SSO Identity Store API Reference. Either IdentityName or IdentityId must be specified.
     */
    identityName: pulumi.Input<string>;
    /**
     * Specifies whether the identity to map to the Studio is a user or a group.
     */
    identityType: pulumi.Input<enums.emr.StudioSessionMappingIdentityType>;
    /**
     * The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group. Session policies refine Studio user permissions without the need to use multiple IAM user roles.
     */
    sessionPolicyArn: pulumi.Input<string>;
    /**
     * The ID of the Amazon EMR Studio to which the user or group will be mapped.
     */
    studioId: pulumi.Input<string>;
}
