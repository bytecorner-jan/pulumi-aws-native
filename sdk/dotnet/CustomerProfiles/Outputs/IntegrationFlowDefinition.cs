// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-flowdefinition.html
    /// </summary>
    [OutputType]
    public sealed class IntegrationFlowDefinition
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-flowdefinition.html#cfn-customerprofiles-integration-flowdefinition-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-flowdefinition.html#cfn-customerprofiles-integration-flowdefinition-flowname
        /// </summary>
        public readonly string FlowName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-flowdefinition.html#cfn-customerprofiles-integration-flowdefinition-kmsarn
        /// </summary>
        public readonly string KmsArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-flowdefinition.html#cfn-customerprofiles-integration-flowdefinition-sourceflowconfig
        /// </summary>
        public readonly Outputs.IntegrationSourceFlowConfig SourceFlowConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-flowdefinition.html#cfn-customerprofiles-integration-flowdefinition-tasks
        /// </summary>
        public readonly ImmutableArray<Outputs.IntegrationTask> Tasks;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-flowdefinition.html#cfn-customerprofiles-integration-flowdefinition-triggerconfig
        /// </summary>
        public readonly Outputs.IntegrationTriggerConfig TriggerConfig;

        [OutputConstructor]
        private IntegrationFlowDefinition(
            string? description,

            string flowName,

            string kmsArn,

            Outputs.IntegrationSourceFlowConfig sourceFlowConfig,

            ImmutableArray<Outputs.IntegrationTask> tasks,

            Outputs.IntegrationTriggerConfig triggerConfig)
        {
            Description = description;
            FlowName = flowName;
            KmsArn = kmsArn;
            SourceFlowConfig = sourceFlowConfig;
            Tasks = tasks;
            TriggerConfig = triggerConfig;
        }
    }
}
