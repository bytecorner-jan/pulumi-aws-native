// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Represents a studio that contains other Nimble Studio resources
 */
export class Studio extends pulumi.CustomResource {
    /**
     * Get an existing Studio resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Studio {
        return new Studio(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:nimblestudio:Studio';

    /**
     * Returns true if the given object is an instance of Studio.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Studio {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Studio.__pulumiType;
    }

    /**
     * <p>The IAM role that Studio Admins will assume when logging in to the Nimble Studio portal.</p>
     */
    public readonly adminRoleArn!: pulumi.Output<string>;
    /**
     * <p>A friendly name for the studio.</p>
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * <p>The Amazon Web Services Region where the studio resource is located.</p>
     */
    public /*out*/ readonly homeRegion!: pulumi.Output<string>;
    /**
     * <p>The Amazon Web Services SSO application client ID used to integrate with Amazon Web Services SSO to enable Amazon Web Services SSO users to log in to Nimble Studio portal.</p>
     */
    public /*out*/ readonly ssoClientId!: pulumi.Output<string>;
    public readonly studioEncryptionConfiguration!: pulumi.Output<outputs.nimblestudio.StudioEncryptionConfiguration | undefined>;
    public /*out*/ readonly studioId!: pulumi.Output<string>;
    /**
     * <p>The studio name that is used in the URL of the Nimble Studio portal when accessed by Nimble Studio users.</p>
     */
    public readonly studioName!: pulumi.Output<string>;
    /**
     * <p>The address of the web page for the studio.</p>
     */
    public /*out*/ readonly studioUrl!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.nimblestudio.StudioTags | undefined>;
    /**
     * <p>The IAM role that Studio Users will assume when logging in to the Nimble Studio portal.</p>
     */
    public readonly userRoleArn!: pulumi.Output<string>;

    /**
     * Create a Studio resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StudioArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.adminRoleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adminRoleArn'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.userRoleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userRoleArn'");
            }
            inputs["adminRoleArn"] = args ? args.adminRoleArn : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["studioEncryptionConfiguration"] = args ? args.studioEncryptionConfiguration : undefined;
            inputs["studioName"] = args ? args.studioName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["userRoleArn"] = args ? args.userRoleArn : undefined;
            inputs["homeRegion"] = undefined /*out*/;
            inputs["ssoClientId"] = undefined /*out*/;
            inputs["studioId"] = undefined /*out*/;
            inputs["studioUrl"] = undefined /*out*/;
        } else {
            inputs["adminRoleArn"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["homeRegion"] = undefined /*out*/;
            inputs["ssoClientId"] = undefined /*out*/;
            inputs["studioEncryptionConfiguration"] = undefined /*out*/;
            inputs["studioId"] = undefined /*out*/;
            inputs["studioName"] = undefined /*out*/;
            inputs["studioUrl"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["userRoleArn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Studio.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Studio resource.
 */
export interface StudioArgs {
    /**
     * <p>The IAM role that Studio Admins will assume when logging in to the Nimble Studio portal.</p>
     */
    adminRoleArn: pulumi.Input<string>;
    /**
     * <p>A friendly name for the studio.</p>
     */
    displayName: pulumi.Input<string>;
    studioEncryptionConfiguration?: pulumi.Input<inputs.nimblestudio.StudioEncryptionConfigurationArgs>;
    /**
     * <p>The studio name that is used in the URL of the Nimble Studio portal when accessed by Nimble Studio users.</p>
     */
    studioName?: pulumi.Input<string>;
    tags?: pulumi.Input<inputs.nimblestudio.StudioTagsArgs>;
    /**
     * <p>The IAM role that Studio Users will assume when logging in to the Nimble Studio portal.</p>
     */
    userRoleArn: pulumi.Input<string>;
}
