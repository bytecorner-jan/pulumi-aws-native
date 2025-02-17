// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Cloud9::EnvironmentEC2
 *
 * @deprecated EnvironmentEC2 is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class EnvironmentEC2 extends pulumi.CustomResource {
    /**
     * Get an existing EnvironmentEC2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EnvironmentEC2 {
        pulumi.log.warn("EnvironmentEC2 is deprecated: EnvironmentEC2 is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new EnvironmentEC2(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cloud9:EnvironmentEC2';

    /**
     * Returns true if the given object is an instance of EnvironmentEC2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EnvironmentEC2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EnvironmentEC2.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly automaticStopTimeMinutes!: pulumi.Output<number | undefined>;
    public readonly connectionType!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly imageId!: pulumi.Output<string | undefined>;
    public readonly instanceType!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly ownerArn!: pulumi.Output<string | undefined>;
    public readonly repositories!: pulumi.Output<outputs.cloud9.EnvironmentEC2Repository[] | undefined>;
    public readonly subnetId!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<outputs.cloud9.EnvironmentEC2Tag[] | undefined>;

    /**
     * Create a EnvironmentEC2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated EnvironmentEC2 is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: EnvironmentEC2Args, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("EnvironmentEC2 is deprecated: EnvironmentEC2 is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            inputs["automaticStopTimeMinutes"] = args ? args.automaticStopTimeMinutes : undefined;
            inputs["connectionType"] = args ? args.connectionType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["imageId"] = args ? args.imageId : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["ownerArn"] = args ? args.ownerArn : undefined;
            inputs["repositories"] = args ? args.repositories : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        } else {
            inputs["arn"] = undefined /*out*/;
            inputs["automaticStopTimeMinutes"] = undefined /*out*/;
            inputs["connectionType"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["imageId"] = undefined /*out*/;
            inputs["instanceType"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["ownerArn"] = undefined /*out*/;
            inputs["repositories"] = undefined /*out*/;
            inputs["subnetId"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EnvironmentEC2.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a EnvironmentEC2 resource.
 */
export interface EnvironmentEC2Args {
    automaticStopTimeMinutes?: pulumi.Input<number>;
    connectionType?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    imageId?: pulumi.Input<string>;
    instanceType: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    ownerArn?: pulumi.Input<string>;
    repositories?: pulumi.Input<pulumi.Input<inputs.cloud9.EnvironmentEC2RepositoryArgs>[]>;
    subnetId?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.cloud9.EnvironmentEC2TagArgs>[]>;
}
