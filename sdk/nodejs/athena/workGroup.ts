// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html
 */
export class WorkGroup extends pulumi.CustomResource {
    /**
     * Get an existing WorkGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WorkGroup {
        return new WorkGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:Athena:WorkGroup';

    /**
     * Returns true if the given object is an instance of WorkGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkGroup.__pulumiType;
    }

    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-recursivedeleteoption
     */
    public readonly recursiveDeleteOption!: pulumi.Output<boolean | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-state
     */
    public readonly state!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-workgroupconfiguration
     */
    public readonly workGroupConfiguration!: pulumi.Output<outputs.Athena.WorkGroupWorkGroupConfiguration | undefined>;
    public /*out*/ readonly workGroupConfigurationEngineVersionEffectiveEngineVersion!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-workgroupconfigurationupdates
     */
    public readonly workGroupConfigurationUpdates!: pulumi.Output<outputs.Athena.WorkGroupWorkGroupConfigurationUpdates | undefined>;
    public /*out*/ readonly workGroupConfigurationUpdatesEngineVersionEffectiveEngineVersion!: pulumi.Output<string>;

    /**
     * Create a WorkGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["recursiveDeleteOption"] = args ? args.recursiveDeleteOption : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["workGroupConfiguration"] = args ? args.workGroupConfiguration : undefined;
            inputs["workGroupConfigurationUpdates"] = args ? args.workGroupConfigurationUpdates : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["workGroupConfigurationEngineVersionEffectiveEngineVersion"] = undefined /*out*/;
            inputs["workGroupConfigurationUpdatesEngineVersionEffectiveEngineVersion"] = undefined /*out*/;
        } else {
            inputs["creationTime"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["recursiveDeleteOption"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["workGroupConfiguration"] = undefined /*out*/;
            inputs["workGroupConfigurationEngineVersionEffectiveEngineVersion"] = undefined /*out*/;
            inputs["workGroupConfigurationUpdates"] = undefined /*out*/;
            inputs["workGroupConfigurationUpdatesEngineVersionEffectiveEngineVersion"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WorkGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WorkGroup resource.
 */
export interface WorkGroupArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-description
     */
    description?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-name
     */
    name: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-recursivedeleteoption
     */
    recursiveDeleteOption?: pulumi.Input<boolean>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-state
     */
    state?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-tags
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-workgroupconfiguration
     */
    workGroupConfiguration?: pulumi.Input<inputs.Athena.WorkGroupWorkGroupConfigurationArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-workgroupconfigurationupdates
     */
    workGroupConfigurationUpdates?: pulumi.Input<inputs.Athena.WorkGroupWorkGroupConfigurationUpdatesArgs>;
}
