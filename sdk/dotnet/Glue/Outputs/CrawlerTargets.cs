// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Outputs
{

    [OutputType]
    public sealed class CrawlerTargets
    {
        public readonly ImmutableArray<Outputs.CrawlerCatalogTarget> CatalogTargets;
        public readonly ImmutableArray<Outputs.CrawlerDynamoDBTarget> DynamoDBTargets;
        public readonly ImmutableArray<Outputs.CrawlerJdbcTarget> JdbcTargets;
        public readonly ImmutableArray<Outputs.CrawlerS3Target> S3Targets;

        [OutputConstructor]
        private CrawlerTargets(
            ImmutableArray<Outputs.CrawlerCatalogTarget> catalogTargets,

            ImmutableArray<Outputs.CrawlerDynamoDBTarget> dynamoDBTargets,

            ImmutableArray<Outputs.CrawlerJdbcTarget> jdbcTargets,

            ImmutableArray<Outputs.CrawlerS3Target> s3Targets)
        {
            CatalogTargets = catalogTargets;
            DynamoDBTargets = dynamoDBTargets;
            JdbcTargets = jdbcTargets;
            S3Targets = s3Targets;
        }
    }
}
