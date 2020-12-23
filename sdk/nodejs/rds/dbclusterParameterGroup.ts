// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html
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
        return new DBClusterParameterGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudformation:RDS:DBClusterParameterGroup';

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

    /**
     * The attributes associated with the resource
     */
    public /*out*/ readonly attributes!: pulumi.Output<outputs.RDS.DBClusterParameterGroupAttributes>;
    /**
     * An explicit logical ID for the resource
     */
    public readonly logicalId!: pulumi.Output<string | undefined>;
    /**
     * Arbitrary structured data associated with the resource
     */
    public readonly metadata!: pulumi.Output<any | string | undefined>;
    /**
     * The input properties associated with the resource
     */
    public readonly properties!: pulumi.Output<outputs.RDS.DBClusterParameterGroupProperties>;

    /**
     * Create a DBClusterParameterGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DBClusterParameterGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.properties === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'properties'");
            }
            inputs["deletionPolicy"] = args ? args.deletionPolicy : undefined;
            inputs["logicalId"] = args ? args.logicalId : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["updateReplacePolicy"] = args ? args.updateReplacePolicy : undefined;
            inputs["attributes"] = undefined /*out*/;
        } else {
            inputs["attributes"] = undefined /*out*/;
            inputs["logicalId"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DBClusterParameterGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DBClusterParameterGroup resource.
 */
export interface DBClusterParameterGroupArgs {
    /**
     * With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
     */
    readonly deletionPolicy?: pulumi.Input<string>;
    /**
     * An explicit logical ID for the resource
     */
    readonly logicalId?: pulumi.Input<string>;
    /**
     * Arbitrary structured data associated with the resource
     */
    readonly metadata?: pulumi.Input<any | string>;
    /**
     * The input properties associated with the resource
     */
    readonly properties: pulumi.Input<inputs.RDS.DBClusterParameterGroupProperties>;
    /**
     * Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
     */
    readonly updateReplacePolicy?: pulumi.Input<string>;
}
