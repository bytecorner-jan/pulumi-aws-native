// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Neptune::DBClusterParameterGroup
 *
 * @deprecated DBClusterParameterGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class DBClusterParameterGroup extends pulumi.CustomResource {
    /**
     * Get an existing DBClusterParameterGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DBClusterParameterGroup {
        pulumi.log.warn("DBClusterParameterGroup is deprecated: DBClusterParameterGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new DBClusterParameterGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:neptune:DBClusterParameterGroup';

    /**
     * Returns true if the given object is an instance of DBClusterParameterGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DBClusterParameterGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DBClusterParameterGroup.__pulumiType;
    }

    public readonly description!: pulumi.Output<string>;
    public readonly family!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly parameters!: pulumi.Output<any>;
    public readonly tags!: pulumi.Output<outputs.neptune.DBClusterParameterGroupTag[] | undefined>;

    /**
     * Create a DBClusterParameterGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated DBClusterParameterGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: DBClusterParameterGroupArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DBClusterParameterGroup is deprecated: DBClusterParameterGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.family === undefined) && !opts.urn) {
                throw new Error("Missing required property 'family'");
            }
            if ((!args || args.parameters === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parameters'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["family"] = args ? args.family : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["family"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["parameters"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DBClusterParameterGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DBClusterParameterGroup resource.
 */
export interface DBClusterParameterGroupArgs {
    description: pulumi.Input<string>;
    family: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    parameters: any;
    tags?: pulumi.Input<pulumi.Input<inputs.neptune.DBClusterParameterGroupTagArgs>[]>;
}
