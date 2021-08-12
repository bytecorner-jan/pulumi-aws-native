// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDB.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaglobalsecondaryindexspecification.html
    /// </summary>
    public sealed class GlobalTableReplicaGlobalSecondaryIndexSpecificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaglobalsecondaryindexspecification.html#cfn-dynamodb-globaltable-replicaglobalsecondaryindexspecification-contributorinsightsspecification
        /// </summary>
        [Input("contributorInsightsSpecification")]
        public Input<Inputs.GlobalTableContributorInsightsSpecificationArgs>? ContributorInsightsSpecification { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaglobalsecondaryindexspecification.html#cfn-dynamodb-globaltable-replicaglobalsecondaryindexspecification-indexname
        /// </summary>
        [Input("indexName", required: true)]
        public Input<string> IndexName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaglobalsecondaryindexspecification.html#cfn-dynamodb-globaltable-replicaglobalsecondaryindexspecification-readprovisionedthroughputsettings
        /// </summary>
        [Input("readProvisionedThroughputSettings")]
        public Input<Inputs.GlobalTableReadProvisionedThroughputSettingsArgs>? ReadProvisionedThroughputSettings { get; set; }

        public GlobalTableReplicaGlobalSecondaryIndexSpecificationArgs()
        {
        }
    }
}
