// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html
 */
export class AssessmentTemplate extends pulumi.CustomResource {
    /**
     * Get an existing AssessmentTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AssessmentTemplate {
        return new AssessmentTemplate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:inspector:AssessmentTemplate';

    /**
     * Returns true if the given object is an instance of AssessmentTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AssessmentTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AssessmentTemplate.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html#cfn-inspector-assessmenttemplate-assessmenttargetarn
     */
    public readonly assessmentTargetArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html#cfn-inspector-assessmenttemplate-assessmenttemplatename
     */
    public readonly assessmentTemplateName!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html#cfn-inspector-assessmenttemplate-durationinseconds
     */
    public readonly durationInSeconds!: pulumi.Output<number>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html#cfn-inspector-assessmenttemplate-rulespackagearns
     */
    public readonly rulesPackageArns!: pulumi.Output<string[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html#cfn-inspector-assessmenttemplate-userattributesforfindings
     */
    public readonly userAttributesForFindings!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a AssessmentTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssessmentTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.assessmentTargetArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'assessmentTargetArn'");
            }
            if ((!args || args.durationInSeconds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'durationInSeconds'");
            }
            if ((!args || args.rulesPackageArns === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rulesPackageArns'");
            }
            inputs["assessmentTargetArn"] = args ? args.assessmentTargetArn : undefined;
            inputs["assessmentTemplateName"] = args ? args.assessmentTemplateName : undefined;
            inputs["durationInSeconds"] = args ? args.durationInSeconds : undefined;
            inputs["rulesPackageArns"] = args ? args.rulesPackageArns : undefined;
            inputs["userAttributesForFindings"] = args ? args.userAttributesForFindings : undefined;
            inputs["arn"] = undefined /*out*/;
        } else {
            inputs["arn"] = undefined /*out*/;
            inputs["assessmentTargetArn"] = undefined /*out*/;
            inputs["assessmentTemplateName"] = undefined /*out*/;
            inputs["durationInSeconds"] = undefined /*out*/;
            inputs["rulesPackageArns"] = undefined /*out*/;
            inputs["userAttributesForFindings"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AssessmentTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AssessmentTemplate resource.
 */
export interface AssessmentTemplateArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html#cfn-inspector-assessmenttemplate-assessmenttargetarn
     */
    assessmentTargetArn: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html#cfn-inspector-assessmenttemplate-assessmenttemplatename
     */
    assessmentTemplateName?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html#cfn-inspector-assessmenttemplate-durationinseconds
     */
    durationInSeconds: pulumi.Input<number>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html#cfn-inspector-assessmenttemplate-rulespackagearns
     */
    rulesPackageArns: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html#cfn-inspector-assessmenttemplate-userattributesforfindings
     */
    userAttributesForFindings?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
