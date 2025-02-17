// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway.Outputs
{

    [OutputType]
    public sealed class DeploymentMethodSetting
    {
        public readonly bool? CacheDataEncrypted;
        public readonly int? CacheTtlInSeconds;
        public readonly bool? CachingEnabled;
        public readonly bool? DataTraceEnabled;
        public readonly string? HttpMethod;
        public readonly string? LoggingLevel;
        public readonly bool? MetricsEnabled;
        public readonly string? ResourcePath;
        public readonly int? ThrottlingBurstLimit;
        public readonly double? ThrottlingRateLimit;

        [OutputConstructor]
        private DeploymentMethodSetting(
            bool? cacheDataEncrypted,

            int? cacheTtlInSeconds,

            bool? cachingEnabled,

            bool? dataTraceEnabled,

            string? httpMethod,

            string? loggingLevel,

            bool? metricsEnabled,

            string? resourcePath,

            int? throttlingBurstLimit,

            double? throttlingRateLimit)
        {
            CacheDataEncrypted = cacheDataEncrypted;
            CacheTtlInSeconds = cacheTtlInSeconds;
            CachingEnabled = cachingEnabled;
            DataTraceEnabled = dataTraceEnabled;
            HttpMethod = httpMethod;
            LoggingLevel = loggingLevel;
            MetricsEnabled = metricsEnabled;
            ResourcePath = resourcePath;
            ThrottlingBurstLimit = throttlingBurstLimit;
            ThrottlingRateLimit = throttlingRateLimit;
        }
    }
}
