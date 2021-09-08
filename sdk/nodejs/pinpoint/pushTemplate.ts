// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html
 */
export class PushTemplate extends pulumi.CustomResource {
    /**
     * Get an existing PushTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PushTemplate {
        return new PushTemplate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:pinpoint:PushTemplate';

    /**
     * Returns true if the given object is an instance of PushTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PushTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PushTemplate.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-adm
     */
    public readonly aDM!: pulumi.Output<outputs.pinpoint.PushTemplateAndroidPushNotificationTemplate | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-apns
     */
    public readonly aPNS!: pulumi.Output<outputs.pinpoint.PushTemplateAPNSPushNotificationTemplate | undefined>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-baidu
     */
    public readonly baidu!: pulumi.Output<outputs.pinpoint.PushTemplateAndroidPushNotificationTemplate | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-default
     */
    public readonly default!: pulumi.Output<outputs.pinpoint.PushTemplateDefaultPushNotificationTemplate | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-defaultsubstitutions
     */
    public readonly defaultSubstitutions!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-gcm
     */
    public readonly gCM!: pulumi.Output<outputs.pinpoint.PushTemplateAndroidPushNotificationTemplate | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-tags
     */
    public readonly tags!: pulumi.Output<any | string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-templatedescription
     */
    public readonly templateDescription!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-templatename
     */
    public readonly templateName!: pulumi.Output<string>;

    /**
     * Create a PushTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PushTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.templateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateName'");
            }
            inputs["aDM"] = args ? args.aDM : undefined;
            inputs["aPNS"] = args ? args.aPNS : undefined;
            inputs["baidu"] = args ? args.baidu : undefined;
            inputs["default"] = args ? args.default : undefined;
            inputs["defaultSubstitutions"] = args ? args.defaultSubstitutions : undefined;
            inputs["gCM"] = args ? args.gCM : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["templateDescription"] = args ? args.templateDescription : undefined;
            inputs["templateName"] = args ? args.templateName : undefined;
            inputs["arn"] = undefined /*out*/;
        } else {
            inputs["aDM"] = undefined /*out*/;
            inputs["aPNS"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["baidu"] = undefined /*out*/;
            inputs["default"] = undefined /*out*/;
            inputs["defaultSubstitutions"] = undefined /*out*/;
            inputs["gCM"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["templateDescription"] = undefined /*out*/;
            inputs["templateName"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(PushTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PushTemplate resource.
 */
export interface PushTemplateArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-adm
     */
    aDM?: pulumi.Input<inputs.pinpoint.PushTemplateAndroidPushNotificationTemplateArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-apns
     */
    aPNS?: pulumi.Input<inputs.pinpoint.PushTemplateAPNSPushNotificationTemplateArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-baidu
     */
    baidu?: pulumi.Input<inputs.pinpoint.PushTemplateAndroidPushNotificationTemplateArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-default
     */
    default?: pulumi.Input<inputs.pinpoint.PushTemplateDefaultPushNotificationTemplateArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-defaultsubstitutions
     */
    defaultSubstitutions?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-gcm
     */
    gCM?: pulumi.Input<inputs.pinpoint.PushTemplateAndroidPushNotificationTemplateArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-tags
     */
    tags?: pulumi.Input<any | string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-templatedescription
     */
    templateDescription?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-templatename
     */
    templateName: pulumi.Input<string>;
}
