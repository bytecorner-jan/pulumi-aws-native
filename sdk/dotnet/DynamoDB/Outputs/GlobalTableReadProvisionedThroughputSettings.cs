// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDB.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-readprovisionedthroughputsettings.html
    /// </summary>
    [OutputType]
    public sealed class GlobalTableReadProvisionedThroughputSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-readprovisionedthroughputsettings.html#cfn-dynamodb-globaltable-readprovisionedthroughputsettings-readcapacityautoscalingsettings
        /// </summary>
        public readonly Outputs.GlobalTableCapacityAutoScalingSettings? ReadCapacityAutoScalingSettings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-readprovisionedthroughputsettings.html#cfn-dynamodb-globaltable-readprovisionedthroughputsettings-readcapacityunits
        /// </summary>
        public readonly int? ReadCapacityUnits;

        [OutputConstructor]
        private GlobalTableReadProvisionedThroughputSettings(
            Outputs.GlobalTableCapacityAutoScalingSettings? readCapacityAutoScalingSettings,

            int? readCapacityUnits)
        {
            ReadCapacityAutoScalingSettings = readCapacityAutoScalingSettings;
            ReadCapacityUnits = readCapacityUnits;
        }
    }
}
