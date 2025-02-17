// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDB.Inputs
{

    public sealed class GlobalTableReplicaSpecificationArgs : Pulumi.ResourceArgs
    {
        [Input("contributorInsightsSpecification")]
        public Input<Inputs.GlobalTableContributorInsightsSpecificationArgs>? ContributorInsightsSpecification { get; set; }

        [Input("globalSecondaryIndexes")]
        private InputList<Inputs.GlobalTableReplicaGlobalSecondaryIndexSpecificationArgs>? _globalSecondaryIndexes;
        public InputList<Inputs.GlobalTableReplicaGlobalSecondaryIndexSpecificationArgs> GlobalSecondaryIndexes
        {
            get => _globalSecondaryIndexes ?? (_globalSecondaryIndexes = new InputList<Inputs.GlobalTableReplicaGlobalSecondaryIndexSpecificationArgs>());
            set => _globalSecondaryIndexes = value;
        }

        [Input("pointInTimeRecoverySpecification")]
        public Input<Inputs.GlobalTablePointInTimeRecoverySpecificationArgs>? PointInTimeRecoverySpecification { get; set; }

        [Input("readProvisionedThroughputSettings")]
        public Input<Inputs.GlobalTableReadProvisionedThroughputSettingsArgs>? ReadProvisionedThroughputSettings { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("sSESpecification")]
        public Input<Inputs.GlobalTableReplicaSSESpecificationArgs>? SSESpecification { get; set; }

        [Input("tags")]
        private InputList<Inputs.GlobalTableTagArgs>? _tags;
        public InputList<Inputs.GlobalTableTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GlobalTableTagArgs>());
            set => _tags = value;
        }

        public GlobalTableReplicaSpecificationArgs()
        {
        }
    }
}
