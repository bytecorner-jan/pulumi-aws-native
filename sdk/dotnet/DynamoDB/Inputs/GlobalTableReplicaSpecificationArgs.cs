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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html
    /// </summary>
    public sealed class GlobalTableReplicaSpecificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-contributorinsightsspecification
        /// </summary>
        [Input("contributorInsightsSpecification")]
        public Input<Inputs.GlobalTableContributorInsightsSpecificationArgs>? ContributorInsightsSpecification { get; set; }

        [Input("globalSecondaryIndexes")]
        private InputList<Inputs.GlobalTableReplicaGlobalSecondaryIndexSpecificationArgs>? _globalSecondaryIndexes;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-globalsecondaryindexes
        /// </summary>
        public InputList<Inputs.GlobalTableReplicaGlobalSecondaryIndexSpecificationArgs> GlobalSecondaryIndexes
        {
            get => _globalSecondaryIndexes ?? (_globalSecondaryIndexes = new InputList<Inputs.GlobalTableReplicaGlobalSecondaryIndexSpecificationArgs>());
            set => _globalSecondaryIndexes = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-pointintimerecoveryspecification
        /// </summary>
        [Input("pointInTimeRecoverySpecification")]
        public Input<Inputs.GlobalTablePointInTimeRecoverySpecificationArgs>? PointInTimeRecoverySpecification { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-readprovisionedthroughputsettings
        /// </summary>
        [Input("readProvisionedThroughputSettings")]
        public Input<Inputs.GlobalTableReadProvisionedThroughputSettingsArgs>? ReadProvisionedThroughputSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-region
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-ssespecification
        /// </summary>
        [Input("sSESpecification")]
        public Input<Inputs.GlobalTableReplicaSSESpecificationArgs>? SSESpecification { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-tags
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public GlobalTableReplicaSpecificationArgs()
        {
        }
    }
}
