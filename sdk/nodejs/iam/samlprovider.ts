// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IAM::SAMLProvider
 */
export class SAMLProvider extends pulumi.CustomResource {
    /**
     * Get an existing SAMLProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SAMLProvider {
        return new SAMLProvider(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iam:SAMLProvider';

    /**
     * Returns true if the given object is an instance of SAMLProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SAMLProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SAMLProvider.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the SAML provider
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly samlMetadataDocument!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.iam.SAMLProviderTag[] | undefined>;

    /**
     * Create a SAMLProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SAMLProviderArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.samlMetadataDocument === undefined) && !opts.urn) {
                throw new Error("Missing required property 'samlMetadataDocument'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["samlMetadataDocument"] = args ? args.samlMetadataDocument : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        } else {
            inputs["arn"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["samlMetadataDocument"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SAMLProvider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SAMLProvider resource.
 */
export interface SAMLProviderArgs {
    name?: pulumi.Input<string>;
    samlMetadataDocument: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.iam.SAMLProviderTagArgs>[]>;
}
