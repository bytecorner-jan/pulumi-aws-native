// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::ClientVpnEndpoint
 *
 * @deprecated ClientVpnEndpoint is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class ClientVpnEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing ClientVpnEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ClientVpnEndpoint {
        pulumi.log.warn("ClientVpnEndpoint is deprecated: ClientVpnEndpoint is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new ClientVpnEndpoint(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:ClientVpnEndpoint';

    /**
     * Returns true if the given object is an instance of ClientVpnEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClientVpnEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClientVpnEndpoint.__pulumiType;
    }

    public readonly authenticationOptions!: pulumi.Output<outputs.ec2.ClientVpnEndpointClientAuthenticationRequest[]>;
    public readonly clientCidrBlock!: pulumi.Output<string>;
    public readonly clientConnectOptions!: pulumi.Output<outputs.ec2.ClientVpnEndpointClientConnectOptions | undefined>;
    public readonly connectionLogOptions!: pulumi.Output<outputs.ec2.ClientVpnEndpointConnectionLogOptions>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly dnsServers!: pulumi.Output<string[] | undefined>;
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    public readonly selfServicePortal!: pulumi.Output<string | undefined>;
    public readonly serverCertificateArn!: pulumi.Output<string>;
    public readonly splitTunnel!: pulumi.Output<boolean | undefined>;
    public readonly tagSpecifications!: pulumi.Output<outputs.ec2.ClientVpnEndpointTagSpecification[] | undefined>;
    public readonly transportProtocol!: pulumi.Output<string | undefined>;
    public readonly vpcId!: pulumi.Output<string | undefined>;
    public readonly vpnPort!: pulumi.Output<number | undefined>;

    /**
     * Create a ClientVpnEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated ClientVpnEndpoint is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: ClientVpnEndpointArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ClientVpnEndpoint is deprecated: ClientVpnEndpoint is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.authenticationOptions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authenticationOptions'");
            }
            if ((!args || args.clientCidrBlock === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientCidrBlock'");
            }
            if ((!args || args.connectionLogOptions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionLogOptions'");
            }
            if ((!args || args.serverCertificateArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverCertificateArn'");
            }
            inputs["authenticationOptions"] = args ? args.authenticationOptions : undefined;
            inputs["clientCidrBlock"] = args ? args.clientCidrBlock : undefined;
            inputs["clientConnectOptions"] = args ? args.clientConnectOptions : undefined;
            inputs["connectionLogOptions"] = args ? args.connectionLogOptions : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["dnsServers"] = args ? args.dnsServers : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["selfServicePortal"] = args ? args.selfServicePortal : undefined;
            inputs["serverCertificateArn"] = args ? args.serverCertificateArn : undefined;
            inputs["splitTunnel"] = args ? args.splitTunnel : undefined;
            inputs["tagSpecifications"] = args ? args.tagSpecifications : undefined;
            inputs["transportProtocol"] = args ? args.transportProtocol : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["vpnPort"] = args ? args.vpnPort : undefined;
        } else {
            inputs["authenticationOptions"] = undefined /*out*/;
            inputs["clientCidrBlock"] = undefined /*out*/;
            inputs["clientConnectOptions"] = undefined /*out*/;
            inputs["connectionLogOptions"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["dnsServers"] = undefined /*out*/;
            inputs["securityGroupIds"] = undefined /*out*/;
            inputs["selfServicePortal"] = undefined /*out*/;
            inputs["serverCertificateArn"] = undefined /*out*/;
            inputs["splitTunnel"] = undefined /*out*/;
            inputs["tagSpecifications"] = undefined /*out*/;
            inputs["transportProtocol"] = undefined /*out*/;
            inputs["vpcId"] = undefined /*out*/;
            inputs["vpnPort"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ClientVpnEndpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ClientVpnEndpoint resource.
 */
export interface ClientVpnEndpointArgs {
    authenticationOptions: pulumi.Input<pulumi.Input<inputs.ec2.ClientVpnEndpointClientAuthenticationRequestArgs>[]>;
    clientCidrBlock: pulumi.Input<string>;
    clientConnectOptions?: pulumi.Input<inputs.ec2.ClientVpnEndpointClientConnectOptionsArgs>;
    connectionLogOptions: pulumi.Input<inputs.ec2.ClientVpnEndpointConnectionLogOptionsArgs>;
    description?: pulumi.Input<string>;
    dnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    selfServicePortal?: pulumi.Input<string>;
    serverCertificateArn: pulumi.Input<string>;
    splitTunnel?: pulumi.Input<boolean>;
    tagSpecifications?: pulumi.Input<pulumi.Input<inputs.ec2.ClientVpnEndpointTagSpecificationArgs>[]>;
    transportProtocol?: pulumi.Input<string>;
    vpcId?: pulumi.Input<string>;
    vpnPort?: pulumi.Input<number>;
}
