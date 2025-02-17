// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppStream::DirectoryConfig
 *
 * @deprecated DirectoryConfig is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class DirectoryConfig extends pulumi.CustomResource {
    /**
     * Get an existing DirectoryConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DirectoryConfig {
        pulumi.log.warn("DirectoryConfig is deprecated: DirectoryConfig is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new DirectoryConfig(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:appstream:DirectoryConfig';

    /**
     * Returns true if the given object is an instance of DirectoryConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DirectoryConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DirectoryConfig.__pulumiType;
    }

    public readonly directoryName!: pulumi.Output<string>;
    public readonly organizationalUnitDistinguishedNames!: pulumi.Output<string[]>;
    public readonly serviceAccountCredentials!: pulumi.Output<outputs.appstream.DirectoryConfigServiceAccountCredentials>;

    /**
     * Create a DirectoryConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated DirectoryConfig is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: DirectoryConfigArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DirectoryConfig is deprecated: DirectoryConfig is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.directoryName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'directoryName'");
            }
            if ((!args || args.organizationalUnitDistinguishedNames === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationalUnitDistinguishedNames'");
            }
            if ((!args || args.serviceAccountCredentials === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceAccountCredentials'");
            }
            inputs["directoryName"] = args ? args.directoryName : undefined;
            inputs["organizationalUnitDistinguishedNames"] = args ? args.organizationalUnitDistinguishedNames : undefined;
            inputs["serviceAccountCredentials"] = args ? args.serviceAccountCredentials : undefined;
        } else {
            inputs["directoryName"] = undefined /*out*/;
            inputs["organizationalUnitDistinguishedNames"] = undefined /*out*/;
            inputs["serviceAccountCredentials"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DirectoryConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DirectoryConfig resource.
 */
export interface DirectoryConfigArgs {
    directoryName: pulumi.Input<string>;
    organizationalUnitDistinguishedNames: pulumi.Input<pulumi.Input<string>[]>;
    serviceAccountCredentials: pulumi.Input<inputs.appstream.DirectoryConfigServiceAccountCredentialsArgs>;
}
