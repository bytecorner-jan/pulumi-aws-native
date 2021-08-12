// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTWireless.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-lorawanupdategatewaytaskentry.html
    /// </summary>
    public sealed class TaskDefinitionLoRaWANUpdateGatewayTaskEntryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-lorawanupdategatewaytaskentry.html#cfn-iotwireless-taskdefinition-lorawanupdategatewaytaskentry-currentversion
        /// </summary>
        [Input("currentVersion")]
        public Input<Inputs.TaskDefinitionLoRaWANGatewayVersionArgs>? CurrentVersion { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-lorawanupdategatewaytaskentry.html#cfn-iotwireless-taskdefinition-lorawanupdategatewaytaskentry-updateversion
        /// </summary>
        [Input("updateVersion")]
        public Input<Inputs.TaskDefinitionLoRaWANGatewayVersionArgs>? UpdateVersion { get; set; }

        public TaskDefinitionLoRaWANUpdateGatewayTaskEntryArgs()
        {
        }
    }
}
