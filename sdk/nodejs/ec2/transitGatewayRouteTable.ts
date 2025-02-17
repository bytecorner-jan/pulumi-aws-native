// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::TransitGatewayRouteTable
 *
 * @deprecated TransitGatewayRouteTable is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class TransitGatewayRouteTable extends pulumi.CustomResource {
    /**
     * Get an existing TransitGatewayRouteTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TransitGatewayRouteTable {
        pulumi.log.warn("TransitGatewayRouteTable is deprecated: TransitGatewayRouteTable is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new TransitGatewayRouteTable(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:TransitGatewayRouteTable';

    /**
     * Returns true if the given object is an instance of TransitGatewayRouteTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransitGatewayRouteTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransitGatewayRouteTable.__pulumiType;
    }

    public readonly tags!: pulumi.Output<outputs.ec2.TransitGatewayRouteTableTag[] | undefined>;
    public readonly transitGatewayId!: pulumi.Output<string>;

    /**
     * Create a TransitGatewayRouteTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated TransitGatewayRouteTable is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: TransitGatewayRouteTableArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("TransitGatewayRouteTable is deprecated: TransitGatewayRouteTable is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.transitGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitGatewayId'");
            }
            inputs["tags"] = args ? args.tags : undefined;
            inputs["transitGatewayId"] = args ? args.transitGatewayId : undefined;
        } else {
            inputs["tags"] = undefined /*out*/;
            inputs["transitGatewayId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TransitGatewayRouteTable.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a TransitGatewayRouteTable resource.
 */
export interface TransitGatewayRouteTableArgs {
    tags?: pulumi.Input<pulumi.Input<inputs.ec2.TransitGatewayRouteTableTagArgs>[]>;
    transitGatewayId: pulumi.Input<string>;
}
