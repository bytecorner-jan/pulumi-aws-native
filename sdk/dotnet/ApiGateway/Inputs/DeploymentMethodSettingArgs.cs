// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway.Inputs
{

    public sealed class DeploymentMethodSettingArgs : Pulumi.ResourceArgs
    {
        [Input("cacheDataEncrypted")]
        public Input<bool>? CacheDataEncrypted { get; set; }

        [Input("cacheTtlInSeconds")]
        public Input<int>? CacheTtlInSeconds { get; set; }

        [Input("cachingEnabled")]
        public Input<bool>? CachingEnabled { get; set; }

        [Input("dataTraceEnabled")]
        public Input<bool>? DataTraceEnabled { get; set; }

        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        [Input("loggingLevel")]
        public Input<string>? LoggingLevel { get; set; }

        [Input("metricsEnabled")]
        public Input<bool>? MetricsEnabled { get; set; }

        [Input("resourcePath")]
        public Input<string>? ResourcePath { get; set; }

        [Input("throttlingBurstLimit")]
        public Input<int>? ThrottlingBurstLimit { get; set; }

        [Input("throttlingRateLimit")]
        public Input<double>? ThrottlingRateLimit { get; set; }

        public DeploymentMethodSettingArgs()
        {
        }
    }
}
