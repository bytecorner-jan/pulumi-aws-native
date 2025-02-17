// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    /// <summary>
    /// Describes the notification configuration for an Amazon S3 bucket.
    /// </summary>
    public sealed class BucketNotificationConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("lambdaConfigurations")]
        private InputList<Inputs.BucketLambdaConfigurationArgs>? _lambdaConfigurations;
        public InputList<Inputs.BucketLambdaConfigurationArgs> LambdaConfigurations
        {
            get => _lambdaConfigurations ?? (_lambdaConfigurations = new InputList<Inputs.BucketLambdaConfigurationArgs>());
            set => _lambdaConfigurations = value;
        }

        [Input("queueConfigurations")]
        private InputList<Inputs.BucketQueueConfigurationArgs>? _queueConfigurations;
        public InputList<Inputs.BucketQueueConfigurationArgs> QueueConfigurations
        {
            get => _queueConfigurations ?? (_queueConfigurations = new InputList<Inputs.BucketQueueConfigurationArgs>());
            set => _queueConfigurations = value;
        }

        [Input("topicConfigurations")]
        private InputList<Inputs.BucketTopicConfigurationArgs>? _topicConfigurations;
        public InputList<Inputs.BucketTopicConfigurationArgs> TopicConfigurations
        {
            get => _topicConfigurations ?? (_topicConfigurations = new InputList<Inputs.BucketTopicConfigurationArgs>());
            set => _topicConfigurations = value;
        }

        public BucketNotificationConfigurationArgs()
        {
        }
    }
}
