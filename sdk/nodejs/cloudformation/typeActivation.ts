// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Enable a resource that has been published in the CloudFormation Registry.
 */
export class TypeActivation extends pulumi.CustomResource {
    /**
     * Get an existing TypeActivation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TypeActivation {
        return new TypeActivation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cloudformation:TypeActivation';

    /**
     * Returns true if the given object is an instance of TypeActivation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TypeActivation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TypeActivation.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the extension.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Whether to automatically update the extension in this account and region when a new minor version is published by the extension publisher. Major versions released by the publisher must be manually updated.
     */
    public readonly autoUpdate!: pulumi.Output<boolean | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
     */
    public readonly executionRoleArn!: pulumi.Output<string | undefined>;
    /**
     * Specifies logging configuration information for a type.
     */
    public readonly loggingConfig!: pulumi.Output<outputs.cloudformation.TypeActivationLoggingConfig | undefined>;
    /**
     * The Major Version of the type you want to enable
     */
    public readonly majorVersion!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Number (ARN) assigned to the public extension upon publication
     */
    public readonly publicTypeArn!: pulumi.Output<string | undefined>;
    /**
     * The publisher id assigned by CloudFormation for publishing in this region.
     */
    public readonly publisherId!: pulumi.Output<string | undefined>;
    /**
     * The kind of extension
     */
    public readonly type!: pulumi.Output<enums.cloudformation.TypeActivationType | undefined>;
    /**
     * The name of the type being registered.
     *
     * We recommend that type names adhere to the following pattern: company_or_organization::service::type.
     */
    public readonly typeName!: pulumi.Output<string | undefined>;
    /**
     * An alias to assign to the public extension in this account and region. If you specify an alias for the extension, you must then use the alias to refer to the extension in your templates.
     */
    public readonly typeNameAlias!: pulumi.Output<string | undefined>;
    /**
     * Manually updates a previously-enabled type to a new major or minor version, if available. You can also use this parameter to update the value of AutoUpdateEnabled
     */
    public readonly versionBump!: pulumi.Output<enums.cloudformation.TypeActivationVersionBump | undefined>;

    /**
     * Create a TypeActivation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TypeActivationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["autoUpdate"] = args ? args.autoUpdate : undefined;
            inputs["executionRoleArn"] = args ? args.executionRoleArn : undefined;
            inputs["loggingConfig"] = args ? args.loggingConfig : undefined;
            inputs["majorVersion"] = args ? args.majorVersion : undefined;
            inputs["publicTypeArn"] = args ? args.publicTypeArn : undefined;
            inputs["publisherId"] = args ? args.publisherId : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["typeName"] = args ? args.typeName : undefined;
            inputs["typeNameAlias"] = args ? args.typeNameAlias : undefined;
            inputs["versionBump"] = args ? args.versionBump : undefined;
            inputs["arn"] = undefined /*out*/;
        } else {
            inputs["arn"] = undefined /*out*/;
            inputs["autoUpdate"] = undefined /*out*/;
            inputs["executionRoleArn"] = undefined /*out*/;
            inputs["loggingConfig"] = undefined /*out*/;
            inputs["majorVersion"] = undefined /*out*/;
            inputs["publicTypeArn"] = undefined /*out*/;
            inputs["publisherId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["typeName"] = undefined /*out*/;
            inputs["typeNameAlias"] = undefined /*out*/;
            inputs["versionBump"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TypeActivation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a TypeActivation resource.
 */
export interface TypeActivationArgs {
    /**
     * Whether to automatically update the extension in this account and region when a new minor version is published by the extension publisher. Major versions released by the publisher must be manually updated.
     */
    autoUpdate?: pulumi.Input<boolean>;
    /**
     * The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
     */
    executionRoleArn?: pulumi.Input<string>;
    /**
     * Specifies logging configuration information for a type.
     */
    loggingConfig?: pulumi.Input<inputs.cloudformation.TypeActivationLoggingConfigArgs>;
    /**
     * The Major Version of the type you want to enable
     */
    majorVersion?: pulumi.Input<string>;
    /**
     * The Amazon Resource Number (ARN) assigned to the public extension upon publication
     */
    publicTypeArn?: pulumi.Input<string>;
    /**
     * The publisher id assigned by CloudFormation for publishing in this region.
     */
    publisherId?: pulumi.Input<string>;
    /**
     * The kind of extension
     */
    type?: pulumi.Input<enums.cloudformation.TypeActivationType>;
    /**
     * The name of the type being registered.
     *
     * We recommend that type names adhere to the following pattern: company_or_organization::service::type.
     */
    typeName?: pulumi.Input<string>;
    /**
     * An alias to assign to the public extension in this account and region. If you specify an alias for the extension, you must then use the alias to refer to the extension in your templates.
     */
    typeNameAlias?: pulumi.Input<string>;
    /**
     * Manually updates a previously-enabled type to a new major or minor version, if available. You can also use this parameter to update the value of AutoUpdateEnabled
     */
    versionBump?: pulumi.Input<enums.cloudformation.TypeActivationVersionBump>;
}
