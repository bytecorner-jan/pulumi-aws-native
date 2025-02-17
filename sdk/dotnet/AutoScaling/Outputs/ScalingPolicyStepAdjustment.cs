// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Outputs
{

    [OutputType]
    public sealed class ScalingPolicyStepAdjustment
    {
        public readonly double? MetricIntervalLowerBound;
        public readonly double? MetricIntervalUpperBound;
        public readonly int ScalingAdjustment;

        [OutputConstructor]
        private ScalingPolicyStepAdjustment(
            double? metricIntervalLowerBound,

            double? metricIntervalUpperBound,

            int scalingAdjustment)
        {
            MetricIntervalLowerBound = metricIntervalLowerBound;
            MetricIntervalUpperBound = metricIntervalUpperBound;
            ScalingAdjustment = scalingAdjustment;
        }
    }
}
