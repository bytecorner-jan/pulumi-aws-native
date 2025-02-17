// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::ImageBuilder::Component
 */
export class Component extends pulumi.CustomResource {
    /**
     * Get an existing Component resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Component {
        return new Component(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:imagebuilder:Component';

    /**
     * Returns true if the given object is an instance of Component.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Component {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Component.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the component.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The change description of the component.
     */
    public readonly changeDescription!: pulumi.Output<string | undefined>;
    /**
     * The data of the component.
     */
    public readonly data!: pulumi.Output<string | undefined>;
    /**
     * The description of the component.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The encryption status of the component.
     */
    public /*out*/ readonly encrypted!: pulumi.Output<boolean>;
    /**
     * The KMS key identifier used to encrypt the component.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * The name of the component.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The platform of the component.
     */
    public readonly platform!: pulumi.Output<enums.imagebuilder.ComponentPlatform>;
    /**
     * The operating system (OS) version supported by the component.
     */
    public readonly supportedOsVersions!: pulumi.Output<string[] | undefined>;
    /**
     * The tags associated with the component.
     */
    public readonly tags!: pulumi.Output<any | undefined>;
    /**
     * The type of the component denotes whether the component is used to build the image or only to test it. 
     */
    public /*out*/ readonly type!: pulumi.Output<enums.imagebuilder.ComponentType>;
    /**
     * The uri of the component.
     */
    public readonly uri!: pulumi.Output<string | undefined>;
    /**
     * The version of the component.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a Component resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComponentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.platform === undefined) && !opts.urn) {
                throw new Error("Missing required property 'platform'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            inputs["changeDescription"] = args ? args.changeDescription : undefined;
            inputs["data"] = args ? args.data : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["platform"] = args ? args.platform : undefined;
            inputs["supportedOsVersions"] = args ? args.supportedOsVersions : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["uri"] = args ? args.uri : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["encrypted"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["arn"] = undefined /*out*/;
            inputs["changeDescription"] = undefined /*out*/;
            inputs["data"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["encrypted"] = undefined /*out*/;
            inputs["kmsKeyId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["platform"] = undefined /*out*/;
            inputs["supportedOsVersions"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uri"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Component.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Component resource.
 */
export interface ComponentArgs {
    /**
     * The change description of the component.
     */
    changeDescription?: pulumi.Input<string>;
    /**
     * The data of the component.
     */
    data?: pulumi.Input<string>;
    /**
     * The description of the component.
     */
    description?: pulumi.Input<string>;
    /**
     * The KMS key identifier used to encrypt the component.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The name of the component.
     */
    name?: pulumi.Input<string>;
    /**
     * The platform of the component.
     */
    platform: pulumi.Input<enums.imagebuilder.ComponentPlatform>;
    /**
     * The operating system (OS) version supported by the component.
     */
    supportedOsVersions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The tags associated with the component.
     */
    tags?: any;
    /**
     * The uri of the component.
     */
    uri?: pulumi.Input<string>;
    /**
     * The version of the component.
     */
    version: pulumi.Input<string>;
}
