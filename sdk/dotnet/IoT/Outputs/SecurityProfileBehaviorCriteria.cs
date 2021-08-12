// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behaviorcriteria.html
    /// </summary>
    [OutputType]
    public sealed class SecurityProfileBehaviorCriteria
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behaviorcriteria.html#cfn-iot-securityprofile-behaviorcriteria-comparisonoperator
        /// </summary>
        public readonly string? ComparisonOperator;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behaviorcriteria.html#cfn-iot-securityprofile-behaviorcriteria-consecutivedatapointstoalarm
        /// </summary>
        public readonly int? ConsecutiveDatapointsToAlarm;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behaviorcriteria.html#cfn-iot-securityprofile-behaviorcriteria-consecutivedatapointstoclear
        /// </summary>
        public readonly int? ConsecutiveDatapointsToClear;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behaviorcriteria.html#cfn-iot-securityprofile-behaviorcriteria-durationseconds
        /// </summary>
        public readonly int? DurationSeconds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behaviorcriteria.html#cfn-iot-securityprofile-behaviorcriteria-mldetectionconfig
        /// </summary>
        public readonly Outputs.SecurityProfileMachineLearningDetectionConfig? MlDetectionConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behaviorcriteria.html#cfn-iot-securityprofile-behaviorcriteria-statisticalthreshold
        /// </summary>
        public readonly Outputs.SecurityProfileStatisticalThreshold? StatisticalThreshold;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behaviorcriteria.html#cfn-iot-securityprofile-behaviorcriteria-value
        /// </summary>
        public readonly Outputs.SecurityProfileMetricValue? Value;

        [OutputConstructor]
        private SecurityProfileBehaviorCriteria(
            string? comparisonOperator,

            int? consecutiveDatapointsToAlarm,

            int? consecutiveDatapointsToClear,

            int? durationSeconds,

            Outputs.SecurityProfileMachineLearningDetectionConfig? mlDetectionConfig,

            Outputs.SecurityProfileStatisticalThreshold? statisticalThreshold,

            Outputs.SecurityProfileMetricValue? value)
        {
            ComparisonOperator = comparisonOperator;
            ConsecutiveDatapointsToAlarm = consecutiveDatapointsToAlarm;
            ConsecutiveDatapointsToClear = consecutiveDatapointsToClear;
            DurationSeconds = durationSeconds;
            MlDetectionConfig = mlDetectionConfig;
            StatisticalThreshold = statisticalThreshold;
            Value = value;
        }
    }
}
