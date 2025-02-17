// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SNS::Topic
 *
 * @deprecated Topic is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Topic extends pulumi.CustomResource {
    /**
     * Get an existing Topic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Topic {
        pulumi.log.warn("Topic is deprecated: Topic is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Topic(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sns:Topic';

    /**
     * Returns true if the given object is an instance of Topic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Topic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Topic.__pulumiType;
    }

    public readonly contentBasedDeduplication!: pulumi.Output<boolean | undefined>;
    public readonly displayName!: pulumi.Output<string | undefined>;
    public readonly fifoTopic!: pulumi.Output<boolean | undefined>;
    public readonly kmsMasterKeyId!: pulumi.Output<string | undefined>;
    public readonly subscription!: pulumi.Output<outputs.sns.TopicSubscription[] | undefined>;
    public readonly tags!: pulumi.Output<outputs.sns.TopicTag[] | undefined>;
    public readonly topicName!: pulumi.Output<string | undefined>;

    /**
     * Create a Topic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Topic is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: TopicArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Topic is deprecated: Topic is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["contentBasedDeduplication"] = args ? args.contentBasedDeduplication : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["fifoTopic"] = args ? args.fifoTopic : undefined;
            inputs["kmsMasterKeyId"] = args ? args.kmsMasterKeyId : undefined;
            inputs["subscription"] = args ? args.subscription : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["topicName"] = args ? args.topicName : undefined;
        } else {
            inputs["contentBasedDeduplication"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["fifoTopic"] = undefined /*out*/;
            inputs["kmsMasterKeyId"] = undefined /*out*/;
            inputs["subscription"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["topicName"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Topic.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Topic resource.
 */
export interface TopicArgs {
    contentBasedDeduplication?: pulumi.Input<boolean>;
    displayName?: pulumi.Input<string>;
    fifoTopic?: pulumi.Input<boolean>;
    kmsMasterKeyId?: pulumi.Input<string>;
    subscription?: pulumi.Input<pulumi.Input<inputs.sns.TopicSubscriptionArgs>[]>;
    tags?: pulumi.Input<pulumi.Input<inputs.sns.TopicTagArgs>[]>;
    topicName?: pulumi.Input<string>;
}
