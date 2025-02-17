// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EMR::Cluster
 *
 * @deprecated Cluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Cluster {
        pulumi.log.warn("Cluster is deprecated: Cluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Cluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:emr:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    public readonly additionalInfo!: pulumi.Output<any | undefined>;
    public readonly applications!: pulumi.Output<outputs.emr.ClusterApplication[] | undefined>;
    public readonly autoScalingRole!: pulumi.Output<string | undefined>;
    public readonly bootstrapActions!: pulumi.Output<outputs.emr.ClusterBootstrapActionConfig[] | undefined>;
    public readonly configurations!: pulumi.Output<outputs.emr.ClusterConfiguration[] | undefined>;
    public readonly customAmiId!: pulumi.Output<string | undefined>;
    public readonly ebsRootVolumeSize!: pulumi.Output<number | undefined>;
    public readonly instances!: pulumi.Output<outputs.emr.ClusterJobFlowInstancesConfig>;
    public readonly jobFlowRole!: pulumi.Output<string>;
    public readonly kerberosAttributes!: pulumi.Output<outputs.emr.ClusterKerberosAttributes | undefined>;
    public readonly logEncryptionKmsKeyId!: pulumi.Output<string | undefined>;
    public readonly logUri!: pulumi.Output<string | undefined>;
    public readonly managedScalingPolicy!: pulumi.Output<outputs.emr.ClusterManagedScalingPolicy | undefined>;
    public /*out*/ readonly masterPublicDNS!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly releaseLabel!: pulumi.Output<string | undefined>;
    public readonly scaleDownBehavior!: pulumi.Output<string | undefined>;
    public readonly securityConfiguration!: pulumi.Output<string | undefined>;
    public readonly serviceRole!: pulumi.Output<string>;
    public readonly stepConcurrencyLevel!: pulumi.Output<number | undefined>;
    public readonly steps!: pulumi.Output<outputs.emr.ClusterStepConfig[] | undefined>;
    public readonly tags!: pulumi.Output<outputs.emr.ClusterTag[] | undefined>;
    public readonly visibleToAllUsers!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Cluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Cluster is deprecated: Cluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instances === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instances'");
            }
            if ((!args || args.jobFlowRole === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobFlowRole'");
            }
            if ((!args || args.serviceRole === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceRole'");
            }
            inputs["additionalInfo"] = args ? args.additionalInfo : undefined;
            inputs["applications"] = args ? args.applications : undefined;
            inputs["autoScalingRole"] = args ? args.autoScalingRole : undefined;
            inputs["bootstrapActions"] = args ? args.bootstrapActions : undefined;
            inputs["configurations"] = args ? args.configurations : undefined;
            inputs["customAmiId"] = args ? args.customAmiId : undefined;
            inputs["ebsRootVolumeSize"] = args ? args.ebsRootVolumeSize : undefined;
            inputs["instances"] = args ? args.instances : undefined;
            inputs["jobFlowRole"] = args ? args.jobFlowRole : undefined;
            inputs["kerberosAttributes"] = args ? args.kerberosAttributes : undefined;
            inputs["logEncryptionKmsKeyId"] = args ? args.logEncryptionKmsKeyId : undefined;
            inputs["logUri"] = args ? args.logUri : undefined;
            inputs["managedScalingPolicy"] = args ? args.managedScalingPolicy : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["releaseLabel"] = args ? args.releaseLabel : undefined;
            inputs["scaleDownBehavior"] = args ? args.scaleDownBehavior : undefined;
            inputs["securityConfiguration"] = args ? args.securityConfiguration : undefined;
            inputs["serviceRole"] = args ? args.serviceRole : undefined;
            inputs["stepConcurrencyLevel"] = args ? args.stepConcurrencyLevel : undefined;
            inputs["steps"] = args ? args.steps : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["visibleToAllUsers"] = args ? args.visibleToAllUsers : undefined;
            inputs["masterPublicDNS"] = undefined /*out*/;
        } else {
            inputs["additionalInfo"] = undefined /*out*/;
            inputs["applications"] = undefined /*out*/;
            inputs["autoScalingRole"] = undefined /*out*/;
            inputs["bootstrapActions"] = undefined /*out*/;
            inputs["configurations"] = undefined /*out*/;
            inputs["customAmiId"] = undefined /*out*/;
            inputs["ebsRootVolumeSize"] = undefined /*out*/;
            inputs["instances"] = undefined /*out*/;
            inputs["jobFlowRole"] = undefined /*out*/;
            inputs["kerberosAttributes"] = undefined /*out*/;
            inputs["logEncryptionKmsKeyId"] = undefined /*out*/;
            inputs["logUri"] = undefined /*out*/;
            inputs["managedScalingPolicy"] = undefined /*out*/;
            inputs["masterPublicDNS"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["releaseLabel"] = undefined /*out*/;
            inputs["scaleDownBehavior"] = undefined /*out*/;
            inputs["securityConfiguration"] = undefined /*out*/;
            inputs["serviceRole"] = undefined /*out*/;
            inputs["stepConcurrencyLevel"] = undefined /*out*/;
            inputs["steps"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["visibleToAllUsers"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    additionalInfo?: any;
    applications?: pulumi.Input<pulumi.Input<inputs.emr.ClusterApplicationArgs>[]>;
    autoScalingRole?: pulumi.Input<string>;
    bootstrapActions?: pulumi.Input<pulumi.Input<inputs.emr.ClusterBootstrapActionConfigArgs>[]>;
    configurations?: pulumi.Input<pulumi.Input<inputs.emr.ClusterConfigurationArgs>[]>;
    customAmiId?: pulumi.Input<string>;
    ebsRootVolumeSize?: pulumi.Input<number>;
    instances: pulumi.Input<inputs.emr.ClusterJobFlowInstancesConfigArgs>;
    jobFlowRole: pulumi.Input<string>;
    kerberosAttributes?: pulumi.Input<inputs.emr.ClusterKerberosAttributesArgs>;
    logEncryptionKmsKeyId?: pulumi.Input<string>;
    logUri?: pulumi.Input<string>;
    managedScalingPolicy?: pulumi.Input<inputs.emr.ClusterManagedScalingPolicyArgs>;
    name?: pulumi.Input<string>;
    releaseLabel?: pulumi.Input<string>;
    scaleDownBehavior?: pulumi.Input<string>;
    securityConfiguration?: pulumi.Input<string>;
    serviceRole: pulumi.Input<string>;
    stepConcurrencyLevel?: pulumi.Input<number>;
    steps?: pulumi.Input<pulumi.Input<inputs.emr.ClusterStepConfigArgs>[]>;
    tags?: pulumi.Input<pulumi.Input<inputs.emr.ClusterTagArgs>[]>;
    visibleToAllUsers?: pulumi.Input<boolean>;
}
