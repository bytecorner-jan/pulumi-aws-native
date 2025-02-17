// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppSync::GraphQLApi
 *
 * @deprecated GraphQLApi is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class GraphQLApi extends pulumi.CustomResource {
    /**
     * Get an existing GraphQLApi resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GraphQLApi {
        pulumi.log.warn("GraphQLApi is deprecated: GraphQLApi is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new GraphQLApi(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:appsync:GraphQLApi';

    /**
     * Returns true if the given object is an instance of GraphQLApi.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GraphQLApi {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GraphQLApi.__pulumiType;
    }

    public readonly additionalAuthenticationProviders!: pulumi.Output<outputs.appsync.GraphQLApiAdditionalAuthenticationProviders | undefined>;
    public /*out*/ readonly apiId!: pulumi.Output<string>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly authenticationType!: pulumi.Output<string>;
    public /*out*/ readonly graphQLUrl!: pulumi.Output<string>;
    public readonly lambdaAuthorizerConfig!: pulumi.Output<outputs.appsync.GraphQLApiLambdaAuthorizerConfig | undefined>;
    public readonly logConfig!: pulumi.Output<outputs.appsync.GraphQLApiLogConfig | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly openIDConnectConfig!: pulumi.Output<outputs.appsync.GraphQLApiOpenIDConnectConfig | undefined>;
    public readonly tags!: pulumi.Output<outputs.appsync.GraphQLApiTags | undefined>;
    public readonly userPoolConfig!: pulumi.Output<outputs.appsync.GraphQLApiUserPoolConfig | undefined>;
    public readonly xrayEnabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a GraphQLApi resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated GraphQLApi is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: GraphQLApiArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("GraphQLApi is deprecated: GraphQLApi is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.authenticationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authenticationType'");
            }
            inputs["additionalAuthenticationProviders"] = args ? args.additionalAuthenticationProviders : undefined;
            inputs["authenticationType"] = args ? args.authenticationType : undefined;
            inputs["lambdaAuthorizerConfig"] = args ? args.lambdaAuthorizerConfig : undefined;
            inputs["logConfig"] = args ? args.logConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["openIDConnectConfig"] = args ? args.openIDConnectConfig : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["userPoolConfig"] = args ? args.userPoolConfig : undefined;
            inputs["xrayEnabled"] = args ? args.xrayEnabled : undefined;
            inputs["apiId"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["graphQLUrl"] = undefined /*out*/;
        } else {
            inputs["additionalAuthenticationProviders"] = undefined /*out*/;
            inputs["apiId"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["authenticationType"] = undefined /*out*/;
            inputs["graphQLUrl"] = undefined /*out*/;
            inputs["lambdaAuthorizerConfig"] = undefined /*out*/;
            inputs["logConfig"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["openIDConnectConfig"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["userPoolConfig"] = undefined /*out*/;
            inputs["xrayEnabled"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(GraphQLApi.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a GraphQLApi resource.
 */
export interface GraphQLApiArgs {
    additionalAuthenticationProviders?: pulumi.Input<inputs.appsync.GraphQLApiAdditionalAuthenticationProvidersArgs>;
    authenticationType: pulumi.Input<string>;
    lambdaAuthorizerConfig?: pulumi.Input<inputs.appsync.GraphQLApiLambdaAuthorizerConfigArgs>;
    logConfig?: pulumi.Input<inputs.appsync.GraphQLApiLogConfigArgs>;
    name?: pulumi.Input<string>;
    openIDConnectConfig?: pulumi.Input<inputs.appsync.GraphQLApiOpenIDConnectConfigArgs>;
    tags?: pulumi.Input<inputs.appsync.GraphQLApiTagsArgs>;
    userPoolConfig?: pulumi.Input<inputs.appsync.GraphQLApiUserPoolConfigArgs>;
    xrayEnabled?: pulumi.Input<boolean>;
}
