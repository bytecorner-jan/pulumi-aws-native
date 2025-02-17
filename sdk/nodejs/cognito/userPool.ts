// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Cognito::UserPool
 *
 * @deprecated UserPool is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class UserPool extends pulumi.CustomResource {
    /**
     * Get an existing UserPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): UserPool {
        pulumi.log.warn("UserPool is deprecated: UserPool is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new UserPool(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cognito:UserPool';

    /**
     * Returns true if the given object is an instance of UserPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserPool.__pulumiType;
    }

    public readonly accountRecoverySetting!: pulumi.Output<outputs.cognito.UserPoolAccountRecoverySetting | undefined>;
    public readonly adminCreateUserConfig!: pulumi.Output<outputs.cognito.UserPoolAdminCreateUserConfig | undefined>;
    public readonly aliasAttributes!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly autoVerifiedAttributes!: pulumi.Output<string[] | undefined>;
    public readonly deviceConfiguration!: pulumi.Output<outputs.cognito.UserPoolDeviceConfiguration | undefined>;
    public readonly emailConfiguration!: pulumi.Output<outputs.cognito.UserPoolEmailConfiguration | undefined>;
    public readonly emailVerificationMessage!: pulumi.Output<string | undefined>;
    public readonly emailVerificationSubject!: pulumi.Output<string | undefined>;
    public readonly enabledMfas!: pulumi.Output<string[] | undefined>;
    public readonly lambdaConfig!: pulumi.Output<outputs.cognito.UserPoolLambdaConfig | undefined>;
    public readonly mfaConfiguration!: pulumi.Output<string | undefined>;
    public readonly policies!: pulumi.Output<outputs.cognito.UserPoolPolicies | undefined>;
    public /*out*/ readonly providerName!: pulumi.Output<string>;
    public /*out*/ readonly providerURL!: pulumi.Output<string>;
    public readonly schema!: pulumi.Output<outputs.cognito.UserPoolSchemaAttribute[] | undefined>;
    public readonly smsAuthenticationMessage!: pulumi.Output<string | undefined>;
    public readonly smsConfiguration!: pulumi.Output<outputs.cognito.UserPoolSmsConfiguration | undefined>;
    public readonly smsVerificationMessage!: pulumi.Output<string | undefined>;
    public readonly userPoolAddOns!: pulumi.Output<outputs.cognito.UserPoolAddOns | undefined>;
    public readonly userPoolName!: pulumi.Output<string | undefined>;
    public readonly userPoolTags!: pulumi.Output<any | undefined>;
    public readonly usernameAttributes!: pulumi.Output<string[] | undefined>;
    public readonly usernameConfiguration!: pulumi.Output<outputs.cognito.UserPoolUsernameConfiguration | undefined>;
    public readonly verificationMessageTemplate!: pulumi.Output<outputs.cognito.UserPoolVerificationMessageTemplate | undefined>;

    /**
     * Create a UserPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated UserPool is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: UserPoolArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("UserPool is deprecated: UserPool is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["accountRecoverySetting"] = args ? args.accountRecoverySetting : undefined;
            inputs["adminCreateUserConfig"] = args ? args.adminCreateUserConfig : undefined;
            inputs["aliasAttributes"] = args ? args.aliasAttributes : undefined;
            inputs["autoVerifiedAttributes"] = args ? args.autoVerifiedAttributes : undefined;
            inputs["deviceConfiguration"] = args ? args.deviceConfiguration : undefined;
            inputs["emailConfiguration"] = args ? args.emailConfiguration : undefined;
            inputs["emailVerificationMessage"] = args ? args.emailVerificationMessage : undefined;
            inputs["emailVerificationSubject"] = args ? args.emailVerificationSubject : undefined;
            inputs["enabledMfas"] = args ? args.enabledMfas : undefined;
            inputs["lambdaConfig"] = args ? args.lambdaConfig : undefined;
            inputs["mfaConfiguration"] = args ? args.mfaConfiguration : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["schema"] = args ? args.schema : undefined;
            inputs["smsAuthenticationMessage"] = args ? args.smsAuthenticationMessage : undefined;
            inputs["smsConfiguration"] = args ? args.smsConfiguration : undefined;
            inputs["smsVerificationMessage"] = args ? args.smsVerificationMessage : undefined;
            inputs["userPoolAddOns"] = args ? args.userPoolAddOns : undefined;
            inputs["userPoolName"] = args ? args.userPoolName : undefined;
            inputs["userPoolTags"] = args ? args.userPoolTags : undefined;
            inputs["usernameAttributes"] = args ? args.usernameAttributes : undefined;
            inputs["usernameConfiguration"] = args ? args.usernameConfiguration : undefined;
            inputs["verificationMessageTemplate"] = args ? args.verificationMessageTemplate : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["providerName"] = undefined /*out*/;
            inputs["providerURL"] = undefined /*out*/;
        } else {
            inputs["accountRecoverySetting"] = undefined /*out*/;
            inputs["adminCreateUserConfig"] = undefined /*out*/;
            inputs["aliasAttributes"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["autoVerifiedAttributes"] = undefined /*out*/;
            inputs["deviceConfiguration"] = undefined /*out*/;
            inputs["emailConfiguration"] = undefined /*out*/;
            inputs["emailVerificationMessage"] = undefined /*out*/;
            inputs["emailVerificationSubject"] = undefined /*out*/;
            inputs["enabledMfas"] = undefined /*out*/;
            inputs["lambdaConfig"] = undefined /*out*/;
            inputs["mfaConfiguration"] = undefined /*out*/;
            inputs["policies"] = undefined /*out*/;
            inputs["providerName"] = undefined /*out*/;
            inputs["providerURL"] = undefined /*out*/;
            inputs["schema"] = undefined /*out*/;
            inputs["smsAuthenticationMessage"] = undefined /*out*/;
            inputs["smsConfiguration"] = undefined /*out*/;
            inputs["smsVerificationMessage"] = undefined /*out*/;
            inputs["userPoolAddOns"] = undefined /*out*/;
            inputs["userPoolName"] = undefined /*out*/;
            inputs["userPoolTags"] = undefined /*out*/;
            inputs["usernameAttributes"] = undefined /*out*/;
            inputs["usernameConfiguration"] = undefined /*out*/;
            inputs["verificationMessageTemplate"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(UserPool.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a UserPool resource.
 */
export interface UserPoolArgs {
    accountRecoverySetting?: pulumi.Input<inputs.cognito.UserPoolAccountRecoverySettingArgs>;
    adminCreateUserConfig?: pulumi.Input<inputs.cognito.UserPoolAdminCreateUserConfigArgs>;
    aliasAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    autoVerifiedAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    deviceConfiguration?: pulumi.Input<inputs.cognito.UserPoolDeviceConfigurationArgs>;
    emailConfiguration?: pulumi.Input<inputs.cognito.UserPoolEmailConfigurationArgs>;
    emailVerificationMessage?: pulumi.Input<string>;
    emailVerificationSubject?: pulumi.Input<string>;
    enabledMfas?: pulumi.Input<pulumi.Input<string>[]>;
    lambdaConfig?: pulumi.Input<inputs.cognito.UserPoolLambdaConfigArgs>;
    mfaConfiguration?: pulumi.Input<string>;
    policies?: pulumi.Input<inputs.cognito.UserPoolPoliciesArgs>;
    schema?: pulumi.Input<pulumi.Input<inputs.cognito.UserPoolSchemaAttributeArgs>[]>;
    smsAuthenticationMessage?: pulumi.Input<string>;
    smsConfiguration?: pulumi.Input<inputs.cognito.UserPoolSmsConfigurationArgs>;
    smsVerificationMessage?: pulumi.Input<string>;
    userPoolAddOns?: pulumi.Input<inputs.cognito.UserPoolAddOnsArgs>;
    userPoolName?: pulumi.Input<string>;
    userPoolTags?: any;
    usernameAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    usernameConfiguration?: pulumi.Input<inputs.cognito.UserPoolUsernameConfigurationArgs>;
    verificationMessageTemplate?: pulumi.Input<inputs.cognito.UserPoolVerificationMessageTemplateArgs>;
}
