// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECS.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedscaling.html
    /// </summary>
    [OutputType]
    public sealed class CapacityProviderManagedScaling
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedscaling.html#cfn-ecs-capacityprovider-managedscaling-instancewarmupperiod
        /// </summary>
        public readonly int? InstanceWarmupPeriod;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedscaling.html#cfn-ecs-capacityprovider-managedscaling-maximumscalingstepsize
        /// </summary>
        public readonly int? MaximumScalingStepSize;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedscaling.html#cfn-ecs-capacityprovider-managedscaling-minimumscalingstepsize
        /// </summary>
        public readonly int? MinimumScalingStepSize;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedscaling.html#cfn-ecs-capacityprovider-managedscaling-status
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedscaling.html#cfn-ecs-capacityprovider-managedscaling-targetcapacity
        /// </summary>
        public readonly int? TargetCapacity;

        [OutputConstructor]
        private CapacityProviderManagedScaling(
            int? instanceWarmupPeriod,

            int? maximumScalingStepSize,

            int? minimumScalingStepSize,

            string? status,

            int? targetCapacity)
        {
            InstanceWarmupPeriod = instanceWarmupPeriod;
            MaximumScalingStepSize = maximumScalingStepSize;
            MinimumScalingStepSize = minimumScalingStepSize;
            Status = status;
            TargetCapacity = targetCapacity;
        }
    }
}
