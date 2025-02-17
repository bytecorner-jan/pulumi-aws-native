// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::GameLift::Build
 *
 * @deprecated Build is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Build extends pulumi.CustomResource {
    /**
     * Get an existing Build resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Build {
        pulumi.log.warn("Build is deprecated: Build is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Build(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:gamelift:Build';

    /**
     * Returns true if the given object is an instance of Build.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Build {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Build.__pulumiType;
    }

    public readonly name!: pulumi.Output<string | undefined>;
    public readonly operatingSystem!: pulumi.Output<string | undefined>;
    public readonly storageLocation!: pulumi.Output<outputs.gamelift.BuildS3Location | undefined>;
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a Build resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Build is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: BuildArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Build is deprecated: Build is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["name"] = args ? args.name : undefined;
            inputs["operatingSystem"] = args ? args.operatingSystem : undefined;
            inputs["storageLocation"] = args ? args.storageLocation : undefined;
            inputs["version"] = args ? args.version : undefined;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["operatingSystem"] = undefined /*out*/;
            inputs["storageLocation"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Build.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Build resource.
 */
export interface BuildArgs {
    name?: pulumi.Input<string>;
    operatingSystem?: pulumi.Input<string>;
    storageLocation?: pulumi.Input<inputs.gamelift.BuildS3LocationArgs>;
    version?: pulumi.Input<string>;
}
