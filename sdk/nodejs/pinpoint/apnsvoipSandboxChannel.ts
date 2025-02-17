// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Pinpoint::APNSVoipSandboxChannel
 *
 * @deprecated APNSVoipSandboxChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class APNSVoipSandboxChannel extends pulumi.CustomResource {
    /**
     * Get an existing APNSVoipSandboxChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): APNSVoipSandboxChannel {
        pulumi.log.warn("APNSVoipSandboxChannel is deprecated: APNSVoipSandboxChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new APNSVoipSandboxChannel(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:pinpoint:APNSVoipSandboxChannel';

    /**
     * Returns true if the given object is an instance of APNSVoipSandboxChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is APNSVoipSandboxChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === APNSVoipSandboxChannel.__pulumiType;
    }

    public readonly applicationId!: pulumi.Output<string>;
    public readonly bundleId!: pulumi.Output<string | undefined>;
    public readonly certificate!: pulumi.Output<string | undefined>;
    public readonly defaultAuthenticationMethod!: pulumi.Output<string | undefined>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly privateKey!: pulumi.Output<string | undefined>;
    public readonly teamId!: pulumi.Output<string | undefined>;
    public readonly tokenKey!: pulumi.Output<string | undefined>;
    public readonly tokenKeyId!: pulumi.Output<string | undefined>;

    /**
     * Create a APNSVoipSandboxChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated APNSVoipSandboxChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: APNSVoipSandboxChannelArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("APNSVoipSandboxChannel is deprecated: APNSVoipSandboxChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            inputs["applicationId"] = args ? args.applicationId : undefined;
            inputs["bundleId"] = args ? args.bundleId : undefined;
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["defaultAuthenticationMethod"] = args ? args.defaultAuthenticationMethod : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["privateKey"] = args ? args.privateKey : undefined;
            inputs["teamId"] = args ? args.teamId : undefined;
            inputs["tokenKey"] = args ? args.tokenKey : undefined;
            inputs["tokenKeyId"] = args ? args.tokenKeyId : undefined;
        } else {
            inputs["applicationId"] = undefined /*out*/;
            inputs["bundleId"] = undefined /*out*/;
            inputs["certificate"] = undefined /*out*/;
            inputs["defaultAuthenticationMethod"] = undefined /*out*/;
            inputs["enabled"] = undefined /*out*/;
            inputs["privateKey"] = undefined /*out*/;
            inputs["teamId"] = undefined /*out*/;
            inputs["tokenKey"] = undefined /*out*/;
            inputs["tokenKeyId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(APNSVoipSandboxChannel.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a APNSVoipSandboxChannel resource.
 */
export interface APNSVoipSandboxChannelArgs {
    applicationId: pulumi.Input<string>;
    bundleId?: pulumi.Input<string>;
    certificate?: pulumi.Input<string>;
    defaultAuthenticationMethod?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    privateKey?: pulumi.Input<string>;
    teamId?: pulumi.Input<string>;
    tokenKey?: pulumi.Input<string>;
    tokenKeyId?: pulumi.Input<string>;
}
