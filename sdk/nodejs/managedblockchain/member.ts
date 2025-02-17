// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ManagedBlockchain::Member
 *
 * @deprecated Member is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Member extends pulumi.CustomResource {
    /**
     * Get an existing Member resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Member {
        pulumi.log.warn("Member is deprecated: Member is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Member(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:managedblockchain:Member';

    /**
     * Returns true if the given object is an instance of Member.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Member {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Member.__pulumiType;
    }

    public readonly invitationId!: pulumi.Output<string | undefined>;
    public readonly memberConfiguration!: pulumi.Output<outputs.managedblockchain.MemberConfiguration>;
    public /*out*/ readonly memberId!: pulumi.Output<string>;
    public readonly networkConfiguration!: pulumi.Output<outputs.managedblockchain.MemberNetworkConfiguration | undefined>;
    public readonly networkId!: pulumi.Output<string | undefined>;

    /**
     * Create a Member resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Member is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: MemberArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Member is deprecated: Member is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.memberConfiguration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memberConfiguration'");
            }
            inputs["invitationId"] = args ? args.invitationId : undefined;
            inputs["memberConfiguration"] = args ? args.memberConfiguration : undefined;
            inputs["networkConfiguration"] = args ? args.networkConfiguration : undefined;
            inputs["networkId"] = args ? args.networkId : undefined;
            inputs["memberId"] = undefined /*out*/;
        } else {
            inputs["invitationId"] = undefined /*out*/;
            inputs["memberConfiguration"] = undefined /*out*/;
            inputs["memberId"] = undefined /*out*/;
            inputs["networkConfiguration"] = undefined /*out*/;
            inputs["networkId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Member.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Member resource.
 */
export interface MemberArgs {
    invitationId?: pulumi.Input<string>;
    memberConfiguration: pulumi.Input<inputs.managedblockchain.MemberConfigurationArgs>;
    networkConfiguration?: pulumi.Input<inputs.managedblockchain.MemberNetworkConfigurationArgs>;
    networkId?: pulumi.Input<string>;
}
