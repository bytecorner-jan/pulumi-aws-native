// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGateway::GatewayResponse
 *
 * @deprecated GatewayResponse is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class GatewayResponse extends pulumi.CustomResource {
    /**
     * Get an existing GatewayResponse resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GatewayResponse {
        pulumi.log.warn("GatewayResponse is deprecated: GatewayResponse is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new GatewayResponse(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:apigateway:GatewayResponse';

    /**
     * Returns true if the given object is an instance of GatewayResponse.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GatewayResponse {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GatewayResponse.__pulumiType;
    }

    /**
     * The response parameters (paths, query strings, and headers) for the response.
     */
    public readonly responseParameters!: pulumi.Output<any | undefined>;
    /**
     * The response templates for the response.
     */
    public readonly responseTemplates!: pulumi.Output<any | undefined>;
    /**
     * The type of the Gateway Response.
     */
    public readonly responseType!: pulumi.Output<string>;
    /**
     * The identifier of the API.
     */
    public readonly restApiId!: pulumi.Output<string>;
    /**
     * The HTTP status code for the response.
     */
    public readonly statusCode!: pulumi.Output<string | undefined>;

    /**
     * Create a GatewayResponse resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated GatewayResponse is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: GatewayResponseArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("GatewayResponse is deprecated: GatewayResponse is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.responseType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'responseType'");
            }
            if ((!args || args.restApiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'restApiId'");
            }
            inputs["responseParameters"] = args ? args.responseParameters : undefined;
            inputs["responseTemplates"] = args ? args.responseTemplates : undefined;
            inputs["responseType"] = args ? args.responseType : undefined;
            inputs["restApiId"] = args ? args.restApiId : undefined;
            inputs["statusCode"] = args ? args.statusCode : undefined;
        } else {
            inputs["responseParameters"] = undefined /*out*/;
            inputs["responseTemplates"] = undefined /*out*/;
            inputs["responseType"] = undefined /*out*/;
            inputs["restApiId"] = undefined /*out*/;
            inputs["statusCode"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(GatewayResponse.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a GatewayResponse resource.
 */
export interface GatewayResponseArgs {
    /**
     * The response parameters (paths, query strings, and headers) for the response.
     */
    responseParameters?: any;
    /**
     * The response templates for the response.
     */
    responseTemplates?: any;
    /**
     * The type of the Gateway Response.
     */
    responseType: pulumi.Input<string>;
    /**
     * The identifier of the API.
     */
    restApiId: pulumi.Input<string>;
    /**
     * The HTTP status code for the response.
     */
    statusCode?: pulumi.Input<string>;
}
