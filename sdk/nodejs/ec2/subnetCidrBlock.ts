// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::SubnetCidrBlock
 *
 * @deprecated SubnetCidrBlock is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class SubnetCidrBlock extends pulumi.CustomResource {
    /**
     * Get an existing SubnetCidrBlock resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SubnetCidrBlock {
        pulumi.log.warn("SubnetCidrBlock is deprecated: SubnetCidrBlock is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new SubnetCidrBlock(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:SubnetCidrBlock';

    /**
     * Returns true if the given object is an instance of SubnetCidrBlock.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SubnetCidrBlock {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SubnetCidrBlock.__pulumiType;
    }

    public readonly ipv6CidrBlock!: pulumi.Output<string>;
    public readonly subnetId!: pulumi.Output<string>;

    /**
     * Create a SubnetCidrBlock resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated SubnetCidrBlock is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: SubnetCidrBlockArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("SubnetCidrBlock is deprecated: SubnetCidrBlock is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.ipv6CidrBlock === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipv6CidrBlock'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            inputs["ipv6CidrBlock"] = args ? args.ipv6CidrBlock : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
        } else {
            inputs["ipv6CidrBlock"] = undefined /*out*/;
            inputs["subnetId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SubnetCidrBlock.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SubnetCidrBlock resource.
 */
export interface SubnetCidrBlockArgs {
    ipv6CidrBlock: pulumi.Input<string>;
    subnetId: pulumi.Input<string>;
}
