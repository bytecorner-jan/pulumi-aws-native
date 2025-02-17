// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SSM::Parameter
 *
 * @deprecated Parameter is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Parameter extends pulumi.CustomResource {
    /**
     * Get an existing Parameter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Parameter {
        pulumi.log.warn("Parameter is deprecated: Parameter is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Parameter(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ssm:Parameter';

    /**
     * Returns true if the given object is an instance of Parameter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Parameter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Parameter.__pulumiType;
    }

    public readonly allowedPattern!: pulumi.Output<string | undefined>;
    public readonly dataType!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly policies!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<any | undefined>;
    public readonly tier!: pulumi.Output<string | undefined>;
    public readonly type!: pulumi.Output<string>;
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a Parameter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Parameter is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: ParameterArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Parameter is deprecated: Parameter is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            inputs["allowedPattern"] = args ? args.allowedPattern : undefined;
            inputs["dataType"] = args ? args.dataType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tier"] = args ? args.tier : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["value"] = args ? args.value : undefined;
        } else {
            inputs["allowedPattern"] = undefined /*out*/;
            inputs["dataType"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["policies"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["tier"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["value"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Parameter.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Parameter resource.
 */
export interface ParameterArgs {
    allowedPattern?: pulumi.Input<string>;
    dataType?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    policies?: pulumi.Input<string>;
    tags?: any;
    tier?: pulumi.Input<string>;
    type: pulumi.Input<string>;
    value: pulumi.Input<string>;
}
