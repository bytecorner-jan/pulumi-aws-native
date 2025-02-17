// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::ModelPackageGroup
 */
export class ModelPackageGroup extends pulumi.CustomResource {
    /**
     * Get an existing ModelPackageGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ModelPackageGroup {
        return new ModelPackageGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sagemaker:ModelPackageGroup';

    /**
     * Returns true if the given object is an instance of ModelPackageGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ModelPackageGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ModelPackageGroup.__pulumiType;
    }

    /**
     * The time at which the model package group was created.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    public /*out*/ readonly modelPackageGroupArn!: pulumi.Output<string>;
    public readonly modelPackageGroupDescription!: pulumi.Output<string | undefined>;
    public readonly modelPackageGroupName!: pulumi.Output<string>;
    public readonly modelPackageGroupPolicy!: pulumi.Output<any | undefined>;
    /**
     * The status of a modelpackage group job.
     */
    public /*out*/ readonly modelPackageGroupStatus!: pulumi.Output<enums.sagemaker.ModelPackageGroupStatus>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.sagemaker.ModelPackageGroupTag[] | undefined>;

    /**
     * Create a ModelPackageGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ModelPackageGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["modelPackageGroupDescription"] = args ? args.modelPackageGroupDescription : undefined;
            inputs["modelPackageGroupName"] = args ? args.modelPackageGroupName : undefined;
            inputs["modelPackageGroupPolicy"] = args ? args.modelPackageGroupPolicy : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["modelPackageGroupArn"] = undefined /*out*/;
            inputs["modelPackageGroupStatus"] = undefined /*out*/;
        } else {
            inputs["creationTime"] = undefined /*out*/;
            inputs["modelPackageGroupArn"] = undefined /*out*/;
            inputs["modelPackageGroupDescription"] = undefined /*out*/;
            inputs["modelPackageGroupName"] = undefined /*out*/;
            inputs["modelPackageGroupPolicy"] = undefined /*out*/;
            inputs["modelPackageGroupStatus"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ModelPackageGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ModelPackageGroup resource.
 */
export interface ModelPackageGroupArgs {
    modelPackageGroupDescription?: pulumi.Input<string>;
    modelPackageGroupName?: pulumi.Input<string>;
    modelPackageGroupPolicy?: any;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.sagemaker.ModelPackageGroupTagArgs>[]>;
}
