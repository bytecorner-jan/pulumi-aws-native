// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Outputs
{

    [OutputType]
    public sealed class DeliveryStreamSplunkDestinationConfiguration
    {
        public readonly Outputs.DeliveryStreamCloudWatchLoggingOptions? CloudWatchLoggingOptions;
        public readonly int? HECAcknowledgmentTimeoutInSeconds;
        public readonly string HECEndpoint;
        public readonly Pulumi.AwsNative.KinesisFirehose.DeliveryStreamSplunkDestinationConfigurationHECEndpointType HECEndpointType;
        public readonly string HECToken;
        public readonly Outputs.DeliveryStreamProcessingConfiguration? ProcessingConfiguration;
        public readonly Outputs.DeliveryStreamSplunkRetryOptions? RetryOptions;
        public readonly string? S3BackupMode;
        public readonly Outputs.DeliveryStreamS3DestinationConfiguration S3Configuration;

        [OutputConstructor]
        private DeliveryStreamSplunkDestinationConfiguration(
            Outputs.DeliveryStreamCloudWatchLoggingOptions? cloudWatchLoggingOptions,

            int? hECAcknowledgmentTimeoutInSeconds,

            string hECEndpoint,

            Pulumi.AwsNative.KinesisFirehose.DeliveryStreamSplunkDestinationConfigurationHECEndpointType hECEndpointType,

            string hECToken,

            Outputs.DeliveryStreamProcessingConfiguration? processingConfiguration,

            Outputs.DeliveryStreamSplunkRetryOptions? retryOptions,

            string? s3BackupMode,

            Outputs.DeliveryStreamS3DestinationConfiguration s3Configuration)
        {
            CloudWatchLoggingOptions = cloudWatchLoggingOptions;
            HECAcknowledgmentTimeoutInSeconds = hECAcknowledgmentTimeoutInSeconds;
            HECEndpoint = hECEndpoint;
            HECEndpointType = hECEndpointType;
            HECToken = hECToken;
            ProcessingConfiguration = processingConfiguration;
            RetryOptions = retryOptions;
            S3BackupMode = s3BackupMode;
            S3Configuration = s3Configuration;
        }
    }
}
