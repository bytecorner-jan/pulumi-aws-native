// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElastiCache.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-globalreplicationgroup-reshardingconfiguration.html
    /// </summary>
    public sealed class GlobalReplicationGroupReshardingConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-globalreplicationgroup-reshardingconfiguration.html#cfn-elasticache-globalreplicationgroup-reshardingconfiguration-nodegroupid
        /// </summary>
        [Input("nodeGroupId")]
        public Input<string>? NodeGroupId { get; set; }

        [Input("preferredAvailabilityZones")]
        private InputList<string>? _preferredAvailabilityZones;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-globalreplicationgroup-reshardingconfiguration.html#cfn-elasticache-globalreplicationgroup-reshardingconfiguration-preferredavailabilityzones
        /// </summary>
        public InputList<string> PreferredAvailabilityZones
        {
            get => _preferredAvailabilityZones ?? (_preferredAvailabilityZones = new InputList<string>());
            set => _preferredAvailabilityZones = value;
        }

        public GlobalReplicationGroupReshardingConfigurationArgs()
        {
        }
    }
}
