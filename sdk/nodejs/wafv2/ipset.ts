// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Contains a list of IP addresses. This can be either IPV4 or IPV6. The list will be mutually
 */
export class IPSet extends pulumi.CustomResource {
    /**
     * Get an existing IPSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IPSet {
        return new IPSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:wafv2:IPSet';

    /**
     * Returns true if the given object is an instance of IPSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IPSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IPSet.__pulumiType;
    }

    /**
     * List of IPAddresses.
     */
    public readonly addresses!: pulumi.Output<string[]>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly iPAddressVersion!: pulumi.Output<enums.wafv2.IPSetIPAddressVersion>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly scope!: pulumi.Output<enums.wafv2.IPSetScope>;
    public readonly tags!: pulumi.Output<outputs.wafv2.IPSetTag[] | undefined>;

    /**
     * Create a IPSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IPSetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.addresses === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addresses'");
            }
            if ((!args || args.iPAddressVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'iPAddressVersion'");
            }
            if ((!args || args.scope === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scope'");
            }
            inputs["addresses"] = args ? args.addresses : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["iPAddressVersion"] = args ? args.iPAddressVersion : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        } else {
            inputs["addresses"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["iPAddressVersion"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["scope"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(IPSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a IPSet resource.
 */
export interface IPSetArgs {
    /**
     * List of IPAddresses.
     */
    addresses: pulumi.Input<pulumi.Input<string>[]>;
    description?: pulumi.Input<string>;
    iPAddressVersion: pulumi.Input<enums.wafv2.IPSetIPAddressVersion>;
    name?: pulumi.Input<string>;
    scope: pulumi.Input<enums.wafv2.IPSetScope>;
    tags?: pulumi.Input<pulumi.Input<inputs.wafv2.IPSetTagArgs>[]>;
}
